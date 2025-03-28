// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogs"
	"blog/internal/ent/tags"
	"blog/internal/ent/tagsrelation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagsRelationCreate is the builder for creating a TagsRelation entity.
type TagsRelationCreate struct {
	config
	mutation *TagsRelationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTagID sets the "tag_id" field.
func (trc *TagsRelationCreate) SetTagID(i int) *TagsRelationCreate {
	trc.mutation.SetTagID(i)
	return trc
}

// SetRelation sets the "relation" field.
func (trc *TagsRelationCreate) SetRelation(s string) *TagsRelationCreate {
	trc.mutation.SetRelation(s)
	return trc
}

// SetRelationID sets the "relation_id" field.
func (trc *TagsRelationCreate) SetRelationID(i int) *TagsRelationCreate {
	trc.mutation.SetRelationID(i)
	return trc
}

// SetBlogID sets the "blog" edge to the Blogs entity by ID.
func (trc *TagsRelationCreate) SetBlogID(id int) *TagsRelationCreate {
	trc.mutation.SetBlogID(id)
	return trc
}

// SetBlog sets the "blog" edge to the Blogs entity.
func (trc *TagsRelationCreate) SetBlog(b *Blogs) *TagsRelationCreate {
	return trc.SetBlogID(b.ID)
}

// SetTag sets the "tag" edge to the Tags entity.
func (trc *TagsRelationCreate) SetTag(t *Tags) *TagsRelationCreate {
	return trc.SetTagID(t.ID)
}

// Mutation returns the TagsRelationMutation object of the builder.
func (trc *TagsRelationCreate) Mutation() *TagsRelationMutation {
	return trc.mutation
}

// Save creates the TagsRelation in the database.
func (trc *TagsRelationCreate) Save(ctx context.Context) (*TagsRelation, error) {
	return withHooks(ctx, trc.sqlSave, trc.mutation, trc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TagsRelationCreate) SaveX(ctx context.Context) *TagsRelation {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TagsRelationCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TagsRelationCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TagsRelationCreate) check() error {
	if _, ok := trc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "TagsRelation.tag_id"`)}
	}
	if _, ok := trc.mutation.Relation(); !ok {
		return &ValidationError{Name: "relation", err: errors.New(`ent: missing required field "TagsRelation.relation"`)}
	}
	if _, ok := trc.mutation.RelationID(); !ok {
		return &ValidationError{Name: "relation_id", err: errors.New(`ent: missing required field "TagsRelation.relation_id"`)}
	}
	if len(trc.mutation.BlogIDs()) == 0 {
		return &ValidationError{Name: "blog", err: errors.New(`ent: missing required edge "TagsRelation.blog"`)}
	}
	if len(trc.mutation.TagIDs()) == 0 {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required edge "TagsRelation.tag"`)}
	}
	return nil
}

func (trc *TagsRelationCreate) sqlSave(ctx context.Context) (*TagsRelation, error) {
	if err := trc.check(); err != nil {
		return nil, err
	}
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	trc.mutation.id = &_node.ID
	trc.mutation.done = true
	return _node, nil
}

func (trc *TagsRelationCreate) createSpec() (*TagsRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &TagsRelation{config: trc.config}
		_spec = sqlgraph.NewCreateSpec(tagsrelation.Table, sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt))
	)
	_spec.OnConflict = trc.conflict
	if value, ok := trc.mutation.Relation(); ok {
		_spec.SetField(tagsrelation.FieldRelation, field.TypeString, value)
		_node.Relation = value
	}
	if nodes := trc.mutation.BlogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.BlogTable,
			Columns: []string{tagsrelation.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.TagTable,
			Columns: []string{tagsrelation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TagsRelation.Create().
//		SetTagID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagsRelationUpsert) {
//			SetTagID(v+v).
//		}).
//		Exec(ctx)
func (trc *TagsRelationCreate) OnConflict(opts ...sql.ConflictOption) *TagsRelationUpsertOne {
	trc.conflict = opts
	return &TagsRelationUpsertOne{
		create: trc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (trc *TagsRelationCreate) OnConflictColumns(columns ...string) *TagsRelationUpsertOne {
	trc.conflict = append(trc.conflict, sql.ConflictColumns(columns...))
	return &TagsRelationUpsertOne{
		create: trc,
	}
}

type (
	// TagsRelationUpsertOne is the builder for "upsert"-ing
	//  one TagsRelation node.
	TagsRelationUpsertOne struct {
		create *TagsRelationCreate
	}

	// TagsRelationUpsert is the "OnConflict" setter.
	TagsRelationUpsert struct {
		*sql.UpdateSet
	}
)

// SetTagID sets the "tag_id" field.
func (u *TagsRelationUpsert) SetTagID(v int) *TagsRelationUpsert {
	u.Set(tagsrelation.FieldTagID, v)
	return u
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TagsRelationUpsert) UpdateTagID() *TagsRelationUpsert {
	u.SetExcluded(tagsrelation.FieldTagID)
	return u
}

// SetRelation sets the "relation" field.
func (u *TagsRelationUpsert) SetRelation(v string) *TagsRelationUpsert {
	u.Set(tagsrelation.FieldRelation, v)
	return u
}

// UpdateRelation sets the "relation" field to the value that was provided on create.
func (u *TagsRelationUpsert) UpdateRelation() *TagsRelationUpsert {
	u.SetExcluded(tagsrelation.FieldRelation)
	return u
}

// SetRelationID sets the "relation_id" field.
func (u *TagsRelationUpsert) SetRelationID(v int) *TagsRelationUpsert {
	u.Set(tagsrelation.FieldRelationID, v)
	return u
}

// UpdateRelationID sets the "relation_id" field to the value that was provided on create.
func (u *TagsRelationUpsert) UpdateRelationID() *TagsRelationUpsert {
	u.SetExcluded(tagsrelation.FieldRelationID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TagsRelationUpsertOne) UpdateNewValues() *TagsRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TagsRelationUpsertOne) Ignore() *TagsRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagsRelationUpsertOne) DoNothing() *TagsRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagsRelationCreate.OnConflict
// documentation for more info.
func (u *TagsRelationUpsertOne) Update(set func(*TagsRelationUpsert)) *TagsRelationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagsRelationUpsert{UpdateSet: update})
	}))
	return u
}

// SetTagID sets the "tag_id" field.
func (u *TagsRelationUpsertOne) SetTagID(v int) *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TagsRelationUpsertOne) UpdateTagID() *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateTagID()
	})
}

// SetRelation sets the "relation" field.
func (u *TagsRelationUpsertOne) SetRelation(v string) *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetRelation(v)
	})
}

// UpdateRelation sets the "relation" field to the value that was provided on create.
func (u *TagsRelationUpsertOne) UpdateRelation() *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateRelation()
	})
}

// SetRelationID sets the "relation_id" field.
func (u *TagsRelationUpsertOne) SetRelationID(v int) *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetRelationID(v)
	})
}

// UpdateRelationID sets the "relation_id" field to the value that was provided on create.
func (u *TagsRelationUpsertOne) UpdateRelationID() *TagsRelationUpsertOne {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateRelationID()
	})
}

// Exec executes the query.
func (u *TagsRelationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TagsRelationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagsRelationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TagsRelationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TagsRelationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TagsRelationCreateBulk is the builder for creating many TagsRelation entities in bulk.
type TagsRelationCreateBulk struct {
	config
	err      error
	builders []*TagsRelationCreate
	conflict []sql.ConflictOption
}

// Save creates the TagsRelation entities in the database.
func (trcb *TagsRelationCreateBulk) Save(ctx context.Context) ([]*TagsRelation, error) {
	if trcb.err != nil {
		return nil, trcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TagsRelation, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagsRelationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = trcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TagsRelationCreateBulk) SaveX(ctx context.Context) []*TagsRelation {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TagsRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TagsRelationCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TagsRelation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagsRelationUpsert) {
//			SetTagID(v+v).
//		}).
//		Exec(ctx)
func (trcb *TagsRelationCreateBulk) OnConflict(opts ...sql.ConflictOption) *TagsRelationUpsertBulk {
	trcb.conflict = opts
	return &TagsRelationUpsertBulk{
		create: trcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (trcb *TagsRelationCreateBulk) OnConflictColumns(columns ...string) *TagsRelationUpsertBulk {
	trcb.conflict = append(trcb.conflict, sql.ConflictColumns(columns...))
	return &TagsRelationUpsertBulk{
		create: trcb,
	}
}

// TagsRelationUpsertBulk is the builder for "upsert"-ing
// a bulk of TagsRelation nodes.
type TagsRelationUpsertBulk struct {
	create *TagsRelationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TagsRelationUpsertBulk) UpdateNewValues() *TagsRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TagsRelation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TagsRelationUpsertBulk) Ignore() *TagsRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagsRelationUpsertBulk) DoNothing() *TagsRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagsRelationCreateBulk.OnConflict
// documentation for more info.
func (u *TagsRelationUpsertBulk) Update(set func(*TagsRelationUpsert)) *TagsRelationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagsRelationUpsert{UpdateSet: update})
	}))
	return u
}

// SetTagID sets the "tag_id" field.
func (u *TagsRelationUpsertBulk) SetTagID(v int) *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TagsRelationUpsertBulk) UpdateTagID() *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateTagID()
	})
}

// SetRelation sets the "relation" field.
func (u *TagsRelationUpsertBulk) SetRelation(v string) *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetRelation(v)
	})
}

// UpdateRelation sets the "relation" field to the value that was provided on create.
func (u *TagsRelationUpsertBulk) UpdateRelation() *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateRelation()
	})
}

// SetRelationID sets the "relation_id" field.
func (u *TagsRelationUpsertBulk) SetRelationID(v int) *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.SetRelationID(v)
	})
}

// UpdateRelationID sets the "relation_id" field to the value that was provided on create.
func (u *TagsRelationUpsertBulk) UpdateRelationID() *TagsRelationUpsertBulk {
	return u.Update(func(s *TagsRelationUpsert) {
		s.UpdateRelationID()
	})
}

// Exec executes the query.
func (u *TagsRelationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TagsRelationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TagsRelationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagsRelationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
