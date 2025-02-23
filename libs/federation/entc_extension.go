package federation

import (
	"entgo.io/ent/entc/gen"
	"errors"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	RemoveNodeGoModel = func(_ *gen.Graph, s *ast.Schema) error {
		n, ok := s.Types["Node"]
		if !ok {
			return errors.New("failed to find node interface in schema")
		}

		dirs := ast.DirectiveList{}

		for _, d := range n.Directives {
			switch d.Name {
			case "goModel":
				continue
			default:
				dirs = append(dirs, d)
			}
		}
		n.Directives = dirs

		return nil
	}

	RemoveNodeQueries = func(_ *gen.Graph, s *ast.Schema) error {
		q, ok := s.Types["Query"]
		if !ok {
			return errors.New("failed to find query definition in schema")
		}

		fields := ast.FieldList{}

		for _, f := range q.Fields {
			switch f.Name {
			case "node":
			case "nodes":
				continue
			default:
				fields = append(fields, f)
			}
		}
		q.Fields = fields

		return nil
	}

	SetPageInfoShareable = func(_ *gen.Graph, s *ast.Schema) error {
		q, ok := s.Types["PageInfo"]
		if !ok {
			return nil
		}

		q.Directives = append(q.Directives, &ast.Directive{Name: "shareable"})

		return nil
	}
)
