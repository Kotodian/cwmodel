// Code generated by ent, DO NOT EDIT.

package manufacturer

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the manufacturer type in the database.
	Label = "manufacturer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the manufacturer in the database.
	Table = "equip_manufacturer"
)

// Columns holds all SQL columns for manufacturer fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldCode,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int64
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy datasource.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy datasource.UUID
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
