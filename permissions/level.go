package permissions

const (
	LevelVisitor  PermissionLevel = iota
	LevelMember                   = 1
	LevelOperator                 = 2
	LevelCustom                   = 3
)

// A Permission level is used to connect groups with permissions.
type PermissionLevel byte
