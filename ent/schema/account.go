package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().MaxLen(320).NotEmpty(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("token", Token.Type),
	}
}
