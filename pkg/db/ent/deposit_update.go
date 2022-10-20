// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// DepositUpdate is the builder for updating Deposit entities.
type DepositUpdate struct {
	config
	hooks     []Hook
	mutation  *DepositMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DepositUpdate builder.
func (du *DepositUpdate) Where(ps ...predicate.Deposit) *DepositUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DepositUpdate) SetCreatedAt(u uint32) *DepositUpdate {
	du.mutation.ResetCreatedAt()
	du.mutation.SetCreatedAt(u)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DepositUpdate) SetNillableCreatedAt(u *uint32) *DepositUpdate {
	if u != nil {
		du.SetCreatedAt(*u)
	}
	return du
}

// AddCreatedAt adds u to the "created_at" field.
func (du *DepositUpdate) AddCreatedAt(u int32) *DepositUpdate {
	du.mutation.AddCreatedAt(u)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DepositUpdate) SetUpdatedAt(u uint32) *DepositUpdate {
	du.mutation.ResetUpdatedAt()
	du.mutation.SetUpdatedAt(u)
	return du
}

// AddUpdatedAt adds u to the "updated_at" field.
func (du *DepositUpdate) AddUpdatedAt(u int32) *DepositUpdate {
	du.mutation.AddUpdatedAt(u)
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DepositUpdate) SetDeletedAt(u uint32) *DepositUpdate {
	du.mutation.ResetDeletedAt()
	du.mutation.SetDeletedAt(u)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DepositUpdate) SetNillableDeletedAt(u *uint32) *DepositUpdate {
	if u != nil {
		du.SetDeletedAt(*u)
	}
	return du
}

// AddDeletedAt adds u to the "deleted_at" field.
func (du *DepositUpdate) AddDeletedAt(u int32) *DepositUpdate {
	du.mutation.AddDeletedAt(u)
	return du
}

// SetAppID sets the "app_id" field.
func (du *DepositUpdate) SetAppID(u uuid.UUID) *DepositUpdate {
	du.mutation.SetAppID(u)
	return du
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (du *DepositUpdate) SetNillableAppID(u *uuid.UUID) *DepositUpdate {
	if u != nil {
		du.SetAppID(*u)
	}
	return du
}

// ClearAppID clears the value of the "app_id" field.
func (du *DepositUpdate) ClearAppID() *DepositUpdate {
	du.mutation.ClearAppID()
	return du
}

// SetUserID sets the "user_id" field.
func (du *DepositUpdate) SetUserID(u uuid.UUID) *DepositUpdate {
	du.mutation.SetUserID(u)
	return du
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (du *DepositUpdate) SetNillableUserID(u *uuid.UUID) *DepositUpdate {
	if u != nil {
		du.SetUserID(*u)
	}
	return du
}

// ClearUserID clears the value of the "user_id" field.
func (du *DepositUpdate) ClearUserID() *DepositUpdate {
	du.mutation.ClearUserID()
	return du
}

// SetCoinTypeID sets the "coin_type_id" field.
func (du *DepositUpdate) SetCoinTypeID(u uuid.UUID) *DepositUpdate {
	du.mutation.SetCoinTypeID(u)
	return du
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (du *DepositUpdate) SetNillableCoinTypeID(u *uuid.UUID) *DepositUpdate {
	if u != nil {
		du.SetCoinTypeID(*u)
	}
	return du
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (du *DepositUpdate) ClearCoinTypeID() *DepositUpdate {
	du.mutation.ClearCoinTypeID()
	return du
}

// SetAccountID sets the "account_id" field.
func (du *DepositUpdate) SetAccountID(u uuid.UUID) *DepositUpdate {
	du.mutation.SetAccountID(u)
	return du
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (du *DepositUpdate) SetNillableAccountID(u *uuid.UUID) *DepositUpdate {
	if u != nil {
		du.SetAccountID(*u)
	}
	return du
}

// ClearAccountID clears the value of the "account_id" field.
func (du *DepositUpdate) ClearAccountID() *DepositUpdate {
	du.mutation.ClearAccountID()
	return du
}

// SetIncoming sets the "incoming" field.
func (du *DepositUpdate) SetIncoming(d decimal.Decimal) *DepositUpdate {
	du.mutation.SetIncoming(d)
	return du
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (du *DepositUpdate) SetNillableIncoming(d *decimal.Decimal) *DepositUpdate {
	if d != nil {
		du.SetIncoming(*d)
	}
	return du
}

// ClearIncoming clears the value of the "incoming" field.
func (du *DepositUpdate) ClearIncoming() *DepositUpdate {
	du.mutation.ClearIncoming()
	return du
}

// SetOutcoming sets the "outcoming" field.
func (du *DepositUpdate) SetOutcoming(d decimal.Decimal) *DepositUpdate {
	du.mutation.SetOutcoming(d)
	return du
}

// SetNillableOutcoming sets the "outcoming" field if the given value is not nil.
func (du *DepositUpdate) SetNillableOutcoming(d *decimal.Decimal) *DepositUpdate {
	if d != nil {
		du.SetOutcoming(*d)
	}
	return du
}

// ClearOutcoming clears the value of the "outcoming" field.
func (du *DepositUpdate) ClearOutcoming() *DepositUpdate {
	du.mutation.ClearOutcoming()
	return du
}

// SetCollectingTid sets the "collecting_tid" field.
func (du *DepositUpdate) SetCollectingTid(u uuid.UUID) *DepositUpdate {
	du.mutation.SetCollectingTid(u)
	return du
}

// SetNillableCollectingTid sets the "collecting_tid" field if the given value is not nil.
func (du *DepositUpdate) SetNillableCollectingTid(u *uuid.UUID) *DepositUpdate {
	if u != nil {
		du.SetCollectingTid(*u)
	}
	return du
}

// ClearCollectingTid clears the value of the "collecting_tid" field.
func (du *DepositUpdate) ClearCollectingTid() *DepositUpdate {
	du.mutation.ClearCollectingTid()
	return du
}

// SetScannableAt sets the "scannable_at" field.
func (du *DepositUpdate) SetScannableAt(u uint32) *DepositUpdate {
	du.mutation.ResetScannableAt()
	du.mutation.SetScannableAt(u)
	return du
}

// SetNillableScannableAt sets the "scannable_at" field if the given value is not nil.
func (du *DepositUpdate) SetNillableScannableAt(u *uint32) *DepositUpdate {
	if u != nil {
		du.SetScannableAt(*u)
	}
	return du
}

// AddScannableAt adds u to the "scannable_at" field.
func (du *DepositUpdate) AddScannableAt(u int32) *DepositUpdate {
	du.mutation.AddScannableAt(u)
	return du
}

// ClearScannableAt clears the value of the "scannable_at" field.
func (du *DepositUpdate) ClearScannableAt() *DepositUpdate {
	du.mutation.ClearScannableAt()
	return du
}

// Mutation returns the DepositMutation object of the builder.
func (du *DepositUpdate) Mutation() *DepositMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DepositUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := du.defaults(); err != nil {
		return 0, err
	}
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepositMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepositUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepositUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepositUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DepositUpdate) defaults() error {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		if deposit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deposit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deposit.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DepositUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DepositUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DepositUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deposit.Table,
			Columns: deposit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deposit.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldAppID,
		})
	}
	if du.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldAppID,
		})
	}
	if value, ok := du.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldUserID,
		})
	}
	if du.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldUserID,
		})
	}
	if value, ok := du.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldCoinTypeID,
		})
	}
	if du.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldCoinTypeID,
		})
	}
	if value, ok := du.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldAccountID,
		})
	}
	if du.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldAccountID,
		})
	}
	if value, ok := du.mutation.Incoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: deposit.FieldIncoming,
		})
	}
	if du.mutation.IncomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: deposit.FieldIncoming,
		})
	}
	if value, ok := du.mutation.Outcoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: deposit.FieldOutcoming,
		})
	}
	if du.mutation.OutcomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: deposit.FieldOutcoming,
		})
	}
	if value, ok := du.mutation.CollectingTid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldCollectingTid,
		})
	}
	if du.mutation.CollectingTidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldCollectingTid,
		})
	}
	if value, ok := du.mutation.ScannableAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldScannableAt,
		})
	}
	if value, ok := du.mutation.AddedScannableAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldScannableAt,
		})
	}
	if du.mutation.ScannableAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deposit.FieldScannableAt,
		})
	}
	_spec.Modifiers = du.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deposit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DepositUpdateOne is the builder for updating a single Deposit entity.
type DepositUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DepositMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (duo *DepositUpdateOne) SetCreatedAt(u uint32) *DepositUpdateOne {
	duo.mutation.ResetCreatedAt()
	duo.mutation.SetCreatedAt(u)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableCreatedAt(u *uint32) *DepositUpdateOne {
	if u != nil {
		duo.SetCreatedAt(*u)
	}
	return duo
}

// AddCreatedAt adds u to the "created_at" field.
func (duo *DepositUpdateOne) AddCreatedAt(u int32) *DepositUpdateOne {
	duo.mutation.AddCreatedAt(u)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DepositUpdateOne) SetUpdatedAt(u uint32) *DepositUpdateOne {
	duo.mutation.ResetUpdatedAt()
	duo.mutation.SetUpdatedAt(u)
	return duo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (duo *DepositUpdateOne) AddUpdatedAt(u int32) *DepositUpdateOne {
	duo.mutation.AddUpdatedAt(u)
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DepositUpdateOne) SetDeletedAt(u uint32) *DepositUpdateOne {
	duo.mutation.ResetDeletedAt()
	duo.mutation.SetDeletedAt(u)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableDeletedAt(u *uint32) *DepositUpdateOne {
	if u != nil {
		duo.SetDeletedAt(*u)
	}
	return duo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (duo *DepositUpdateOne) AddDeletedAt(u int32) *DepositUpdateOne {
	duo.mutation.AddDeletedAt(u)
	return duo
}

// SetAppID sets the "app_id" field.
func (duo *DepositUpdateOne) SetAppID(u uuid.UUID) *DepositUpdateOne {
	duo.mutation.SetAppID(u)
	return duo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableAppID(u *uuid.UUID) *DepositUpdateOne {
	if u != nil {
		duo.SetAppID(*u)
	}
	return duo
}

// ClearAppID clears the value of the "app_id" field.
func (duo *DepositUpdateOne) ClearAppID() *DepositUpdateOne {
	duo.mutation.ClearAppID()
	return duo
}

// SetUserID sets the "user_id" field.
func (duo *DepositUpdateOne) SetUserID(u uuid.UUID) *DepositUpdateOne {
	duo.mutation.SetUserID(u)
	return duo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableUserID(u *uuid.UUID) *DepositUpdateOne {
	if u != nil {
		duo.SetUserID(*u)
	}
	return duo
}

// ClearUserID clears the value of the "user_id" field.
func (duo *DepositUpdateOne) ClearUserID() *DepositUpdateOne {
	duo.mutation.ClearUserID()
	return duo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (duo *DepositUpdateOne) SetCoinTypeID(u uuid.UUID) *DepositUpdateOne {
	duo.mutation.SetCoinTypeID(u)
	return duo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *DepositUpdateOne {
	if u != nil {
		duo.SetCoinTypeID(*u)
	}
	return duo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (duo *DepositUpdateOne) ClearCoinTypeID() *DepositUpdateOne {
	duo.mutation.ClearCoinTypeID()
	return duo
}

// SetAccountID sets the "account_id" field.
func (duo *DepositUpdateOne) SetAccountID(u uuid.UUID) *DepositUpdateOne {
	duo.mutation.SetAccountID(u)
	return duo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableAccountID(u *uuid.UUID) *DepositUpdateOne {
	if u != nil {
		duo.SetAccountID(*u)
	}
	return duo
}

// ClearAccountID clears the value of the "account_id" field.
func (duo *DepositUpdateOne) ClearAccountID() *DepositUpdateOne {
	duo.mutation.ClearAccountID()
	return duo
}

// SetIncoming sets the "incoming" field.
func (duo *DepositUpdateOne) SetIncoming(d decimal.Decimal) *DepositUpdateOne {
	duo.mutation.SetIncoming(d)
	return duo
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableIncoming(d *decimal.Decimal) *DepositUpdateOne {
	if d != nil {
		duo.SetIncoming(*d)
	}
	return duo
}

// ClearIncoming clears the value of the "incoming" field.
func (duo *DepositUpdateOne) ClearIncoming() *DepositUpdateOne {
	duo.mutation.ClearIncoming()
	return duo
}

// SetOutcoming sets the "outcoming" field.
func (duo *DepositUpdateOne) SetOutcoming(d decimal.Decimal) *DepositUpdateOne {
	duo.mutation.SetOutcoming(d)
	return duo
}

// SetNillableOutcoming sets the "outcoming" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableOutcoming(d *decimal.Decimal) *DepositUpdateOne {
	if d != nil {
		duo.SetOutcoming(*d)
	}
	return duo
}

// ClearOutcoming clears the value of the "outcoming" field.
func (duo *DepositUpdateOne) ClearOutcoming() *DepositUpdateOne {
	duo.mutation.ClearOutcoming()
	return duo
}

// SetCollectingTid sets the "collecting_tid" field.
func (duo *DepositUpdateOne) SetCollectingTid(u uuid.UUID) *DepositUpdateOne {
	duo.mutation.SetCollectingTid(u)
	return duo
}

// SetNillableCollectingTid sets the "collecting_tid" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableCollectingTid(u *uuid.UUID) *DepositUpdateOne {
	if u != nil {
		duo.SetCollectingTid(*u)
	}
	return duo
}

// ClearCollectingTid clears the value of the "collecting_tid" field.
func (duo *DepositUpdateOne) ClearCollectingTid() *DepositUpdateOne {
	duo.mutation.ClearCollectingTid()
	return duo
}

// SetScannableAt sets the "scannable_at" field.
func (duo *DepositUpdateOne) SetScannableAt(u uint32) *DepositUpdateOne {
	duo.mutation.ResetScannableAt()
	duo.mutation.SetScannableAt(u)
	return duo
}

// SetNillableScannableAt sets the "scannable_at" field if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableScannableAt(u *uint32) *DepositUpdateOne {
	if u != nil {
		duo.SetScannableAt(*u)
	}
	return duo
}

// AddScannableAt adds u to the "scannable_at" field.
func (duo *DepositUpdateOne) AddScannableAt(u int32) *DepositUpdateOne {
	duo.mutation.AddScannableAt(u)
	return duo
}

// ClearScannableAt clears the value of the "scannable_at" field.
func (duo *DepositUpdateOne) ClearScannableAt() *DepositUpdateOne {
	duo.mutation.ClearScannableAt()
	return duo
}

// Mutation returns the DepositMutation object of the builder.
func (duo *DepositUpdateOne) Mutation() *DepositMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DepositUpdateOne) Select(field string, fields ...string) *DepositUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deposit entity.
func (duo *DepositUpdateOne) Save(ctx context.Context) (*Deposit, error) {
	var (
		err  error
		node *Deposit
	)
	if err := duo.defaults(); err != nil {
		return nil, err
	}
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepositMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Deposit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DepositMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepositUpdateOne) SaveX(ctx context.Context) *Deposit {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DepositUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepositUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DepositUpdateOne) defaults() error {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		if deposit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deposit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deposit.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DepositUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DepositUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DepositUpdateOne) sqlSave(ctx context.Context) (_node *Deposit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deposit.Table,
			Columns: deposit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deposit.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Deposit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deposit.FieldID)
		for _, f := range fields {
			if !deposit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deposit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldAppID,
		})
	}
	if duo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldAppID,
		})
	}
	if value, ok := duo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldUserID,
		})
	}
	if duo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldUserID,
		})
	}
	if value, ok := duo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldCoinTypeID,
		})
	}
	if duo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldCoinTypeID,
		})
	}
	if value, ok := duo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldAccountID,
		})
	}
	if duo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldAccountID,
		})
	}
	if value, ok := duo.mutation.Incoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: deposit.FieldIncoming,
		})
	}
	if duo.mutation.IncomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: deposit.FieldIncoming,
		})
	}
	if value, ok := duo.mutation.Outcoming(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: deposit.FieldOutcoming,
		})
	}
	if duo.mutation.OutcomingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: deposit.FieldOutcoming,
		})
	}
	if value, ok := duo.mutation.CollectingTid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deposit.FieldCollectingTid,
		})
	}
	if duo.mutation.CollectingTidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: deposit.FieldCollectingTid,
		})
	}
	if value, ok := duo.mutation.ScannableAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldScannableAt,
		})
	}
	if value, ok := duo.mutation.AddedScannableAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deposit.FieldScannableAt,
		})
	}
	if duo.mutation.ScannableAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deposit.FieldScannableAt,
		})
	}
	_spec.Modifiers = duo.modifiers
	_node = &Deposit{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deposit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
