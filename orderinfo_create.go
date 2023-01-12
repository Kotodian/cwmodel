// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/orderevent"
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/gokit/datasource"
)

// OrderInfoCreate is the builder for creating a OrderInfo entity.
type OrderInfoCreate struct {
	config
	mutation *OrderInfoMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (oic *OrderInfoCreate) SetVersion(i int64) *OrderInfoCreate {
	oic.mutation.SetVersion(i)
	return oic
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableVersion(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetVersion(*i)
	}
	return oic
}

// SetCreatedBy sets the "created_by" field.
func (oic *OrderInfoCreate) SetCreatedBy(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetCreatedBy(d)
	return oic
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrderInfoCreate) SetCreatedAt(i int64) *OrderInfoCreate {
	oic.mutation.SetCreatedAt(i)
	return oic
}

// SetUpdatedBy sets the "updated_by" field.
func (oic *OrderInfoCreate) SetUpdatedBy(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetUpdatedBy(d)
	return oic
}

// SetUpdatedAt sets the "updated_at" field.
func (oic *OrderInfoCreate) SetUpdatedAt(i int64) *OrderInfoCreate {
	oic.mutation.SetUpdatedAt(i)
	return oic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableUpdatedAt(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetUpdatedAt(*i)
	}
	return oic
}

// SetEquipmentID sets the "equipment_id" field.
func (oic *OrderInfoCreate) SetEquipmentID(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetEquipmentID(d)
	return oic
}

// SetConnectorID sets the "connector_id" field.
func (oic *OrderInfoCreate) SetConnectorID(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetConnectorID(d)
	return oic
}

// SetRemoteStartID sets the "remote_start_id" field.
func (oic *OrderInfoCreate) SetRemoteStartID(i int64) *OrderInfoCreate {
	oic.mutation.SetRemoteStartID(i)
	return oic
}

// SetNillableRemoteStartID sets the "remote_start_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableRemoteStartID(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetRemoteStartID(*i)
	}
	return oic
}

// SetTransactionID sets the "transaction_id" field.
func (oic *OrderInfoCreate) SetTransactionID(s string) *OrderInfoCreate {
	oic.mutation.SetTransactionID(s)
	return oic
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableTransactionID(s *string) *OrderInfoCreate {
	if s != nil {
		oic.SetTransactionID(*s)
	}
	return oic
}

// SetAuthorizationID sets the "authorization_id" field.
func (oic *OrderInfoCreate) SetAuthorizationID(s string) *OrderInfoCreate {
	oic.mutation.SetAuthorizationID(s)
	return oic
}

// SetNillableAuthorizationID sets the "authorization_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableAuthorizationID(s *string) *OrderInfoCreate {
	if s != nil {
		oic.SetAuthorizationID(*s)
	}
	return oic
}

// SetAuthorizationMode sets the "authorization_mode" field.
func (oic *OrderInfoCreate) SetAuthorizationMode(i int) *OrderInfoCreate {
	oic.mutation.SetAuthorizationMode(i)
	return oic
}

// SetNillableAuthorizationMode sets the "authorization_mode" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableAuthorizationMode(i *int) *OrderInfoCreate {
	if i != nil {
		oic.SetAuthorizationMode(*i)
	}
	return oic
}

// SetCustomerID sets the "customer_id" field.
func (oic *OrderInfoCreate) SetCustomerID(s string) *OrderInfoCreate {
	oic.mutation.SetCustomerID(s)
	return oic
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableCustomerID(s *string) *OrderInfoCreate {
	if s != nil {
		oic.SetCustomerID(*s)
	}
	return oic
}

// SetCallerOrderID sets the "caller_order_id" field.
func (oic *OrderInfoCreate) SetCallerOrderID(s string) *OrderInfoCreate {
	oic.mutation.SetCallerOrderID(s)
	return oic
}

// SetNillableCallerOrderID sets the "caller_order_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableCallerOrderID(s *string) *OrderInfoCreate {
	if s != nil {
		oic.SetCallerOrderID(*s)
	}
	return oic
}

// SetTotalElectricity sets the "total_electricity" field.
func (oic *OrderInfoCreate) SetTotalElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetTotalElectricity(f)
	return oic
}

// SetNillableTotalElectricity sets the "total_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableTotalElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetTotalElectricity(*f)
	}
	return oic
}

// SetChargeStartElectricity sets the "charge_start_electricity" field.
func (oic *OrderInfoCreate) SetChargeStartElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetChargeStartElectricity(f)
	return oic
}

// SetNillableChargeStartElectricity sets the "charge_start_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableChargeStartElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetChargeStartElectricity(*f)
	}
	return oic
}

// SetChargeFinalElectricity sets the "charge_final_electricity" field.
func (oic *OrderInfoCreate) SetChargeFinalElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetChargeFinalElectricity(f)
	return oic
}

// SetNillableChargeFinalElectricity sets the "charge_final_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableChargeFinalElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetChargeFinalElectricity(*f)
	}
	return oic
}

// SetSharpElectricity sets the "sharp_electricity" field.
func (oic *OrderInfoCreate) SetSharpElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetSharpElectricity(f)
	return oic
}

// SetNillableSharpElectricity sets the "sharp_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableSharpElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetSharpElectricity(*f)
	}
	return oic
}

// SetPeakElectricity sets the "peak_electricity" field.
func (oic *OrderInfoCreate) SetPeakElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetPeakElectricity(f)
	return oic
}

// SetNillablePeakElectricity sets the "peak_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillablePeakElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetPeakElectricity(*f)
	}
	return oic
}

// SetFlatElectricity sets the "flat_electricity" field.
func (oic *OrderInfoCreate) SetFlatElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetFlatElectricity(f)
	return oic
}

// SetNillableFlatElectricity sets the "flat_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableFlatElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetFlatElectricity(*f)
	}
	return oic
}

// SetValleyElectricity sets the "valley_electricity" field.
func (oic *OrderInfoCreate) SetValleyElectricity(f float64) *OrderInfoCreate {
	oic.mutation.SetValleyElectricity(f)
	return oic
}

// SetNillableValleyElectricity sets the "valley_electricity" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableValleyElectricity(f *float64) *OrderInfoCreate {
	if f != nil {
		oic.SetValleyElectricity(*f)
	}
	return oic
}

// SetStopReasonCode sets the "stop_reason_code" field.
func (oic *OrderInfoCreate) SetStopReasonCode(i int32) *OrderInfoCreate {
	oic.mutation.SetStopReasonCode(i)
	return oic
}

// SetNillableStopReasonCode sets the "stop_reason_code" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableStopReasonCode(i *int32) *OrderInfoCreate {
	if i != nil {
		oic.SetStopReasonCode(*i)
	}
	return oic
}

// SetVIN sets the "VIN" field.
func (oic *OrderInfoCreate) SetVIN(s string) *OrderInfoCreate {
	oic.mutation.SetVIN(s)
	return oic
}

// SetNillableVIN sets the "VIN" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableVIN(s *string) *OrderInfoCreate {
	if s != nil {
		oic.SetVIN(*s)
	}
	return oic
}

// SetState sets the "state" field.
func (oic *OrderInfoCreate) SetState(i int) *OrderInfoCreate {
	oic.mutation.SetState(i)
	return oic
}

// SetNillableState sets the "state" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableState(i *int) *OrderInfoCreate {
	if i != nil {
		oic.SetState(*i)
	}
	return oic
}

// SetOffline sets the "offline" field.
func (oic *OrderInfoCreate) SetOffline(b bool) *OrderInfoCreate {
	oic.mutation.SetOffline(b)
	return oic
}

// SetPriceSchemeReleaseID sets the "price_scheme_release_id" field.
func (oic *OrderInfoCreate) SetPriceSchemeReleaseID(i int64) *OrderInfoCreate {
	oic.mutation.SetPriceSchemeReleaseID(i)
	return oic
}

// SetNillablePriceSchemeReleaseID sets the "price_scheme_release_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillablePriceSchemeReleaseID(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetPriceSchemeReleaseID(*i)
	}
	return oic
}

// SetOrderStartTime sets the "order_start_time" field.
func (oic *OrderInfoCreate) SetOrderStartTime(i int64) *OrderInfoCreate {
	oic.mutation.SetOrderStartTime(i)
	return oic
}

// SetNillableOrderStartTime sets the "order_start_time" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableOrderStartTime(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetOrderStartTime(*i)
	}
	return oic
}

// SetOrderFinalTime sets the "order_final_time" field.
func (oic *OrderInfoCreate) SetOrderFinalTime(i int64) *OrderInfoCreate {
	oic.mutation.SetOrderFinalTime(i)
	return oic
}

// SetNillableOrderFinalTime sets the "order_final_time" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableOrderFinalTime(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetOrderFinalTime(*i)
	}
	return oic
}

// SetChargeStartTime sets the "charge_start_time" field.
func (oic *OrderInfoCreate) SetChargeStartTime(i int64) *OrderInfoCreate {
	oic.mutation.SetChargeStartTime(i)
	return oic
}

// SetNillableChargeStartTime sets the "charge_start_time" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableChargeStartTime(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetChargeStartTime(*i)
	}
	return oic
}

// SetChargeFinalTime sets the "charge_final_time" field.
func (oic *OrderInfoCreate) SetChargeFinalTime(i int64) *OrderInfoCreate {
	oic.mutation.SetChargeFinalTime(i)
	return oic
}

// SetNillableChargeFinalTime sets the "charge_final_time" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableChargeFinalTime(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetChargeFinalTime(*i)
	}
	return oic
}

// SetIntellectID sets the "intellect_id" field.
func (oic *OrderInfoCreate) SetIntellectID(i int64) *OrderInfoCreate {
	oic.mutation.SetIntellectID(i)
	return oic
}

// SetNillableIntellectID sets the "intellect_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableIntellectID(i *int64) *OrderInfoCreate {
	if i != nil {
		oic.SetIntellectID(*i)
	}
	return oic
}

// SetStationID sets the "station_id" field.
func (oic *OrderInfoCreate) SetStationID(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetStationID(d)
	return oic
}

// SetNillableStationID sets the "station_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableStationID(d *datasource.UUID) *OrderInfoCreate {
	if d != nil {
		oic.SetStationID(*d)
	}
	return oic
}

// SetOperatorID sets the "operator_id" field.
func (oic *OrderInfoCreate) SetOperatorID(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetOperatorID(d)
	return oic
}

// SetNillableOperatorID sets the "operator_id" field if the given value is not nil.
func (oic *OrderInfoCreate) SetNillableOperatorID(d *datasource.UUID) *OrderInfoCreate {
	if d != nil {
		oic.SetOperatorID(*d)
	}
	return oic
}

// SetID sets the "id" field.
func (oic *OrderInfoCreate) SetID(d datasource.UUID) *OrderInfoCreate {
	oic.mutation.SetID(d)
	return oic
}

// SetConnector sets the "connector" edge to the Connector entity.
func (oic *OrderInfoCreate) SetConnector(c *Connector) *OrderInfoCreate {
	return oic.SetConnectorID(c.ID)
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (oic *OrderInfoCreate) SetEquipment(e *Equipment) *OrderInfoCreate {
	return oic.SetEquipmentID(e.ID)
}

// AddOrderEventIDs adds the "order_event" edge to the OrderEvent entity by IDs.
func (oic *OrderInfoCreate) AddOrderEventIDs(ids ...datasource.UUID) *OrderInfoCreate {
	oic.mutation.AddOrderEventIDs(ids...)
	return oic
}

// AddOrderEvent adds the "order_event" edges to the OrderEvent entity.
func (oic *OrderInfoCreate) AddOrderEvent(o ...*OrderEvent) *OrderInfoCreate {
	ids := make([]datasource.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oic.AddOrderEventIDs(ids...)
}

// Mutation returns the OrderInfoMutation object of the builder.
func (oic *OrderInfoCreate) Mutation() *OrderInfoMutation {
	return oic.mutation
}

// Save creates the OrderInfo in the database.
func (oic *OrderInfoCreate) Save(ctx context.Context) (*OrderInfo, error) {
	var (
		err  error
		node *OrderInfo
	)
	oic.defaults()
	if len(oic.hooks) == 0 {
		if err = oic.check(); err != nil {
			return nil, err
		}
		node, err = oic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oic.check(); err != nil {
				return nil, err
			}
			oic.mutation = mutation
			if node, err = oic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oic.hooks) - 1; i >= 0; i-- {
			if oic.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = oic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderInfoCreate) SaveX(ctx context.Context) *OrderInfo {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderInfoCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderInfoCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderInfoCreate) defaults() {
	if _, ok := oic.mutation.Version(); !ok {
		v := orderinfo.DefaultVersion
		oic.mutation.SetVersion(v)
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		v := orderinfo.DefaultUpdatedAt
		oic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderInfoCreate) check() error {
	if _, ok := oic.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`cwmodel: missing required field "OrderInfo.version"`)}
	}
	if _, ok := oic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cwmodel: missing required field "OrderInfo.created_by"`)}
	}
	if _, ok := oic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cwmodel: missing required field "OrderInfo.created_at"`)}
	}
	if _, ok := oic.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cwmodel: missing required field "OrderInfo.updated_by"`)}
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cwmodel: missing required field "OrderInfo.updated_at"`)}
	}
	if _, ok := oic.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment_id", err: errors.New(`cwmodel: missing required field "OrderInfo.equipment_id"`)}
	}
	if _, ok := oic.mutation.ConnectorID(); !ok {
		return &ValidationError{Name: "connector_id", err: errors.New(`cwmodel: missing required field "OrderInfo.connector_id"`)}
	}
	if _, ok := oic.mutation.Offline(); !ok {
		return &ValidationError{Name: "offline", err: errors.New(`cwmodel: missing required field "OrderInfo.offline"`)}
	}
	if _, ok := oic.mutation.ConnectorID(); !ok {
		return &ValidationError{Name: "connector", err: errors.New(`cwmodel: missing required edge "OrderInfo.connector"`)}
	}
	if _, ok := oic.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`cwmodel: missing required edge "OrderInfo.equipment"`)}
	}
	return nil
}

func (oic *OrderInfoCreate) sqlSave(ctx context.Context) (*OrderInfo, error) {
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = datasource.UUID(id)
	}
	return _node, nil
}

func (oic *OrderInfoCreate) createSpec() (*OrderInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderInfo{config: oic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderinfo.FieldID,
			},
		}
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oic.mutation.Version(); ok {
		_spec.SetField(orderinfo.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := oic.mutation.CreatedBy(); ok {
		_spec.SetField(orderinfo.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.SetField(orderinfo.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := oic.mutation.UpdatedBy(); ok {
		_spec.SetField(orderinfo.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := oic.mutation.UpdatedAt(); ok {
		_spec.SetField(orderinfo.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := oic.mutation.RemoteStartID(); ok {
		_spec.SetField(orderinfo.FieldRemoteStartID, field.TypeInt64, value)
		_node.RemoteStartID = &value
	}
	if value, ok := oic.mutation.TransactionID(); ok {
		_spec.SetField(orderinfo.FieldTransactionID, field.TypeString, value)
		_node.TransactionID = &value
	}
	if value, ok := oic.mutation.AuthorizationID(); ok {
		_spec.SetField(orderinfo.FieldAuthorizationID, field.TypeString, value)
		_node.AuthorizationID = &value
	}
	if value, ok := oic.mutation.AuthorizationMode(); ok {
		_spec.SetField(orderinfo.FieldAuthorizationMode, field.TypeInt, value)
		_node.AuthorizationMode = &value
	}
	if value, ok := oic.mutation.CustomerID(); ok {
		_spec.SetField(orderinfo.FieldCustomerID, field.TypeString, value)
		_node.CustomerID = &value
	}
	if value, ok := oic.mutation.CallerOrderID(); ok {
		_spec.SetField(orderinfo.FieldCallerOrderID, field.TypeString, value)
		_node.CallerOrderID = &value
	}
	if value, ok := oic.mutation.TotalElectricity(); ok {
		_spec.SetField(orderinfo.FieldTotalElectricity, field.TypeFloat64, value)
		_node.TotalElectricity = &value
	}
	if value, ok := oic.mutation.ChargeStartElectricity(); ok {
		_spec.SetField(orderinfo.FieldChargeStartElectricity, field.TypeFloat64, value)
		_node.ChargeStartElectricity = &value
	}
	if value, ok := oic.mutation.ChargeFinalElectricity(); ok {
		_spec.SetField(orderinfo.FieldChargeFinalElectricity, field.TypeFloat64, value)
		_node.ChargeFinalElectricity = &value
	}
	if value, ok := oic.mutation.SharpElectricity(); ok {
		_spec.SetField(orderinfo.FieldSharpElectricity, field.TypeFloat64, value)
		_node.SharpElectricity = &value
	}
	if value, ok := oic.mutation.PeakElectricity(); ok {
		_spec.SetField(orderinfo.FieldPeakElectricity, field.TypeFloat64, value)
		_node.PeakElectricity = &value
	}
	if value, ok := oic.mutation.FlatElectricity(); ok {
		_spec.SetField(orderinfo.FieldFlatElectricity, field.TypeFloat64, value)
		_node.FlatElectricity = &value
	}
	if value, ok := oic.mutation.ValleyElectricity(); ok {
		_spec.SetField(orderinfo.FieldValleyElectricity, field.TypeFloat64, value)
		_node.ValleyElectricity = &value
	}
	if value, ok := oic.mutation.StopReasonCode(); ok {
		_spec.SetField(orderinfo.FieldStopReasonCode, field.TypeInt32, value)
		_node.StopReasonCode = &value
	}
	if value, ok := oic.mutation.VIN(); ok {
		_spec.SetField(orderinfo.FieldVIN, field.TypeString, value)
		_node.VIN = &value
	}
	if value, ok := oic.mutation.State(); ok {
		_spec.SetField(orderinfo.FieldState, field.TypeInt, value)
		_node.State = value
	}
	if value, ok := oic.mutation.Offline(); ok {
		_spec.SetField(orderinfo.FieldOffline, field.TypeBool, value)
		_node.Offline = value
	}
	if value, ok := oic.mutation.PriceSchemeReleaseID(); ok {
		_spec.SetField(orderinfo.FieldPriceSchemeReleaseID, field.TypeInt64, value)
		_node.PriceSchemeReleaseID = &value
	}
	if value, ok := oic.mutation.OrderStartTime(); ok {
		_spec.SetField(orderinfo.FieldOrderStartTime, field.TypeInt64, value)
		_node.OrderStartTime = &value
	}
	if value, ok := oic.mutation.OrderFinalTime(); ok {
		_spec.SetField(orderinfo.FieldOrderFinalTime, field.TypeInt64, value)
		_node.OrderFinalTime = &value
	}
	if value, ok := oic.mutation.ChargeStartTime(); ok {
		_spec.SetField(orderinfo.FieldChargeStartTime, field.TypeInt64, value)
		_node.ChargeStartTime = &value
	}
	if value, ok := oic.mutation.ChargeFinalTime(); ok {
		_spec.SetField(orderinfo.FieldChargeFinalTime, field.TypeInt64, value)
		_node.ChargeFinalTime = &value
	}
	if value, ok := oic.mutation.IntellectID(); ok {
		_spec.SetField(orderinfo.FieldIntellectID, field.TypeInt64, value)
		_node.IntellectID = &value
	}
	if value, ok := oic.mutation.StationID(); ok {
		_spec.SetField(orderinfo.FieldStationID, field.TypeUint64, value)
		_node.StationID = &value
	}
	if value, ok := oic.mutation.OperatorID(); ok {
		_spec.SetField(orderinfo.FieldOperatorID, field.TypeUint64, value)
		_node.OperatorID = &value
	}
	if nodes := oic.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderinfo.ConnectorTable,
			Columns: []string{orderinfo.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ConnectorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderinfo.EquipmentTable,
			Columns: []string{orderinfo.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EquipmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.OrderEventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderinfo.OrderEventTable,
			Columns: []string{orderinfo.OrderEventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderevent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderInfoCreateBulk is the builder for creating many OrderInfo entities in bulk.
type OrderInfoCreateBulk struct {
	config
	builders []*OrderInfoCreate
}

// Save creates the OrderInfo entities in the database.
func (oicb *OrderInfoCreateBulk) Save(ctx context.Context) ([]*OrderInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderInfo, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = datasource.UUID(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderInfoCreateBulk) SaveX(ctx context.Context) []*OrderInfo {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderInfoCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
