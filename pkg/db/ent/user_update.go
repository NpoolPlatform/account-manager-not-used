// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(u uint32) *UserUpdate {
	uu.mutation.ResetCreatedAt()
	uu.mutation.SetCreatedAt(u)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetCreatedAt(*u)
	}
	return uu
}

// AddCreatedAt adds u to the "created_at" field.
func (uu *UserUpdate) AddCreatedAt(u int32) *UserUpdate {
	uu.mutation.AddCreatedAt(u)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(u uint32) *UserUpdate {
	uu.mutation.ResetUpdatedAt()
	uu.mutation.SetUpdatedAt(u)
	return uu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (uu *UserUpdate) AddUpdatedAt(u int32) *UserUpdate {
	uu.mutation.AddUpdatedAt(u)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(u uint32) *UserUpdate {
	uu.mutation.ResetDeletedAt()
	uu.mutation.SetDeletedAt(u)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetDeletedAt(*u)
	}
	return uu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (uu *UserUpdate) AddDeletedAt(u int32) *UserUpdate {
	uu.mutation.AddDeletedAt(u)
	return uu
}

// SetAppID sets the "app_id" field.
func (uu *UserUpdate) SetAppID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetAppID(u)
	return uu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAppID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetAppID(*u)
	}
	return uu
}

// ClearAppID clears the value of the "app_id" field.
func (uu *UserUpdate) ClearAppID() *UserUpdate {
	uu.mutation.ClearAppID()
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UserUpdate) SetUserID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetUserID(u)
	return uu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetUserID(*u)
	}
	return uu
}

// ClearUserID clears the value of the "user_id" field.
func (uu *UserUpdate) ClearUserID() *UserUpdate {
	uu.mutation.ClearUserID()
	return uu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uu *UserUpdate) SetCoinTypeID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetCoinTypeID(u)
	return uu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCoinTypeID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetCoinTypeID(*u)
	}
	return uu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (uu *UserUpdate) ClearCoinTypeID() *UserUpdate {
	uu.mutation.ClearCoinTypeID()
	return uu
}

// SetAccountID sets the "account_id" field.
func (uu *UserUpdate) SetAccountID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetAccountID(u)
	return uu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAccountID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetAccountID(*u)
	}
	return uu
}

// ClearAccountID clears the value of the "account_id" field.
func (uu *UserUpdate) ClearAccountID() *UserUpdate {
	uu.mutation.ClearAccountID()
	return uu
}

// SetUsedFor sets the "used_for" field.
func (uu *UserUpdate) SetUsedFor(s string) *UserUpdate {
	uu.mutation.SetUsedFor(s)
	return uu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsedFor(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsedFor(*s)
	}
	return uu
}

// ClearUsedFor clears the value of the "used_for" field.
func (uu *UserUpdate) ClearUsedFor() *UserUpdate {
	uu.mutation.ClearUsedFor()
	return uu
}

// SetLabels sets the "labels" field.
func (uu *UserUpdate) SetLabels(s []string) *UserUpdate {
	uu.mutation.SetLabels(s)
	return uu
}

// ClearLabels clears the value of the "labels" field.
func (uu *UserUpdate) ClearLabels() *UserUpdate {
	uu.mutation.ClearLabels()
	return uu
}

// SetMemo sets the "memo" field.
func (uu *UserUpdate) SetMemo(s string) *UserUpdate {
	uu.mutation.SetMemo(s)
	return uu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMemo(s *string) *UserUpdate {
	if s != nil {
		uu.SetMemo(*s)
	}
	return uu
}

// ClearMemo clears the value of the "memo" field.
func (uu *UserUpdate) ClearMemo() *UserUpdate {
	uu.mutation.ClearMemo()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
	}
	if value, ok := uu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
	}
	if value, ok := uu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAppID,
		})
	}
	if uu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldAppID,
		})
	}
	if value, ok := uu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldUserID,
		})
	}
	if uu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldUserID,
		})
	}
	if value, ok := uu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCoinTypeID,
		})
	}
	if uu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldCoinTypeID,
		})
	}
	if value, ok := uu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAccountID,
		})
	}
	if uu.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldAccountID,
		})
	}
	if value, ok := uu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsedFor,
		})
	}
	if uu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsedFor,
		})
	}
	if value, ok := uu.mutation.Labels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldLabels,
		})
	}
	if uu.mutation.LabelsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldLabels,
		})
	}
	if value, ok := uu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMemo,
		})
	}
	if uu.mutation.MemoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldMemo,
		})
	}
	_spec.Modifiers = uu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(u uint32) *UserUpdateOne {
	uuo.mutation.ResetCreatedAt()
	uuo.mutation.SetCreatedAt(u)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetCreatedAt(*u)
	}
	return uuo
}

// AddCreatedAt adds u to the "created_at" field.
func (uuo *UserUpdateOne) AddCreatedAt(u int32) *UserUpdateOne {
	uuo.mutation.AddCreatedAt(u)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(u uint32) *UserUpdateOne {
	uuo.mutation.ResetUpdatedAt()
	uuo.mutation.SetUpdatedAt(u)
	return uuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (uuo *UserUpdateOne) AddUpdatedAt(u int32) *UserUpdateOne {
	uuo.mutation.AddUpdatedAt(u)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(u uint32) *UserUpdateOne {
	uuo.mutation.ResetDeletedAt()
	uuo.mutation.SetDeletedAt(u)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetDeletedAt(*u)
	}
	return uuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (uuo *UserUpdateOne) AddDeletedAt(u int32) *UserUpdateOne {
	uuo.mutation.AddDeletedAt(u)
	return uuo
}

// SetAppID sets the "app_id" field.
func (uuo *UserUpdateOne) SetAppID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAppID(u)
	return uuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAppID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetAppID(*u)
	}
	return uuo
}

// ClearAppID clears the value of the "app_id" field.
func (uuo *UserUpdateOne) ClearAppID() *UserUpdateOne {
	uuo.mutation.ClearAppID()
	return uuo
}

// SetUserID sets the "user_id" field.
func (uuo *UserUpdateOne) SetUserID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetUserID(u)
	return uuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetUserID(*u)
	}
	return uuo
}

// ClearUserID clears the value of the "user_id" field.
func (uuo *UserUpdateOne) ClearUserID() *UserUpdateOne {
	uuo.mutation.ClearUserID()
	return uuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (uuo *UserUpdateOne) SetCoinTypeID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetCoinTypeID(u)
	return uuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetCoinTypeID(*u)
	}
	return uuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (uuo *UserUpdateOne) ClearCoinTypeID() *UserUpdateOne {
	uuo.mutation.ClearCoinTypeID()
	return uuo
}

// SetAccountID sets the "account_id" field.
func (uuo *UserUpdateOne) SetAccountID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAccountID(u)
	return uuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAccountID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetAccountID(*u)
	}
	return uuo
}

// ClearAccountID clears the value of the "account_id" field.
func (uuo *UserUpdateOne) ClearAccountID() *UserUpdateOne {
	uuo.mutation.ClearAccountID()
	return uuo
}

// SetUsedFor sets the "used_for" field.
func (uuo *UserUpdateOne) SetUsedFor(s string) *UserUpdateOne {
	uuo.mutation.SetUsedFor(s)
	return uuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsedFor(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsedFor(*s)
	}
	return uuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (uuo *UserUpdateOne) ClearUsedFor() *UserUpdateOne {
	uuo.mutation.ClearUsedFor()
	return uuo
}

// SetLabels sets the "labels" field.
func (uuo *UserUpdateOne) SetLabels(s []string) *UserUpdateOne {
	uuo.mutation.SetLabels(s)
	return uuo
}

// ClearLabels clears the value of the "labels" field.
func (uuo *UserUpdateOne) ClearLabels() *UserUpdateOne {
	uuo.mutation.ClearLabels()
	return uuo
}

// SetMemo sets the "memo" field.
func (uuo *UserUpdateOne) SetMemo(s string) *UserUpdateOne {
	uuo.mutation.SetMemo(s)
	return uuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMemo(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMemo(*s)
	}
	return uuo
}

// ClearMemo clears the value of the "memo" field.
func (uuo *UserUpdateOne) ClearMemo() *UserUpdateOne {
	uuo.mutation.ClearMemo()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
	}
	if value, ok := uuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
	}
	if value, ok := uuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAppID,
		})
	}
	if uuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldAppID,
		})
	}
	if value, ok := uuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldUserID,
		})
	}
	if uuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldUserID,
		})
	}
	if value, ok := uuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCoinTypeID,
		})
	}
	if uuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldCoinTypeID,
		})
	}
	if value, ok := uuo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldAccountID,
		})
	}
	if uuo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: user.FieldAccountID,
		})
	}
	if value, ok := uuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsedFor,
		})
	}
	if uuo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsedFor,
		})
	}
	if value, ok := uuo.mutation.Labels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldLabels,
		})
	}
	if uuo.mutation.LabelsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldLabels,
		})
	}
	if value, ok := uuo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMemo,
		})
	}
	if uuo.mutation.MemoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldMemo,
		})
	}
	_spec.Modifiers = uuo.modifiers
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
