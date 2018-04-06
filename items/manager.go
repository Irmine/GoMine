package items

import "github.com/irmine/gonbt"

// Manager supplies helper functions for item type registering.
// Item types get registered by their string ID,
// and can be retrieved using these.
type Manager struct {
	// stringIds is a map containing item types,
	// indexed by string IDs.
	// Example: "minecraft:golden_apple": Type
	stringIds map[string]Type
	// creativeItems is a map containing item types,
	// indexed by string IDs, similarly to stringIds.
	// This map contains all items,
	// that should be displayed in the creative inventory.
	creativeItems map[string]Type
}

// DefaultManager is the default item manager.
// The default items are registered upon the init function.
var DefaultManager = NewManager()

// init initializes all default item types,
// of the default item manager.
func init() {
	DefaultManager.RegisterDefaults()
}

// NewManager returns a new item registry.
// New registries will not have default items registered.
// Default registries should be registered using RegisterDefaults.
func NewManager() *Manager {
	return &Manager{make(map[string]Type), make(map[string]Type)}
}

// Register registers a new item type.
// The item type will be registered to the stringIds map.
// Registered item types can be deregistered,
// using the Deregister functions.
// If registerCreative is set to true,
// the item will also be registered as creative item.
func (registry *Manager) Register(t Type, registerCreative bool) {
	registry.stringIds[t.GetId()] = t
	if registerCreative {
		registry.RegisterCreativeType(t)
	}
}

// RegisterMultiple registers multiple types at once.
// Item types will be registered to the stringIds map,
// and can be deregistered separately from each other.
// If registerCreative is set to true,
// the items will also be registered as creative item.
func (registry *Manager) RegisterMultiple(types []Type, registerCreative bool) {
	for _, t := range types {
		registry.stringIds[t.GetId()] = t
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
	registry.creativeItems[t.GetId()] = t
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

// IsRegistered checks if an item type is registered,
// by its string ID in the stringIds map.
// Returns true if the string ID is registered.
func (registry *Manager) IsRegistered(stringId string) bool {
	_, ok := registry.stringIds[stringId]
	return ok
}

// Deregister deregisters an item type,
// by its string ID in the stringIds map.
// Returns true if the item type was deregistered successfully.
func (registry *Manager) Deregister(stringId string) bool {
	_, ok := registry.stringIds[stringId]
	if !ok {
		return false
	}
	delete(registry.stringIds, stringId)
	return true
}

// Get attempts to return a new item stack by a string ID,
// and sets the stack's count to the count given.
// A bool gets returned to indicate whether any item was found.
// If no item type could be found with the given string ID,
// a default air item and a bool false gets returned.
// It is recommended to use this function over Get where possible.
func (registry *Manager) Get(stringId string, count byte) (*Stack, bool) {
	t, ok := registry.stringIds[stringId]
	if !ok {
		t = registry.stringIds["minecraft:air"]
	}
	return &Stack{Type: t, Count: count, DisplayName: t.name, cachedNBT: gonbt.NewCompound("", make(map[string]gonbt.INamedTag))}, ok
}

// GetTypes returns all registered item types.
// Item types are returned in a map of the form stringId => Type.
func (registry *Manager) GetTypes() map[string]Type {
	return registry.stringIds
}

// RegisterDefaults registers all default items.
// This function should be called immediately after NewManager,
// in order to register the proper default items.
func (registry *Manager) RegisterDefaults() {
	registry.Register(NewType("minecraft:air"), false)
	registry.Register(NewType("minecraft:glass_bottle"), true)
}
