package mapsecure

type SafeMap2 struct {
	operation chan operation
	m         map[string]interface{}
}

type operation struct {
	key    string
	value  interface{}
	action string
	result chan result
}

type result struct {
	value interface{}
	ok    bool
}

func NewSafeMap2() *SafeMap2 {
	sm := &SafeMap2{
		operation: make(chan operation),
		m:         make(map[string]interface{}),
	}
	go sm.run()
	return sm
}

func (sm *SafeMap2) run() {
	for op := range sm.operation {
		switch op.action {
		case "get":
			val, ok := sm.m[op.key]
			op.result <- result{value: val, ok: ok}
		case "set":
			sm.m[op.key] = op.value
			op.result <- result{}
		case "delete":
			delete(sm.m, op.key)
			op.result <- result{}
		}
	}
}

func (sm *SafeMap2) Get2(key string) (interface{}, bool) {
	resultChan := make(chan result)
	sm.operation <- operation{
		action: "Get",
		key:    key,
		result: resultChan,
	}
	res := <-resultChan
	return res.value, res.ok
}

func (sm *SafeMap2) Set2(key string, value interface{}) {
	resultChan := make(chan result)
	sm.operation <- operation{
		action: "set",
		key:    key,
		value:  value,
		result: resultChan,
	}
	<-resultChan
}

func (sm *SafeMap2) Delete2(key string) {
	resultChan := make(chan result)
	sm.operation <- operation{
		key:    key,
		action: "delete",
		result: resultChan,
	}
	<-resultChan
}
