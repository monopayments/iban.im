package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Merge func to merge all the GraphQL files
// Arguments:
// - indent string : the padding to generate schema eg. "\t" or " "
// - paths : A relative path to find *.graphql or *.gql files recursively
func Merge(indent string, paths ...string) *string {
	schemas := make([]Schema, 0, len(paths))

	for _, path := range paths {
		if sc := getSchema(path); sc != nil {
			schemas = append(schemas, *sc)
		}
	}

	if len(schemas) == 0 {
		return nil
	}

	schema := joinSchemas(schemas)
	ms := MergedSchema{Indent: indent}
	ss := ms.StitchSchema(schema)
	return &ss
}

func getSchema(path string) *Schema {
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	sc := &Schema{}
	// at this moment, path should be an absolute path
	sc.GetSchema(abs)

	if len(sc.Files) == 0 {
		return nil
	}

	for _, file := range sc.Files {
		l := NewLexer(file)
		sc.ParseSchema(l)
	}

	return sc
}

func joinSchemas(schemas []Schema) *Schema {
	schema := Schema{}

	for _, s := range schemas {
		schema.Files = append(schema.Files, s.Files...)
		schema.Mutations = append(schema.Mutations, s.Mutations...)
		schema.Queries = append(schema.Queries, s.Queries...)
		schema.Subscriptions = append(schema.Subscriptions, s.Subscriptions...)
		schema.TypeNames = append(schema.TypeNames, s.TypeNames...)
		schema.Scalars = append(schema.Scalars, s.Scalars...)
		schema.Enums = append(schema.Enums, s.Enums...)
		schema.Interfaces = append(schema.Interfaces, s.Interfaces...)
		schema.Unions = append(schema.Unions, s.Unions...)
		schema.Inputs = append(schema.Inputs, s.Inputs...)
	}

	wg := sync.WaitGroup{}
	wg.Add(8)

	go schema.UniqueMutation(&wg)
	go schema.UniqueQuery(&wg)
	go schema.UniqueTypeName(&wg)
	go schema.UniqueScalar(&wg)
	go schema.UniqueEnum(&wg)
	go schema.UniqueInterface(&wg)
	go schema.UniqueUnion(&wg)
	go schema.UniqueInput(&wg)

	wg.Wait()

	return &schema
}
