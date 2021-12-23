package circle

import (
	"testing"
)

func checkEquality(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCircle(t *testing.T) {
	testCases := []struct {
		name     string
		left     int
		top      int
		square   int
		number   int
		expected interface{}
	}{
		{
			name:     "Calculate Square",
			left:     2,
			top:      5,
			square:   0,
			number:   0,
			expected: 10,
		},
		{
			name:     "Calculate Circle Length",
			left:     0,
			top:      0,
			square:   15,
			number:   0,
			expected: 13.729368492956533,
		},
		{
			name:     "Calculate Circle Dia",
			left:     0,
			top:      0,
			square:   15,
			number:   0,
			expected: 4.370193722368317,
		},
		{
			name:     "Divide By Digits",
			left:     0,
			top:      0,
			square:   0,
			number:   321,
			expected: []int{3, 2, 1},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			switch test.name {
			case "Calculate Square":
				if calcSquare(test.left, test.top) != test.expected {
					t.Errorf("Expected %d got %d", test.expected, calcSquare(test.left, test.top))
				}
				break
			case "Calculate Circle Length":
				_, l := calcCircleDL(test.square)
				if l != test.expected {
					t.Errorf("Expected %f got %f", test.expected, l)
				}
				break
			case "Calculate Circle Dia":
				d, _ := calcCircleDL(test.square)
				if d != test.expected {
					t.Errorf("Expected %f got %f", test.expected, d)
				}
				break
			case "Divide By Digits":
				res := make([]int, 3, 3)
				res[0], res[1], res[2] = divideByDigits(test.number)
				if !checkEquality(res, test.expected.([]int)) {
					t.Errorf("Expected %s got %d,%d,%d", test.expected, res[0], res[1], res[2])
				}
				break
			}
		})
	}
}
