package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
		wantNil bool
	}{
		{
			name:    "zero size",
			size:    0,
			wantNil: true,
		},
		{
			name:    "negative size",
			size:    -15,
			wantNil: true,
		},
		{
			name:    "positive size",
			size:    10,
			wantLen: 10,
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if tt.wantNil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if len(result) != tt.wantLen {
				t.Errorf("expected slice length %d, got %d", tt.wantLen, len(result))
			}

			for _, v := range result {
				if v < 1 || v > 100_000_000 {
					t.Errorf("value out of range: %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty slice",
			input: nil,
			want:  0,
		},
		{
			name:  "one element",
			input: []int{55},
			want:  55,
		},
		{
			name:  "multiple elements",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)

			if result != tt.want {
				t.Errorf("expected %d, got %d", tt.want, result)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty slice",
			input: nil,
			want:  0,
		},
		{
			name:  "one element",
			input: []int{42},
			want:  42,
		},
		{
			name:  "small slice",
			input: []int{3, 7, 2, 9, 5, 1},
			want:  9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)

			if result != tt.want {
				t.Errorf("expected %d, got %d", tt.want, result)
			}
		})
	}
}

func TestMaxChunks_EqualsMaximum(t *testing.T) {
	data := generateRandomElements(1_000_000)
	expected := maximum(data)
	result := maxChunks(data)

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
