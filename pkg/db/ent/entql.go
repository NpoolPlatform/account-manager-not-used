// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/payment"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
		Type: "Account",
		Fields: map[string]*sqlgraph.FieldSpec{
			account.FieldCreatedAt:              {Type: field.TypeUint32, Column: account.FieldCreatedAt},
			account.FieldUpdatedAt:              {Type: field.TypeUint32, Column: account.FieldUpdatedAt},
			account.FieldDeletedAt:              {Type: field.TypeUint32, Column: account.FieldDeletedAt},
			account.FieldCoinTypeID:             {Type: field.TypeUUID, Column: account.FieldCoinTypeID},
			account.FieldAddress:                {Type: field.TypeString, Column: account.FieldAddress},
			account.FieldUsedFor:                {Type: field.TypeString, Column: account.FieldUsedFor},
			account.FieldPlatformHoldPrivateKey: {Type: field.TypeBool, Column: account.FieldPlatformHoldPrivateKey},
			account.FieldActive:                 {Type: field.TypeBool, Column: account.FieldActive},
			account.FieldLocked:                 {Type: field.TypeBool, Column: account.FieldLocked},
			account.FieldBlocked:                {Type: field.TypeBool, Column: account.FieldBlocked},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   goodbenefit.Table,
			Columns: goodbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodbenefit.FieldID,
			},
		},
		Type: "GoodBenefit",
		Fields: map[string]*sqlgraph.FieldSpec{
			goodbenefit.FieldCreatedAt: {Type: field.TypeUint32, Column: goodbenefit.FieldCreatedAt},
			goodbenefit.FieldUpdatedAt: {Type: field.TypeUint32, Column: goodbenefit.FieldUpdatedAt},
			goodbenefit.FieldDeletedAt: {Type: field.TypeUint32, Column: goodbenefit.FieldDeletedAt},
			goodbenefit.FieldGoodID:    {Type: field.TypeUUID, Column: goodbenefit.FieldGoodID},
			goodbenefit.FieldAccountID: {Type: field.TypeUUID, Column: goodbenefit.FieldAccountID},
			goodbenefit.FieldBackup:    {Type: field.TypeBool, Column: goodbenefit.FieldBackup},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: payment.FieldID,
			},
		},
		Type: "Payment",
		Fields: map[string]*sqlgraph.FieldSpec{
			payment.FieldCreatedAt:     {Type: field.TypeUint32, Column: payment.FieldCreatedAt},
			payment.FieldUpdatedAt:     {Type: field.TypeUint32, Column: payment.FieldUpdatedAt},
			payment.FieldDeletedAt:     {Type: field.TypeUint32, Column: payment.FieldDeletedAt},
			payment.FieldCoinTypeID:    {Type: field.TypeUUID, Column: payment.FieldCoinTypeID},
			payment.FieldAccountID:     {Type: field.TypeUUID, Column: payment.FieldAccountID},
			payment.FieldIdle:          {Type: field.TypeBool, Column: payment.FieldIdle},
			payment.FieldOccupiedBy:    {Type: field.TypeString, Column: payment.FieldOccupiedBy},
			payment.FieldCollectingTid: {Type: field.TypeUUID, Column: payment.FieldCollectingTid},
			payment.FieldAvailableAt:   {Type: field.TypeUint32, Column: payment.FieldAvailableAt},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   platform.Table,
			Columns: platform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: platform.FieldID,
			},
		},
		Type: "Platform",
		Fields: map[string]*sqlgraph.FieldSpec{
			platform.FieldCreatedAt:  {Type: field.TypeUint32, Column: platform.FieldCreatedAt},
			platform.FieldUpdatedAt:  {Type: field.TypeUint32, Column: platform.FieldUpdatedAt},
			platform.FieldDeletedAt:  {Type: field.TypeUint32, Column: platform.FieldDeletedAt},
			platform.FieldCoinTypeID: {Type: field.TypeUUID, Column: platform.FieldCoinTypeID},
			platform.FieldAccountID:  {Type: field.TypeUUID, Column: platform.FieldAccountID},
			platform.FieldUsedFor:    {Type: field.TypeString, Column: platform.FieldUsedFor},
			platform.FieldBackup:     {Type: field.TypeBool, Column: platform.FieldBackup},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreatedAt:  {Type: field.TypeUint32, Column: user.FieldCreatedAt},
			user.FieldUpdatedAt:  {Type: field.TypeUint32, Column: user.FieldUpdatedAt},
			user.FieldDeletedAt:  {Type: field.TypeUint32, Column: user.FieldDeletedAt},
			user.FieldAppID:      {Type: field.TypeUUID, Column: user.FieldAppID},
			user.FieldUserID:     {Type: field.TypeUUID, Column: user.FieldUserID},
			user.FieldCoinTypeID: {Type: field.TypeUUID, Column: user.FieldCoinTypeID},
			user.FieldAccountID:  {Type: field.TypeUUID, Column: user.FieldAccountID},
			user.FieldUsedFor:    {Type: field.TypeString, Column: user.FieldUsedFor},
			user.FieldLabels:     {Type: field.TypeJSON, Column: user.FieldLabels},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *AccountQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AccountQuery builder.
func (aq *AccountQuery) Filter() *AccountFilter {
	return &AccountFilter{aq}
}

// addPredicate implements the predicateAdder interface.
func (m *AccountMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AccountMutation builder.
func (m *AccountMutation) Filter() *AccountFilter {
	return &AccountFilter{m}
}

// AccountFilter provides a generic filtering capability at runtime for AccountQuery.
type AccountFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *AccountFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AccountFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(account.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AccountFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AccountFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AccountFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *AccountFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(account.FieldCoinTypeID))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *AccountFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(account.FieldAddress))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *AccountFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(account.FieldUsedFor))
}

// WherePlatformHoldPrivateKey applies the entql bool predicate on the platform_hold_private_key field.
func (f *AccountFilter) WherePlatformHoldPrivateKey(p entql.BoolP) {
	f.Where(p.Field(account.FieldPlatformHoldPrivateKey))
}

// WhereActive applies the entql bool predicate on the active field.
func (f *AccountFilter) WhereActive(p entql.BoolP) {
	f.Where(p.Field(account.FieldActive))
}

// WhereLocked applies the entql bool predicate on the locked field.
func (f *AccountFilter) WhereLocked(p entql.BoolP) {
	f.Where(p.Field(account.FieldLocked))
}

// WhereBlocked applies the entql bool predicate on the blocked field.
func (f *AccountFilter) WhereBlocked(p entql.BoolP) {
	f.Where(p.Field(account.FieldBlocked))
}

// addPredicate implements the predicateAdder interface.
func (gbq *GoodBenefitQuery) addPredicate(pred func(s *sql.Selector)) {
	gbq.predicates = append(gbq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GoodBenefitQuery builder.
func (gbq *GoodBenefitQuery) Filter() *GoodBenefitFilter {
	return &GoodBenefitFilter{gbq}
}

// addPredicate implements the predicateAdder interface.
func (m *GoodBenefitMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GoodBenefitMutation builder.
func (m *GoodBenefitMutation) Filter() *GoodBenefitFilter {
	return &GoodBenefitFilter{m}
}

// GoodBenefitFilter provides a generic filtering capability at runtime for GoodBenefitQuery.
type GoodBenefitFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *GoodBenefitFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *GoodBenefitFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(goodbenefit.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *GoodBenefitFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodbenefit.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *GoodBenefitFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodbenefit.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *GoodBenefitFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(goodbenefit.FieldDeletedAt))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *GoodBenefitFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(goodbenefit.FieldGoodID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *GoodBenefitFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(goodbenefit.FieldAccountID))
}

// WhereBackup applies the entql bool predicate on the backup field.
func (f *GoodBenefitFilter) WhereBackup(p entql.BoolP) {
	f.Where(p.Field(goodbenefit.FieldBackup))
}

// addPredicate implements the predicateAdder interface.
func (pq *PaymentQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PaymentQuery builder.
func (pq *PaymentQuery) Filter() *PaymentFilter {
	return &PaymentFilter{pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PaymentMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PaymentMutation builder.
func (m *PaymentMutation) Filter() *PaymentFilter {
	return &PaymentFilter{m}
}

// PaymentFilter provides a generic filtering capability at runtime for PaymentQuery.
type PaymentFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *PaymentFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PaymentFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PaymentFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PaymentFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PaymentFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *PaymentFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldCoinTypeID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *PaymentFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(payment.FieldAccountID))
}

// WhereIdle applies the entql bool predicate on the idle field.
func (f *PaymentFilter) WhereIdle(p entql.BoolP) {
	f.Where(p.Field(payment.FieldIdle))
}

// WhereOccupiedBy applies the entql string predicate on the occupied_by field.
func (f *PaymentFilter) WhereOccupiedBy(p entql.StringP) {
	f.Where(p.Field(payment.FieldOccupiedBy))
}

// WhereCollectingTid applies the entql [16]byte predicate on the collecting_tid field.
func (f *PaymentFilter) WhereCollectingTid(p entql.ValueP) {
	f.Where(p.Field(payment.FieldCollectingTid))
}

// WhereAvailableAt applies the entql uint32 predicate on the available_at field.
func (f *PaymentFilter) WhereAvailableAt(p entql.Uint32P) {
	f.Where(p.Field(payment.FieldAvailableAt))
}

// addPredicate implements the predicateAdder interface.
func (pq *PlatformQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlatformQuery builder.
func (pq *PlatformQuery) Filter() *PlatformFilter {
	return &PlatformFilter{pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlatformMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlatformMutation builder.
func (m *PlatformMutation) Filter() *PlatformFilter {
	return &PlatformFilter{m}
}

// PlatformFilter provides a generic filtering capability at runtime for PlatformQuery.
type PlatformFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *PlatformFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PlatformFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(platform.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PlatformFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(platform.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PlatformFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(platform.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PlatformFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(platform.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *PlatformFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(platform.FieldCoinTypeID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *PlatformFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(platform.FieldAccountID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *PlatformFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(platform.FieldUsedFor))
}

// WhereBackup applies the entql bool predicate on the backup field.
func (f *PlatformFilter) WhereBackup(p entql.BoolP) {
	f.Where(p.Field(platform.FieldBackup))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *UserFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(user.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *UserFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(user.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *UserFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(user.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *UserFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(user.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *UserFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(user.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *UserFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(user.FieldCoinTypeID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *UserFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(user.FieldAccountID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *UserFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(user.FieldUsedFor))
}

// WhereLabels applies the entql json.RawMessage predicate on the labels field.
func (f *UserFilter) WhereLabels(p entql.BytesP) {
	f.Where(p.Field(user.FieldLabels))
}
