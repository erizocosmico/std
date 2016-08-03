package conv

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestToSliceOf(t *testing.T) {
	cases := []struct {
		from     interface{}
		to       interface{}
		expected interface{}
		err      bool
	}{
		{[]int{1, 2, 3}, int64(0), []int64{1, 2, 3}, false},
		{ // Deep is not allowed yet
			[][]int{[]int{1, 2}, []int{1, 2}},
			[]int64{},
			nil,
			true,
		},
		{45, 4, nil, true},
		{
			[]string{"a", "b"},
			([]byte)(nil),
			[][]byte{[]byte("a"), []byte("b")},
			false,
		},
	}

	for _, c := range cases {
		result, err := ToSliceOf(c.from, c.to)
		if (err != nil) != c.err {
			t.Fatalf("expected error %q to be %v", err, c.err)
		}

		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("expecting %#v to be %#v", result, c.expected)
		}
	}
}

var result1000000 interface{}
var baseSlice1000000 = func() []int {
	var res []int
	for i := 0; i < 1000000; i++ {
		res = append(res, i)
	}
	return res
}()

func BenchmarkToSliceOf1000000(b *testing.B) {
	var r interface{}
	var err error
	for n := 0; n < b.N; n++ {
		r, err = ToSliceOf(baseSlice1000000, int64(0))
		if err != nil {
			panic(err)
		}
	}
	result1000000 = r
}

var result100 interface{}
var baseSlice100 = func() []int {
	var res []int
	for i := 0; i < 100; i++ {
		res = append(res, i)
	}
	return res
}()

func BenchmarkToSliceOf100(b *testing.B) {
	var r interface{}
	var err error
	for n := 0; n < b.N; n++ {
		r, err = ToSliceOf(baseSlice100, int64(0))
		if err != nil {
			panic(err)
		}
	}
	result100 = r
}

var result100Str interface{}
var baseSlice100Str = func() []string {
	var res []string
	for i := 0; i < 100; i++ {
		var bytes = make([]byte, 100)
		rand.Read(bytes)
		res = append(res, string(bytes))
	}
	return res
}()

func BenchmarkToSliceOf100Str(b *testing.B) {
	var r interface{}
	var err error
	for n := 0; n < b.N; n++ {
		r, err = ToSliceOf(baseSlice100Str, ([]byte)(nil))
		if err != nil {
			panic(err)
		}
	}
	result100Str = r
}
