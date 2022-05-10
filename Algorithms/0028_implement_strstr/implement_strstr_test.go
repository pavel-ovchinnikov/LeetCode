package implement_strstr_0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	haystack string
	needle   string
}

type OutputData struct {
	result int
}

type TestCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []TestCase{
		{
			InputData{"hello", "ll"},
			OutputData{2},
		},
		{
			InputData{"hello", "all"},
			OutputData{-1},
		},
	}

	for _, q := range qs {
		in, out := q.in, q.out
		ast.Equal(out.result, strStr(in.haystack, in.needle), "Input data:%v", in)
	}
}
