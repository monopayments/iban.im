package lib

import (
	"sync"
	"text/scanner"
)

func (s *Schema) ParseSchema(l *Lexer) {
	l.ConsumeWhitespace()

	for l.Peek() != scanner.EOF {

		switch x := l.ConsumeIdent(); x {

		case "schema":
			// skip the schema { ... }
			// it will be generated after parsing all
			for {
				l.next = l.sc.Scan()

				if l.next == '}' {
					break
				}
			}
			l.ConsumeToken('}')

		case "scalar":
			c := Scalar{}
			c.Name = l.ConsumeIdent()
			s.Scalars = append(s.Scalars, &c)

		case "enum":
			e := Enum{}
			e.Name = l.ConsumeIdent()
			l.ConsumeToken('{')
			for l.Peek() != '}' {
				e.Fields = append(e.Fields, l.ConsumeIdent())
			}
			l.ConsumeToken('}')
			s.Enums = append(s.Enums, &e)

		case "interface":
			i := Interface{}
			i.Name = l.ConsumeIdent()

			l.ConsumeToken('{')

			for l.Peek() != '}' {
				p := Prop{}
				p.Name = l.ConsumeIdent()

				p.Args = ParseArgument(l)

				l.ConsumeToken(':')

				if l.Peek() == '[' {
					p.IsList = true
					l.ConsumeToken('[')
					p.Type = l.ConsumeIdent()
					if x := l.sc.TokenText(); x == "!" {
						p.Null = false
						l.ConsumeToken('!')
					} else {
						p.Null = true
					}
					l.ConsumeToken(']')
					if x := l.sc.TokenText(); x == "!" {
						p.IsListNull = false
						l.ConsumeToken('!')
					} else {
						p.IsListNull = true
					}
				} else {
					p.IsList = false
					p.IsListNull = false
					p.Type = l.ConsumeIdent()
					if x := l.sc.TokenText(); x == "!" {
						p.Null = false
						l.ConsumeToken('!')
					} else {
						p.Null = true
					}
				}

				if l.Peek() == '@' {
					l.ConsumeToken('@')
					l.ConsumeDirective()
					d := Directive{l.GetBuffer()}
					l.ConsumeWhitespace()
					p.Directive = &d
				}

				i.Props = append(i.Props, &p)
			}

			s.Interfaces = append(s.Interfaces, &i)
			l.ConsumeToken('}')

		case "union":
			u := Union{}
			u.Name = l.ConsumeIdent()
			l.ConsumeToken('=')
			for l.Peek() != '\r' || l.Peek() != '\n' || l.Peek() != scanner.EOF {
				u.Fields = append(u.Fields, l.ConsumeIdent())
				if l.Peek() == '|' {
					l.ConsumeToken('|')
				} else {
					break
				}
			}
			s.Unions = append(s.Unions, &u)

		case "input":
			i := Input{}
			i.Name = l.ConsumeIdent()

			l.ConsumeToken('{')

			for l.Peek() != '}' {
				p := Prop{}
				p.Name = l.ConsumeIdent()
				l.ConsumeToken(':')

				if l.Peek() == '[' {
					p.IsList = true
					l.ConsumeToken('[')
					p.Type = l.ConsumeIdent()
					if x := l.sc.TokenText(); x == "!" {
						p.Null = false
						l.ConsumeToken('!')
					} else {
						p.Null = true
					}
					l.ConsumeToken(']')
					if x := l.sc.TokenText(); x == "!" {
						p.IsListNull = false
						l.ConsumeToken('!')
					} else {
						p.IsListNull = true
					}
				} else {
					p.IsList = false
					p.IsListNull = false
					p.Type = l.ConsumeIdent()
					if x := l.sc.TokenText(); x == "!" {
						p.Null = false
						l.ConsumeToken('!')
					} else {
						p.Null = true
					}
				}

				if l.Peek() == '@' {
					l.ConsumeToken('@')
					l.ConsumeDirective()
					d := Directive{l.GetBuffer()}
					l.ConsumeWhitespace()
					p.Directive = &d
				}

				i.Props = append(i.Props, &p)
			}

			s.Inputs = append(s.Inputs, &i)
			l.ConsumeToken('}')

		case "type":

			switch x := l.ConsumeIdent(); x {

			case "Query":
				l.ConsumeToken('{')

				for l.Peek() != '}' {
					q := Query{}
					q.Name = l.ConsumeIdent()

					q.Args = ParseArgument(l)

					l.ConsumeToken(':')
					r := Resp{}
					if l.Peek() == '[' {
						r.IsList = true
						l.ConsumeToken('[')
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
						l.ConsumeToken(']')
						if x := l.sc.TokenText(); x == "!" {
							r.IsListNull = false
							l.ConsumeToken('!')
						} else {
							r.IsListNull = true
						}
					} else {
						r.IsList = false
						r.IsListNull = false
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
					}
					q.Resp = r

					if l.Peek() == '@' {
						l.ConsumeToken('@')
						l.ConsumeDirective()
						d := Directive{l.GetBuffer()}
						l.ConsumeWhitespace()
						q.Directive = &d
					}

					s.Queries = append(s.Queries, &q)
				}
				l.ConsumeToken('}')

			case "Mutation":
				l.ConsumeToken('{')

				for l.Peek() != '}' {
					m := Mutation{}
					m.Name = l.ConsumeIdent()

					m.Args = ParseArgument(l)

					l.ConsumeToken(':')
					r := Resp{}
					if l.Peek() == '[' {
						r.IsList = true
						l.ConsumeToken('[')
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
						l.ConsumeToken(']')
						if x := l.sc.TokenText(); x == "!" {
							r.IsListNull = false
							l.ConsumeToken('!')
						} else {
							r.IsListNull = true
						}
					} else {
						r.IsList = false
						r.IsListNull = false
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
					}

					m.Resp = r

					if l.Peek() == '@' {
						l.ConsumeToken('@')
						l.ConsumeDirective()
						d := Directive{l.GetBuffer()}
						l.ConsumeWhitespace()
						m.Directive = &d
					}

					s.Mutations = append(s.Mutations, &m)
				}
				l.ConsumeToken('}')

			case "Subscription":
				l.ConsumeToken('{')

				for l.Peek() != '}' {
					c := Subscription{}
					c.Name = l.ConsumeIdent()

					c.Args = ParseArgument(l)

					l.ConsumeToken(':')
					r := Resp{}
					if l.Peek() == '[' {
						r.IsList = true
						l.ConsumeToken('[')
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
						l.ConsumeToken(']')
						if x := l.sc.TokenText(); x == "!" {
							r.IsListNull = false
							l.ConsumeToken('!')
						} else {
							r.IsListNull = true
						}
					} else {
						r.IsList = false
						r.IsListNull = false
						r.Name = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							r.Null = false
							l.ConsumeToken('!')
						} else {
							r.Null = true
						}
					}
					c.Resp = r

					if l.Peek() == '@' {
						l.ConsumeToken('@')
						l.ConsumeDirective()
						d := Directive{l.GetBuffer()}
						l.ConsumeWhitespace()
						c.Directive = &d
					}

					s.Subscriptions = append(s.Subscriptions, &c)
				}
				l.ConsumeToken('}')

			default:
				t := TypeName{}
				t.Name = x

				// handling in case of type has implements
				if l.Peek() == scanner.Ident {
					l.ConsumeIdent()
					t.Impl = true
					x := l.ConsumeIdent()
					t.ImplType = &x
				} else {
					t.Impl = false
				}

				l.ConsumeToken('{')

				for l.Peek() != '}' {
					p := Prop{}
					p.Name = l.ConsumeIdent()

					p.Args = ParseArgument(l)

					l.ConsumeToken(':')

					if l.Peek() == '[' {
						p.IsList = true
						l.ConsumeToken('[')
						p.Type = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							p.Null = false
							l.ConsumeToken('!')
						} else {
							p.Null = true
						}
						l.ConsumeToken(']')
						if x := l.sc.TokenText(); x == "!" {
							p.IsListNull = false
							l.ConsumeToken('!')
						} else {
							p.IsListNull = true
						}
					} else {
						p.IsList = false
						p.IsListNull = false
						p.Type = l.ConsumeIdent()
						if x := l.sc.TokenText(); x == "!" {
							p.Null = false
							l.ConsumeToken('!')
						} else {
							p.Null = true
						}
					}

					if l.Peek() == '@' {
						l.ConsumeToken('@')
						l.ConsumeDirective()
						d := Directive{l.GetBuffer()}
						l.ConsumeWhitespace()
						p.Directive = &d
					}

					t.Props = append(t.Props, &p)
				}

				s.TypeNames = append(s.TypeNames, &t)
				l.ConsumeToken('}')
			}
		}
	}
}

func (s *Schema) UniqueMutation(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Mutations))
	for _, v := range s.Mutations {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Mutations[j] = v
		j++
	}
	s.Mutations = s.Mutations[:j]
}

func (s *Schema) UniqueQuery(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Queries))
	for _, v := range s.Queries {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Queries[j] = v
		j++
	}
	s.Queries = s.Queries[:j]
}

func (s *Schema) UniqueTypeName(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.TypeNames))
	for _, v := range s.TypeNames {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.TypeNames[j] = v
		j++
	}
	s.TypeNames = s.TypeNames[:j]
}

func (s *Schema) UniqueScalar(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Scalars))
	for _, v := range s.Scalars {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Scalars[j] = v
		j++
	}
	s.Scalars = s.Scalars[:j]
}

func (s *Schema) UniqueEnum(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Enums))
	for _, v := range s.Enums {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Enums[j] = v
		j++
	}
	s.Enums = s.Enums[:j]
}

func (s *Schema) UniqueInterface(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Interfaces))
	for _, v := range s.Interfaces {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Interfaces[j] = v
		j++
	}
	s.Interfaces = s.Interfaces[:j]
}

func (s *Schema) UniqueUnion(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Unions))
	for _, v := range s.Unions {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Unions[j] = v
		j++
	}
	s.Unions = s.Unions[:j]
}

func (s *Schema) UniqueInput(wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	seen := make(map[string]struct{}, len(s.Inputs))
	for _, v := range s.Inputs {
		if _, ok := seen[v.Name]; ok {
			continue
		}
		seen[v.Name] = struct{}{}
		s.Inputs[j] = v
		j++
	}
	s.Inputs = s.Inputs[:j]
}

func ParseArgument(l *Lexer) []*Arg {
	args := []*Arg{}

	for l.Peek() == '(' {
		l.ConsumeToken('(')
		for l.Peek() != ')' {
			arg := Arg{}
			arg.Param = l.ConsumeIdent()
			l.ConsumeToken(':')

			if l.Peek() == '[' {
				arg.IsList = true
				l.ConsumeToken('[')
				arg.Type = l.ConsumeIdent()
				if l.Peek() == '!' {
					arg.Null = false
					l.ConsumeToken('!')
				} else {
					arg.Null = true
				}
				l.ConsumeToken(']')

				if x := l.sc.TokenText(); x == "!" {
					arg.IsListNull = false
					l.ConsumeToken('!')
				} else {
					arg.IsListNull = true
				}
			} else {
				arg.Type = l.ConsumeIdent()

				if l.Peek() == '=' {
					l.ConsumeToken('=')
					ext := l.ConsumeIdent()
					arg.TypeExt = &ext
				}

				if x := l.sc.TokenText(); x == "!" {
					arg.Null = false
					l.ConsumeToken('!')
				} else {
					arg.Null = true
				}
			}

			args = append(args, &arg)
		}
		l.ConsumeToken(')')
	}
	return args
}
