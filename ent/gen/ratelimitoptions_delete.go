// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isikhi/go-rate-limiter/ent/gen/predicate"
	"github.com/isikhi/go-rate-limiter/ent/gen/ratelimitoptions"
)

// RateLimitOptionsDelete is the builder for deleting a RateLimitOptions entity.
type RateLimitOptionsDelete struct {
	config
	hooks    []Hook
	mutation *RateLimitOptionsMutation
}

// Where appends a list predicates to the RateLimitOptionsDelete builder.
func (rlod *RateLimitOptionsDelete) Where(ps ...predicate.RateLimitOptions) *RateLimitOptionsDelete {
	rlod.mutation.Where(ps...)
	return rlod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rlod *RateLimitOptionsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rlod.sqlExec, rlod.mutation, rlod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rlod *RateLimitOptionsDelete) ExecX(ctx context.Context) int {
	n, err := rlod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rlod *RateLimitOptionsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ratelimitoptions.Table, sqlgraph.NewFieldSpec(ratelimitoptions.FieldID, field.TypeUint64))
	if ps := rlod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rlod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rlod.mutation.done = true
	return affected, err
}

// RateLimitOptionsDeleteOne is the builder for deleting a single RateLimitOptions entity.
type RateLimitOptionsDeleteOne struct {
	rlod *RateLimitOptionsDelete
}

// Where appends a list predicates to the RateLimitOptionsDelete builder.
func (rlodo *RateLimitOptionsDeleteOne) Where(ps ...predicate.RateLimitOptions) *RateLimitOptionsDeleteOne {
	rlodo.rlod.mutation.Where(ps...)
	return rlodo
}

// Exec executes the deletion query.
func (rlodo *RateLimitOptionsDeleteOne) Exec(ctx context.Context) error {
	n, err := rlodo.rlod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ratelimitoptions.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rlodo *RateLimitOptionsDeleteOne) ExecX(ctx context.Context) {
	if err := rlodo.Exec(ctx); err != nil {
		panic(err)
	}
}
