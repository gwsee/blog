// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/palacesmemo"
	"blog/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PalacesMemoDelete is the builder for deleting a PalacesMemo entity.
type PalacesMemoDelete struct {
	config
	hooks    []Hook
	mutation *PalacesMemoMutation
}

// Where appends a list predicates to the PalacesMemoDelete builder.
func (pmd *PalacesMemoDelete) Where(ps ...predicate.PalacesMemo) *PalacesMemoDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *PalacesMemoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *PalacesMemoDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *PalacesMemoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(palacesmemo.Table, sqlgraph.NewFieldSpec(palacesmemo.FieldID, field.TypeInt))
	if ps := pmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pmd.mutation.done = true
	return affected, err
}

// PalacesMemoDeleteOne is the builder for deleting a single PalacesMemo entity.
type PalacesMemoDeleteOne struct {
	pmd *PalacesMemoDelete
}

// Where appends a list predicates to the PalacesMemoDelete builder.
func (pmdo *PalacesMemoDeleteOne) Where(ps ...predicate.PalacesMemo) *PalacesMemoDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *PalacesMemoDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{palacesmemo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *PalacesMemoDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
