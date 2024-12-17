package mapsecure

import (
	"testing"
)

func TestRwLock(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", "value1")
	val, ok := sm.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("预期应该是value1, 实际是%v", val)
	}
	sm.Delete("key1")
	_, ok = sm.Get("key1")
	if ok {
		t.Errorf("预期应该是false, 实际是true")
	}

}
