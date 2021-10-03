package split

import (
	"reflect"
	"testing"
)

// 测试函数：单个测试
// func TestSplit(t *testing.T) {
// 	got := Split("我爱中国", "爱")
// 	ret := []string{"我", "中国"}
// 	if !reflect.DeepEqual(got, ret) {
// 		t.Errorf("got: %v, ret: %v", got, ret)
// 	}
// }

// 测试函数：测试组
/*
使用命令：
	1、执行测试：go test -v
	2、执行组中某一个测试：go test -run=split/test3
	3、测试覆盖率：go test -cover
	4、输出测试覆盖率文件，并使用html方式查看：go test -cover -coverprofile=c.out | go tool cover -html=c.out
*/
func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		ret   []string
	}
	tests := map[string]test{
		"test1": {input: "xiangshangdexinyang", sep: "de", ret: []string{"xiangshang", "xinyang"}},
		"test":  {input: "向上的信仰", sep: "的", ret: []string{"向上", "信仰"}},
		"test3": {input: "a:b:c", sep: ":", ret: []string{"a", "b", "c"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.ret) {
				t.Errorf("got: %v, ret: %v", got, tc.ret)
			}
		})
	}
}

// 性能基准测试
/*
使用命令：
1、go test -bench=Split
结果：
BenchmarkSplit-8        10975743               106.8 ns/op	// BenchmarkSplit-8:真正工作的进程数为8；10975743：执行次数， 106.8 ns/op:每一次操作耗费106.8纳秒
PASS
ok      senior/split    2.119s

2、查看内存申请：go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: senior/split
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkSplit-8        11256488               106.6 ns/op            48 B/op          2 allocs/op	// 48 B/op: 每次操作申请48字节的内存; 2 allocs/op:每次操作申请两次内存
PASS
ok      senior/split    1.866s
*/
func BenchmarkSplit(b *testing.B) {
	// b.N 为不固定的数
	for i := 0; i < b.N; i++ {
		Split("我爱中国", "爱")
	}
}
