package calculator

import "testing"

func Test_Multiply_Input_5_Multiply_By_10_Should_Be_50(t *testing.T) {
	number1 := 5
	number2 := 10
	expected := 50
	resultNumber := Multiply(number1, number2)
	isEqual(t, expected, resultNumber)
}

func Test_Multiply_Input_2_Multiply_By_3_Should_Be_6(t *testing.T) {
	number1 := 2
	number2 := 3
	expected := 6
	resultNumber := Multiply(number1, number2)
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_3_Plus_Minus_2_Should_Be_1(t *testing.T) {
	number1 := 3
	number2 := -2
	expected := 1
	resultNumber := Add(number1, number2)
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_Minus_1_Plus_1_Should_Be_0(t *testing.T) {
	number1 := -1
	number2 := 1
	expected := 0
	resultNumber := Add(number1, number2)
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_1_plus_1_Should_Be_2(t *testing.T) {
	number1 := 1
	number2 := 1
	expected := 2
	resultNumber := Add(number1, number2)
	isEqual(t, expected, resultNumber)
}

func isEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("expected %d, but %d\n", expected, actual)
	}
}
