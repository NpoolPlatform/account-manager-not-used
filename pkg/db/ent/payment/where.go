// Code generated by ent, DO NOT EDIT.

package payment

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// CollectingTid applies equality check predicate on the "collecting_tid" field. It's identical to CollectingTidEQ.
func CollectingTid(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectingTid), v))
	})
}

// AvailableAt applies equality check predicate on the "available_at" field. It's identical to AvailableAtEQ.
func AvailableAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// AccountIDIsNil applies the IsNil predicate on the "account_id" field.
func AccountIDIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAccountID)))
	})
}

// AccountIDNotNil applies the NotNil predicate on the "account_id" field.
func AccountIDNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAccountID)))
	})
}

// CollectingTidEQ applies the EQ predicate on the "collecting_tid" field.
func CollectingTidEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidNEQ applies the NEQ predicate on the "collecting_tid" field.
func CollectingTidNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidIn applies the In predicate on the "collecting_tid" field.
func CollectingTidIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCollectingTid), v...))
	})
}

// CollectingTidNotIn applies the NotIn predicate on the "collecting_tid" field.
func CollectingTidNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCollectingTid), v...))
	})
}

// CollectingTidGT applies the GT predicate on the "collecting_tid" field.
func CollectingTidGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidGTE applies the GTE predicate on the "collecting_tid" field.
func CollectingTidGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidLT applies the LT predicate on the "collecting_tid" field.
func CollectingTidLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidLTE applies the LTE predicate on the "collecting_tid" field.
func CollectingTidLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCollectingTid), v))
	})
}

// CollectingTidIsNil applies the IsNil predicate on the "collecting_tid" field.
func CollectingTidIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCollectingTid)))
	})
}

// CollectingTidNotNil applies the NotNil predicate on the "collecting_tid" field.
func CollectingTidNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCollectingTid)))
	})
}

// AvailableAtEQ applies the EQ predicate on the "available_at" field.
func AvailableAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtNEQ applies the NEQ predicate on the "available_at" field.
func AvailableAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtIn applies the In predicate on the "available_at" field.
func AvailableAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAvailableAt), v...))
	})
}

// AvailableAtNotIn applies the NotIn predicate on the "available_at" field.
func AvailableAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAvailableAt), v...))
	})
}

// AvailableAtGT applies the GT predicate on the "available_at" field.
func AvailableAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtGTE applies the GTE predicate on the "available_at" field.
func AvailableAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtLT applies the LT predicate on the "available_at" field.
func AvailableAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtLTE applies the LTE predicate on the "available_at" field.
func AvailableAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtIsNil applies the IsNil predicate on the "available_at" field.
func AvailableAtIsNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvailableAt)))
	})
}

// AvailableAtNotNil applies the NotNil predicate on the "available_at" field.
func AvailableAtNotNil() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvailableAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
