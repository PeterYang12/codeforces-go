// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc110/submit?taskScreenName=arc110_a
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`3`,
			`7`,
		},
		{
			`10`,
			`39916801`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc110/tasks/arc110_a