package main

import "testing"

func TestValidate(t *testing.T) {
	t.Run("test with f param", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "qwerty ui op"
		intFlags["f"] = 1
		result := "ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with s param", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "qwerty ui op"
		intFlags["s"] = 5
		result := "y ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with i param", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "QWERTY ui op"
		boolFlags["i"] = true
		result := "qwerty ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with f and s params", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "QWERTY ui op"
		intFlags["f"], intFlags["s"] = 1, 2
		result := " op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with f and i params", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "qwerty UI OP"
		intFlags["f"], boolFlags["i"] = 1, true
		result := "ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with s and i params", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "qwErTY UI OP"
		intFlags["s"], boolFlags["i"] = 2, true
		result := "erty ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
	t.Run("test with s, f and i params", func(t *testing.T) {
		intFlags, boolFlags, str := map[string]uint{}, map[string]bool{}, "qwErTY UI OP ui OP"
		intFlags["s"], intFlags["f"], boolFlags["i"] = 2, 1, true
		result := " op ui op"
		realResult := validate(str, intFlags, boolFlags)

		if realResult != result || len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}
	})
}

func TestPreparing(t *testing.T) {
	t.Run("testing preparing without params", func(t *testing.T) {
		var stdin = InputParams{
			boolFlags: map[string]bool{},
			intFlags:  map[string]uint{},
			data: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		}

		result := []PreparedForOutput{
			PreparedForOutput{"I love music.", 3},
			PreparedForOutput{"", 1},
			PreparedForOutput{"I love music of Kartik.", 2},
			PreparedForOutput{"Thanks.", 1},
			PreparedForOutput{"I love music of Kartik.", 2},
		}

		realResult := preparing(stdin)

		if len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %v\nto equal\n    %v", result, realResult)
		}

		for ind, _ := range realResult {
			if result[ind] != realResult[ind] {
				t.Errorf("\nExpected\n    %v\nto equal\n    %v", result[ind], realResult[ind])
			}
		}
	})
	t.Run("testing preparing with i param", func(t *testing.T) {
		var stdin = InputParams{
			boolFlags: map[string]bool{
				"i": true,
			},
			intFlags: map[string]uint{},
			data: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
		}

		result := []PreparedForOutput{
			PreparedForOutput{"I LOVE MUSIC.", 3},
			PreparedForOutput{"", 1},
			PreparedForOutput{"I love MuSIC of Kartik.", 2},
			PreparedForOutput{"Thanks.", 1},
			PreparedForOutput{"I love music of kartik.", 2},
		}

		realResult := preparing(stdin)

		if len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %v\nto equal\n    %v", result, realResult)
		}

		for ind, _ := range realResult {
			if result[ind] != realResult[ind] {
				t.Errorf("\nExpected\n    %v\nto equal\n    %v", result[ind], realResult[ind])
			}
		}
	})
	t.Run("testing preparing with f param", func(t *testing.T) {
		var stdin = InputParams{
			boolFlags: map[string]bool{},
			intFlags: map[string]uint{
				"f": 1,
			},
			data: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		}

		result := []PreparedForOutput{
			PreparedForOutput{"We love music.", 3},
			PreparedForOutput{"", 1},
			PreparedForOutput{"I love music of Kartik.", 2},
			PreparedForOutput{"Thanks.", 1},
		}

		realResult := preparing(stdin)

		if len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %v\nto equal\n    %v", result, realResult)
		}

		for ind, _ := range realResult {
			if result[ind] != realResult[ind] {
				t.Errorf("\nExpected\n    %v\nto equal\n    %v", result[ind], realResult[ind])
			}
		}
	})
	t.Run("testing preparing with f param", func(t *testing.T) {
		var stdin = InputParams{
			boolFlags: map[string]bool{},
			intFlags: map[string]uint{
				"s": 1,
			},
			data: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		}

		result := []PreparedForOutput{
			PreparedForOutput{"I love music.", 3},
			PreparedForOutput{"", 1},
			PreparedForOutput{"I love music of Kartik.", 1},
			PreparedForOutput{"We love music of Kartik.", 1},
			PreparedForOutput{"Thanks.", 1},
		}

		realResult := preparing(stdin)

		if len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %v\nto equal\n    %v", result, realResult)
		}

		for ind, _ := range realResult {
			if result[ind] != realResult[ind] {
				t.Errorf("\nExpected\n    %v\nto equal\n    %v", result[ind], realResult[ind])
			}
		}
	})
	t.Run("testing preparing with i, f, s params", func(t *testing.T) {
		var stdin = InputParams{
			boolFlags: map[string]bool{
				"i": true,
			},
			intFlags: map[string]uint{
				"f": 1,
				"s": 2,
			},
			data: []string{" UI OP ui OP",
				"UI OP ui OP",
				"asd Fgh,!!! jkL",
				"asd Fgh,!!! jkL",
				""},
		}

		result := []PreparedForOutput{
			PreparedForOutput{" UI OP ui OP", 1},
			PreparedForOutput{"UI OP ui OP", 1},
			PreparedForOutput{"asd Fgh,!!! jkL", 2},
			PreparedForOutput{"", 1},
		}

		realResult := preparing(stdin)

		if len(result) != len(realResult) {
			t.Errorf("\nExpected\n    %v\nto equal\n    %v", result, realResult)
		}

		for ind, _ := range realResult {
			if result[ind] != realResult[ind] {
				t.Errorf("\nExpected\n    %v\nto equal\n    %v", result[ind], realResult[ind])
			}
		}
	})
}
