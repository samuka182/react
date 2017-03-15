package main

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/myitcv/immutable"
)

type immStruct struct {
	fset *token.FileSet

	name string
	dec  *ast.GenDecl
	st   *ast.StructType

	special bool
}

func (o *output) genImmStructs(structs []immStruct) {
	type genField struct {
		Name string
		Type string
		f    *ast.Field
	}

	for _, s := range structs {

		o.printCommentGroup(s.dec.Doc)
		o.printImmPreamble(s.name, s.st)

		// start of struct
		o.pfln("type %v struct {", s.name)

		o.printLeadSpecCommsFor(s.st)

		o.pln("")

		var fields []genField

		for _, f := range s.st.Fields.List {
			names := ""
			sep := ""

			typ := o.exprString(f.Type)

			tag := ""

			if f.Tag != nil {
				tag = f.Tag.Value
			}

			if len(f.Names) == 0 {
				n := strings.TrimPrefix(typ, "*")
				names = fieldHidingPrefix + n

				fields = append(fields, genField{
					Name: n,
					Type: typ,
					f:    f,
				})
			} else {
				for _, n := range f.Names {
					names = names + sep + fieldHidingPrefix + n.Name
					sep = ", "

					fields = append(fields, genField{
						Name: n.Name,
						Type: typ,
						f:    f,
					})
				}
			}
			o.pfln("%v %v %v", names, typ, tag)
		}

		o.pln("")
		o.pln("mutable bool")
		o.pfln("__tmpl %v%v", immutable.ImmTypeTmplPrefix, s.name)

		// end of struct
		o.pfln("}")

		o.pln()

		o.pfln("var _ immutable.Immutable = &%v{}", s.name)
		o.pln()

		exp := exporter(s.name)

		o.pt(`
		func (s *{{.}}) AsMutable() *{{.}} {
			if s.Mutable() {
				return s
			}

			res := *s
		`, exp, s.name)
		if s.special {
			o.pt(`
			res._Key.Version++
			`, exp, nil)
		}
		o.pt(`
			res.mutable = true
			return &res
		}

		func (s *{{.}}) AsImmutable(v *{{.}}) *{{.}} {
			if s == nil {
				return nil
			}

			if s == v {
				return s
			}

			s.mutable = false
			return s
		}

		func (s *{{.}}) Mutable() bool {
			return s.mutable
		}

		func (s *{{.}}) WithMutable(f func(si *{{.}})) *{{.}} {
			res := s.AsMutable()
			f(res)
			res = res.AsImmutable(s)

			return res
		}

		func (s *{{.}}) WithImmutable(f func(si *{{.}})) *{{.}} {
			prev := s.mutable
			s.mutable = false
			f(s)
			s.mutable = prev

			return s
		}
		`, exp, s.name)

		for _, f := range fields {
			tmpl := struct {
				TypeName string
				Field    genField
			}{
				TypeName: s.name,
				Field:    f,
			}

			exp := exporter(f.Name)

			o.printCommentGroup(f.f.Doc)

			o.pt(`
			func (s *{{.TypeName}}) {{.Field.Name}}() {{.Field.Type}} {
				return s.`+fieldHidingPrefix+`{{.Field.Name}}
			}

			// {{Export "Set"}}{{Capitalise .Field.Name}} is the setter for {{Capitalise .Field.Name}}()
			func (s *{{.TypeName}}) {{Export "Set"}}{{Capitalise .Field.Name}}(n {{.Field.Type}}) *{{.TypeName}} {
				if s.mutable {
					s.`+fieldHidingPrefix+`{{.Field.Name}} = n
					return s
				}

				res := *s
			`, exp, tmpl)
			if s.special {
				o.pt(`
				res._Key.Version++
				`, exp, tmpl)
			}
			o.pt(`
				res.`+fieldHidingPrefix+`{{.Field.Name}} = n
				return &res
			}
			`, exp, tmpl)
		}
	}
}

func (o *output) printLeadSpecCommsFor(st *ast.StructType) {

	var end token.Pos

	// we are looking for comments before the first field (if there is one)

	if f := st.Fields; f != nil && len(f.List) > 0 {
		end = f.List[0].End()
	} else {
		end = st.End()
	}

	for _, cg := range o.curFile.Comments {
		if cg.Pos() > st.Pos() && cg.End() < end {
			for _, c := range cg.List {
				if strings.HasPrefix(c.Text, "//") && !strings.HasPrefix(c.Text, "// ") {
					o.pln(c.Text)
				}
			}
		}
	}

}