// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Kotodian/ent-practice/ent/appmoduleinfo"
)

// AppModuleInfo is the model entity for the AppModuleInfo schema.
type AppModuleInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppModuleInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appmoduleinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case appmoduleinfo.FieldName, appmoduleinfo.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppModuleInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppModuleInfo fields.
func (ami *AppModuleInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appmoduleinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ami.ID = int(value.Int64)
		case appmoduleinfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ami.Name = value.String
			}
		case appmoduleinfo.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				ami.Desc = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppModuleInfo.
// Note that you need to call AppModuleInfo.Unwrap() before calling this method if this AppModuleInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ami *AppModuleInfo) Update() *AppModuleInfoUpdateOne {
	return (&AppModuleInfoClient{config: ami.config}).UpdateOne(ami)
}

// Unwrap unwraps the AppModuleInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ami *AppModuleInfo) Unwrap() *AppModuleInfo {
	_tx, ok := ami.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppModuleInfo is not a transactional entity")
	}
	ami.config.driver = _tx.drv
	return ami
}

// String implements the fmt.Stringer.
func (ami *AppModuleInfo) String() string {
	var builder strings.Builder
	builder.WriteString("AppModuleInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ami.ID))
	builder.WriteString("name=")
	builder.WriteString(ami.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(ami.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// AppModuleInfos is a parsable slice of AppModuleInfo.
type AppModuleInfos []*AppModuleInfo

func (ami AppModuleInfos) config(cfg config) {
	for _i := range ami {
		ami[_i].config = cfg
	}
}