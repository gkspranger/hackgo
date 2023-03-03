package calculator_test

import (
	"calculator"
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func RunFunctionTestCases(tcs []testCase, fn func(a, b float64) float64, t *testing.T) {
	fnName := GetFunctionName(fn)
	for _, tc := range tcs {
		got := fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s(%f, %f): want %f, got %f",
				fnName, tc.a, tc.b, tc.want, got)
		}
	}
}

func RunFunctionTestCasesWithErrors(tcs []testCase, fn func(a, b float64) (float64, error), t *testing.T) {
	fnName := GetFunctionName(fn)
	for _, tc := range tcs {
		got, err := fn(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("%s(%f, %f): want %f, got %f",
				fnName, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 1, want: 3},
		{a: 3, b: 2, want: 5},
		{a: 4, b: 3, want: 7},
	}
	RunFunctionTestCases(testCases, calculator.Add, t)
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 1, want: 1},
		{a: -3, b: -2, want: -1},
		{a: -4, b: 3, want: -7},
	}
	RunFunctionTestCases(testCases, calculator.Subtract, t)
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 1, want: 2},
		{a: -3, b: -2, want: 6},
		{a: -4, b: 3, want: -12},
	}
	RunFunctionTestCases(testCases, calculator.Multiply, t)
}

func TestDivideValidTestCases(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -3, b: 1, want: -3},
		{a: 12, b: 3, want: 4},
	}
	RunFunctionTestCasesWithErrors(testCases, calculator.Divide, t)
}
