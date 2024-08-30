package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建并初始化 map
	m := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 20,
	}

	// 提取 map 的所有键
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// 对键进行排序
	sort.Strings(keys)

	// 按排序后的键顺序遍历 map
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
