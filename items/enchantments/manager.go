package enchantments

// Manager provides helper functions for managing enchantments,
// such as registering, deregistering and checks for those.
type Manager struct {
	// stringIds is a map of enchantment types,
	// indexed with the string ID.
	// Example: "minecraft:absorption": Type
	stringIds map[string]Type
	// byteIds is a map of enchantment types,
	// indexed with the byte ID.
	// Example: 3: Type
	byteIds map[byte]Type
}

// DefaultManager is the default enchantment manager.
// The init function registers the default enchantments.
var DefaultManager = NewManager()

// init registers default enchantments of the manager.
func init() {
	DefaultManager.RegisterDefaults()
}

// NewManager returns a new enchantment manager.
// Maps are allocated, but no default enchantments
// are registered yet.
func NewManager() *Manager {
	return &Manager{make(map[string]Type), make(map[byte]Type)}
}

// RegisterDefaults registers all default enchantments.
// This function should be called whenever a new manager
// is made, in order to have all default enchantments registered.
func (manager *Manager) RegisterDefaults() {

}
