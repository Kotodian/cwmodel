package repository

import (
	"context"

	"github.com/Kotodian/ent-practice/ent"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/gokit/datasource"
)

func FindAllEquipment(ctx context.Context, client *ent.Client, withEquipmentInfo bool, fields ...string) ([]*ent.Equipment, error) {
	query := client.Equipment.Query()
	var sel *ent.EquipmentSelect
	if len(fields) > 0 {
		sel = query.Select(fields...)
	} else {
		sel = query.Select(equipment.Columns...)
	}
	if withEquipmentInfo {
		query = sel.Where(equipment.HasEquipmentInfo())
		query = query.WithEquipmentInfo()
	}

	return query.All(ctx)
}

func GetEquipment(ctx context.Context, client *ent.Client, id datasource.UUID, withEquipmentInfo bool, fields ...string) (*ent.Equipment, error) {
	query := client.Equipment.Query()
	var sel *ent.EquipmentSelect
	if len(fields) > 0 {
		sel = query.Select(fields...)
	} else {
		sel = query.Select(equipment.Columns...)
	}
	if withEquipmentInfo {
		query = sel.Where(equipment.HasEquipmentInfo()).Where(equipment.IDEQ(id))
		query = query.WithEquipmentInfo()
	}
	return query.Only(ctx)
}

func CreateEquipmentInfo(ctx context.Context, client *ent.Client, equipment *ent.Equipment, info *ent.EquipmentInfo) error {
	create := client.EquipmentInfo.Create()
	return create.SetEquipment(equipment).
		SetEquipmentID(equipment.ID).
		SetEquipmentSn(equipment.Sn).
		SetAccessPod(info.AccessPod).
		SetState(info.State).Exec(ctx)
}
