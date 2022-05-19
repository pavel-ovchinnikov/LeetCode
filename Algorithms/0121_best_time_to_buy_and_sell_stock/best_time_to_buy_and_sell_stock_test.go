package best_time_to_buy_and_sell_stock_0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	value []int
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

	testCase := []TestCase{
		{
			InputData{[]int{7, 1, 5, 3, 6, 4}},
			OutputData{5},
		},
		{
			InputData{[]int{7, 6, 4, 3, 1}},
			OutputData{0},
		},
	}

	for _, test := range testCase {
		in, out := test.in, test.out

		ast.Equal(out.result, maxProfit(in.value), "Input data:%v", in)
	}
}
