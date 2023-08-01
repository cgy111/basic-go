package main

func Map() {
	m1 := map[string]int{
		"key1": 123,
		"key2": 789,
		"key3": 101112,
	}
	m1["hello"] = 345

	for k, v := range m1 {
		println(k, v)
	}

	for k := range m1 {
		println(k)
	}

	delete(m1, "hello")
	//容量
	m2 := make(map[string]int, 12)
	m2["key2"] = 12

	val, ok := m1["cgy"]
	if ok {
		//有这个键值对
		println(val)
	}

	val = m1["Cgy"]
	println("Cgy对应的值:", val)

	delete(m1, "key1")
}

func UseKey() {
	m := map[string]int{
		"key1": 123,
	}
	keys := Keys(m)
	println(keys)
}

func Keys(m map[string]int) []string {
	return []string{"helllo"}
}
