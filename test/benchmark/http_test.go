package benchmark

import (
	http2 "sec-tools/fasthttp"
	"sec-tools/http"
	"testing"
)

/**
 * @Description
 * @Author r0cky
 * @Date 2022/3/9 14:16
 */

func BenchmarkHttp(b *testing.B) {

	for i := 0; i < b.N; i++ {
		req := http.NewRequest()
		_, _, _ = req.GetH("https://baidu.com")
	}
}

func BenchmarkFastHttp(b *testing.B) {
	req := http2.NewRequest()
	for i := 0; i < b.N; i++ {
		_, _, _ = req.GetH("https://baidu.com")
	}
}
