// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sosljuk8/analytics/orm/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDirection sets the "Direction" field.
func (mc *MessageCreate) SetDirection(s string) *MessageCreate {
	mc.mutation.SetDirection(s)
	return mc
}

// SetSent sets the "Sent" field.
func (mc *MessageCreate) SetSent(s string) *MessageCreate {
	mc.mutation.SetSent(s)
	return mc
}

// SetMailboxId sets the "MailboxId" field.
func (mc *MessageCreate) SetMailboxId(i int) *MessageCreate {
	mc.mutation.SetMailboxId(i)
	return mc
}

// SetCrmRid sets the "CrmRid" field.
func (mc *MessageCreate) SetCrmRid(i int) *MessageCreate {
	mc.mutation.SetCrmRid(i)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.Direction(); !ok {
		return &ValidationError{Name: "Direction", err: errors.New(`ent: missing required field "Message.Direction"`)}
	}
	if _, ok := mc.mutation.Sent(); !ok {
		return &ValidationError{Name: "Sent", err: errors.New(`ent: missing required field "Message.Sent"`)}
	}
	if _, ok := mc.mutation.MailboxId(); !ok {
		return &ValidationError{Name: "MailboxId", err: errors.New(`ent: missing required field "Message.MailboxId"`)}
	}
	if _, ok := mc.mutation.CrmRid(); !ok {
		return &ValidationError{Name: "CrmRid", err: errors.New(`ent: missing required field "Message.CrmRid"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.Direction(); ok {
		_spec.SetField(message.FieldDirection, field.TypeString, value)
		_node.Direction = value
	}
	if value, ok := mc.mutation.Sent(); ok {
		_spec.SetField(message.FieldSent, field.TypeString, value)
		_node.Sent = value
	}
	if value, ok := mc.mutation.MailboxId(); ok {
		_spec.SetField(message.FieldMailboxId, field.TypeInt, value)
		_node.MailboxId = value
	}
	if value, ok := mc.mutation.CrmRid(); ok {
		_spec.SetField(message.FieldCrmRid, field.TypeInt, value)
		_node.CrmRid = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetDirection(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetDirection(v+v).
//		}).
//		Exec(ctx)
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetDirection sets the "Direction" field.
func (u *MessageUpsert) SetDirection(v string) *MessageUpsert {
	u.Set(message.FieldDirection, v)
	return u
}

// UpdateDirection sets the "Direction" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDirection() *MessageUpsert {
	u.SetExcluded(message.FieldDirection)
	return u
}

// SetSent sets the "Sent" field.
func (u *MessageUpsert) SetSent(v string) *MessageUpsert {
	u.Set(message.FieldSent, v)
	return u
}

// UpdateSent sets the "Sent" field to the value that was provided on create.
func (u *MessageUpsert) UpdateSent() *MessageUpsert {
	u.SetExcluded(message.FieldSent)
	return u
}

// SetMailboxId sets the "MailboxId" field.
func (u *MessageUpsert) SetMailboxId(v int) *MessageUpsert {
	u.Set(message.FieldMailboxId, v)
	return u
}

// UpdateMailboxId sets the "MailboxId" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMailboxId() *MessageUpsert {
	u.SetExcluded(message.FieldMailboxId)
	return u
}

// AddMailboxId adds v to the "MailboxId" field.
func (u *MessageUpsert) AddMailboxId(v int) *MessageUpsert {
	u.Add(message.FieldMailboxId, v)
	return u
}

// SetCrmRid sets the "CrmRid" field.
func (u *MessageUpsert) SetCrmRid(v int) *MessageUpsert {
	u.Set(message.FieldCrmRid, v)
	return u
}

// UpdateCrmRid sets the "CrmRid" field to the value that was provided on create.
func (u *MessageUpsert) UpdateCrmRid() *MessageUpsert {
	u.SetExcluded(message.FieldCrmRid)
	return u
}

// AddCrmRid adds v to the "CrmRid" field.
func (u *MessageUpsert) AddCrmRid(v int) *MessageUpsert {
	u.Add(message.FieldCrmRid, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetDirection sets the "Direction" field.
func (u *MessageUpsertOne) SetDirection(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDirection(v)
	})
}

// UpdateDirection sets the "Direction" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDirection() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDirection()
	})
}

// SetSent sets the "Sent" field.
func (u *MessageUpsertOne) SetSent(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetSent(v)
	})
}

// UpdateSent sets the "Sent" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateSent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSent()
	})
}

// SetMailboxId sets the "MailboxId" field.
func (u *MessageUpsertOne) SetMailboxId(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMailboxId(v)
	})
}

// AddMailboxId adds v to the "MailboxId" field.
func (u *MessageUpsertOne) AddMailboxId(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddMailboxId(v)
	})
}

// UpdateMailboxId sets the "MailboxId" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMailboxId() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMailboxId()
	})
}

// SetCrmRid sets the "CrmRid" field.
func (u *MessageUpsertOne) SetCrmRid(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetCrmRid(v)
	})
}

// AddCrmRid adds v to the "CrmRid" field.
func (u *MessageUpsertOne) AddCrmRid(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddCrmRid(v)
	})
}

// UpdateCrmRid sets the "CrmRid" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateCrmRid() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCrmRid()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetDirection(v+v).
//		}).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetDirection sets the "Direction" field.
func (u *MessageUpsertBulk) SetDirection(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDirection(v)
	})
}

// UpdateDirection sets the "Direction" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDirection() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDirection()
	})
}

// SetSent sets the "Sent" field.
func (u *MessageUpsertBulk) SetSent(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetSent(v)
	})
}

// UpdateSent sets the "Sent" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateSent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSent()
	})
}

// SetMailboxId sets the "MailboxId" field.
func (u *MessageUpsertBulk) SetMailboxId(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMailboxId(v)
	})
}

// AddMailboxId adds v to the "MailboxId" field.
func (u *MessageUpsertBulk) AddMailboxId(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddMailboxId(v)
	})
}

// UpdateMailboxId sets the "MailboxId" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMailboxId() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMailboxId()
	})
}

// SetCrmRid sets the "CrmRid" field.
func (u *MessageUpsertBulk) SetCrmRid(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetCrmRid(v)
	})
}

// AddCrmRid adds v to the "CrmRid" field.
func (u *MessageUpsertBulk) AddCrmRid(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddCrmRid(v)
	})
}

// UpdateCrmRid sets the "CrmRid" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateCrmRid() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCrmRid()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
