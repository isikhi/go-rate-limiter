package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type RateLimitOptions struct {
	ent.Schema
}

func (RateLimitOptions) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("client_id"),
		field.String("token_count").Optional(),
		field.String("duration"),
		field.Time("created_at").Optional().StructTag(`json:"-"`),
	}
}
