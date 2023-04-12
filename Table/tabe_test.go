package table

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		description string
		num1        int
		num2        int
		expected    int
	}{
		{
			description: "1 + 2",
			num1:        1,
			num2:        2,
			expected:    3,
		},
		{
			description: "3 + 4",
			num1:        3,
			num2:        4,
			expected:    7,
		},
		{
			description: "10 + 45",
			num1:        10,
			num2:        45,
			expected:    70,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := Sum(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
