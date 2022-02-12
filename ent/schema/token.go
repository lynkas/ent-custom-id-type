package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").Unique().MaxLen(320).NotEmpty(),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("token").Required().Unique(),
	}
}
