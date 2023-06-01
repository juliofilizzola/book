package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("email").
			Unique(),
		field.String("nick").
			Unique(),
		field.String("password"),
		field.String("following"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("following", User.Type).Ref("following"),
	}
}
