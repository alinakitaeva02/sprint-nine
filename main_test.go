package main

import "testing"

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestGenerateRandomElements_Size(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	if len(result) != size {
		t.Errorf("expected slice length %d, got %d", size, len(result))
	}
}

func TestGenerateRandomElements_Range(t *testing.T) {
	size := 100
	result := generateRandomElements(size)

	for _, v := range result {
		if v < 1 || v > 100_000_000 {
			t.Errorf("value out of range: %d", v)
		}
	}
}

func TestGenerateRandomElements_PositiveNumber(t *testing.T) {
	result := generateRandomElements(-15)

	if result != nil {
		t.Errorf("expected nil for negative size, got %d", result)
	}
}

func TestMaximum_ZeroSize(t *testing.T) {
	result := maximum(nil)

	if result != 0 {
		t.Errorf("expectged 0 for empty slice, got %d", result)
	}
}

func TestMaximum_OneElement(t *testing.T) {
	data := []int{55}
	result := maximum(data)

	if result != 55 {
		t.Errorf("expected %d, got %d", 55, result)
	}
}

func TestMaximum_MultipleElements(t *testing.T) {
	result := maximum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	if result != 10 {
		t.Errorf("expected maximum value %d, got %d", 10, result)
	}
}

func TestMaxChunks_EmptySlice(t *testing.T) {
	var data []int
	result := maxChunks(data)

	if result != 0 {
		t.Errorf("expected 0 for empty slice, got %d", result)
	}
}

func TestMaxChunks_OneElement(t *testing.T) {
	data := []int{42}
	result := maxChunks(data)

	if result != 42 {
		t.Errorf("expected %d, got %d", 42, result)
	}
}

func TestMaxChunks_SmallSlice(t *testing.T) {
	data := []int{3, 7, 2, 9, 5, 1}
	result := maxChunks(data)

	if result != 9 {
		t.Errorf("expected 9, got %d", result)
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
