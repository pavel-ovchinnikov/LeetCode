package sqrt_0069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value int
}

type OutputData struct {
	result int
}

type TastCase struct {
	in  InputData
	out OutputData
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	tastCase := []TastCase{
		{
			InputData{4},
			OutputData{2},
		},
		{
			InputData{8},
			OutputData{2},
		},
		{
			InputData{9},
			OutputData{3},
		},
		{
			InputData{11},
			OutputData{3},
		},
	}

	for _, test := range tastCase {
		in, out := test.in, test.out
		ast.Equal(out.result, mySqrt(in.value), "Input data:%v", in)
	}
}
