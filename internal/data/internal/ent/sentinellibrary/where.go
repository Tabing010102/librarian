// Code generated by ent, DO NOT EDIT.

package sentinellibrary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldID, id))
}

// SentinelID applies equality check predicate on the "sentinel_id" field. It's identical to SentinelIDEQ.
func SentinelID(v model.InternalID) predicate.SentinelLibrary {
	vc := int64(v)
	return predicate.SentinelLibrary(sql.FieldEQ(FieldSentinelID, vc))
}

// ReportedID applies equality check predicate on the "reported_id" field. It's identical to ReportedIDEQ.
func ReportedID(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldReportedID, v))
}

// DownloadBasePath applies equality check predicate on the "download_base_path" field. It's identical to DownloadBasePathEQ.
func DownloadBasePath(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldDownloadBasePath, v))
}

// ActiveSnapshot applies equality check predicate on the "active_snapshot" field. It's identical to ActiveSnapshotEQ.
func ActiveSnapshot(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldActiveSnapshot, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldCreatedAt, v))
}

// LibraryReportSequence applies equality check predicate on the "library_report_sequence" field. It's identical to LibraryReportSequenceEQ.
func LibraryReportSequence(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldLibraryReportSequence, v))
}

// SentinelIDEQ applies the EQ predicate on the "sentinel_id" field.
func SentinelIDEQ(v model.InternalID) predicate.SentinelLibrary {
	vc := int64(v)
	return predicate.SentinelLibrary(sql.FieldEQ(FieldSentinelID, vc))
}

// SentinelIDNEQ applies the NEQ predicate on the "sentinel_id" field.
func SentinelIDNEQ(v model.InternalID) predicate.SentinelLibrary {
	vc := int64(v)
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldSentinelID, vc))
}

// SentinelIDIn applies the In predicate on the "sentinel_id" field.
func SentinelIDIn(vs ...model.InternalID) predicate.SentinelLibrary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.SentinelLibrary(sql.FieldIn(FieldSentinelID, v...))
}

// SentinelIDNotIn applies the NotIn predicate on the "sentinel_id" field.
func SentinelIDNotIn(vs ...model.InternalID) predicate.SentinelLibrary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldSentinelID, v...))
}

// ReportedIDEQ applies the EQ predicate on the "reported_id" field.
func ReportedIDEQ(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldReportedID, v))
}

// ReportedIDNEQ applies the NEQ predicate on the "reported_id" field.
func ReportedIDNEQ(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldReportedID, v))
}

// ReportedIDIn applies the In predicate on the "reported_id" field.
func ReportedIDIn(vs ...int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldReportedID, vs...))
}

// ReportedIDNotIn applies the NotIn predicate on the "reported_id" field.
func ReportedIDNotIn(vs ...int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldReportedID, vs...))
}

// ReportedIDGT applies the GT predicate on the "reported_id" field.
func ReportedIDGT(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldReportedID, v))
}

// ReportedIDGTE applies the GTE predicate on the "reported_id" field.
func ReportedIDGTE(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldReportedID, v))
}

// ReportedIDLT applies the LT predicate on the "reported_id" field.
func ReportedIDLT(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldReportedID, v))
}

// ReportedIDLTE applies the LTE predicate on the "reported_id" field.
func ReportedIDLTE(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldReportedID, v))
}

// DownloadBasePathEQ applies the EQ predicate on the "download_base_path" field.
func DownloadBasePathEQ(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldDownloadBasePath, v))
}

// DownloadBasePathNEQ applies the NEQ predicate on the "download_base_path" field.
func DownloadBasePathNEQ(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldDownloadBasePath, v))
}

// DownloadBasePathIn applies the In predicate on the "download_base_path" field.
func DownloadBasePathIn(vs ...string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldDownloadBasePath, vs...))
}

// DownloadBasePathNotIn applies the NotIn predicate on the "download_base_path" field.
func DownloadBasePathNotIn(vs ...string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldDownloadBasePath, vs...))
}

// DownloadBasePathGT applies the GT predicate on the "download_base_path" field.
func DownloadBasePathGT(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldDownloadBasePath, v))
}

// DownloadBasePathGTE applies the GTE predicate on the "download_base_path" field.
func DownloadBasePathGTE(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldDownloadBasePath, v))
}

// DownloadBasePathLT applies the LT predicate on the "download_base_path" field.
func DownloadBasePathLT(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldDownloadBasePath, v))
}

// DownloadBasePathLTE applies the LTE predicate on the "download_base_path" field.
func DownloadBasePathLTE(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldDownloadBasePath, v))
}

// DownloadBasePathContains applies the Contains predicate on the "download_base_path" field.
func DownloadBasePathContains(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldContains(FieldDownloadBasePath, v))
}

// DownloadBasePathHasPrefix applies the HasPrefix predicate on the "download_base_path" field.
func DownloadBasePathHasPrefix(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldHasPrefix(FieldDownloadBasePath, v))
}

// DownloadBasePathHasSuffix applies the HasSuffix predicate on the "download_base_path" field.
func DownloadBasePathHasSuffix(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldHasSuffix(FieldDownloadBasePath, v))
}

// DownloadBasePathEqualFold applies the EqualFold predicate on the "download_base_path" field.
func DownloadBasePathEqualFold(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEqualFold(FieldDownloadBasePath, v))
}

// DownloadBasePathContainsFold applies the ContainsFold predicate on the "download_base_path" field.
func DownloadBasePathContainsFold(v string) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldContainsFold(FieldDownloadBasePath, v))
}

// ActiveSnapshotEQ applies the EQ predicate on the "active_snapshot" field.
func ActiveSnapshotEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldActiveSnapshot, v))
}

// ActiveSnapshotNEQ applies the NEQ predicate on the "active_snapshot" field.
func ActiveSnapshotNEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldActiveSnapshot, v))
}

// ActiveSnapshotIn applies the In predicate on the "active_snapshot" field.
func ActiveSnapshotIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldActiveSnapshot, vs...))
}

// ActiveSnapshotNotIn applies the NotIn predicate on the "active_snapshot" field.
func ActiveSnapshotNotIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldActiveSnapshot, vs...))
}

// ActiveSnapshotGT applies the GT predicate on the "active_snapshot" field.
func ActiveSnapshotGT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldActiveSnapshot, v))
}

// ActiveSnapshotGTE applies the GTE predicate on the "active_snapshot" field.
func ActiveSnapshotGTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldActiveSnapshot, v))
}

// ActiveSnapshotLT applies the LT predicate on the "active_snapshot" field.
func ActiveSnapshotLT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldActiveSnapshot, v))
}

// ActiveSnapshotLTE applies the LTE predicate on the "active_snapshot" field.
func ActiveSnapshotLTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldActiveSnapshot, v))
}

// ActiveSnapshotIsNil applies the IsNil predicate on the "active_snapshot" field.
func ActiveSnapshotIsNil() predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIsNull(FieldActiveSnapshot))
}

// ActiveSnapshotNotNil applies the NotNil predicate on the "active_snapshot" field.
func ActiveSnapshotNotNil() predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotNull(FieldActiveSnapshot))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldCreatedAt, v))
}

// LibraryReportSequenceEQ applies the EQ predicate on the "library_report_sequence" field.
func LibraryReportSequenceEQ(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldEQ(FieldLibraryReportSequence, v))
}

// LibraryReportSequenceNEQ applies the NEQ predicate on the "library_report_sequence" field.
func LibraryReportSequenceNEQ(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNEQ(FieldLibraryReportSequence, v))
}

// LibraryReportSequenceIn applies the In predicate on the "library_report_sequence" field.
func LibraryReportSequenceIn(vs ...int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldIn(FieldLibraryReportSequence, vs...))
}

// LibraryReportSequenceNotIn applies the NotIn predicate on the "library_report_sequence" field.
func LibraryReportSequenceNotIn(vs ...int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldNotIn(FieldLibraryReportSequence, vs...))
}

// LibraryReportSequenceGT applies the GT predicate on the "library_report_sequence" field.
func LibraryReportSequenceGT(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGT(FieldLibraryReportSequence, v))
}

// LibraryReportSequenceGTE applies the GTE predicate on the "library_report_sequence" field.
func LibraryReportSequenceGTE(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldGTE(FieldLibraryReportSequence, v))
}

// LibraryReportSequenceLT applies the LT predicate on the "library_report_sequence" field.
func LibraryReportSequenceLT(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLT(FieldLibraryReportSequence, v))
}

// LibraryReportSequenceLTE applies the LTE predicate on the "library_report_sequence" field.
func LibraryReportSequenceLTE(v int64) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.FieldLTE(FieldLibraryReportSequence, v))
}

// HasSentinel applies the HasEdge predicate on the "sentinel" edge.
func HasSentinel() predicate.SentinelLibrary {
	return predicate.SentinelLibrary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SentinelTable, SentinelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSentinelWith applies the HasEdge predicate on the "sentinel" edge with a given conditions (other predicates).
func HasSentinelWith(preds ...predicate.Sentinel) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(func(s *sql.Selector) {
		step := newSentinelStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SentinelLibrary) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SentinelLibrary) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SentinelLibrary) predicate.SentinelLibrary {
	return predicate.SentinelLibrary(sql.NotPredicates(p))
}
