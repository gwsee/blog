// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/palacesmemory"
	"blog/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PalacesMemoryDelete is the builder for deleting a PalacesMemory entity.
type PalacesMemoryDelete struct {
	config
	hooks    []Hook
	mutation *PalacesMemoryMutation
}

// Where appends a list predicates to the PalacesMemoryDelete builder.
func (pmd *PalacesMemoryDelete) Where(ps ...predicate.PalacesMemory) *PalacesMemoryDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *PalacesMemoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *PalacesMemoryDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *PalacesMemoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(palacesmemory.Table, sqlgraph.NewFieldSpec(palacesmemory.FieldID, field.TypeInt))
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

// PalacesMemoryDeleteOne is the builder for deleting a single PalacesMemory entity.
type PalacesMemoryDeleteOne struct {
	pmd *PalacesMemoryDelete
}

// Where appends a list predicates to the PalacesMemoryDelete builder.
func (pmdo *PalacesMemoryDeleteOne) Where(ps ...predicate.PalacesMemory) *PalacesMemoryDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *PalacesMemoryDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{palacesmemory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *PalacesMemoryDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
