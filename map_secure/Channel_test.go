package mapsecure

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeMapConcurrentAccess(t *testing.T) {
	// 创建 SafeMap 实例
	sm := NewSafeMap2()

	// 使用 WaitGroup 来等待所有 Goroutine 完成
	var wg sync.WaitGroup

	// 并发执行 Set 和 Get 操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)

			// 设置值
			sm.Set2(key, value)

			// 获取值并检查
			val, ok := sm.Get2(key)
			if !ok || val != value {
				t.Errorf("key: %s, expected value: %s, got: %v", key, value, val)
			}
		}(i)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()

	// 删除一些键值
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)

			// 删除键
			sm.Delete2(key)

			// 确认键是否已删除
			_, ok := sm.Get2(key)
			if ok {
				t.Errorf("key: %s should be deleted", key)
			}
		}(i)
	}

	// 等待所有删除操作完成
	wg.Wait()
}
