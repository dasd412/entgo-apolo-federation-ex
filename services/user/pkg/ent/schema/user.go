package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/vektah/gqlparser/v2/ast"
	"time"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Comment("이메일").
			Unique().
			Annotations(
			//	GraphExternalDirective(),
			),
		field.String("password").
			Comment("해시화된 비밀 번호"),
		field.String("name").
			Comment("이름"),
		field.Time("created_at").
			Comment("생성 날짜").
			Default(time.Now).
			Immutable(),
		field.Enum("role").
			Values("admin", "author", "guest").
			Comment("인가 권한"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),

		GraphKeyDirective("id"),
	}
}

func GraphKeyDirective(fields string) entgql.Annotation {
	return entgql.Directives(keyDirective(fields))
}

func keyDirective(fields string) entgql.Directive {
	var args []*ast.Argument
	if fields != "" {
		args = append(args, &ast.Argument{
			Name: "fields",
			Value: &ast.Value{
				Raw:  fields,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.NewDirective("key", args...)
}

func GraphExternalDirective() entgql.Annotation {
	return entgql.Directives(entgql.NewDirective("external"))
}

func GraphProvidesDirective(fields string) entgql.Annotation {
	return entgql.Directives(keyDirective(fields))
}

func providesDirective(fields string) entgql.Directive {
	var args []*ast.Argument
	if fields != "" {
		args = append(args, &ast.Argument{
			Name: "fields",
			Value: &ast.Value{
				Raw:  fields,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.NewDirective("provides", args...)
}
