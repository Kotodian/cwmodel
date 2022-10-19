// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/Kotodian/ent-practice/ent"
)

// The AppModuleInfoFunc type is an adapter to allow the use of ordinary
// function as AppModuleInfo mutator.
type AppModuleInfoFunc func(context.Context, *ent.AppModuleInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppModuleInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppModuleInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppModuleInfoMutation", m)
	}
	return f(ctx, mv)
}

// The ConnectorFunc type is an adapter to allow the use of ordinary
// function as Connector mutator.
type ConnectorFunc func(context.Context, *ent.ConnectorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ConnectorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ConnectorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ConnectorMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentFunc type is an adapter to allow the use of ordinary
// function as Equipment mutator.
type EquipmentFunc func(context.Context, *ent.EquipmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentAlarmFunc type is an adapter to allow the use of ordinary
// function as EquipmentAlarm mutator.
type EquipmentAlarmFunc func(context.Context, *ent.EquipmentAlarmMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentAlarmFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentAlarmMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentAlarmMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentFirmwareEffectFunc type is an adapter to allow the use of ordinary
// function as EquipmentFirmwareEffect mutator.
type EquipmentFirmwareEffectFunc func(context.Context, *ent.EquipmentFirmwareEffectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentFirmwareEffectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentFirmwareEffectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentFirmwareEffectMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentInfoFunc type is an adapter to allow the use of ordinary
// function as EquipmentInfo mutator.
type EquipmentInfoFunc func(context.Context, *ent.EquipmentInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentInfoMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentIotFunc type is an adapter to allow the use of ordinary
// function as EquipmentIot mutator.
type EquipmentIotFunc func(context.Context, *ent.EquipmentIotMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentIotFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentIotMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentIotMutation", m)
	}
	return f(ctx, mv)
}

// The EquipmentLogFunc type is an adapter to allow the use of ordinary
// function as EquipmentLog mutator.
type EquipmentLogFunc func(context.Context, *ent.EquipmentLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EquipmentLogMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentLogMutation", m)
	}
	return f(ctx, mv)
}

// The EvseFunc type is an adapter to allow the use of ordinary
// function as Evse mutator.
type EvseFunc func(context.Context, *ent.EvseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EvseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EvseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EvseMutation", m)
	}
	return f(ctx, mv)
}

// The FirmwareFunc type is an adapter to allow the use of ordinary
// function as Firmware mutator.
type FirmwareFunc func(context.Context, *ent.FirmwareMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FirmwareFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FirmwareMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FirmwareMutation", m)
	}
	return f(ctx, mv)
}

// The ManufacturerFunc type is an adapter to allow the use of ordinary
// function as Manufacturer mutator.
type ManufacturerFunc func(context.Context, *ent.ManufacturerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ManufacturerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ManufacturerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ManufacturerMutation", m)
	}
	return f(ctx, mv)
}

// The ModelFunc type is an adapter to allow the use of ordinary
// function as Model mutator.
type ModelFunc func(context.Context, *ent.ModelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ModelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModelMutation", m)
	}
	return f(ctx, mv)
}

// The OrderEventFunc type is an adapter to allow the use of ordinary
// function as OrderEvent mutator.
type OrderEventFunc func(context.Context, *ent.OrderEventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderEventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderEventMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderEventMutation", m)
	}
	return f(ctx, mv)
}

// The OrderInfoFunc type is an adapter to allow the use of ordinary
// function as OrderInfo mutator.
type OrderInfoFunc func(context.Context, *ent.OrderInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderInfoMutation", m)
	}
	return f(ctx, mv)
}

// The ReservationFunc type is an adapter to allow the use of ordinary
// function as Reservation mutator.
type ReservationFunc func(context.Context, *ent.ReservationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReservationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ReservationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReservationMutation", m)
	}
	return f(ctx, mv)
}

// The SmartChargingEffectFunc type is an adapter to allow the use of ordinary
// function as SmartChargingEffect mutator.
type SmartChargingEffectFunc func(context.Context, *ent.SmartChargingEffectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SmartChargingEffectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SmartChargingEffectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SmartChargingEffectMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
