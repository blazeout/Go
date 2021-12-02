package _2splitString

import (
	"reflect"
	"testing"
)

//func TestSplitString(t *testing.T) {
//	got := SplitString("abcdbef","b")
//	want := []string{"a","cd","ef"}
//	if !reflect.DeepEqual(got,want) {
//		// 测试用例失败了
//		t.Errorf("want : %v but got : %v",want,got)
//	}
//}
//
//func TestSplitString2(t *testing.T) {
//	got := SplitString("a:b:c",":")
//	want := []string{"a","b","c"}
//	if !reflect.DeepEqual(got,want) {
//		// 测试用例失败了
//		t.Errorf("want : %v but got : %v",want,got)
//	}
//}
//
//func TestSplitString3(t *testing.T) {
//	ret := SplitString("abcef","bc")
//	want := []string{"a","ef"}
//	if !reflect.DeepEqual(ret,want) {
//		t.Fatalf("want : %v but got : %v",want,ret)
//	}
//}

//func TestSplitString4(t *testing.T) {
//	type testCase struct {
//		str   string
//		sep   string
//		wangt []string
//	}
//	testGroup := []testCase{
//		{str: "abcdbef", sep: "b", wangt: []string{"a", "cd", "ef"}},
//		{str: "a:b:c", sep: ":", wangt: []string{"a", "b", "c"}},
//		{str: "abcdef", sep: "bc", wangt: []string{"a", "def"}},
//		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
//	}
//	for _, tc := range testGroup {
//		got := SplitString(tc.str, tc.sep)
//		if !reflect.DeepEqual(got, tc.wangt) {
//			t.Fatalf("want : %v but got : %v", tc.wangt, got)
//		}
//	}
//}

// 子测试
// 不需要重复写了,相当于C语言的函数指针这种
func TestSplitString4(t *testing.T) {
	type testCase struct {
		str   string
		sep   string
		wangt []string
	}
	testGroup := map[string]testCase{
		"case1" : {str: "abcdbef", sep: "b", wangt: []string{"a", "cd", "ef"}},
		"case2" : {str: "a:b:c", sep: ":", wangt: []string{"a", "b", "c"}},
		"case3" : {str: "abcdef", sep: "bc", wangt: []string{"a", "def"}},
		"case4" : {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	// 切片
	//testGroup := []testCase{
	//	{str: "abcdbef", sep: "b", wangt: []string{"a", "cd", "ef"}},
	//	{str: "a:b:c", sep: ":", wangt: []string{"a", "b", "c"}},
	//	{str: "abcdef", sep: "bc", wangt: []string{"a", "def"}},
	//	{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	//}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := SplitString(tc.str,tc.sep)
			if !reflect.DeepEqual(got,tc.wangt) {
				t.Fatalf("want : %#v but got : %#v",tc.wangt,got)
			}
		})
	}
}

// 基准测试
func BenchmarkSplitString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitString("a:b:c:d:c",":")
	}
}

