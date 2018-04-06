package inventory

import (
	"fmt"
	"github.com/irmine/gomine/items"
	"testing"
)

func Test(t *testing.T) {
	manager := items.NewManager()
	manager.Register(items.NewType("minecraft:emerald"), true)
	manager.Register(items.NewType("minecraft:glass_bottle"), true)

	inv := NewInventory(9)
	item, _ := manager.Get("minecraft:emerald", 8)
	inv.SetItem(item, 6)
	inv.SetItem(item, 4)
	item, _ = manager.Get("minecraft:glass_bottle", 34)
	inv.SetItem(item, 8)

	if err := inv.SetItem(item, 9); err != nil {
		fmt.Println("Inventory size check works:", err)
	}

	fmt.Println(inv)

	item, _ = manager.Get("minecraft:emerald", 16)
	if inv.Contains(item) {
		fmt.Println("Inventory contains 16 emeralds.")
	} else {
		fmt.Println("Inventory does not contain 16 emeralds.")
	}
	item, _ = manager.Get("minecraft:glass_bottle", 35)
	if inv.Contains(item) {
		fmt.Println("Inventory contains 35 glass bottles.")
	} else {
		fmt.Println("Inventory does not contain 35 glass bottles.")
	}
}
