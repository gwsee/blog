// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/predicate"
	"blog/internal/ent/tagsrelation"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagsRelationDelete is the builder for deleting a TagsRelation entity.
type TagsRelationDelete struct {
	config
	hooks    []Hook
	mutation *TagsRelationMutation
}

// Where appends a list predicates to the TagsRelationDelete builder.
func (trd *TagsRelationDelete) Where(ps ...predicate.TagsRelation) *TagsRelationDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TagsRelationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, trd.sqlExec, trd.mutation, trd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TagsRelationDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TagsRelationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tagsrelation.Table, sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt))
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	trd.mutation.done = true
	return affected, err
}

// TagsRelationDeleteOne is the builder for deleting a single TagsRelation entity.
type TagsRelationDeleteOne struct {
	trd *TagsRelationDelete
}

// Where appends a list predicates to the TagsRelationDelete builder.
func (trdo *TagsRelationDeleteOne) Where(ps ...predicate.TagsRelation) *TagsRelationDeleteOne {
	trdo.trd.mutation.Where(ps...)
	return trdo
}

// Exec executes the deletion query.
func (trdo *TagsRelationDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tagsrelation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TagsRelationDeleteOne) ExecX(ctx context.Context) {
	if err := trdo.Exec(ctx); err != nil {
		panic(err)
	}
}
