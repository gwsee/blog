// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/predicate"
	"blog/internal/ent/travelextends"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TravelExtendsDelete is the builder for deleting a TravelExtends entity.
type TravelExtendsDelete struct {
	config
	hooks    []Hook
	mutation *TravelExtendsMutation
}

// Where appends a list predicates to the TravelExtendsDelete builder.
func (ted *TravelExtendsDelete) Where(ps ...predicate.TravelExtends) *TravelExtendsDelete {
	ted.mutation.Where(ps...)
	return ted
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ted *TravelExtendsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ted.sqlExec, ted.mutation, ted.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ted *TravelExtendsDelete) ExecX(ctx context.Context) int {
	n, err := ted.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ted *TravelExtendsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(travelextends.Table, sqlgraph.NewFieldSpec(travelextends.FieldID, field.TypeInt))
	if ps := ted.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ted.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ted.mutation.done = true
	return affected, err
}

// TravelExtendsDeleteOne is the builder for deleting a single TravelExtends entity.
type TravelExtendsDeleteOne struct {
	ted *TravelExtendsDelete
}

// Where appends a list predicates to the TravelExtendsDelete builder.
func (tedo *TravelExtendsDeleteOne) Where(ps ...predicate.TravelExtends) *TravelExtendsDeleteOne {
	tedo.ted.mutation.Where(ps...)
	return tedo
}

// Exec executes the deletion query.
func (tedo *TravelExtendsDeleteOne) Exec(ctx context.Context) error {
	n, err := tedo.ted.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{travelextends.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tedo *TravelExtendsDeleteOne) ExecX(ctx context.Context) {
	if err := tedo.Exec(ctx); err != nil {
		panic(err)
	}
}
