package items

import "fmt"

// Manager supplies helper functions for item type registering.
// Item types get registered by both string ID and ID + meta keys,
// and can be retrieved using either of those.
type Manager struct {
	// stringIds is a map containing item types,
	// indexed by string IDs.
	// Example: "minecraft:golden_apple": Type
	stringIds map[string]Type
	// idDataIds is a map containing item types,
	// indexed by ID + meta combinations.
	// Example: "388:0": Type
	idDataIds map[string]Type
	// creativeItems is a map containing item types,
	// indexed by string IDs, similarly to stringIds.
	// This map contains all items,
	// that should be displayed in the creative inventory.
	creativeItems map[string]Type
}

// NewManager returns a new item registry.
// New registries will not have default items registered.
// Default registries should be registered using RegisterDefaults.
func NewManager() *Manager {
	return &Manager{make(map[string]Type), make(map[string]Type), make(map[string]Type)}
}

// RegisterDefaults registers all default items.
// This function should be called immediately after NewManager,
// in order to register the proper default items.
func (registry *Manager) RegisterDefaults() {
	registry.Register(NewType("Air", "minecraft:air", 0, 0), false)
	registry.RegisterMultiple([]Type{
		NewType("Stone", "minecraft:stone", 1, 0),
		NewType("Granite", "minecraft:granite", 1, 1),
		NewType("Polished Granite", "minecraft:polished_granite", 1, 2),
		NewType("Diorite", "minecraft:diorite", 1, 3),
		NewType("Polished Diorite", "minecraft:polished_diorite", 1, 4),
		NewType("Andesite", "minecraft:andesite", 1, 5),
		NewType("Polished Andesite", "minecraft:polished_andesite", 1, 6),
	}, true)
}

// Register registers a new item type.
// The item type will be registered to both
// the stringIds map and the idDataIds map.
// Registered item types can be deregistered,
// using the Deregister functions.
// If registerCreative is set to true,
// the item will also be registered as creative item.
func (registry *Manager) Register(t Type, registerCreative bool) {
	registry.stringIds[t.GetStringId()] = t
	registry.idDataIds[getKey(t.id, t.data)] = t
	if registerCreative {
		registry.RegisterCreativeType(t)
	}
}

// RegisterMultiple registers multiple types at once.
// Item types will be registered to both
// the stringIds map and the idDataIds map,
// and can be deregistered separately from each other.
// If registerCreative is set to true,
// the items will also be registered as creative item.
func (registry *Manager) RegisterMultiple(types []Type, registerCreative bool) {
	for _, t := range types {
		registry.stringIds[t.GetStringId()] = t
		registry.idDataIds[getKey(t.id, t.data)] = t
		if registerCreative {
			registry.RegisterCreativeType(t)
		}
	}
}

// RegisterCreativeType registers an item type,
// to the creative items map.
// All creative items will be displayed,
// in the creative inventory.
func (registry *Manager) RegisterCreativeType(t Type) {
	registry.creativeItems[t.GetStringId()] = t
}

// IsCreativeTypeRegistered checks if an item type
// is registered to the creative inventory map.
func (registry *Manager) IsCreativeTypeRegistered(stringId string) bool {
	_, ok := registry.creativeItems[stringId]
	return ok
}

// DeregisterCreativeType deregisters a creative item.
// Creative items can be deregistered using the string ID
// of that particular item.
// A bool gets returned to indicate success of the action.
func (registry *Manager) DeregisterCreativeType(stringId string) bool {
	_, ok := registry.creativeItems[stringId]
	delete(registry.creativeItems, stringId)
	return ok
}

// GetCreativeTypes returns all creative items.
// A map gets returned in the form of stringId => Type.
func (registry *Manager) GetCreativeTypes() map[string]Type {
	return registry.creativeItems
}

// IsStringIdRegistered checks if an item type is registered,
// by its string ID in the stringIds map.
// Returns true if the string ID is registered.
func (registry *Manager) IsStringIdRegistered(stringId string) bool {
	_, ok := registry.stringIds[stringId]
	return ok
}

// IsRegistered checks if an item type is registered,
// by its item ID + data combination in the idDataIds map.
// Returns true if the combination is registered.
func (registry *Manager) IsRegistered(id int16, data int16) bool {
	_, ok := registry.idDataIds[getKey(id, data)]
	return ok
}

// DeregisterByStringId deregisters an item type,
// by its string ID in the stringIds map.
// Returns true if the item type was deregistered successfully.
func (registry *Manager) DeregisterByStringId(stringId string) bool {
	t, ok := registry.stringIds[stringId]
	if !ok {
		return false
	}
	delete(registry.stringIds, stringId)
	delete(registry.idDataIds, getKey(t.id, t.data))
	return true
}

// Deregister deregisters an item type,
// by its item ID + data combination in the idDataIds map.
// Returns true if the item type was deregistered successfully.
func (registry *Manager) Deregister(id int16, data int16) bool {
	key := getKey(id, data)
	t, ok := registry.idDataIds[key]
	if !ok {
		return false
	}
	delete(registry.stringIds, t.stringId)
	delete(registry.idDataIds, key)
	return true
}

// Get attempts to return a new item stack by an ID and data,
// and sets the stack's count to the count given.
// A bool gets returned to indicate whether any item type was found.
// If no item type could be found with the given ID + data combination,
// an attempt will be made to retrieve an item by the item ID only.
// Should an item be found after that, and should it be breakable,
// then the item gets retrieved with data 0, and durability gets set.
// If still no item could be found after this,
// a default air item and a bool false gets returned.
func (registry *Manager) Get(id int16, data int16, count byte) (*Stack, bool) {
	var t Type
	var ok bool
	var dur int16
	t, ok = registry.idDataIds[getKey(id, data)]
	if !ok {
		var tTemp Type
		tTemp, ok = registry.idDataIds[getKey(id, 0)]
		if t.IsBreakable() {
			t = tTemp
			dur = data
		} else {
			t, _ = registry.idDataIds["0:0"]
		}
	}
	return &Stack{t, count, "", dur}, ok
}

func (registry *Manager) GetByStringId(stringId string, count byte) (*Stack, bool) {

}

// GetTypes returns all registered item types.
// Item types are returned in a map of the form stringId => Type.
func (registry *Manager) GetTypes() map[string]Type {
	return registry.stringIds
}

// getKey returns the idDataIds map key for the given ID + data.
func getKey(id int16, data int16) string {
	return fmt.Sprint(id, ":", data)
}
