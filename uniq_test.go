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
	t.Run("testing preparing", func(t *testing.T) {
		var stdin = Std{
			boolFlags: map[string]bool{},
			intFlags:  map[string]uint{},
			data:      []string{},
		}
		stdin.intFlags["f"], stdin.intFlags["s"], stdin.boolFlags["i"] = 1, 2, true
		stdin.data = append(stdin.data,
			" UI OP ui OP",
			"UI OP ui OP",
			"asd Fgh,!!! jkL",
			"asd Fgh,!!! jkL",
			"")

		result := make([]El, 0)
		result = append(result,
			El{" UI OP ui OP", 1},
			El{"UI OP ui OP", 1},
			El{"asd Fgh,!!! jkL", 2},
			El{"", 1})

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

/*
этот тест не получается, хз почему разные ссылки у двух переменных, ссылающихся на одну функцию
*/

//func TestProgramBehaviour(t *testing.T) {
//	t.Run("testing prog behaviour with c param", func(t *testing.T) {
//		boolFlags := map[string]bool{"c": true}
//
//		result := strWithNumOfRepeat
//		realResult := programBehaviour(boolFlags)
//
//		if &result != &realResult {
//			t.Errorf("\nExpected\n    %v\nto equal\n    %v", &result, &realResult)
//		}
//	})
//}
