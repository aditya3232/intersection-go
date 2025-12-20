package intersection

import "testing"

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Data normal ada intersection",
			input:    []string{"1,3,4,7,13", "1,2,4,13,15"},
			expected: "1,4,13",
		},
		{
			name:     "Tidak ada intersection",
			input:    []string{"1,2,3", "4,5,6"},
			expected: "",
		},
		{
			name:     "Satu list kosong",
			input:    []string{"", "1,2,3"},
			expected: "",
		},
		{
			name:     "Dua list kosong",
			input:    []string{"", ""},
			expected: "",
		},
		{
			name:     "Angka dengan spasi",
			input:    []string{"1, 2, 3", "2, 3, 4"},
			expected: "2,3",
		},
		{
			name:     "Hanya satu angka yang sama",
			input:    []string{"10,20,30", "5,7,10,99"},
			expected: "10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindIntersection(tt.input)
			if result != tt.expected {
				t.Errorf("%s: expected %q, got %q", tt.name, tt.expected, result)
			}
		})
	}
}
