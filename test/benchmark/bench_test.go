package benchmark

import (
	"sec-tools/bin/misc"
	"testing"
)

/**
 * @Description
 * @Author r0cky
 * @Date 2022/2/13 02:18
 */

func BenchmarkToMap(b *testing.B) {
	m := make(map[string]string)
	m["a"] = "111"
	for i := 0; i < b.N; i++ {
		misc.ToMap(m)
	}
}

func BenchmarkToMap2(b *testing.B) {
	m := make(map[string]string)
	m["a"] = "111"
	for i := 0; i < b.N; i++ {
		//misc.ToMap2(m)
	}
}
