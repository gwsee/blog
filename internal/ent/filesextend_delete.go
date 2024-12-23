// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/filesextend"
	"blog/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FilesExtendDelete is the builder for deleting a FilesExtend entity.
type FilesExtendDelete struct {
	config
	hooks    []Hook
	mutation *FilesExtendMutation
}

// Where appends a list predicates to the FilesExtendDelete builder.
func (fed *FilesExtendDelete) Where(ps ...predicate.FilesExtend) *FilesExtendDelete {
	fed.mutation.Where(ps...)
	return fed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fed *FilesExtendDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fed.sqlExec, fed.mutation, fed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fed *FilesExtendDelete) ExecX(ctx context.Context) int {
	n, err := fed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fed *FilesExtendDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(filesextend.Table, sqlgraph.NewFieldSpec(filesextend.FieldID, field.TypeInt))
	if ps := fed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fed.mutation.done = true
	return affected, err
}

// FilesExtendDeleteOne is the builder for deleting a single FilesExtend entity.
type FilesExtendDeleteOne struct {
	fed *FilesExtendDelete
}

// Where appends a list predicates to the FilesExtendDelete builder.
func (fedo *FilesExtendDeleteOne) Where(ps ...predicate.FilesExtend) *FilesExtendDeleteOne {
	fedo.fed.mutation.Where(ps...)
	return fedo
}

// Exec executes the deletion query.
func (fedo *FilesExtendDeleteOne) Exec(ctx context.Context) error {
	n, err := fedo.fed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{filesextend.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fedo *FilesExtendDeleteOne) ExecX(ctx context.Context) {
	if err := fedo.Exec(ctx); err != nil {
		panic(err)
	}
}
