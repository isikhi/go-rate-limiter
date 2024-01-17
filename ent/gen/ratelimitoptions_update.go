// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isikhi/go-rate-limiter/ent/gen/predicate"
	"github.com/isikhi/go-rate-limiter/ent/gen/ratelimitoptions"
)

// RateLimitOptionsUpdate is the builder for updating RateLimitOptions entities.
type RateLimitOptionsUpdate struct {
	config
	hooks    []Hook
	mutation *RateLimitOptionsMutation
}

// Where appends a list predicates to the RateLimitOptionsUpdate builder.
func (rlou *RateLimitOptionsUpdate) Where(ps ...predicate.RateLimitOptions) *RateLimitOptionsUpdate {
	rlou.mutation.Where(ps...)
	return rlou
}

// SetClientID sets the "client_id" field.
func (rlou *RateLimitOptionsUpdate) SetClientID(s string) *RateLimitOptionsUpdate {
	rlou.mutation.SetClientID(s)
	return rlou
}

// SetTokenCount sets the "token_count" field.
func (rlou *RateLimitOptionsUpdate) SetTokenCount(s string) *RateLimitOptionsUpdate {
	rlou.mutation.SetTokenCount(s)
	return rlou
}

// SetNillableTokenCount sets the "token_count" field if the given value is not nil.
func (rlou *RateLimitOptionsUpdate) SetNillableTokenCount(s *string) *RateLimitOptionsUpdate {
	if s != nil {
		rlou.SetTokenCount(*s)
	}
	return rlou
}

// ClearTokenCount clears the value of the "token_count" field.
func (rlou *RateLimitOptionsUpdate) ClearTokenCount() *RateLimitOptionsUpdate {
	rlou.mutation.ClearTokenCount()
	return rlou
}

// SetDuration sets the "duration" field.
func (rlou *RateLimitOptionsUpdate) SetDuration(s string) *RateLimitOptionsUpdate {
	rlou.mutation.SetDuration(s)
	return rlou
}

// SetCreatedAt sets the "created_at" field.
func (rlou *RateLimitOptionsUpdate) SetCreatedAt(t time.Time) *RateLimitOptionsUpdate {
	rlou.mutation.SetCreatedAt(t)
	return rlou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rlou *RateLimitOptionsUpdate) SetNillableCreatedAt(t *time.Time) *RateLimitOptionsUpdate {
	if t != nil {
		rlou.SetCreatedAt(*t)
	}
	return rlou
}

// ClearCreatedAt clears the value of the "created_at" field.
func (rlou *RateLimitOptionsUpdate) ClearCreatedAt() *RateLimitOptionsUpdate {
	rlou.mutation.ClearCreatedAt()
	return rlou
}

// Mutation returns the RateLimitOptionsMutation object of the builder.
func (rlou *RateLimitOptionsUpdate) Mutation() *RateLimitOptionsMutation {
	return rlou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rlou *RateLimitOptionsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rlou.sqlSave, rlou.mutation, rlou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rlou *RateLimitOptionsUpdate) SaveX(ctx context.Context) int {
	affected, err := rlou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rlou *RateLimitOptionsUpdate) Exec(ctx context.Context) error {
	_, err := rlou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlou *RateLimitOptionsUpdate) ExecX(ctx context.Context) {
	if err := rlou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rlou *RateLimitOptionsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ratelimitoptions.Table, ratelimitoptions.Columns, sqlgraph.NewFieldSpec(ratelimitoptions.FieldID, field.TypeUint64))
	if ps := rlou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rlou.mutation.ClientID(); ok {
		_spec.SetField(ratelimitoptions.FieldClientID, field.TypeString, value)
	}
	if value, ok := rlou.mutation.TokenCount(); ok {
		_spec.SetField(ratelimitoptions.FieldTokenCount, field.TypeString, value)
	}
	if rlou.mutation.TokenCountCleared() {
		_spec.ClearField(ratelimitoptions.FieldTokenCount, field.TypeString)
	}
	if value, ok := rlou.mutation.Duration(); ok {
		_spec.SetField(ratelimitoptions.FieldDuration, field.TypeString, value)
	}
	if value, ok := rlou.mutation.CreatedAt(); ok {
		_spec.SetField(ratelimitoptions.FieldCreatedAt, field.TypeTime, value)
	}
	if rlou.mutation.CreatedAtCleared() {
		_spec.ClearField(ratelimitoptions.FieldCreatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rlou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ratelimitoptions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rlou.mutation.done = true
	return n, nil
}

// RateLimitOptionsUpdateOne is the builder for updating a single RateLimitOptions entity.
type RateLimitOptionsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RateLimitOptionsMutation
}

// SetClientID sets the "client_id" field.
func (rlouo *RateLimitOptionsUpdateOne) SetClientID(s string) *RateLimitOptionsUpdateOne {
	rlouo.mutation.SetClientID(s)
	return rlouo
}

// SetTokenCount sets the "token_count" field.
func (rlouo *RateLimitOptionsUpdateOne) SetTokenCount(s string) *RateLimitOptionsUpdateOne {
	rlouo.mutation.SetTokenCount(s)
	return rlouo
}

// SetNillableTokenCount sets the "token_count" field if the given value is not nil.
func (rlouo *RateLimitOptionsUpdateOne) SetNillableTokenCount(s *string) *RateLimitOptionsUpdateOne {
	if s != nil {
		rlouo.SetTokenCount(*s)
	}
	return rlouo
}

// ClearTokenCount clears the value of the "token_count" field.
func (rlouo *RateLimitOptionsUpdateOne) ClearTokenCount() *RateLimitOptionsUpdateOne {
	rlouo.mutation.ClearTokenCount()
	return rlouo
}

// SetDuration sets the "duration" field.
func (rlouo *RateLimitOptionsUpdateOne) SetDuration(s string) *RateLimitOptionsUpdateOne {
	rlouo.mutation.SetDuration(s)
	return rlouo
}

// SetCreatedAt sets the "created_at" field.
func (rlouo *RateLimitOptionsUpdateOne) SetCreatedAt(t time.Time) *RateLimitOptionsUpdateOne {
	rlouo.mutation.SetCreatedAt(t)
	return rlouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rlouo *RateLimitOptionsUpdateOne) SetNillableCreatedAt(t *time.Time) *RateLimitOptionsUpdateOne {
	if t != nil {
		rlouo.SetCreatedAt(*t)
	}
	return rlouo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (rlouo *RateLimitOptionsUpdateOne) ClearCreatedAt() *RateLimitOptionsUpdateOne {
	rlouo.mutation.ClearCreatedAt()
	return rlouo
}

// Mutation returns the RateLimitOptionsMutation object of the builder.
func (rlouo *RateLimitOptionsUpdateOne) Mutation() *RateLimitOptionsMutation {
	return rlouo.mutation
}

// Where appends a list predicates to the RateLimitOptionsUpdate builder.
func (rlouo *RateLimitOptionsUpdateOne) Where(ps ...predicate.RateLimitOptions) *RateLimitOptionsUpdateOne {
	rlouo.mutation.Where(ps...)
	return rlouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rlouo *RateLimitOptionsUpdateOne) Select(field string, fields ...string) *RateLimitOptionsUpdateOne {
	rlouo.fields = append([]string{field}, fields...)
	return rlouo
}

// Save executes the query and returns the updated RateLimitOptions entity.
func (rlouo *RateLimitOptionsUpdateOne) Save(ctx context.Context) (*RateLimitOptions, error) {
	return withHooks(ctx, rlouo.sqlSave, rlouo.mutation, rlouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rlouo *RateLimitOptionsUpdateOne) SaveX(ctx context.Context) *RateLimitOptions {
	node, err := rlouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rlouo *RateLimitOptionsUpdateOne) Exec(ctx context.Context) error {
	_, err := rlouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlouo *RateLimitOptionsUpdateOne) ExecX(ctx context.Context) {
	if err := rlouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rlouo *RateLimitOptionsUpdateOne) sqlSave(ctx context.Context) (_node *RateLimitOptions, err error) {
	_spec := sqlgraph.NewUpdateSpec(ratelimitoptions.Table, ratelimitoptions.Columns, sqlgraph.NewFieldSpec(ratelimitoptions.FieldID, field.TypeUint64))
	id, ok := rlouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "RateLimitOptions.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rlouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ratelimitoptions.FieldID)
		for _, f := range fields {
			if !ratelimitoptions.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != ratelimitoptions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rlouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rlouo.mutation.ClientID(); ok {
		_spec.SetField(ratelimitoptions.FieldClientID, field.TypeString, value)
	}
	if value, ok := rlouo.mutation.TokenCount(); ok {
		_spec.SetField(ratelimitoptions.FieldTokenCount, field.TypeString, value)
	}
	if rlouo.mutation.TokenCountCleared() {
		_spec.ClearField(ratelimitoptions.FieldTokenCount, field.TypeString)
	}
	if value, ok := rlouo.mutation.Duration(); ok {
		_spec.SetField(ratelimitoptions.FieldDuration, field.TypeString, value)
	}
	if value, ok := rlouo.mutation.CreatedAt(); ok {
		_spec.SetField(ratelimitoptions.FieldCreatedAt, field.TypeTime, value)
	}
	if rlouo.mutation.CreatedAtCleared() {
		_spec.ClearField(ratelimitoptions.FieldCreatedAt, field.TypeTime)
	}
	_node = &RateLimitOptions{config: rlouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rlouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ratelimitoptions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rlouo.mutation.done = true
	return _node, nil
}
