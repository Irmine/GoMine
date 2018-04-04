package items

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	manager := NewManager()
	manager.RegisterDefaults()
	manager.Register(NewType("minecraft:emerald"), true)

	emerald, ok := manager.Get("minecraft:emerald", 5)
	if !ok {
		panic("item not registered")
	}

	fmt.Println(emerald.name, emerald.Count)
}
