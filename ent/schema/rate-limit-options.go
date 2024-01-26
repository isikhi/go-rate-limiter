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
		field.String("token_count"),
		field.String("duration_in_seconds"),
		field.String("throttle_percentage"),
		field.Time("created_at").Optional().StructTag(`json:"-"`),
	}
}
