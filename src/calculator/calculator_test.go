package calculator_test

import (
	. "calculator"
	"testing"
)

func Test_Divide_Input_100_Divide_By_11_Should_Be_0(t *testing.T) {
	expected := 0
	calculator := Calculator{NumberOne: 100, NumberTwo: 11}
	resultNumber := calculator.Divide()
	isEqual(t, expected, resultNumber)
}

func Test_Divide_Input_10_Divide_By_0_Should_Be_0(t *testing.T) {
	expected := 0
	calculator := Calculator{NumberOne: 10, NumberTwo: 0}
	resultNumber := calculator.Divide()
	isEqual(t, expected, resultNumber)
}

func Test_Divide_Input_99_Divide_By_11_Should_Be_9(t *testing.T) {
	expected := 9
	calculator := Calculator{NumberOne: 99, NumberTwo: 11}
	resultNumber := calculator.Divide()
	isEqual(t, expected, resultNumber)
}

func Test_Multiply_Input_5_Multiply_By_10_Should_Be_50(t *testing.T) {
	expected := 50
	calculator := Calculator{NumberOne: 5, NumberTwo: 10}
	resultNumber := calculator.Multiply()
	isEqual(t, expected, resultNumber)
}

func Test_Multiply_Input_2_Multiply_By_3_Should_Be_6(t *testing.T) {
	expected := 6
	calculator := Calculator{NumberOne: 2, NumberTwo: 3}
	resultNumber := calculator.Multiply()
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_3_Plus_Minus_2_Should_Be_1(t *testing.T) {
	expected := 1
	calculator := Calculator{NumberOne: 3, NumberTwo: -2}
	resultNumber := calculator.Add()
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_Minus_1_Plus_1_Should_Be_0(t *testing.T) {
	expected := 0
	calculator := Calculator{NumberOne: -1, NumberTwo: 1}
	resultNumber := calculator.Add()
	isEqual(t, expected, resultNumber)
}

func Test_Add_Input_1_plus_1_Should_Be_2(t *testing.T) {
	expected := 2
	calculator := Calculator{NumberOne: 1, NumberTwo: 1}
	resultNumber := calculator.Add()
	isEqual(t, expected, resultNumber)
}

func isEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Fatalf("expected %d, but %d\n", expected, actual)
	}
}
