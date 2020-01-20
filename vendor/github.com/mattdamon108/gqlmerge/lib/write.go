package lib

import (
	"strings"
)

type MergedSchema struct {
	buf    strings.Builder
	Indent string
}

func (ms *MergedSchema) StitchSchema(s *Schema) string {
	numOfQurs := len(s.Queries)
	numOfMuts := len(s.Mutations)
	numOfSubs := len(s.Subscriptions)

	ms.buf.WriteString("schema {\n")
	if numOfQurs > 0 {
		ms.addIndent(1)
		ms.buf.WriteString("query: Query\n")
	}
	if numOfMuts > 0 {
		ms.addIndent(1)
		ms.buf.WriteString("mutation: Mutation\n")
	}
	if numOfSubs > 0 {
		ms.addIndent(1)
		ms.buf.WriteString("subscription: Subscription\n")
	}
	ms.buf.WriteString("}\n\n")

	if numOfQurs > 0 {
		ms.buf.WriteString(`type Query {
`)
		for _, q := range s.Queries {
			ms.addIndent(1)
			ms.buf.WriteString(q.Name)
			if l := len(q.Args); l > 0 {
				ms.buf.WriteString("(")
				if l > 2 {
					ms.buf.WriteString("\n")
				}

				for i, a := range q.Args {
					ms.stitchArgument(a, l, i)
				}

				if l > 2 {
					ms.buf.WriteString("\n")
					ms.addIndent(1)
				}
				ms.buf.WriteString(")")
			}
			ms.buf.WriteString(": ")
			if q.Resp.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(q.Resp.Name)
			if !q.Resp.Null {
				ms.buf.WriteString("!")
			}
			if q.Resp.IsList {
				ms.buf.WriteString("]")
			}
			if q.Resp.IsList && !q.Resp.IsListNull {
				ms.buf.WriteString("!")
			}

			if q.Directive != nil {
				ms.buf.WriteString(" @" + q.Directive.string)
			}

			ms.buf.WriteString("\n")
		}
		ms.buf.WriteString("}\n\n")
	}

	if numOfMuts > 0 {
		ms.buf.WriteString(`type Mutation {
`)
		for _, m := range s.Mutations {
			ms.addIndent(1)
			ms.buf.WriteString(m.Name)
			if l := len(m.Args); l > 0 {
				ms.buf.WriteString("(")
				if l > 2 {
					ms.buf.WriteString("\n")
				}

				for i, a := range m.Args {
					ms.stitchArgument(a, l, i)
				}

				if l > 2 {
					ms.buf.WriteString("\n")
					ms.addIndent(1)
				}
				ms.buf.WriteString(")")
			}
			ms.buf.WriteString(": ")
			if m.Resp.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(m.Resp.Name)
			if !m.Resp.Null {
				ms.buf.WriteString("!")
			}
			if m.Resp.IsList {
				ms.buf.WriteString("]")
			}
			if m.Resp.IsList && !m.Resp.IsListNull {
				ms.buf.WriteString("!")
			}

			if m.Directive != nil {
				ms.buf.WriteString(" @" + m.Directive.string)
			}

			ms.buf.WriteString("\n")
		}
		ms.buf.WriteString("}\n\n")
	}

	if numOfSubs > 0 {
		ms.buf.WriteString(`type Subscription {
`)
		for _, c := range s.Subscriptions {
			ms.addIndent(1)
			ms.buf.WriteString(c.Name)
			if l := len(c.Args); l > 0 {
				ms.buf.WriteString("(")
				if l > 2 {
					ms.buf.WriteString("\n")
				}

				for i, a := range c.Args {
					ms.stitchArgument(a, l, i)
				}

				if l > 2 {
					ms.buf.WriteString("\n")
					ms.addIndent(1)
				}
				ms.buf.WriteString(")")
			}
			ms.buf.WriteString(": ")
			if c.Resp.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(c.Resp.Name)
			if !c.Resp.Null {
				ms.buf.WriteString("!")
			}
			if c.Resp.IsList {
				ms.buf.WriteString("]")
			}
			if c.Resp.IsList && !c.Resp.IsListNull {
				ms.buf.WriteString("!")
			}

			if c.Directive != nil {
				ms.buf.WriteString(" @" + c.Directive.string)
			}

			ms.buf.WriteString("\n")
		}
		ms.buf.WriteString("}\n\n")
	}

	for i, t := range s.TypeNames {
		ms.buf.WriteString("type ")
		ms.buf.WriteString(t.Name)
		if t.Impl {
			ms.buf.WriteString(" implements " + *t.ImplType)
		}
		ms.buf.WriteString(" {\n")
		for _, p := range t.Props {
			ms.addIndent(1)
			ms.buf.WriteString(p.Name)

			if l := len(p.Args); l > 0 {
				ms.buf.WriteString("(")
				if l > 2 {
					ms.buf.WriteString("\n")
				}
				for i, a := range p.Args {
					ms.stitchArgument(a, l, i)
				}
				if l > 2 {
					ms.buf.WriteString("\n")
					ms.addIndent(1)
				}
				ms.buf.WriteString(")")
			}

			ms.buf.WriteString(": ")
			if p.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(p.Type)
			if !p.Null {
				ms.buf.WriteString("!")
			}
			if p.IsList {
				ms.buf.WriteString("]")
			}
			if p.IsList && !p.IsListNull {
				ms.buf.WriteString("!")
			}

			if p.Directive != nil {
				ms.buf.WriteString(" @" + p.Directive.string)
			}

			ms.buf.WriteString("\n")
		}
		ms.buf.WriteString("}\n")
		if i != len(s.TypeNames)-1 {
			ms.buf.WriteString("\n")
		}
	}
	ms.buf.WriteString("\n")

	for i, c := range s.Scalars {
		ms.buf.WriteString("scalar " + c.Name + "\n")
		if i != len(s.Scalars)-1 {
			ms.buf.WriteString("\n")
		}
	}
	ms.buf.WriteString("\n")

	for i, e := range s.Enums {
		ms.buf.WriteString("enum " + e.Name + " {\n")
		for _, n := range e.Fields {
			ms.addIndent(1)
			ms.buf.WriteString(n + "\n")
		}
		ms.buf.WriteString("}\n")
		if i != len(s.Enums)-1 {
			ms.buf.WriteString("\n")
		}
	}
	ms.buf.WriteString("\n")

	for j, i := range s.Interfaces {
		ms.buf.WriteString("interface " + i.Name + " {\n")

		for _, p := range i.Props {
			ms.addIndent(1)
			ms.buf.WriteString(p.Name)

			if l := len(p.Args); l > 0 {
				ms.buf.WriteString("(")
				if l > 2 {
					ms.buf.WriteString("\n")
				}
				for i, a := range p.Args {
					ms.stitchArgument(a, l, i)
				}
				if l > 2 {
					ms.buf.WriteString("\n")
					ms.addIndent(1)
				}
				ms.buf.WriteString(")")
			}

			ms.buf.WriteString(": ")
			if p.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(p.Type)
			if !p.Null {
				ms.buf.WriteString("!")
			}
			if p.IsList {
				ms.buf.WriteString("]")
			}
			if p.IsList && !p.IsListNull {
				ms.buf.WriteString("!")
			}

			if p.Directive != nil {
				ms.buf.WriteString(" @" + p.Directive.string)
			}

			ms.buf.WriteString("\n")
		}
		ms.buf.WriteString("}\n")
		if j < len(s.Interfaces)-1 {
			ms.buf.WriteString("\n")
		}
	}
	ms.buf.WriteString("\n")

	for _, u := range s.Unions {
		ms.buf.WriteString("union " + u.Name + " = ")
		fields := strings.Join(u.Fields, " | ")
		ms.buf.WriteString(fields + "\n\n")
	}

	for j, i := range s.Inputs {
		ms.buf.WriteString("input " + i.Name + " {\n")

		for _, p := range i.Props {
			ms.addIndent(1)
			ms.buf.WriteString(p.Name + ": ")
			if p.IsList {
				ms.buf.WriteString("[")
			}
			ms.buf.WriteString(p.Type)
			if !p.Null {
				ms.buf.WriteString("!")
			}
			if p.IsList {
				ms.buf.WriteString("]")
			}
			if p.IsList && !p.IsListNull {
				ms.buf.WriteString("!")
			}

			if p.Directive != nil {
				ms.buf.WriteString(" @" + p.Directive.string)
			}

			ms.buf.WriteString("\n")
		}

		ms.buf.WriteString("}\n")
		if j < len(s.Inputs)-1 {
			ms.buf.WriteString("\n")
		}
	}

	return ms.buf.String()
}

func (ms *MergedSchema) addIndent(n int) {
	i := strings.Repeat(ms.Indent, n)
	ms.buf.WriteString(i)
}

func (ms *MergedSchema) stitchArgument(a *Arg, l int, i int) {
	if l > 2 {
		ms.addIndent(2)
	}
	ms.buf.WriteString(a.Param + ": ")

	if a.IsList {
		ms.buf.WriteString("[")
		ms.buf.WriteString(a.Type)

		if !a.Null {
			ms.buf.WriteString("!")
		}
		ms.buf.WriteString("]")
		if !a.IsListNull {
			ms.buf.WriteString("!")
		}
	} else {
		ms.buf.WriteString(a.Type)
		if a.TypeExt != nil {
			ms.buf.WriteString(" = " + *a.TypeExt)
		}
		if !a.Null {
			ms.buf.WriteString("!")
		}
	}

	if l <= 2 && i != l-1 {
		ms.buf.WriteString(", ")
	}
	if l > 2 && i != l-1 {
		ms.buf.WriteString("\n")
	}
}
