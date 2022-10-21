// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"github.com/Kotodian/cwmodel/appmoduleinfo"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/equipmentalarm"
	"github.com/Kotodian/cwmodel/equipmentfirmwareeffect"
	"github.com/Kotodian/cwmodel/equipmentinfo"
	"github.com/Kotodian/cwmodel/equipmentiot"
	"github.com/Kotodian/cwmodel/equipmentlog"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/cwmodel/firmware"
	"github.com/Kotodian/cwmodel/manufacturer"
	"github.com/Kotodian/cwmodel/model"
	"github.com/Kotodian/cwmodel/orderevent"
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/cwmodel/reservation"
	"github.com/Kotodian/cwmodel/schema"
	"github.com/Kotodian/cwmodel/smartchargingeffect"
	"github.com/Kotodian/gokit/datasource"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appmoduleinfoFields := schema.AppModuleInfo{}.Fields()
	_ = appmoduleinfoFields
	// appmoduleinfoDescID is the schema descriptor for id field.
	appmoduleinfoDescID := appmoduleinfoFields[0].Descriptor()
	// appmoduleinfo.DefaultID holds the default value on creation for the id field.
	appmoduleinfo.DefaultID = datasource.UUID(appmoduleinfoDescID.Default.(uint64))
	connectorMixin := schema.Connector{}.Mixin()
	connectorMixinFields0 := connectorMixin[0].Fields()
	_ = connectorMixinFields0
	connectorFields := schema.Connector{}.Fields()
	_ = connectorFields
	// connectorDescVersion is the schema descriptor for version field.
	connectorDescVersion := connectorMixinFields0[1].Descriptor()
	// connector.DefaultVersion holds the default value on creation for the version field.
	connector.DefaultVersion = connectorDescVersion.Default.(int64)
	// connectorDescCreatedBy is the schema descriptor for created_by field.
	connectorDescCreatedBy := connectorMixinFields0[2].Descriptor()
	// connector.DefaultCreatedBy holds the default value on creation for the created_by field.
	connector.DefaultCreatedBy = datasource.UUID(connectorDescCreatedBy.Default.(uint64))
	// connectorDescCreatedAt is the schema descriptor for created_at field.
	connectorDescCreatedAt := connectorMixinFields0[3].Descriptor()
	// connector.DefaultCreatedAt holds the default value on creation for the created_at field.
	connector.DefaultCreatedAt = connectorDescCreatedAt.Default.(int64)
	// connectorDescUpdatedBy is the schema descriptor for updated_by field.
	connectorDescUpdatedBy := connectorMixinFields0[4].Descriptor()
	// connector.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	connector.DefaultUpdatedBy = datasource.UUID(connectorDescUpdatedBy.Default.(uint64))
	// connectorDescUpdatedAt is the schema descriptor for updated_at field.
	connectorDescUpdatedAt := connectorMixinFields0[5].Descriptor()
	// connector.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	connector.DefaultUpdatedAt = connectorDescUpdatedAt.Default.(int64)
	// connector.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	connector.UpdateDefaultUpdatedAt = connectorDescUpdatedAt.UpdateDefault.(func() int64)
	// connectorDescParkNo is the schema descriptor for park_no field.
	connectorDescParkNo := connectorFields[10].Descriptor()
	// connector.DefaultParkNo holds the default value on creation for the park_no field.
	connector.DefaultParkNo = connectorDescParkNo.Default.(string)
	// connectorDescID is the schema descriptor for id field.
	connectorDescID := connectorMixinFields0[0].Descriptor()
	// connector.DefaultID holds the default value on creation for the id field.
	connector.DefaultID = datasource.UUID(connectorDescID.Default.(uint64))
	equipmentMixin := schema.Equipment{}.Mixin()
	equipmentMixinFields0 := equipmentMixin[0].Fields()
	_ = equipmentMixinFields0
	equipmentFields := schema.Equipment{}.Fields()
	_ = equipmentFields
	// equipmentDescVersion is the schema descriptor for version field.
	equipmentDescVersion := equipmentMixinFields0[1].Descriptor()
	// equipment.DefaultVersion holds the default value on creation for the version field.
	equipment.DefaultVersion = equipmentDescVersion.Default.(int64)
	// equipmentDescCreatedBy is the schema descriptor for created_by field.
	equipmentDescCreatedBy := equipmentMixinFields0[2].Descriptor()
	// equipment.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipment.DefaultCreatedBy = datasource.UUID(equipmentDescCreatedBy.Default.(uint64))
	// equipmentDescCreatedAt is the schema descriptor for created_at field.
	equipmentDescCreatedAt := equipmentMixinFields0[3].Descriptor()
	// equipment.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipment.DefaultCreatedAt = equipmentDescCreatedAt.Default.(int64)
	// equipmentDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentDescUpdatedBy := equipmentMixinFields0[4].Descriptor()
	// equipment.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipment.DefaultUpdatedBy = datasource.UUID(equipmentDescUpdatedBy.Default.(uint64))
	// equipmentDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentDescUpdatedAt := equipmentMixinFields0[5].Descriptor()
	// equipment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipment.DefaultUpdatedAt = equipmentDescUpdatedAt.Default.(int64)
	// equipment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipment.UpdateDefaultUpdatedAt = equipmentDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentDescSn is the schema descriptor for sn field.
	equipmentDescSn := equipmentFields[0].Descriptor()
	// equipment.SnValidator is a validator for the "sn" field. It is called by the builders before save.
	equipment.SnValidator = equipmentDescSn.Validators[0].(func(string) error)
	// equipmentDescID is the schema descriptor for id field.
	equipmentDescID := equipmentMixinFields0[0].Descriptor()
	// equipment.DefaultID holds the default value on creation for the id field.
	equipment.DefaultID = datasource.UUID(equipmentDescID.Default.(uint64))
	equipmentalarmMixin := schema.EquipmentAlarm{}.Mixin()
	equipmentalarmMixinFields0 := equipmentalarmMixin[0].Fields()
	_ = equipmentalarmMixinFields0
	equipmentalarmFields := schema.EquipmentAlarm{}.Fields()
	_ = equipmentalarmFields
	// equipmentalarmDescVersion is the schema descriptor for version field.
	equipmentalarmDescVersion := equipmentalarmMixinFields0[1].Descriptor()
	// equipmentalarm.DefaultVersion holds the default value on creation for the version field.
	equipmentalarm.DefaultVersion = equipmentalarmDescVersion.Default.(int64)
	// equipmentalarmDescCreatedBy is the schema descriptor for created_by field.
	equipmentalarmDescCreatedBy := equipmentalarmMixinFields0[2].Descriptor()
	// equipmentalarm.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipmentalarm.DefaultCreatedBy = datasource.UUID(equipmentalarmDescCreatedBy.Default.(uint64))
	// equipmentalarmDescCreatedAt is the schema descriptor for created_at field.
	equipmentalarmDescCreatedAt := equipmentalarmMixinFields0[3].Descriptor()
	// equipmentalarm.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipmentalarm.DefaultCreatedAt = equipmentalarmDescCreatedAt.Default.(int64)
	// equipmentalarmDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentalarmDescUpdatedBy := equipmentalarmMixinFields0[4].Descriptor()
	// equipmentalarm.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipmentalarm.DefaultUpdatedBy = datasource.UUID(equipmentalarmDescUpdatedBy.Default.(uint64))
	// equipmentalarmDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentalarmDescUpdatedAt := equipmentalarmMixinFields0[5].Descriptor()
	// equipmentalarm.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipmentalarm.DefaultUpdatedAt = equipmentalarmDescUpdatedAt.Default.(int64)
	// equipmentalarm.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipmentalarm.UpdateDefaultUpdatedAt = equipmentalarmDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentalarmDescCount is the schema descriptor for count field.
	equipmentalarmDescCount := equipmentalarmFields[5].Descriptor()
	// equipmentalarm.DefaultCount holds the default value on creation for the count field.
	equipmentalarm.DefaultCount = equipmentalarmDescCount.Default.(int)
	// equipmentalarmDescID is the schema descriptor for id field.
	equipmentalarmDescID := equipmentalarmMixinFields0[0].Descriptor()
	// equipmentalarm.DefaultID holds the default value on creation for the id field.
	equipmentalarm.DefaultID = datasource.UUID(equipmentalarmDescID.Default.(uint64))
	equipmentfirmwareeffectMixin := schema.EquipmentFirmwareEffect{}.Mixin()
	equipmentfirmwareeffectMixinFields0 := equipmentfirmwareeffectMixin[0].Fields()
	_ = equipmentfirmwareeffectMixinFields0
	equipmentfirmwareeffectFields := schema.EquipmentFirmwareEffect{}.Fields()
	_ = equipmentfirmwareeffectFields
	// equipmentfirmwareeffectDescVersion is the schema descriptor for version field.
	equipmentfirmwareeffectDescVersion := equipmentfirmwareeffectMixinFields0[1].Descriptor()
	// equipmentfirmwareeffect.DefaultVersion holds the default value on creation for the version field.
	equipmentfirmwareeffect.DefaultVersion = equipmentfirmwareeffectDescVersion.Default.(int64)
	// equipmentfirmwareeffectDescCreatedBy is the schema descriptor for created_by field.
	equipmentfirmwareeffectDescCreatedBy := equipmentfirmwareeffectMixinFields0[2].Descriptor()
	// equipmentfirmwareeffect.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipmentfirmwareeffect.DefaultCreatedBy = datasource.UUID(equipmentfirmwareeffectDescCreatedBy.Default.(uint64))
	// equipmentfirmwareeffectDescCreatedAt is the schema descriptor for created_at field.
	equipmentfirmwareeffectDescCreatedAt := equipmentfirmwareeffectMixinFields0[3].Descriptor()
	// equipmentfirmwareeffect.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipmentfirmwareeffect.DefaultCreatedAt = equipmentfirmwareeffectDescCreatedAt.Default.(int64)
	// equipmentfirmwareeffectDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentfirmwareeffectDescUpdatedBy := equipmentfirmwareeffectMixinFields0[4].Descriptor()
	// equipmentfirmwareeffect.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipmentfirmwareeffect.DefaultUpdatedBy = datasource.UUID(equipmentfirmwareeffectDescUpdatedBy.Default.(uint64))
	// equipmentfirmwareeffectDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentfirmwareeffectDescUpdatedAt := equipmentfirmwareeffectMixinFields0[5].Descriptor()
	// equipmentfirmwareeffect.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipmentfirmwareeffect.DefaultUpdatedAt = equipmentfirmwareeffectDescUpdatedAt.Default.(int64)
	// equipmentfirmwareeffect.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipmentfirmwareeffect.UpdateDefaultUpdatedAt = equipmentfirmwareeffectDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentfirmwareeffectDescID is the schema descriptor for id field.
	equipmentfirmwareeffectDescID := equipmentfirmwareeffectMixinFields0[0].Descriptor()
	// equipmentfirmwareeffect.DefaultID holds the default value on creation for the id field.
	equipmentfirmwareeffect.DefaultID = datasource.UUID(equipmentfirmwareeffectDescID.Default.(uint64))
	equipmentinfoMixin := schema.EquipmentInfo{}.Mixin()
	equipmentinfoMixinFields0 := equipmentinfoMixin[0].Fields()
	_ = equipmentinfoMixinFields0
	equipmentinfoFields := schema.EquipmentInfo{}.Fields()
	_ = equipmentinfoFields
	// equipmentinfoDescVersion is the schema descriptor for version field.
	equipmentinfoDescVersion := equipmentinfoMixinFields0[1].Descriptor()
	// equipmentinfo.DefaultVersion holds the default value on creation for the version field.
	equipmentinfo.DefaultVersion = equipmentinfoDescVersion.Default.(int64)
	// equipmentinfoDescCreatedBy is the schema descriptor for created_by field.
	equipmentinfoDescCreatedBy := equipmentinfoMixinFields0[2].Descriptor()
	// equipmentinfo.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipmentinfo.DefaultCreatedBy = datasource.UUID(equipmentinfoDescCreatedBy.Default.(uint64))
	// equipmentinfoDescCreatedAt is the schema descriptor for created_at field.
	equipmentinfoDescCreatedAt := equipmentinfoMixinFields0[3].Descriptor()
	// equipmentinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipmentinfo.DefaultCreatedAt = equipmentinfoDescCreatedAt.Default.(int64)
	// equipmentinfoDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentinfoDescUpdatedBy := equipmentinfoMixinFields0[4].Descriptor()
	// equipmentinfo.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipmentinfo.DefaultUpdatedBy = datasource.UUID(equipmentinfoDescUpdatedBy.Default.(uint64))
	// equipmentinfoDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentinfoDescUpdatedAt := equipmentinfoMixinFields0[5].Descriptor()
	// equipmentinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipmentinfo.DefaultUpdatedAt = equipmentinfoDescUpdatedAt.Default.(int64)
	// equipmentinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipmentinfo.UpdateDefaultUpdatedAt = equipmentinfoDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentinfoDescID is the schema descriptor for id field.
	equipmentinfoDescID := equipmentinfoMixinFields0[0].Descriptor()
	// equipmentinfo.DefaultID holds the default value on creation for the id field.
	equipmentinfo.DefaultID = datasource.UUID(equipmentinfoDescID.Default.(uint64))
	equipmentiotMixin := schema.EquipmentIot{}.Mixin()
	equipmentiotMixinFields0 := equipmentiotMixin[0].Fields()
	_ = equipmentiotMixinFields0
	equipmentiotFields := schema.EquipmentIot{}.Fields()
	_ = equipmentiotFields
	// equipmentiotDescVersion is the schema descriptor for version field.
	equipmentiotDescVersion := equipmentiotMixinFields0[1].Descriptor()
	// equipmentiot.DefaultVersion holds the default value on creation for the version field.
	equipmentiot.DefaultVersion = equipmentiotDescVersion.Default.(int64)
	// equipmentiotDescCreatedBy is the schema descriptor for created_by field.
	equipmentiotDescCreatedBy := equipmentiotMixinFields0[2].Descriptor()
	// equipmentiot.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipmentiot.DefaultCreatedBy = datasource.UUID(equipmentiotDescCreatedBy.Default.(uint64))
	// equipmentiotDescCreatedAt is the schema descriptor for created_at field.
	equipmentiotDescCreatedAt := equipmentiotMixinFields0[3].Descriptor()
	// equipmentiot.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipmentiot.DefaultCreatedAt = equipmentiotDescCreatedAt.Default.(int64)
	// equipmentiotDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentiotDescUpdatedBy := equipmentiotMixinFields0[4].Descriptor()
	// equipmentiot.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipmentiot.DefaultUpdatedBy = datasource.UUID(equipmentiotDescUpdatedBy.Default.(uint64))
	// equipmentiotDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentiotDescUpdatedAt := equipmentiotMixinFields0[5].Descriptor()
	// equipmentiot.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipmentiot.DefaultUpdatedAt = equipmentiotDescUpdatedAt.Default.(int64)
	// equipmentiot.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipmentiot.UpdateDefaultUpdatedAt = equipmentiotDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentiotDescID is the schema descriptor for id field.
	equipmentiotDescID := equipmentiotMixinFields0[0].Descriptor()
	// equipmentiot.DefaultID holds the default value on creation for the id field.
	equipmentiot.DefaultID = datasource.UUID(equipmentiotDescID.Default.(uint64))
	equipmentlogMixin := schema.EquipmentLog{}.Mixin()
	equipmentlogMixinFields0 := equipmentlogMixin[0].Fields()
	_ = equipmentlogMixinFields0
	equipmentlogFields := schema.EquipmentLog{}.Fields()
	_ = equipmentlogFields
	// equipmentlogDescVersion is the schema descriptor for version field.
	equipmentlogDescVersion := equipmentlogMixinFields0[1].Descriptor()
	// equipmentlog.DefaultVersion holds the default value on creation for the version field.
	equipmentlog.DefaultVersion = equipmentlogDescVersion.Default.(int64)
	// equipmentlogDescCreatedBy is the schema descriptor for created_by field.
	equipmentlogDescCreatedBy := equipmentlogMixinFields0[2].Descriptor()
	// equipmentlog.DefaultCreatedBy holds the default value on creation for the created_by field.
	equipmentlog.DefaultCreatedBy = datasource.UUID(equipmentlogDescCreatedBy.Default.(uint64))
	// equipmentlogDescCreatedAt is the schema descriptor for created_at field.
	equipmentlogDescCreatedAt := equipmentlogMixinFields0[3].Descriptor()
	// equipmentlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	equipmentlog.DefaultCreatedAt = equipmentlogDescCreatedAt.Default.(int64)
	// equipmentlogDescUpdatedBy is the schema descriptor for updated_by field.
	equipmentlogDescUpdatedBy := equipmentlogMixinFields0[4].Descriptor()
	// equipmentlog.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	equipmentlog.DefaultUpdatedBy = datasource.UUID(equipmentlogDescUpdatedBy.Default.(uint64))
	// equipmentlogDescUpdatedAt is the schema descriptor for updated_at field.
	equipmentlogDescUpdatedAt := equipmentlogMixinFields0[5].Descriptor()
	// equipmentlog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	equipmentlog.DefaultUpdatedAt = equipmentlogDescUpdatedAt.Default.(int64)
	// equipmentlog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	equipmentlog.UpdateDefaultUpdatedAt = equipmentlogDescUpdatedAt.UpdateDefault.(func() int64)
	// equipmentlogDescID is the schema descriptor for id field.
	equipmentlogDescID := equipmentlogMixinFields0[0].Descriptor()
	// equipmentlog.DefaultID holds the default value on creation for the id field.
	equipmentlog.DefaultID = datasource.UUID(equipmentlogDescID.Default.(uint64))
	evseMixin := schema.Evse{}.Mixin()
	evseMixinFields0 := evseMixin[0].Fields()
	_ = evseMixinFields0
	evseFields := schema.Evse{}.Fields()
	_ = evseFields
	// evseDescVersion is the schema descriptor for version field.
	evseDescVersion := evseMixinFields0[1].Descriptor()
	// evse.DefaultVersion holds the default value on creation for the version field.
	evse.DefaultVersion = evseDescVersion.Default.(int64)
	// evseDescCreatedBy is the schema descriptor for created_by field.
	evseDescCreatedBy := evseMixinFields0[2].Descriptor()
	// evse.DefaultCreatedBy holds the default value on creation for the created_by field.
	evse.DefaultCreatedBy = datasource.UUID(evseDescCreatedBy.Default.(uint64))
	// evseDescCreatedAt is the schema descriptor for created_at field.
	evseDescCreatedAt := evseMixinFields0[3].Descriptor()
	// evse.DefaultCreatedAt holds the default value on creation for the created_at field.
	evse.DefaultCreatedAt = evseDescCreatedAt.Default.(int64)
	// evseDescUpdatedBy is the schema descriptor for updated_by field.
	evseDescUpdatedBy := evseMixinFields0[4].Descriptor()
	// evse.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	evse.DefaultUpdatedBy = datasource.UUID(evseDescUpdatedBy.Default.(uint64))
	// evseDescUpdatedAt is the schema descriptor for updated_at field.
	evseDescUpdatedAt := evseMixinFields0[5].Descriptor()
	// evse.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	evse.DefaultUpdatedAt = evseDescUpdatedAt.Default.(int64)
	// evse.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	evse.UpdateDefaultUpdatedAt = evseDescUpdatedAt.UpdateDefault.(func() int64)
	// evseDescID is the schema descriptor for id field.
	evseDescID := evseMixinFields0[0].Descriptor()
	// evse.DefaultID holds the default value on creation for the id field.
	evse.DefaultID = datasource.UUID(evseDescID.Default.(uint64))
	firmwareMixin := schema.Firmware{}.Mixin()
	firmwareMixinFields0 := firmwareMixin[0].Fields()
	_ = firmwareMixinFields0
	firmwareFields := schema.Firmware{}.Fields()
	_ = firmwareFields
	// firmwareDescVersion is the schema descriptor for version field.
	firmwareDescVersion := firmwareMixinFields0[1].Descriptor()
	// firmware.DefaultVersion holds the default value on creation for the version field.
	firmware.DefaultVersion = firmwareDescVersion.Default.(int64)
	// firmwareDescCreatedBy is the schema descriptor for created_by field.
	firmwareDescCreatedBy := firmwareMixinFields0[2].Descriptor()
	// firmware.DefaultCreatedBy holds the default value on creation for the created_by field.
	firmware.DefaultCreatedBy = datasource.UUID(firmwareDescCreatedBy.Default.(uint64))
	// firmwareDescCreatedAt is the schema descriptor for created_at field.
	firmwareDescCreatedAt := firmwareMixinFields0[3].Descriptor()
	// firmware.DefaultCreatedAt holds the default value on creation for the created_at field.
	firmware.DefaultCreatedAt = firmwareDescCreatedAt.Default.(int64)
	// firmwareDescUpdatedBy is the schema descriptor for updated_by field.
	firmwareDescUpdatedBy := firmwareMixinFields0[4].Descriptor()
	// firmware.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	firmware.DefaultUpdatedBy = datasource.UUID(firmwareDescUpdatedBy.Default.(uint64))
	// firmwareDescUpdatedAt is the schema descriptor for updated_at field.
	firmwareDescUpdatedAt := firmwareMixinFields0[5].Descriptor()
	// firmware.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	firmware.DefaultUpdatedAt = firmwareDescUpdatedAt.Default.(int64)
	// firmware.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	firmware.UpdateDefaultUpdatedAt = firmwareDescUpdatedAt.UpdateDefault.(func() int64)
	// firmwareDescID is the schema descriptor for id field.
	firmwareDescID := firmwareMixinFields0[0].Descriptor()
	// firmware.DefaultID holds the default value on creation for the id field.
	firmware.DefaultID = datasource.UUID(firmwareDescID.Default.(uint64))
	manufacturerMixin := schema.Manufacturer{}.Mixin()
	manufacturerMixinFields0 := manufacturerMixin[0].Fields()
	_ = manufacturerMixinFields0
	manufacturerFields := schema.Manufacturer{}.Fields()
	_ = manufacturerFields
	// manufacturerDescVersion is the schema descriptor for version field.
	manufacturerDescVersion := manufacturerMixinFields0[1].Descriptor()
	// manufacturer.DefaultVersion holds the default value on creation for the version field.
	manufacturer.DefaultVersion = manufacturerDescVersion.Default.(int64)
	// manufacturerDescCreatedBy is the schema descriptor for created_by field.
	manufacturerDescCreatedBy := manufacturerMixinFields0[2].Descriptor()
	// manufacturer.DefaultCreatedBy holds the default value on creation for the created_by field.
	manufacturer.DefaultCreatedBy = datasource.UUID(manufacturerDescCreatedBy.Default.(uint64))
	// manufacturerDescCreatedAt is the schema descriptor for created_at field.
	manufacturerDescCreatedAt := manufacturerMixinFields0[3].Descriptor()
	// manufacturer.DefaultCreatedAt holds the default value on creation for the created_at field.
	manufacturer.DefaultCreatedAt = manufacturerDescCreatedAt.Default.(int64)
	// manufacturerDescUpdatedBy is the schema descriptor for updated_by field.
	manufacturerDescUpdatedBy := manufacturerMixinFields0[4].Descriptor()
	// manufacturer.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	manufacturer.DefaultUpdatedBy = datasource.UUID(manufacturerDescUpdatedBy.Default.(uint64))
	// manufacturerDescUpdatedAt is the schema descriptor for updated_at field.
	manufacturerDescUpdatedAt := manufacturerMixinFields0[5].Descriptor()
	// manufacturer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	manufacturer.DefaultUpdatedAt = manufacturerDescUpdatedAt.Default.(int64)
	// manufacturer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	manufacturer.UpdateDefaultUpdatedAt = manufacturerDescUpdatedAt.UpdateDefault.(func() int64)
	// manufacturerDescID is the schema descriptor for id field.
	manufacturerDescID := manufacturerMixinFields0[0].Descriptor()
	// manufacturer.DefaultID holds the default value on creation for the id field.
	manufacturer.DefaultID = datasource.UUID(manufacturerDescID.Default.(uint64))
	modelMixin := schema.Model{}.Mixin()
	modelMixinFields0 := modelMixin[0].Fields()
	_ = modelMixinFields0
	modelFields := schema.Model{}.Fields()
	_ = modelFields
	// modelDescVersion is the schema descriptor for version field.
	modelDescVersion := modelMixinFields0[1].Descriptor()
	// model.DefaultVersion holds the default value on creation for the version field.
	model.DefaultVersion = modelDescVersion.Default.(int64)
	// modelDescCreatedBy is the schema descriptor for created_by field.
	modelDescCreatedBy := modelMixinFields0[2].Descriptor()
	// model.DefaultCreatedBy holds the default value on creation for the created_by field.
	model.DefaultCreatedBy = datasource.UUID(modelDescCreatedBy.Default.(uint64))
	// modelDescCreatedAt is the schema descriptor for created_at field.
	modelDescCreatedAt := modelMixinFields0[3].Descriptor()
	// model.DefaultCreatedAt holds the default value on creation for the created_at field.
	model.DefaultCreatedAt = modelDescCreatedAt.Default.(int64)
	// modelDescUpdatedBy is the schema descriptor for updated_by field.
	modelDescUpdatedBy := modelMixinFields0[4].Descriptor()
	// model.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	model.DefaultUpdatedBy = datasource.UUID(modelDescUpdatedBy.Default.(uint64))
	// modelDescUpdatedAt is the schema descriptor for updated_at field.
	modelDescUpdatedAt := modelMixinFields0[5].Descriptor()
	// model.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	model.DefaultUpdatedAt = modelDescUpdatedAt.Default.(int64)
	// model.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	model.UpdateDefaultUpdatedAt = modelDescUpdatedAt.UpdateDefault.(func() int64)
	// modelDescID is the schema descriptor for id field.
	modelDescID := modelMixinFields0[0].Descriptor()
	// model.DefaultID holds the default value on creation for the id field.
	model.DefaultID = datasource.UUID(modelDescID.Default.(uint64))
	ordereventFields := schema.OrderEvent{}.Fields()
	_ = ordereventFields
	// ordereventDescID is the schema descriptor for id field.
	ordereventDescID := ordereventFields[0].Descriptor()
	// orderevent.DefaultID holds the default value on creation for the id field.
	orderevent.DefaultID = datasource.UUID(ordereventDescID.Default.(uint64))
	orderinfoMixin := schema.OrderInfo{}.Mixin()
	orderinfoMixinFields0 := orderinfoMixin[0].Fields()
	_ = orderinfoMixinFields0
	orderinfoFields := schema.OrderInfo{}.Fields()
	_ = orderinfoFields
	// orderinfoDescVersion is the schema descriptor for version field.
	orderinfoDescVersion := orderinfoMixinFields0[1].Descriptor()
	// orderinfo.DefaultVersion holds the default value on creation for the version field.
	orderinfo.DefaultVersion = orderinfoDescVersion.Default.(int64)
	// orderinfoDescCreatedBy is the schema descriptor for created_by field.
	orderinfoDescCreatedBy := orderinfoMixinFields0[2].Descriptor()
	// orderinfo.DefaultCreatedBy holds the default value on creation for the created_by field.
	orderinfo.DefaultCreatedBy = datasource.UUID(orderinfoDescCreatedBy.Default.(uint64))
	// orderinfoDescCreatedAt is the schema descriptor for created_at field.
	orderinfoDescCreatedAt := orderinfoMixinFields0[3].Descriptor()
	// orderinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderinfo.DefaultCreatedAt = orderinfoDescCreatedAt.Default.(int64)
	// orderinfoDescUpdatedBy is the schema descriptor for updated_by field.
	orderinfoDescUpdatedBy := orderinfoMixinFields0[4].Descriptor()
	// orderinfo.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	orderinfo.DefaultUpdatedBy = datasource.UUID(orderinfoDescUpdatedBy.Default.(uint64))
	// orderinfoDescUpdatedAt is the schema descriptor for updated_at field.
	orderinfoDescUpdatedAt := orderinfoMixinFields0[5].Descriptor()
	// orderinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderinfo.DefaultUpdatedAt = orderinfoDescUpdatedAt.Default.(int64)
	// orderinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderinfo.UpdateDefaultUpdatedAt = orderinfoDescUpdatedAt.UpdateDefault.(func() int64)
	// orderinfoDescID is the schema descriptor for id field.
	orderinfoDescID := orderinfoMixinFields0[0].Descriptor()
	// orderinfo.DefaultID holds the default value on creation for the id field.
	orderinfo.DefaultID = datasource.UUID(orderinfoDescID.Default.(uint64))
	reservationMixin := schema.Reservation{}.Mixin()
	reservationMixinFields0 := reservationMixin[0].Fields()
	_ = reservationMixinFields0
	reservationFields := schema.Reservation{}.Fields()
	_ = reservationFields
	// reservationDescVersion is the schema descriptor for version field.
	reservationDescVersion := reservationMixinFields0[1].Descriptor()
	// reservation.DefaultVersion holds the default value on creation for the version field.
	reservation.DefaultVersion = reservationDescVersion.Default.(int64)
	// reservationDescCreatedBy is the schema descriptor for created_by field.
	reservationDescCreatedBy := reservationMixinFields0[2].Descriptor()
	// reservation.DefaultCreatedBy holds the default value on creation for the created_by field.
	reservation.DefaultCreatedBy = datasource.UUID(reservationDescCreatedBy.Default.(uint64))
	// reservationDescCreatedAt is the schema descriptor for created_at field.
	reservationDescCreatedAt := reservationMixinFields0[3].Descriptor()
	// reservation.DefaultCreatedAt holds the default value on creation for the created_at field.
	reservation.DefaultCreatedAt = reservationDescCreatedAt.Default.(int64)
	// reservationDescUpdatedBy is the schema descriptor for updated_by field.
	reservationDescUpdatedBy := reservationMixinFields0[4].Descriptor()
	// reservation.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	reservation.DefaultUpdatedBy = datasource.UUID(reservationDescUpdatedBy.Default.(uint64))
	// reservationDescUpdatedAt is the schema descriptor for updated_at field.
	reservationDescUpdatedAt := reservationMixinFields0[5].Descriptor()
	// reservation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reservation.DefaultUpdatedAt = reservationDescUpdatedAt.Default.(int64)
	// reservation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reservation.UpdateDefaultUpdatedAt = reservationDescUpdatedAt.UpdateDefault.(func() int64)
	// reservationDescID is the schema descriptor for id field.
	reservationDescID := reservationMixinFields0[0].Descriptor()
	// reservation.DefaultID holds the default value on creation for the id field.
	reservation.DefaultID = datasource.UUID(reservationDescID.Default.(uint64))
	smartchargingeffectMixin := schema.SmartChargingEffect{}.Mixin()
	smartchargingeffectMixinFields0 := smartchargingeffectMixin[0].Fields()
	_ = smartchargingeffectMixinFields0
	smartchargingeffectFields := schema.SmartChargingEffect{}.Fields()
	_ = smartchargingeffectFields
	// smartchargingeffectDescVersion is the schema descriptor for version field.
	smartchargingeffectDescVersion := smartchargingeffectMixinFields0[1].Descriptor()
	// smartchargingeffect.DefaultVersion holds the default value on creation for the version field.
	smartchargingeffect.DefaultVersion = smartchargingeffectDescVersion.Default.(int64)
	// smartchargingeffectDescCreatedBy is the schema descriptor for created_by field.
	smartchargingeffectDescCreatedBy := smartchargingeffectMixinFields0[2].Descriptor()
	// smartchargingeffect.DefaultCreatedBy holds the default value on creation for the created_by field.
	smartchargingeffect.DefaultCreatedBy = datasource.UUID(smartchargingeffectDescCreatedBy.Default.(uint64))
	// smartchargingeffectDescCreatedAt is the schema descriptor for created_at field.
	smartchargingeffectDescCreatedAt := smartchargingeffectMixinFields0[3].Descriptor()
	// smartchargingeffect.DefaultCreatedAt holds the default value on creation for the created_at field.
	smartchargingeffect.DefaultCreatedAt = smartchargingeffectDescCreatedAt.Default.(int64)
	// smartchargingeffectDescUpdatedBy is the schema descriptor for updated_by field.
	smartchargingeffectDescUpdatedBy := smartchargingeffectMixinFields0[4].Descriptor()
	// smartchargingeffect.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	smartchargingeffect.DefaultUpdatedBy = datasource.UUID(smartchargingeffectDescUpdatedBy.Default.(uint64))
	// smartchargingeffectDescUpdatedAt is the schema descriptor for updated_at field.
	smartchargingeffectDescUpdatedAt := smartchargingeffectMixinFields0[5].Descriptor()
	// smartchargingeffect.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	smartchargingeffect.DefaultUpdatedAt = smartchargingeffectDescUpdatedAt.Default.(int64)
	// smartchargingeffect.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	smartchargingeffect.UpdateDefaultUpdatedAt = smartchargingeffectDescUpdatedAt.UpdateDefault.(func() int64)
	// smartchargingeffectDescID is the schema descriptor for id field.
	smartchargingeffectDescID := smartchargingeffectMixinFields0[0].Descriptor()
	// smartchargingeffect.DefaultID holds the default value on creation for the id field.
	smartchargingeffect.DefaultID = datasource.UUID(smartchargingeffectDescID.Default.(uint64))
}
