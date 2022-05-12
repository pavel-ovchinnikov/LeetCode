package length_of_last_word_0053

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InputData struct {
	str string
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
			InputData{"   fly me   to   the moon  "},
			OutputData{4},
		},
		{
			InputData{"Hello World"},
			OutputData{5},
		},
		{
			InputData{"luffy is still joyboy"},
			OutputData{6},
		},
		{
			InputData{" Hello"},
			OutputData{5},
		},
		{
			InputData{"World"},
			OutputData{5},
		},
		{
			InputData{"  "},
			OutputData{0},
		},
	}

	for _, test := range tastCase {
		in, out := test.in, test.out
		ast.Equal(out.result, lengthOfLastWord(in.str), "Input data:%v", in)
	}
}
