package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetSchema is to parse ./schema/**/*.graphql
func (sc *Schema) GetSchema(path string) {
	// FIX: is there any way to use a relative path?
	// currently, it works only with absolute path
	// in case of using a relative path such as '../schema', it spits out an error
	// the error says invalid memory or nil pointer deference.
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p == "" {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(p)
		if ext != ".graphql" && ext != ".gql" {
			return nil
		}

		file, err := os.Open(p)
		if err != nil {
			fmt.Printf("[Error] There is an error to open %s", p)
			return err
		}

		// TODO: split and get a only filename and print it to user
		// needs to handle in case of OS (windows / unix compatibles)
		sc.Files = append(sc.Files, file)

		return nil
	})
	if err != nil {
		panic(err)
	}

	// try to get relative path
	rel, err := filepath.Rel(os.Getenv("PWD"), path)
	if err == nil {
		// if failed print absolute path
		path = rel
	}

	if len(sc.Files) > 0 {
		fmt.Printf("🎉 [%s] Total %d GraphQL files found!\n", path, len(sc.Files))
	}
}
