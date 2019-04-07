package foldershare

import "testing"

func TestTimeEqual(t *testing.T) {
	type testCase struct {
		aSimpleTime FileTime
		bSimpleTime FileTime
		equal       bool
	}

	testCases := []testCase{
		{
			FileTime{},
			FileTime{},
			true,
		},
		{
			FileTime{2018, 10, 24, 12, 00, 00, 00},
			FileTime{2019, 10, 24, 12, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 12, 00, 00, 00},
			FileTime{2018, 11, 24, 12, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 12, 00, 00, 00},
			FileTime{2018, 10, 25, 12, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 10, 00, 00, 00},
			FileTime{2018, 10, 24, 11, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 10, 00, 00, 00},
			FileTime{2018, 10, 24, 10, 15, 00, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 10, 00, 00, 00},
			FileTime{2018, 10, 24, 10, 00, 30, 00},
			false,
		},
		{
			FileTime{2018, 10, 24, 10, 00, 00, 00},
			FileTime{2018, 10, 24, 10, 00, 00, 50},
			false,
		},
	}

	for _, tt := range testCases {
		if got := tt.aSimpleTime.Equal(tt.bSimpleTime); got != tt.equal {
			t.Errorf("%v.Equal(%v):%v, WANT:%v", tt.aSimpleTime, tt.bSimpleTime, got, tt.equal)
		}
	}
}

func TestTimeIsNewer(t *testing.T) {
	type testCase struct {
		a      FileTime
		b      FileTime
		aNewer bool
	}

	testCases := []testCase{
		{
			FileTime{},
			FileTime{},
			false,
		},
		{
			FileTime{2019, 01, 01, 10, 00, 00, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 02, 01, 10, 00, 00, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 01, 05, 10, 00, 00, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 01, 01, 13, 00, 00, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 01, 01, 10, 30, 00, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 01, 01, 10, 00, 30, 00},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2018, 01, 01, 10, 00, 00, 50},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			true,
		},
		{
			FileTime{2017, 03, 01, 10, 00, 00, 50},
			FileTime{2018, 01, 01, 10, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 10, 00, 00, 50},
			FileTime{2018, 05, 01, 10, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 10, 00, 00, 50},
			FileTime{2018, 03, 30, 10, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 10, 00, 00, 50},
			FileTime{2018, 03, 01, 12, 00, 00, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 12, 00, 00, 50},
			FileTime{2018, 03, 01, 12, 50, 00, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 12, 50, 00, 50},
			FileTime{2018, 03, 01, 12, 50, 30, 00},
			false,
		},
		{
			FileTime{2018, 03, 01, 12, 50, 30, 50},
			FileTime{2018, 03, 01, 12, 50, 30, 55},
			false,
		},
	}

	for _, tt := range testCases {
		if got := tt.a.IsNewer(tt.b); got != tt.aNewer {
			t.Errorf("%v.IsNewer(%v):%v, WANT:%v", tt.a, tt.b, got, tt.aNewer)
		}
	}
}

