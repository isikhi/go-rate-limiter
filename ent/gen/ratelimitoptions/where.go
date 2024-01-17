// Code generated by ent, DO NOT EDIT.

package ratelimitoptions

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/isikhi/go-rate-limiter/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLTE(FieldID, id))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldClientID, v))
}

// TokenCount applies equality check predicate on the "token_count" field. It's identical to TokenCountEQ.
func TokenCount(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldTokenCount, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldDuration, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldCreatedAt, v))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLTE(FieldClientID, v))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContains(FieldClientID, v))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasPrefix(FieldClientID, v))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasSuffix(FieldClientID, v))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEqualFold(FieldClientID, v))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContainsFold(FieldClientID, v))
}

// TokenCountEQ applies the EQ predicate on the "token_count" field.
func TokenCountEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldTokenCount, v))
}

// TokenCountNEQ applies the NEQ predicate on the "token_count" field.
func TokenCountNEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNEQ(FieldTokenCount, v))
}

// TokenCountIn applies the In predicate on the "token_count" field.
func TokenCountIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIn(FieldTokenCount, vs...))
}

// TokenCountNotIn applies the NotIn predicate on the "token_count" field.
func TokenCountNotIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotIn(FieldTokenCount, vs...))
}

// TokenCountGT applies the GT predicate on the "token_count" field.
func TokenCountGT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGT(FieldTokenCount, v))
}

// TokenCountGTE applies the GTE predicate on the "token_count" field.
func TokenCountGTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGTE(FieldTokenCount, v))
}

// TokenCountLT applies the LT predicate on the "token_count" field.
func TokenCountLT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLT(FieldTokenCount, v))
}

// TokenCountLTE applies the LTE predicate on the "token_count" field.
func TokenCountLTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLTE(FieldTokenCount, v))
}

// TokenCountContains applies the Contains predicate on the "token_count" field.
func TokenCountContains(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContains(FieldTokenCount, v))
}

// TokenCountHasPrefix applies the HasPrefix predicate on the "token_count" field.
func TokenCountHasPrefix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasPrefix(FieldTokenCount, v))
}

// TokenCountHasSuffix applies the HasSuffix predicate on the "token_count" field.
func TokenCountHasSuffix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasSuffix(FieldTokenCount, v))
}

// TokenCountIsNil applies the IsNil predicate on the "token_count" field.
func TokenCountIsNil() predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIsNull(FieldTokenCount))
}

// TokenCountNotNil applies the NotNil predicate on the "token_count" field.
func TokenCountNotNil() predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotNull(FieldTokenCount))
}

// TokenCountEqualFold applies the EqualFold predicate on the "token_count" field.
func TokenCountEqualFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEqualFold(FieldTokenCount, v))
}

// TokenCountContainsFold applies the ContainsFold predicate on the "token_count" field.
func TokenCountContainsFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContainsFold(FieldTokenCount, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLTE(FieldDuration, v))
}

// DurationContains applies the Contains predicate on the "duration" field.
func DurationContains(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContains(FieldDuration, v))
}

// DurationHasPrefix applies the HasPrefix predicate on the "duration" field.
func DurationHasPrefix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasPrefix(FieldDuration, v))
}

// DurationHasSuffix applies the HasSuffix predicate on the "duration" field.
func DurationHasSuffix(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldHasSuffix(FieldDuration, v))
}

// DurationEqualFold applies the EqualFold predicate on the "duration" field.
func DurationEqualFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEqualFold(FieldDuration, v))
}

// DurationContainsFold applies the ContainsFold predicate on the "duration" field.
func DurationContainsFold(v string) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldContainsFold(FieldDuration, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.RateLimitOptions {
	return predicate.RateLimitOptions(sql.FieldNotNull(FieldCreatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RateLimitOptions) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RateLimitOptions) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RateLimitOptions) predicate.RateLimitOptions {
	return predicate.RateLimitOptions(func(s *sql.Selector) {
		p(s.Not())
	})
}
