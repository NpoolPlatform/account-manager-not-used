// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodBenefitUpdate is the builder for updating GoodBenefit entities.
type GoodBenefitUpdate struct {
	config
	hooks    []Hook
	mutation *GoodBenefitMutation
}

// Where appends a list predicates to the GoodBenefitUpdate builder.
func (gbu *GoodBenefitUpdate) Where(ps ...predicate.GoodBenefit) *GoodBenefitUpdate {
	gbu.mutation.Where(ps...)
	return gbu
}

// SetCreatedAt sets the "created_at" field.
func (gbu *GoodBenefitUpdate) SetCreatedAt(u uint32) *GoodBenefitUpdate {
	gbu.mutation.ResetCreatedAt()
	gbu.mutation.SetCreatedAt(u)
	return gbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gbu *GoodBenefitUpdate) SetNillableCreatedAt(u *uint32) *GoodBenefitUpdate {
	if u != nil {
		gbu.SetCreatedAt(*u)
	}
	return gbu
}

// AddCreatedAt adds u to the "created_at" field.
func (gbu *GoodBenefitUpdate) AddCreatedAt(u int32) *GoodBenefitUpdate {
	gbu.mutation.AddCreatedAt(u)
	return gbu
}

// SetUpdatedAt sets the "updated_at" field.
func (gbu *GoodBenefitUpdate) SetUpdatedAt(u uint32) *GoodBenefitUpdate {
	gbu.mutation.ResetUpdatedAt()
	gbu.mutation.SetUpdatedAt(u)
	return gbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gbu *GoodBenefitUpdate) AddUpdatedAt(u int32) *GoodBenefitUpdate {
	gbu.mutation.AddUpdatedAt(u)
	return gbu
}

// SetDeletedAt sets the "deleted_at" field.
func (gbu *GoodBenefitUpdate) SetDeletedAt(u uint32) *GoodBenefitUpdate {
	gbu.mutation.ResetDeletedAt()
	gbu.mutation.SetDeletedAt(u)
	return gbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gbu *GoodBenefitUpdate) SetNillableDeletedAt(u *uint32) *GoodBenefitUpdate {
	if u != nil {
		gbu.SetDeletedAt(*u)
	}
	return gbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gbu *GoodBenefitUpdate) AddDeletedAt(u int32) *GoodBenefitUpdate {
	gbu.mutation.AddDeletedAt(u)
	return gbu
}

// SetGoodID sets the "good_id" field.
func (gbu *GoodBenefitUpdate) SetGoodID(u uuid.UUID) *GoodBenefitUpdate {
	gbu.mutation.SetGoodID(u)
	return gbu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gbu *GoodBenefitUpdate) SetNillableGoodID(u *uuid.UUID) *GoodBenefitUpdate {
	if u != nil {
		gbu.SetGoodID(*u)
	}
	return gbu
}

// ClearGoodID clears the value of the "good_id" field.
func (gbu *GoodBenefitUpdate) ClearGoodID() *GoodBenefitUpdate {
	gbu.mutation.ClearGoodID()
	return gbu
}

// SetAccountID sets the "account_id" field.
func (gbu *GoodBenefitUpdate) SetAccountID(u uuid.UUID) *GoodBenefitUpdate {
	gbu.mutation.SetAccountID(u)
	return gbu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (gbu *GoodBenefitUpdate) SetNillableAccountID(u *uuid.UUID) *GoodBenefitUpdate {
	if u != nil {
		gbu.SetAccountID(*u)
	}
	return gbu
}

// ClearAccountID clears the value of the "account_id" field.
func (gbu *GoodBenefitUpdate) ClearAccountID() *GoodBenefitUpdate {
	gbu.mutation.ClearAccountID()
	return gbu
}

// SetBackup sets the "backup" field.
func (gbu *GoodBenefitUpdate) SetBackup(b bool) *GoodBenefitUpdate {
	gbu.mutation.SetBackup(b)
	return gbu
}

// SetNillableBackup sets the "backup" field if the given value is not nil.
func (gbu *GoodBenefitUpdate) SetNillableBackup(b *bool) *GoodBenefitUpdate {
	if b != nil {
		gbu.SetBackup(*b)
	}
	return gbu
}

// ClearBackup clears the value of the "backup" field.
func (gbu *GoodBenefitUpdate) ClearBackup() *GoodBenefitUpdate {
	gbu.mutation.ClearBackup()
	return gbu
}

// Mutation returns the GoodBenefitMutation object of the builder.
func (gbu *GoodBenefitUpdate) Mutation() *GoodBenefitMutation {
	return gbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gbu *GoodBenefitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gbu.defaults(); err != nil {
		return 0, err
	}
	if len(gbu.hooks) == 0 {
		affected, err = gbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gbu.mutation = mutation
			affected, err = gbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gbu.hooks) - 1; i >= 0; i-- {
			if gbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gbu *GoodBenefitUpdate) SaveX(ctx context.Context) int {
	affected, err := gbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gbu *GoodBenefitUpdate) Exec(ctx context.Context) error {
	_, err := gbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbu *GoodBenefitUpdate) ExecX(ctx context.Context) {
	if err := gbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gbu *GoodBenefitUpdate) defaults() error {
	if _, ok := gbu.mutation.UpdatedAt(); !ok {
		if goodbenefit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodbenefit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodbenefit.UpdateDefaultUpdatedAt()
		gbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gbu *GoodBenefitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodbenefit.Table,
			Columns: goodbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodbenefit.FieldID,
			},
		},
	}
	if ps := gbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldCreatedAt,
		})
	}
	if value, ok := gbu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldCreatedAt,
		})
	}
	if value, ok := gbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := gbu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := gbu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldDeletedAt,
		})
	}
	if value, ok := gbu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldDeletedAt,
		})
	}
	if value, ok := gbu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodbenefit.FieldGoodID,
		})
	}
	if gbu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodbenefit.FieldGoodID,
		})
	}
	if value, ok := gbu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodbenefit.FieldAccountID,
		})
	}
	if gbu.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodbenefit.FieldAccountID,
		})
	}
	if value, ok := gbu.mutation.Backup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodbenefit.FieldBackup,
		})
	}
	if gbu.mutation.BackupCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: goodbenefit.FieldBackup,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodbenefit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GoodBenefitUpdateOne is the builder for updating a single GoodBenefit entity.
type GoodBenefitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodBenefitMutation
}

// SetCreatedAt sets the "created_at" field.
func (gbuo *GoodBenefitUpdateOne) SetCreatedAt(u uint32) *GoodBenefitUpdateOne {
	gbuo.mutation.ResetCreatedAt()
	gbuo.mutation.SetCreatedAt(u)
	return gbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gbuo *GoodBenefitUpdateOne) SetNillableCreatedAt(u *uint32) *GoodBenefitUpdateOne {
	if u != nil {
		gbuo.SetCreatedAt(*u)
	}
	return gbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (gbuo *GoodBenefitUpdateOne) AddCreatedAt(u int32) *GoodBenefitUpdateOne {
	gbuo.mutation.AddCreatedAt(u)
	return gbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gbuo *GoodBenefitUpdateOne) SetUpdatedAt(u uint32) *GoodBenefitUpdateOne {
	gbuo.mutation.ResetUpdatedAt()
	gbuo.mutation.SetUpdatedAt(u)
	return gbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gbuo *GoodBenefitUpdateOne) AddUpdatedAt(u int32) *GoodBenefitUpdateOne {
	gbuo.mutation.AddUpdatedAt(u)
	return gbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gbuo *GoodBenefitUpdateOne) SetDeletedAt(u uint32) *GoodBenefitUpdateOne {
	gbuo.mutation.ResetDeletedAt()
	gbuo.mutation.SetDeletedAt(u)
	return gbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gbuo *GoodBenefitUpdateOne) SetNillableDeletedAt(u *uint32) *GoodBenefitUpdateOne {
	if u != nil {
		gbuo.SetDeletedAt(*u)
	}
	return gbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gbuo *GoodBenefitUpdateOne) AddDeletedAt(u int32) *GoodBenefitUpdateOne {
	gbuo.mutation.AddDeletedAt(u)
	return gbuo
}

// SetGoodID sets the "good_id" field.
func (gbuo *GoodBenefitUpdateOne) SetGoodID(u uuid.UUID) *GoodBenefitUpdateOne {
	gbuo.mutation.SetGoodID(u)
	return gbuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gbuo *GoodBenefitUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodBenefitUpdateOne {
	if u != nil {
		gbuo.SetGoodID(*u)
	}
	return gbuo
}

// ClearGoodID clears the value of the "good_id" field.
func (gbuo *GoodBenefitUpdateOne) ClearGoodID() *GoodBenefitUpdateOne {
	gbuo.mutation.ClearGoodID()
	return gbuo
}

// SetAccountID sets the "account_id" field.
func (gbuo *GoodBenefitUpdateOne) SetAccountID(u uuid.UUID) *GoodBenefitUpdateOne {
	gbuo.mutation.SetAccountID(u)
	return gbuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (gbuo *GoodBenefitUpdateOne) SetNillableAccountID(u *uuid.UUID) *GoodBenefitUpdateOne {
	if u != nil {
		gbuo.SetAccountID(*u)
	}
	return gbuo
}

// ClearAccountID clears the value of the "account_id" field.
func (gbuo *GoodBenefitUpdateOne) ClearAccountID() *GoodBenefitUpdateOne {
	gbuo.mutation.ClearAccountID()
	return gbuo
}

// SetBackup sets the "backup" field.
func (gbuo *GoodBenefitUpdateOne) SetBackup(b bool) *GoodBenefitUpdateOne {
	gbuo.mutation.SetBackup(b)
	return gbuo
}

// SetNillableBackup sets the "backup" field if the given value is not nil.
func (gbuo *GoodBenefitUpdateOne) SetNillableBackup(b *bool) *GoodBenefitUpdateOne {
	if b != nil {
		gbuo.SetBackup(*b)
	}
	return gbuo
}

// ClearBackup clears the value of the "backup" field.
func (gbuo *GoodBenefitUpdateOne) ClearBackup() *GoodBenefitUpdateOne {
	gbuo.mutation.ClearBackup()
	return gbuo
}

// Mutation returns the GoodBenefitMutation object of the builder.
func (gbuo *GoodBenefitUpdateOne) Mutation() *GoodBenefitMutation {
	return gbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gbuo *GoodBenefitUpdateOne) Select(field string, fields ...string) *GoodBenefitUpdateOne {
	gbuo.fields = append([]string{field}, fields...)
	return gbuo
}

// Save executes the query and returns the updated GoodBenefit entity.
func (gbuo *GoodBenefitUpdateOne) Save(ctx context.Context) (*GoodBenefit, error) {
	var (
		err  error
		node *GoodBenefit
	)
	if err := gbuo.defaults(); err != nil {
		return nil, err
	}
	if len(gbuo.hooks) == 0 {
		node, err = gbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gbuo.mutation = mutation
			node, err = gbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gbuo.hooks) - 1; i >= 0; i-- {
			if gbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gbuo *GoodBenefitUpdateOne) SaveX(ctx context.Context) *GoodBenefit {
	node, err := gbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gbuo *GoodBenefitUpdateOne) Exec(ctx context.Context) error {
	_, err := gbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbuo *GoodBenefitUpdateOne) ExecX(ctx context.Context) {
	if err := gbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gbuo *GoodBenefitUpdateOne) defaults() error {
	if _, ok := gbuo.mutation.UpdatedAt(); !ok {
		if goodbenefit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodbenefit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodbenefit.UpdateDefaultUpdatedAt()
		gbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gbuo *GoodBenefitUpdateOne) sqlSave(ctx context.Context) (_node *GoodBenefit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodbenefit.Table,
			Columns: goodbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodbenefit.FieldID,
			},
		},
	}
	id, ok := gbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodBenefit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodbenefit.FieldID)
		for _, f := range fields {
			if !goodbenefit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldCreatedAt,
		})
	}
	if value, ok := gbuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldCreatedAt,
		})
	}
	if value, ok := gbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := gbuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := gbuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldDeletedAt,
		})
	}
	if value, ok := gbuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodbenefit.FieldDeletedAt,
		})
	}
	if value, ok := gbuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodbenefit.FieldGoodID,
		})
	}
	if gbuo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodbenefit.FieldGoodID,
		})
	}
	if value, ok := gbuo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodbenefit.FieldAccountID,
		})
	}
	if gbuo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodbenefit.FieldAccountID,
		})
	}
	if value, ok := gbuo.mutation.Backup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodbenefit.FieldBackup,
		})
	}
	if gbuo.mutation.BackupCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: goodbenefit.FieldBackup,
		})
	}
	_node = &GoodBenefit{config: gbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodbenefit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}