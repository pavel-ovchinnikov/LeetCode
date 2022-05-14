package climbing_stairs_0070

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
			InputData{2},
			OutputData{2},
		},
		{
			InputData{3},
			OutputData{3},
		},
	}

	for _, test := range tastCase {
		in, out := test.in, test.out
		ast.Equal(out.result, climbStairs(in.value), "Input data:%v", in)
	}
}
