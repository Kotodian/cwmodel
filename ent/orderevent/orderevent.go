// Code generated by ent, DO NOT EDIT.

package orderevent

import (
	"github.com/Kotodian/gokit/datasource"
)

const (
	// Label holds the string label denoting the orderevent type in the database.
	Label = "order_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldOccurrence holds the string denoting the occurrence field in the database.
	FieldOccurrence = "occurrence"
	// EdgeOrderInfo holds the string denoting the order_info edge name in mutations.
	EdgeOrderInfo = "order_info"
	// Table holds the table name of the orderevent in the database.
	Table = "order_event"
	// OrderInfoTable is the table that holds the order_info relation/edge.
	OrderInfoTable = "order_event"
	// OrderInfoInverseTable is the table name for the OrderInfo entity.
	// It exists in this package in order to avoid circular dependency with the "orderinfo" package.
	OrderInfoInverseTable = "order_info"
	// OrderInfoColumn is the table column denoting the order_info relation/edge.
	OrderInfoColumn = "order_id"
)

// Columns holds all SQL columns for orderevent fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldOccurrence,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_event"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID datasource.UUID
)
