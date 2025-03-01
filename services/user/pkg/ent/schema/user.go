package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"federation"
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
			Unique(),
		field.String("password").
			Comment("해시화된 비밀 번호"),
		field.String("name").
			Comment("이름"),
		field.Time("created_at").
			Comment("생성 날짜").
			Default(time.Now).
			Immutable(),
		field.Enum("role").
			Values("admin", "user").
			Comment("인가 권한"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		federation.GraphKeyDirective("id"),
	}
}
