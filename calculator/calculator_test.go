package calculator_test

import (
	"calculator"
	"math"
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

func RunFunctionValidTestCases(tcs []testCase, fn func(a, b float64) (float64, error), t *testing.T) {
	fnName := GetFunctionName(fn)
	for _, tc := range tcs {
		got, err := fn(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("%s(%f, %f): want %f, got %f",
				fnName, tc.a, tc.b, tc.want, got)
		}
	}
}

func RunFunctionInvalidTestCases(tcs []testCase, fn func(a, b float64) (float64, error), t *testing.T) {
	fnName := GetFunctionName(fn)
	for _, tc := range tcs {
		_, err := fn(tc.a, tc.b)
		if err == nil {
			t.Errorf("%s(%f, %f): want error for invalid input, got nil",
				fnName, tc.a, tc.b)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
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
	RunFunctionValidTestCases(testCases, calculator.Divide, t)
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	// want is meaningless here, so setting to 0
	testCases := []testCase{
		{a: 2, b: 0, want: 0},
	}
	RunFunctionInvalidTestCases(testCases, calculator.Divide, t)
}

func TestValidSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, want: 4},
		{a: 3, want: 9},
		{a: 4, want: 16},
		{a: 0.2, want: 0.04},
	}
	for _, tc := range testCases {
		got, _ := calculator.Sqrt(tc.a)
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("%s(%f): want %f, got %f",
				"calculator.Sqrt", tc.a, tc.want, got)
		}
	}
}

func TestInvalidSqrt(t *testing.T) {
	t.Parallel()
	invalidTestCases := []testCase{
		{a: -3, want: 0},
	}
	for _, tc := range invalidTestCases {
		_, err := calculator.Sqrt(tc.a)
		if err == nil {
			t.Errorf("%s(%f): want error for invalid input, got nil",
				"calculator.Sqrt", tc.a)
		}
	}
}
