package set

import (
	"fmt"
	"testing"
)

type testCase struct {
	init     Set[int]
	query    Set[int]
	expected bool
}

func (c testCase) String() string {
	if c.expected {
		return fmt.Sprintf("A set containing %s, should have %s", c.init.String(), c.query.String())
	}
	return fmt.Sprintf("A set containing %s, should NOT have %s", c.init.String(), c.query.String())
}

var testCases = []testCase{
	// truthy
	{NewValues[int](), NewValues[int](), true},
	{NewValues(1, 2, 3), NewValues[int](), true},
	{NewValues(1, 2, 3), NewValues(1), true},
	{NewValues(1, 2, 3), NewValues(1, 2), true},
	{NewValues(1, 2, 3), NewValues(1, 2, 3), true},
	{NewValues(1, 2, 3), NewValues(2), true},
	{NewValues(1, 2, 3), NewValues(3), true},
	{NewValues(1, 2, 3), NewValues(2, 3), true},
	// falsy
	{NewValues(1, 2, 3), NewValues(4), false},
	{NewValues(1, 2, 3), NewValues(3, 4), false}}

func TestHas(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			if c.init.Has(c.query.ToSlice()...) != c.expected {
				t.Fail()
			}
		})
	}
}

type diffCase struct {
	a, b             Set[int]
	less, same, more Set[int]
}

func (c diffCase) String() string {
	return fmt.Sprintf("a contains %s, b contains %s. Expected %s, %s, %s", c.a, c.b, c.less, c.same, c.more)
}

func TestDiff(t *testing.T) {
	cases := []diffCase{
		{NewValues[int](), NewValues[int](), NewValues[int](), NewValues[int](), NewValues[int]()},
		{NewValues(1, 2, 3), NewValues[int](), NewValues(1, 2, 3), NewValues[int](), NewValues[int]()},
		{NewValues[int](), NewValues(4, 5), NewValues[int](), NewValues[int](), NewValues(4, 5)},
		{NewValues(1, 2, 3), NewValues(4, 5), NewValues(1, 2, 3), NewValues[int](), NewValues(4, 5)},
		{NewValues(1, 2, 3, 4), NewValues(4, 5), NewValues(1, 2, 3), NewValues(4), NewValues(5)}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			t.Logf("Testing that %s", c.String())
			less, same, more := c.a.Diff(c.b)

			if !less.Equals(c.less) {
				t.Errorf("expected 'less' to be %s, but was %s", c.less, less)
			}

			if !same.Equals(c.same) {
				t.Errorf("expected 'same' to be %s, but was %s", c.same, same)
			}

			if !more.Equals(c.more) {
				t.Errorf("expected 'more' to be %s, but was %s", c.more, more)
			}
		})
	}
}
