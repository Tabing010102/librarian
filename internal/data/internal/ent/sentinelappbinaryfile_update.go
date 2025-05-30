// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinaryfile"
	"github.com/tuihub/librarian/internal/model"
)

// SentinelAppBinaryFileUpdate is the builder for updating SentinelAppBinaryFile entities.
type SentinelAppBinaryFileUpdate struct {
	config
	hooks    []Hook
	mutation *SentinelAppBinaryFileMutation
}

// Where appends a list predicates to the SentinelAppBinaryFileUpdate builder.
func (sabfu *SentinelAppBinaryFileUpdate) Where(ps ...predicate.SentinelAppBinaryFile) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.Where(ps...)
	return sabfu
}

// SetSentinelID sets the "sentinel_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSentinelID(mi model.InternalID) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ResetSentinelID()
	sabfu.mutation.SetSentinelID(mi)
	return sabfu
}

// SetNillableSentinelID sets the "sentinel_id" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSentinelID(mi *model.InternalID) *SentinelAppBinaryFileUpdate {
	if mi != nil {
		sabfu.SetSentinelID(*mi)
	}
	return sabfu
}

// AddSentinelID adds mi to the "sentinel_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) AddSentinelID(mi model.InternalID) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.AddSentinelID(mi)
	return sabfu
}

// SetSentinelLibraryReportedID sets the "sentinel_library_reported_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSentinelLibraryReportedID(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ResetSentinelLibraryReportedID()
	sabfu.mutation.SetSentinelLibraryReportedID(i)
	return sabfu
}

// SetNillableSentinelLibraryReportedID sets the "sentinel_library_reported_id" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSentinelLibraryReportedID(i *int64) *SentinelAppBinaryFileUpdate {
	if i != nil {
		sabfu.SetSentinelLibraryReportedID(*i)
	}
	return sabfu
}

// AddSentinelLibraryReportedID adds i to the "sentinel_library_reported_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) AddSentinelLibraryReportedID(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.AddSentinelLibraryReportedID(i)
	return sabfu
}

// SetLibrarySnapshot sets the "library_snapshot" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetLibrarySnapshot(t time.Time) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetLibrarySnapshot(t)
	return sabfu
}

// SetNillableLibrarySnapshot sets the "library_snapshot" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableLibrarySnapshot(t *time.Time) *SentinelAppBinaryFileUpdate {
	if t != nil {
		sabfu.SetLibrarySnapshot(*t)
	}
	return sabfu
}

// SetSentinelAppBinaryGeneratedID sets the "sentinel_app_binary_generated_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSentinelAppBinaryGeneratedID(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetSentinelAppBinaryGeneratedID(s)
	return sabfu
}

// SetNillableSentinelAppBinaryGeneratedID sets the "sentinel_app_binary_generated_id" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSentinelAppBinaryGeneratedID(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetSentinelAppBinaryGeneratedID(*s)
	}
	return sabfu
}

// SetName sets the "name" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetName(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetName(s)
	return sabfu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableName(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetName(*s)
	}
	return sabfu
}

// SetSizeBytes sets the "size_bytes" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSizeBytes(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ResetSizeBytes()
	sabfu.mutation.SetSizeBytes(i)
	return sabfu
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSizeBytes(i *int64) *SentinelAppBinaryFileUpdate {
	if i != nil {
		sabfu.SetSizeBytes(*i)
	}
	return sabfu
}

// AddSizeBytes adds i to the "size_bytes" field.
func (sabfu *SentinelAppBinaryFileUpdate) AddSizeBytes(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.AddSizeBytes(i)
	return sabfu
}

// SetSha256 sets the "sha256" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSha256(b []byte) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetSha256(b)
	return sabfu
}

// SetServerFilePath sets the "server_file_path" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetServerFilePath(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetServerFilePath(s)
	return sabfu
}

// SetNillableServerFilePath sets the "server_file_path" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableServerFilePath(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetServerFilePath(*s)
	}
	return sabfu
}

// SetChunksInfo sets the "chunks_info" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetChunksInfo(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetChunksInfo(s)
	return sabfu
}

// SetNillableChunksInfo sets the "chunks_info" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableChunksInfo(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetChunksInfo(*s)
	}
	return sabfu
}

// ClearChunksInfo clears the value of the "chunks_info" field.
func (sabfu *SentinelAppBinaryFileUpdate) ClearChunksInfo() *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ClearChunksInfo()
	return sabfu
}

// SetUpdatedAt sets the "updated_at" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetUpdatedAt(t time.Time) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetUpdatedAt(t)
	return sabfu
}

// SetCreatedAt sets the "created_at" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetCreatedAt(t time.Time) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetCreatedAt(t)
	return sabfu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableCreatedAt(t *time.Time) *SentinelAppBinaryFileUpdate {
	if t != nil {
		sabfu.SetCreatedAt(*t)
	}
	return sabfu
}

// Mutation returns the SentinelAppBinaryFileMutation object of the builder.
func (sabfu *SentinelAppBinaryFileUpdate) Mutation() *SentinelAppBinaryFileMutation {
	return sabfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sabfu *SentinelAppBinaryFileUpdate) Save(ctx context.Context) (int, error) {
	sabfu.defaults()
	return withHooks(ctx, sabfu.sqlSave, sabfu.mutation, sabfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sabfu *SentinelAppBinaryFileUpdate) SaveX(ctx context.Context) int {
	affected, err := sabfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sabfu *SentinelAppBinaryFileUpdate) Exec(ctx context.Context) error {
	_, err := sabfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabfu *SentinelAppBinaryFileUpdate) ExecX(ctx context.Context) {
	if err := sabfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sabfu *SentinelAppBinaryFileUpdate) defaults() {
	if _, ok := sabfu.mutation.UpdatedAt(); !ok {
		v := sentinelappbinaryfile.UpdateDefaultUpdatedAt()
		sabfu.mutation.SetUpdatedAt(v)
	}
}

func (sabfu *SentinelAppBinaryFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sentinelappbinaryfile.Table, sentinelappbinaryfile.Columns, sqlgraph.NewFieldSpec(sentinelappbinaryfile.FieldID, field.TypeInt64))
	if ps := sabfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sabfu.mutation.SentinelID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelID, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.AddedSentinelID(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSentinelID, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.SentinelLibraryReportedID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelLibraryReportedID, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.AddedSentinelLibraryReportedID(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSentinelLibraryReportedID, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.LibrarySnapshot(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldLibrarySnapshot, field.TypeTime, value)
	}
	if value, ok := sabfu.mutation.SentinelAppBinaryGeneratedID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelAppBinaryGeneratedID, field.TypeString, value)
	}
	if value, ok := sabfu.mutation.Name(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldName, field.TypeString, value)
	}
	if value, ok := sabfu.mutation.SizeBytes(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.AddedSizeBytes(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.Sha256(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSha256, field.TypeBytes, value)
	}
	if value, ok := sabfu.mutation.ServerFilePath(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldServerFilePath, field.TypeString, value)
	}
	if value, ok := sabfu.mutation.ChunksInfo(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString, value)
	}
	if sabfu.mutation.ChunksInfoCleared() {
		_spec.ClearField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString)
	}
	if value, ok := sabfu.mutation.UpdatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sabfu.mutation.CreatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sabfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sentinelappbinaryfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sabfu.mutation.done = true
	return n, nil
}

// SentinelAppBinaryFileUpdateOne is the builder for updating a single SentinelAppBinaryFile entity.
type SentinelAppBinaryFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SentinelAppBinaryFileMutation
}

// SetSentinelID sets the "sentinel_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSentinelID(mi model.InternalID) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ResetSentinelID()
	sabfuo.mutation.SetSentinelID(mi)
	return sabfuo
}

// SetNillableSentinelID sets the "sentinel_id" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSentinelID(mi *model.InternalID) *SentinelAppBinaryFileUpdateOne {
	if mi != nil {
		sabfuo.SetSentinelID(*mi)
	}
	return sabfuo
}

// AddSentinelID adds mi to the "sentinel_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) AddSentinelID(mi model.InternalID) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.AddSentinelID(mi)
	return sabfuo
}

// SetSentinelLibraryReportedID sets the "sentinel_library_reported_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSentinelLibraryReportedID(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ResetSentinelLibraryReportedID()
	sabfuo.mutation.SetSentinelLibraryReportedID(i)
	return sabfuo
}

// SetNillableSentinelLibraryReportedID sets the "sentinel_library_reported_id" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSentinelLibraryReportedID(i *int64) *SentinelAppBinaryFileUpdateOne {
	if i != nil {
		sabfuo.SetSentinelLibraryReportedID(*i)
	}
	return sabfuo
}

// AddSentinelLibraryReportedID adds i to the "sentinel_library_reported_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) AddSentinelLibraryReportedID(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.AddSentinelLibraryReportedID(i)
	return sabfuo
}

// SetLibrarySnapshot sets the "library_snapshot" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetLibrarySnapshot(t time.Time) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetLibrarySnapshot(t)
	return sabfuo
}

// SetNillableLibrarySnapshot sets the "library_snapshot" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableLibrarySnapshot(t *time.Time) *SentinelAppBinaryFileUpdateOne {
	if t != nil {
		sabfuo.SetLibrarySnapshot(*t)
	}
	return sabfuo
}

// SetSentinelAppBinaryGeneratedID sets the "sentinel_app_binary_generated_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSentinelAppBinaryGeneratedID(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetSentinelAppBinaryGeneratedID(s)
	return sabfuo
}

// SetNillableSentinelAppBinaryGeneratedID sets the "sentinel_app_binary_generated_id" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSentinelAppBinaryGeneratedID(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetSentinelAppBinaryGeneratedID(*s)
	}
	return sabfuo
}

// SetName sets the "name" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetName(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetName(s)
	return sabfuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableName(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetName(*s)
	}
	return sabfuo
}

// SetSizeBytes sets the "size_bytes" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSizeBytes(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ResetSizeBytes()
	sabfuo.mutation.SetSizeBytes(i)
	return sabfuo
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSizeBytes(i *int64) *SentinelAppBinaryFileUpdateOne {
	if i != nil {
		sabfuo.SetSizeBytes(*i)
	}
	return sabfuo
}

// AddSizeBytes adds i to the "size_bytes" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) AddSizeBytes(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.AddSizeBytes(i)
	return sabfuo
}

// SetSha256 sets the "sha256" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSha256(b []byte) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetSha256(b)
	return sabfuo
}

// SetServerFilePath sets the "server_file_path" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetServerFilePath(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetServerFilePath(s)
	return sabfuo
}

// SetNillableServerFilePath sets the "server_file_path" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableServerFilePath(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetServerFilePath(*s)
	}
	return sabfuo
}

// SetChunksInfo sets the "chunks_info" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetChunksInfo(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetChunksInfo(s)
	return sabfuo
}

// SetNillableChunksInfo sets the "chunks_info" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableChunksInfo(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetChunksInfo(*s)
	}
	return sabfuo
}

// ClearChunksInfo clears the value of the "chunks_info" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) ClearChunksInfo() *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ClearChunksInfo()
	return sabfuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetUpdatedAt(t time.Time) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetUpdatedAt(t)
	return sabfuo
}

// SetCreatedAt sets the "created_at" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetCreatedAt(t time.Time) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetCreatedAt(t)
	return sabfuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableCreatedAt(t *time.Time) *SentinelAppBinaryFileUpdateOne {
	if t != nil {
		sabfuo.SetCreatedAt(*t)
	}
	return sabfuo
}

// Mutation returns the SentinelAppBinaryFileMutation object of the builder.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Mutation() *SentinelAppBinaryFileMutation {
	return sabfuo.mutation
}

// Where appends a list predicates to the SentinelAppBinaryFileUpdate builder.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Where(ps ...predicate.SentinelAppBinaryFile) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.Where(ps...)
	return sabfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Select(field string, fields ...string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.fields = append([]string{field}, fields...)
	return sabfuo
}

// Save executes the query and returns the updated SentinelAppBinaryFile entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Save(ctx context.Context) (*SentinelAppBinaryFile, error) {
	sabfuo.defaults()
	return withHooks(ctx, sabfuo.sqlSave, sabfuo.mutation, sabfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SaveX(ctx context.Context) *SentinelAppBinaryFile {
	node, err := sabfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Exec(ctx context.Context) error {
	_, err := sabfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabfuo *SentinelAppBinaryFileUpdateOne) ExecX(ctx context.Context) {
	if err := sabfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sabfuo *SentinelAppBinaryFileUpdateOne) defaults() {
	if _, ok := sabfuo.mutation.UpdatedAt(); !ok {
		v := sentinelappbinaryfile.UpdateDefaultUpdatedAt()
		sabfuo.mutation.SetUpdatedAt(v)
	}
}

func (sabfuo *SentinelAppBinaryFileUpdateOne) sqlSave(ctx context.Context) (_node *SentinelAppBinaryFile, err error) {
	_spec := sqlgraph.NewUpdateSpec(sentinelappbinaryfile.Table, sentinelappbinaryfile.Columns, sqlgraph.NewFieldSpec(sentinelappbinaryfile.FieldID, field.TypeInt64))
	id, ok := sabfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SentinelAppBinaryFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sabfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sentinelappbinaryfile.FieldID)
		for _, f := range fields {
			if !sentinelappbinaryfile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sentinelappbinaryfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sabfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sabfuo.mutation.SentinelID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelID, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.AddedSentinelID(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSentinelID, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.SentinelLibraryReportedID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelLibraryReportedID, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.AddedSentinelLibraryReportedID(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSentinelLibraryReportedID, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.LibrarySnapshot(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldLibrarySnapshot, field.TypeTime, value)
	}
	if value, ok := sabfuo.mutation.SentinelAppBinaryGeneratedID(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSentinelAppBinaryGeneratedID, field.TypeString, value)
	}
	if value, ok := sabfuo.mutation.Name(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldName, field.TypeString, value)
	}
	if value, ok := sabfuo.mutation.SizeBytes(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.AddedSizeBytes(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.Sha256(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSha256, field.TypeBytes, value)
	}
	if value, ok := sabfuo.mutation.ServerFilePath(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldServerFilePath, field.TypeString, value)
	}
	if value, ok := sabfuo.mutation.ChunksInfo(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString, value)
	}
	if sabfuo.mutation.ChunksInfoCleared() {
		_spec.ClearField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString)
	}
	if value, ok := sabfuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sabfuo.mutation.CreatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &SentinelAppBinaryFile{config: sabfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sabfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sentinelappbinaryfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sabfuo.mutation.done = true
	return _node, nil
}
