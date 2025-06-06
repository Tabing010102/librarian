// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldEQ(FieldUserID, vc))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldEQ(FieldDeviceID, vc))
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldRefreshToken, v))
}

// ExpireAt applies equality check predicate on the "expire_at" field. It's identical to ExpireAtEQ.
func ExpireAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpireAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...model.InternalID) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Session(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...model.InternalID) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Session(sql.FieldNotIn(FieldUserID, v...))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldEQ(FieldDeviceID, vc))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v model.InternalID) predicate.Session {
	vc := int64(v)
	return predicate.Session(sql.FieldNEQ(FieldDeviceID, vc))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...model.InternalID) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Session(sql.FieldIn(FieldDeviceID, v...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...model.InternalID) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Session(sql.FieldNotIn(FieldDeviceID, v...))
}

// DeviceIDIsNil applies the IsNil predicate on the "device_id" field.
func DeviceIDIsNil() predicate.Session {
	return predicate.Session(sql.FieldIsNull(FieldDeviceID))
}

// DeviceIDNotNil applies the NotNil predicate on the "device_id" field.
func DeviceIDNotNil() predicate.Session {
	return predicate.Session(sql.FieldNotNull(FieldDeviceID))
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldRefreshToken, v))
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldRefreshToken, v))
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldRefreshToken, vs...))
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldRefreshToken, vs...))
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldRefreshToken, v))
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldRefreshToken, v))
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldRefreshToken, v))
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldRefreshToken, v))
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldRefreshToken, v))
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldRefreshToken, v))
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldRefreshToken, v))
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldRefreshToken, v))
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldRefreshToken, v))
}

// ExpireAtEQ applies the EQ predicate on the "expire_at" field.
func ExpireAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpireAt, v))
}

// ExpireAtNEQ applies the NEQ predicate on the "expire_at" field.
func ExpireAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldExpireAt, v))
}

// ExpireAtIn applies the In predicate on the "expire_at" field.
func ExpireAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldExpireAt, vs...))
}

// ExpireAtNotIn applies the NotIn predicate on the "expire_at" field.
func ExpireAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldExpireAt, vs...))
}

// ExpireAtGT applies the GT predicate on the "expire_at" field.
func ExpireAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldExpireAt, v))
}

// ExpireAtGTE applies the GTE predicate on the "expire_at" field.
func ExpireAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldExpireAt, v))
}

// ExpireAtLT applies the LT predicate on the "expire_at" field.
func ExpireAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldExpireAt, v))
}

// ExpireAtLTE applies the LTE predicate on the "expire_at" field.
func ExpireAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldExpireAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(sql.NotPredicates(p))
}
