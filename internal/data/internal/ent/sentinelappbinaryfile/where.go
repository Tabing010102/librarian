// Code generated by ent, DO NOT EDIT.

package sentinelappbinaryfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldID, id))
}

// SentinelID applies equality check predicate on the "sentinel_id" field. It's identical to SentinelIDEQ.
func SentinelID(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelID, vc))
}

// SentinelLibraryReportedID applies equality check predicate on the "sentinel_library_reported_id" field. It's identical to SentinelLibraryReportedIDEQ.
func SentinelLibraryReportedID(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelLibraryReportedID, v))
}

// LibrarySnapshot applies equality check predicate on the "library_snapshot" field. It's identical to LibrarySnapshotEQ.
func LibrarySnapshot(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldLibrarySnapshot, v))
}

// SentinelAppBinaryGeneratedID applies equality check predicate on the "sentinel_app_binary_generated_id" field. It's identical to SentinelAppBinaryGeneratedIDEQ.
func SentinelAppBinaryGeneratedID(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelAppBinaryGeneratedID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldName, v))
}

// SizeBytes applies equality check predicate on the "size_bytes" field. It's identical to SizeBytesEQ.
func SizeBytes(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSizeBytes, v))
}

// Sha256 applies equality check predicate on the "sha256" field. It's identical to Sha256EQ.
func Sha256(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSha256, v))
}

// ServerFilePath applies equality check predicate on the "server_file_path" field. It's identical to ServerFilePathEQ.
func ServerFilePath(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldServerFilePath, v))
}

// ChunksInfo applies equality check predicate on the "chunks_info" field. It's identical to ChunksInfoEQ.
func ChunksInfo(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldChunksInfo, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldCreatedAt, v))
}

// SentinelIDEQ applies the EQ predicate on the "sentinel_id" field.
func SentinelIDEQ(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelID, vc))
}

// SentinelIDNEQ applies the NEQ predicate on the "sentinel_id" field.
func SentinelIDNEQ(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldSentinelID, vc))
}

// SentinelIDIn applies the In predicate on the "sentinel_id" field.
func SentinelIDIn(vs ...model.InternalID) predicate.SentinelAppBinaryFile {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldSentinelID, v...))
}

// SentinelIDNotIn applies the NotIn predicate on the "sentinel_id" field.
func SentinelIDNotIn(vs ...model.InternalID) predicate.SentinelAppBinaryFile {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldSentinelID, v...))
}

// SentinelIDGT applies the GT predicate on the "sentinel_id" field.
func SentinelIDGT(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldSentinelID, vc))
}

// SentinelIDGTE applies the GTE predicate on the "sentinel_id" field.
func SentinelIDGTE(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldSentinelID, vc))
}

// SentinelIDLT applies the LT predicate on the "sentinel_id" field.
func SentinelIDLT(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldSentinelID, vc))
}

// SentinelIDLTE applies the LTE predicate on the "sentinel_id" field.
func SentinelIDLTE(v model.InternalID) predicate.SentinelAppBinaryFile {
	vc := int64(v)
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldSentinelID, vc))
}

// SentinelLibraryReportedIDEQ applies the EQ predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDEQ(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelLibraryReportedID, v))
}

// SentinelLibraryReportedIDNEQ applies the NEQ predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDNEQ(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldSentinelLibraryReportedID, v))
}

// SentinelLibraryReportedIDIn applies the In predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDIn(vs ...int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldSentinelLibraryReportedID, vs...))
}

// SentinelLibraryReportedIDNotIn applies the NotIn predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDNotIn(vs ...int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldSentinelLibraryReportedID, vs...))
}

// SentinelLibraryReportedIDGT applies the GT predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDGT(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldSentinelLibraryReportedID, v))
}

// SentinelLibraryReportedIDGTE applies the GTE predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDGTE(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldSentinelLibraryReportedID, v))
}

// SentinelLibraryReportedIDLT applies the LT predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDLT(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldSentinelLibraryReportedID, v))
}

// SentinelLibraryReportedIDLTE applies the LTE predicate on the "sentinel_library_reported_id" field.
func SentinelLibraryReportedIDLTE(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldSentinelLibraryReportedID, v))
}

// LibrarySnapshotEQ applies the EQ predicate on the "library_snapshot" field.
func LibrarySnapshotEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldLibrarySnapshot, v))
}

// LibrarySnapshotNEQ applies the NEQ predicate on the "library_snapshot" field.
func LibrarySnapshotNEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldLibrarySnapshot, v))
}

// LibrarySnapshotIn applies the In predicate on the "library_snapshot" field.
func LibrarySnapshotIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldLibrarySnapshot, vs...))
}

// LibrarySnapshotNotIn applies the NotIn predicate on the "library_snapshot" field.
func LibrarySnapshotNotIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldLibrarySnapshot, vs...))
}

// LibrarySnapshotGT applies the GT predicate on the "library_snapshot" field.
func LibrarySnapshotGT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldLibrarySnapshot, v))
}

// LibrarySnapshotGTE applies the GTE predicate on the "library_snapshot" field.
func LibrarySnapshotGTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldLibrarySnapshot, v))
}

// LibrarySnapshotLT applies the LT predicate on the "library_snapshot" field.
func LibrarySnapshotLT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldLibrarySnapshot, v))
}

// LibrarySnapshotLTE applies the LTE predicate on the "library_snapshot" field.
func LibrarySnapshotLTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldLibrarySnapshot, v))
}

// SentinelAppBinaryGeneratedIDEQ applies the EQ predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDNEQ applies the NEQ predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDNEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDIn applies the In predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldSentinelAppBinaryGeneratedID, vs...))
}

// SentinelAppBinaryGeneratedIDNotIn applies the NotIn predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDNotIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldSentinelAppBinaryGeneratedID, vs...))
}

// SentinelAppBinaryGeneratedIDGT applies the GT predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDGT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDGTE applies the GTE predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDGTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDLT applies the LT predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDLT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDLTE applies the LTE predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDLTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDContains applies the Contains predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDContains(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContains(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDHasPrefix applies the HasPrefix predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDHasPrefix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasPrefix(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDHasSuffix applies the HasSuffix predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDHasSuffix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasSuffix(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDEqualFold applies the EqualFold predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDEqualFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEqualFold(FieldSentinelAppBinaryGeneratedID, v))
}

// SentinelAppBinaryGeneratedIDContainsFold applies the ContainsFold predicate on the "sentinel_app_binary_generated_id" field.
func SentinelAppBinaryGeneratedIDContainsFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContainsFold(FieldSentinelAppBinaryGeneratedID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContainsFold(FieldName, v))
}

// SizeBytesEQ applies the EQ predicate on the "size_bytes" field.
func SizeBytesEQ(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSizeBytes, v))
}

// SizeBytesNEQ applies the NEQ predicate on the "size_bytes" field.
func SizeBytesNEQ(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldSizeBytes, v))
}

// SizeBytesIn applies the In predicate on the "size_bytes" field.
func SizeBytesIn(vs ...int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldSizeBytes, vs...))
}

// SizeBytesNotIn applies the NotIn predicate on the "size_bytes" field.
func SizeBytesNotIn(vs ...int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldSizeBytes, vs...))
}

// SizeBytesGT applies the GT predicate on the "size_bytes" field.
func SizeBytesGT(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldSizeBytes, v))
}

// SizeBytesGTE applies the GTE predicate on the "size_bytes" field.
func SizeBytesGTE(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldSizeBytes, v))
}

// SizeBytesLT applies the LT predicate on the "size_bytes" field.
func SizeBytesLT(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldSizeBytes, v))
}

// SizeBytesLTE applies the LTE predicate on the "size_bytes" field.
func SizeBytesLTE(v int64) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldSizeBytes, v))
}

// Sha256EQ applies the EQ predicate on the "sha256" field.
func Sha256EQ(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldSha256, v))
}

// Sha256NEQ applies the NEQ predicate on the "sha256" field.
func Sha256NEQ(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldSha256, v))
}

// Sha256In applies the In predicate on the "sha256" field.
func Sha256In(vs ...[]byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldSha256, vs...))
}

// Sha256NotIn applies the NotIn predicate on the "sha256" field.
func Sha256NotIn(vs ...[]byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldSha256, vs...))
}

// Sha256GT applies the GT predicate on the "sha256" field.
func Sha256GT(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldSha256, v))
}

// Sha256GTE applies the GTE predicate on the "sha256" field.
func Sha256GTE(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldSha256, v))
}

// Sha256LT applies the LT predicate on the "sha256" field.
func Sha256LT(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldSha256, v))
}

// Sha256LTE applies the LTE predicate on the "sha256" field.
func Sha256LTE(v []byte) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldSha256, v))
}

// ServerFilePathEQ applies the EQ predicate on the "server_file_path" field.
func ServerFilePathEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldServerFilePath, v))
}

// ServerFilePathNEQ applies the NEQ predicate on the "server_file_path" field.
func ServerFilePathNEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldServerFilePath, v))
}

// ServerFilePathIn applies the In predicate on the "server_file_path" field.
func ServerFilePathIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldServerFilePath, vs...))
}

// ServerFilePathNotIn applies the NotIn predicate on the "server_file_path" field.
func ServerFilePathNotIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldServerFilePath, vs...))
}

// ServerFilePathGT applies the GT predicate on the "server_file_path" field.
func ServerFilePathGT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldServerFilePath, v))
}

// ServerFilePathGTE applies the GTE predicate on the "server_file_path" field.
func ServerFilePathGTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldServerFilePath, v))
}

// ServerFilePathLT applies the LT predicate on the "server_file_path" field.
func ServerFilePathLT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldServerFilePath, v))
}

// ServerFilePathLTE applies the LTE predicate on the "server_file_path" field.
func ServerFilePathLTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldServerFilePath, v))
}

// ServerFilePathContains applies the Contains predicate on the "server_file_path" field.
func ServerFilePathContains(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContains(FieldServerFilePath, v))
}

// ServerFilePathHasPrefix applies the HasPrefix predicate on the "server_file_path" field.
func ServerFilePathHasPrefix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasPrefix(FieldServerFilePath, v))
}

// ServerFilePathHasSuffix applies the HasSuffix predicate on the "server_file_path" field.
func ServerFilePathHasSuffix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasSuffix(FieldServerFilePath, v))
}

// ServerFilePathEqualFold applies the EqualFold predicate on the "server_file_path" field.
func ServerFilePathEqualFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEqualFold(FieldServerFilePath, v))
}

// ServerFilePathContainsFold applies the ContainsFold predicate on the "server_file_path" field.
func ServerFilePathContainsFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContainsFold(FieldServerFilePath, v))
}

// ChunksInfoEQ applies the EQ predicate on the "chunks_info" field.
func ChunksInfoEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldChunksInfo, v))
}

// ChunksInfoNEQ applies the NEQ predicate on the "chunks_info" field.
func ChunksInfoNEQ(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldChunksInfo, v))
}

// ChunksInfoIn applies the In predicate on the "chunks_info" field.
func ChunksInfoIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldChunksInfo, vs...))
}

// ChunksInfoNotIn applies the NotIn predicate on the "chunks_info" field.
func ChunksInfoNotIn(vs ...string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldChunksInfo, vs...))
}

// ChunksInfoGT applies the GT predicate on the "chunks_info" field.
func ChunksInfoGT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldChunksInfo, v))
}

// ChunksInfoGTE applies the GTE predicate on the "chunks_info" field.
func ChunksInfoGTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldChunksInfo, v))
}

// ChunksInfoLT applies the LT predicate on the "chunks_info" field.
func ChunksInfoLT(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldChunksInfo, v))
}

// ChunksInfoLTE applies the LTE predicate on the "chunks_info" field.
func ChunksInfoLTE(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldChunksInfo, v))
}

// ChunksInfoContains applies the Contains predicate on the "chunks_info" field.
func ChunksInfoContains(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContains(FieldChunksInfo, v))
}

// ChunksInfoHasPrefix applies the HasPrefix predicate on the "chunks_info" field.
func ChunksInfoHasPrefix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasPrefix(FieldChunksInfo, v))
}

// ChunksInfoHasSuffix applies the HasSuffix predicate on the "chunks_info" field.
func ChunksInfoHasSuffix(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldHasSuffix(FieldChunksInfo, v))
}

// ChunksInfoIsNil applies the IsNil predicate on the "chunks_info" field.
func ChunksInfoIsNil() predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIsNull(FieldChunksInfo))
}

// ChunksInfoNotNil applies the NotNil predicate on the "chunks_info" field.
func ChunksInfoNotNil() predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotNull(FieldChunksInfo))
}

// ChunksInfoEqualFold applies the EqualFold predicate on the "chunks_info" field.
func ChunksInfoEqualFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEqualFold(FieldChunksInfo, v))
}

// ChunksInfoContainsFold applies the ContainsFold predicate on the "chunks_info" field.
func ChunksInfoContainsFold(v string) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldContainsFold(FieldChunksInfo, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SentinelAppBinaryFile) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SentinelAppBinaryFile) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SentinelAppBinaryFile) predicate.SentinelAppBinaryFile {
	return predicate.SentinelAppBinaryFile(sql.NotPredicates(p))
}
