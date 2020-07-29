package SecondTask

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"nil case",
			nil,
			nil,
		},
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			[]int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice1(&tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", tc.input, tc.expected)
			}
		})
	}
}

func TestAppendFive(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"nil case",
			nil,
			[]int{5},
		},
		{
			"empty case",
			[]int{},
			[]int{5},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			[]int{0, 1, 2, 3, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice2(&tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestPrependFive(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"nil case",
			nil,
			[]int{5},
		},
		{
			"empty case",
			[]int{},
			[]int{5},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			[]int{5, 0, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice3(5, &tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type result struct {
		element int
		slice   []int
	}
	testCases := []struct {
		name     string
		input    []int
		expected result
	}{
		{
			"single case",
			[]int{5},
			result{5, []int{}},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			result{3, []int{0, 1, 2}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			elem := slice4(&tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected.slice) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected.slice)
			}

			if !reflect.DeepEqual(elem, tc.expected.element) {
				t.Errorf("got %v want %v", elem, tc.expected.element)
			}
		})
	}
}

func TestShift(t *testing.T) {
	type result struct {
		element int
		slice   []int
	}
	testCases := []struct {
		name     string
		input    []int
		expected result
	}{
		{
			"single case",
			[]int{5},
			result{5, []int{}},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			result{0, []int{1, 2, 3}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			elem := slice5(&tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected.slice) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected.slice)
			}

			if !reflect.DeepEqual(elem, tc.expected.element) {
				t.Errorf("got %v want %v", elem, tc.expected.element)
			}
		})
	}
}

func TestPopIndex(t *testing.T) {
	type result struct {
		element int
		slice   []int
	}
	testCases := []struct {
		name     string
		input    []int
		inputIdx int
		expected result
	}{
		{
			"single case",
			[]int{5},
			0,
			result{5, []int{}},
		},
		{
			"zero index. even length",
			[]int{0, 1, 2, 3},
			0,
			result{0, []int{1, 2, 3}},
		},
		{
			"less than middle index. even length",
			[]int{0, 1, 2, 3},
			1,
			result{1, []int{0, 2, 3}},
		},
		{
			"middle index. even length",
			[]int{0, 1, 2, 3},
			2,
			result{2, []int{0, 1, 3}},
		},
		{
			"last index. even length",
			[]int{0, 1, 2, 3},
			3,
			result{3, []int{0, 1, 2}},
		},

		{
			"zero index. odd length",
			[]int{0, 1, 2, 3, 4},
			0,
			result{0, []int{1, 2, 3, 4}},
		},
		{
			"less than middle index. even length",
			[]int{0, 1, 2, 3, 4},
			1,
			result{1, []int{0, 2, 3, 4}},
		},
		{
			"middle index. even length",
			[]int{0, 1, 2, 3, 4},
			2,
			result{2, []int{0, 1, 3, 4}},
		},
		{
			"greater than middle index. even length",
			[]int{0, 1, 2, 3, 4},
			3,
			result{3, []int{0, 1, 2, 4}},
		},
		{
			"last index. even length",
			[]int{0, 1, 2, 3},
			3,
			result{3, []int{0, 1, 2}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			elem := slice6(tc.inputIdx, &tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected.slice) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected.slice)
			}

			if !reflect.DeepEqual(res, tc.expected.slice) {
				t.Errorf("got %v want %v", elem, tc.expected.element)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type input struct {
		slice1, slice2 []int
	}
	testCases := []struct {
		name     string
		in       input
		expected []int
	}{
		{
			"first nil case",
			input{nil, []int{1, 2}},
			[]int{1, 2},
		},
		{
			"second nil case",
			input{[]int{1, 2}, nil},
			[]int{1, 2},
		},
		{
			"first empty case",
			input{[]int{}, []int{1, 2}},
			[]int{1, 2},
		},
		{
			"second empty case",
			input{[]int{1, 2}, []int{}},
			[]int{1, 2},
		},
		{
			"basic case",
			input{[]int{0, 1, 2, 3}, []int{3, 4, 5}},
			[]int{0, 1, 2, 3, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := slice7(&tc.in.slice1, &tc.in.slice2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type input struct {
		slice1, slice2 []int
	}
	testCases := []struct {
		name     string
		in       input
		expected []int
	}{
		{
			"first nil case",
			input{nil, []int{1, 2}},
			[]int{},
		},
		{
			"second nil case",
			input{[]int{1, 2}, nil},
			[]int{1, 2},
		},
		{
			"first empty case",
			input{[]int{}, []int{1, 2}},
			[]int{},
		},
		{
			"second empty case",
			input{[]int{1, 2}, []int{}},
			[]int{1, 2},
		},
		{
			"basic case",
			input{[]int{0, 1, 2, 3}, []int{3, 4, 5}},
			[]int{0, 1, 2},
		},
		{
			"duplicates",
			input{[]int{0, 1, 1, 2, 2, 3}, []int{1, 2, 3}},
			[]int{0},
		},
		{
			"preserve order",
			input{[]int{0, 3, 10, 1, 2, 1, 2, -1}, []int{-1, 2, 3}},
			[]int{0, 10, 1, 1},
		},
		{
			"nothing to delete",
			input{[]int{0, 1, 2, 3}, []int{4, 5}},
			[]int{0, 1, 2, 3},
		},
		{
			"everything to delete",
			input{[]int{0, 1, 1, 2, 2, 3}, []int{0, 1, 2, 3}},
			[]int{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice8(&tc.in.slice1, &tc.in.slice2)
			res := tc.in.slice1
			if tc.expected == nil && res != nil {
				t.Fatalf("\ngot\t\t%v\nwant\tnil slice", res)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestRemoveInplace(t *testing.T) {
	type input struct {
		slice1, slice2 []int
	}
	testCases := []struct {
		name     string
		in       input
		expected []int
	}{
		{
			"first nil case",
			input{nil, []int{1, 2}},
			nil,
		},
		{
			"second nil case",
			input{[]int{1, 2}, nil},
			[]int{1, 2},
		},
		{
			"first empty case",
			input{[]int{}, []int{1, 2}},
			[]int{},
		},
		{
			"second empty case",
			input{[]int{1, 2}, []int{}},
			[]int{1, 2},
		},
		{
			"basic case",
			input{[]int{0, 1, 2, 3}, []int{3, 4, 5}},
			[]int{0, 1, 2},
		},
		{
			"duplicates",
			input{[]int{0, 1, 1, 2, 2, 3}, []int{1, 2, 3}},
			[]int{0},
		},
		{
			"preserve order",
			input{[]int{0, 3, 10, 1, 2, 1, 2, -1}, []int{-1, 2, 3}},
			[]int{0, 1, 1, 10},
		},
		{
			"nothing to delete",
			input{[]int{0, 1, 2, 3}, []int{4, 5}},
			[]int{0, 1, 2, 3},
		},
		{
			"everything to delete",
			input{[]int{0, 1, 1, 2, 2, 3}, []int{0, 1, 2, 3}},
			[]int{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice8(&tc.in.slice1, &tc.in.slice2)
			res := tc.in.slice1
			sort.Ints(res)
			if tc.expected == nil && res != nil {
				t.Fatalf("\ngot\t\t%v\nwant\tnil slice", res)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestOffsetLeftOne(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			"nil case",
			nil,
			nil,
		},
		{
			"first empty case",
			[]int{},
			[]int{},
		},
		{
			"single element",
			[]int{3},
			[]int{3},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			[]int{1, 2, 3, 0},
		},
		{
			"duplicates",
			[]int{0, 1, 1, 2, 2, 3},
			[]int{1, 1, 2, 2, 3, 0},
		},
		{
			"preserve order",
			[]int{0, 3, 10, 1, 2, 1, 2, -1},
			[]int{3, 10, 1, 2, 1, 2, -1, 0},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice9(&tc.in)
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", tc.in, tc.expected)
			}
		})
	}
}

func TestOffsetRightOne(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			"nil case",
			nil,
			nil,
		},
		{
			"first empty case",
			[]int{},
			[]int{},
		},
		{
			"single element",
			[]int{3},
			[]int{3},
		},
		{
			"basic case",
			[]int{0, 1, 2, 3},
			[]int{3, 0, 1, 2},
		},
		{
			"duplicates",
			[]int{0, 1, 1, 2, 2, 3},
			[]int{3, 0, 1, 1, 2, 2},
		},
		{
			"preserve order",
			[]int{0, 3, 10, 1, 2, 1, 2, -1},
			[]int{-1, 0, 3, 10, 1, 2, 1, 2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice11(&tc.in)
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", tc.in, tc.expected)
			}
		})
	}
}

func TestOffsetLeft(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		offset   uint64
		expected []int
	}{
		{
			"nil case. zero offset",
			nil,
			0,
			nil,
		},
		{
			"nil case. 1 offset",
			nil,
			1,
			nil,
		},
		{
			"nil case. 10 offset",
			nil,
			10,
			nil,
		},

		{
			"empty case. zero offset",
			[]int{},
			0,
			[]int{},
		},
		{
			"empty case. 1 offset",
			[]int{},
			1,
			[]int{},
		},
		{
			"empty case. 10 offset",
			[]int{},
			10,
			[]int{},
		},

		{
			"single element. zero offset",
			[]int{3},
			0,
			[]int{3},
		},
		{
			"single element. 1 offset",
			[]int{3},
			1,
			[]int{3},
		},
		{
			"single element. 10 offset",
			[]int{3},
			10,
			[]int{3},
		},

		{
			"even length. zero offset",
			[]int{0, 1, 2, 3},
			0,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. 1 offset",
			[]int{0, 1, 2, 3},
			1,
			[]int{1, 2, 3, 0},
		},
		{
			"even length. the offset is less than the half",
			[]int{0, 1, 2, 3},
			2,
			[]int{2, 3, 0, 1},
		},
		{
			"even length. the offset is the half",
			[]int{0, 1, 2, 3},
			3,
			[]int{3, 0, 1, 2},
		},
		{
			"even length. the offset is more than the half",
			[]int{0, 1, 2, 3},
			4,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. the offset is N*len(slice)",
			[]int{0, 1, 2, 3},
			16,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. the offset is N*len(slice)+2",
			[]int{0, 1, 2, 3},
			18,
			[]int{2, 3, 0, 1},
		},

		{
			"odd length. zero offset",
			[]int{0, 1, 2, 3, 4, 5, 6},
			0,
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			"odd length. 1 offset",
			[]int{0, 1, 2, 3, 4, 5, 6},
			1,
			[]int{1, 2, 3, 4, 5, 6, 0},
		},
		{
			"odd length. the offset is less than the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			2,
			[]int{2, 3, 4, 5, 6, 0, 1},
		},
		{
			"odd length. the offset is the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			4,
			[]int{4, 5, 6, 0, 1, 2, 3},
		},
		{
			"odd length. the offset is more than the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			5,
			[]int{5, 6, 0, 1, 2, 3, 4},
		},
		{
			"odd length. the offset is N*len(slice)",
			[]int{0, 1, 2, 3, 4, 5, 6},
			14,
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			"odd length. the offset is N*len(slice)+2",
			[]int{0, 1, 2, 3, 4, 5, 6},
			16,
			[]int{2, 3, 4, 5, 6, 0, 1},
		},

		{
			"duplicates. zero offset",
			[]int{0, 1, 1, 3, 4, 5, 6},
			0,
			[]int{0, 1, 1, 3, 4, 5, 6},
		},
		{
			"duplicates. 1 offset",
			[]int{0, 1, 1, 3, 4, 5, 6},
			1,
			[]int{1, 1, 3, 4, 5, 6, 0},
		},
		{
			"duplicates. the offset is less than the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			2,
			[]int{1, 3, 4, 5, 6, 0, 1},
		},
		{
			"duplicates. the offset is the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			4,
			[]int{4, 5, 6, 0, 1, 1, 3},
		},
		{
			"duplicates. the offset is more than the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			5,
			[]int{5, 6, 0, 1, 1, 3, 4},
		},
		{
			"duplicates. the offset is N*len(slice)",
			[]int{0, 1, 1, 3, 4, 5, 6},
			14,
			[]int{0, 1, 1, 3, 4, 5, 6},
		},
		{
			"duplicates. the offset is N*len(slice)+2",
			[]int{0, 1, 1, 3, 4, 5, 6},
			16,
			[]int{1, 3, 4, 5, 6, 0, 1},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice10(&tc.in, int(tc.offset))
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", tc.in, tc.expected)
			}
		})
	}
}

func TestOffsetRight(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		offset   uint64
		expected []int
	}{
		{
			"nil case. zero offset",
			nil,
			0,
			nil,
		},
		{
			"nil case. 1 offset",
			nil,
			1,
			nil,
		},
		{
			"nil case. 10 offset",
			nil,
			10,
			nil,
		},

		{
			"empty case. zero offset",
			[]int{},
			0,
			[]int{},
		},
		{
			"empty case. 1 offset",
			[]int{},
			1,
			[]int{},
		},
		{
			"empty case. 10 offset",
			[]int{},
			10,
			[]int{},
		},

		{
			"single element. zero offset",
			[]int{3},
			0,
			[]int{3},
		},
		{
			"single element. 1 offset",
			[]int{3},
			1,
			[]int{3},
		},
		{
			"single element. 10 offset",
			[]int{3},
			10,
			[]int{3},
		},

		{
			"even length. zero offset",
			[]int{0, 1, 2, 3},
			0,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. 1 offset",
			[]int{0, 1, 2, 3},
			1,
			[]int{3, 0, 1, 2},
		},
		{
			"even length. the offset is less than the half",
			[]int{0, 1, 2, 3},
			2,
			[]int{2, 3, 0, 1},
		},
		{
			"even length. the offset is the half",
			[]int{0, 1, 2, 3},
			3,
			[]int{1, 2, 3, 0},
		},
		{
			"even length. the offset is more than the half",
			[]int{0, 1, 2, 3},
			4,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. the offset is N*len(slice)",
			[]int{0, 1, 2, 3},
			16,
			[]int{0, 1, 2, 3},
		},
		{
			"even length. the offset is N*len(slice)+2",
			[]int{0, 1, 2, 3},
			18,
			[]int{2, 3, 0, 1},
		},

		{
			"odd length. zero offset",
			[]int{0, 1, 2, 3, 4, 5, 6},
			0,
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			"odd length. 1 offset",
			[]int{0, 1, 2, 3, 4, 5, 6},
			1,
			[]int{6, 0, 1, 2, 3, 4, 5},
		},
		{
			"odd length. the offset is less than the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			2,
			[]int{5, 6, 0, 1, 2, 3, 4},
		},
		{
			"odd length. the offset is the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			4,
			[]int{3, 4, 5, 6, 0, 1, 2},
		},
		{
			"odd length. the offset is more than the half",
			[]int{0, 1, 2, 3, 4, 5, 6},
			5,
			[]int{2, 3, 4, 5, 6, 0, 1},
		},
		{
			"odd length. the offset is N*len(slice)",
			[]int{0, 1, 2, 3, 4, 5, 6},
			14,
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			"odd length. the offset is N*len(slice)+2",
			[]int{0, 1, 2, 3, 4, 5, 6},
			16,
			[]int{5, 6, 0, 1, 2, 3, 4},
		},

		{
			"duplicates. zero offset",
			[]int{0, 1, 1, 3, 4, 5, 6},
			0,
			[]int{0, 1, 1, 3, 4, 5, 6},
		},
		{
			"duplicates. 1 offset",
			[]int{0, 1, 1, 3, 4, 5, 6},
			1,
			[]int{6, 0, 1, 1, 3, 4, 5},
		},
		{
			"duplicates. the offset is less than the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			2,
			[]int{5, 6, 0, 1, 1, 3, 4},
		},
		{
			"duplicates. the offset is the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			4,
			[]int{3, 4, 5, 6, 0, 1, 1},
		},
		{
			"duplicates. the offset is more than the half",
			[]int{0, 1, 1, 3, 4, 5, 6},
			5,
			[]int{1, 3, 4, 5, 6, 0, 1},
		},
		{
			"duplicates. the offset is N*len(slice)",
			[]int{0, 1, 1, 3, 4, 5, 6},
			14,
			[]int{0, 1, 1, 3, 4, 5, 6},
		},
		{
			"duplicates. the offset is N*len(slice)+2",
			[]int{0, 1, 1, 3, 4, 5, 6},
			16,
			[]int{5, 6, 0, 1, 1, 3, 4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice12(&tc.in, int(tc.offset))
			if !reflect.DeepEqual(tc.in, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", tc.in, tc.expected)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{

		{
			"nil case",
			nil,
			[]int{},
		},
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"single case",
			[]int{5},
			[]int{5},
		},
		{
			"even length",
			[]int{0, 1, 2, 3},
			[]int{0, 1, 2, 3},
		},
		{
			"odd length",
			[]int{0, 1, 1, 2, 3},
			[]int{0, 1, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := slice13(&tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestEvenOddSwap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{

		{
			"nil case",
			nil,
			nil,
		},
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"single case",
			[]int{5},
			[]int{5},
		},
		{
			"even length",
			[]int{0, 1, 2, 3},
			[]int{1, 0, 3, 2},
		},
		{
			"odd length",
			[]int{0, 1, 1, 2, 3},
			[]int{1, 0, 2, 1, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			slice14(&tc.input)
			res := tc.input
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{

		{
			"nil case",
			nil,
			nil,
		},
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"single case",
			[]int{5},
			[]int{5},
		},

		{
			"even length. already ordered",
			[]int{0, 1, 2, 3},
			[]int{0, 1, 2, 3},
		},
		{
			"even length",
			[]int{2, 1, 0, 3},
			[]int{0, 1, 2, 3},
		},
		{
			"even length. duplicates",
			[]int{2, 0, 0, 3},
			[]int{0, 0, 2, 3},
		},
		{
			"even length. reverse order",
			[]int{3, 2, 1, 0},
			[]int{0, 1, 2, 3},
		},

		{
			"odd length. already ordered",
			[]int{0, 1, 2, 3, 4},
			[]int{0, 1, 2, 3, 4},
		},
		{
			"odd length",
			[]int{2, -1, 0, 4, 3},
			[]int{-1, 0, 2, 3, 4},
		},
		{
			"odd length. duplicates",
			[]int{2, 0, 0, 4, 3},
			[]int{0, 0, 2, 3, 4},
		},
		{
			"odd length. reverse order",
			[]int{4, 3, 2, 1, 0},
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, _, _ := slice15(&tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestSortReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{

		{
			"nil case",
			nil,
			nil,
		},
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"single case",
			[]int{5},
			[]int{5},
		},

		{
			"even length. already ordered",
			[]int{3, 2, 1, 0},
			[]int{3, 2, 1, 0},
		},
		{
			"even length",
			[]int{2, 1, 0, 3},
			[]int{3, 2, 1, 0},
		},
		{
			"even length. duplicates",
			[]int{2, 0, 0, 3},
			[]int{3, 2, 0, 0},
		},
		{
			"even length. reverse order",
			[]int{0, 1, 2, 3},
			[]int{3, 2, 1, 0},
		},

		{
			"odd length. already ordered",
			[]int{4, 3, 2, 1, 0},
			[]int{4, 3, 2, 1, 0},
		},
		{
			"odd length",
			[]int{2, -1, 0, 4, 3},
			[]int{4, 3, 2, 0, -1},
		},
		{
			"odd length. duplicates",
			[]int{2, 10, 0, 4, 3},
			[]int{10, 4, 3, 2, 0},
		},
		{
			"odd length. reverse order",
			[]int{0, 1, 2, 3, 4},
			[]int{4, 3, 2, 1, 0},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, res, _ := slice15(&tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\ngot\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}
