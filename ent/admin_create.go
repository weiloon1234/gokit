// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/weiloon1234/gokit/ent/admin"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (ac *AdminCreate) SetUsername(s string) *AdminCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUsername(s *string) *AdminCreate {
	if s != nil {
		ac.SetUsername(*s)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AdminCreate) SetName(s string) *AdminCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *AdminCreate) SetNillableName(s *string) *AdminCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetEmail sets the "email" field.
func (ac *AdminCreate) SetEmail(s string) *AdminCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ac *AdminCreate) SetNillableEmail(s *string) *AdminCreate {
	if s != nil {
		ac.SetEmail(*s)
	}
	return ac
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (ac *AdminCreate) SetEmailVerifiedAt(t time.Time) *AdminCreate {
	ac.mutation.SetEmailVerifiedAt(t)
	return ac
}

// SetNillableEmailVerifiedAt sets the "email_verified_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableEmailVerifiedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetEmailVerifiedAt(*t)
	}
	return ac
}

// SetPassword sets the "password" field.
func (ac *AdminCreate) SetPassword(s string) *AdminCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetPassword2 sets the "password2" field.
func (ac *AdminCreate) SetPassword2(s string) *AdminCreate {
	ac.mutation.SetPassword2(s)
	return ac
}

// SetLang sets the "lang" field.
func (ac *AdminCreate) SetLang(s string) *AdminCreate {
	ac.mutation.SetLang(s)
	return ac
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (ac *AdminCreate) SetNillableLang(s *string) *AdminCreate {
	if s != nil {
		ac.SetLang(*s)
	}
	return ac
}

// SetAvatar sets the "avatar" field.
func (ac *AdminCreate) SetAvatar(s string) *AdminCreate {
	ac.mutation.SetAvatar(s)
	return ac
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (ac *AdminCreate) SetNillableAvatar(s *string) *AdminCreate {
	if s != nil {
		ac.SetAvatar(*s)
	}
	return ac
}

// SetType sets the "type" field.
func (ac *AdminCreate) SetType(u uint8) *AdminCreate {
	ac.mutation.SetType(u)
	return ac
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ac *AdminCreate) SetNillableType(u *uint8) *AdminCreate {
	if u != nil {
		ac.SetType(*u)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminCreate) SetCreatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminCreate) SetUpdatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AdminCreate) SetDeletedAt(t time.Time) *AdminCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableDeletedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AdminCreate) SetID(u uint64) *AdminCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.Lang(); !ok {
		v := admin.DefaultLang
		ac.mutation.SetLang(v)
	}
	if _, ok := ac.mutation.GetType(); !ok {
		v := admin.DefaultType
		ac.mutation.SetType(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admin.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admin.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Admin.password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Admin.password": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Password2(); !ok {
		return &ValidationError{Name: "password2", err: errors.New(`ent: missing required field "Admin.password2"`)}
	}
	if v, ok := ac.mutation.Password2(); ok {
		if err := admin.Password2Validator(v); err != nil {
			return &ValidationError{Name: "password2", err: fmt.Errorf(`ent: validator failed for field "Admin.password2": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "Admin.lang"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Admin.type"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Admin.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Admin.updated_at"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(admin.Table, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeUint64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(admin.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(admin.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.EmailVerifiedAt(); ok {
		_spec.SetField(admin.FieldEmailVerifiedAt, field.TypeTime, value)
		_node.EmailVerifiedAt = &value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.Password2(); ok {
		_spec.SetField(admin.FieldPassword2, field.TypeString, value)
		_node.Password2 = value
	}
	if value, ok := ac.mutation.Lang(); ok {
		_spec.SetField(admin.FieldLang, field.TypeString, value)
		_node.Lang = value
	}
	if value, ok := ac.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
		_node.Avatar = &value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(admin.FieldType, field.TypeUint8, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	err      error
	builders []*AdminCreate
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
