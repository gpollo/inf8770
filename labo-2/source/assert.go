package main

import "testing"

func assert2DFloat32ArrayEqual(t *testing.T, got, expected [][]float32) {
	if len(got) != len(expected) {
		t.Fatalf("array Y dimension differs: got=%d, expected=%d", len(got), len(expected))
	}

	if len(got) == 0 {
		return
	}

	if len(got[0]) != len(expected[0]) {
		t.Fatalf("array X dimension differs: got=%d, expected=%d", len(got[0]), len(expected[0]))
	}

	for j := 0; j < len(got); j++ {
		for i := 0; i < len(got[0]); i++ {
			if got[j][i] == expected[j][i] {
				continue
			}

			t.Errorf("value at (%d, %d) differs: got=%f, expected=%f", i, j, got[j][i], expected[j][i])
		}
	}
}
