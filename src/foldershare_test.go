package foldershare

import (
	"testing"
)

func TestEqual(t *testing.T) {
	type testCase = struct {
		local  FileInfo
		remote FileInfo
		eaual  bool
	}

	dateOne := FileTime{2017, 1, 1, 0, 0, 0, 0}
	dateSecond := FileTime{2018, 1, 1, 0, 0, 0, 0}

	testCases := []testCase{
		{
			local:  FileInfo{},
			remote: FileInfo{},
			eaual:  true,
		},
		{
			local:  FileInfo{"noname", dateOne, dateOne},
			remote: FileInfo{"noname", dateOne, dateOne},
			eaual:  false,
		},
		{
			local:  FileInfo{"noname", dateOne, dateOne},
			remote: FileInfo{"noname", dateOne, dateOne},
			eaual:  false,
		},
		{
			local:  FileInfo{"name_a", dateOne, dateOne},
			remote: FileInfo{"name_b", dateOne, dateOne},
			eaual:  false,
		},
		{
			local:  FileInfo{"name_a", dateOne, dateOne},
			remote: FileInfo{"name_a", dateSecond, dateOne},
			eaual:  false,
		},
		{
			local:  FileInfo{"name_a", dateOne, dateOne},
			remote: FileInfo{"name_a", dateOne, dateSecond},
			eaual:  false,
		},
	}

	for _, tt := range testCases {
		if got := Equal(tt.local, tt.remote); got != tt.eaual {
			t.Errorf("Equal(%v, %v): %v, WANT: %v", tt.local, tt.remote, got, tt.eaual)
		}
	}
}

func TestIsNewer(t *testing.T) {
	type testCase = struct {
		a       FileInfo
		b       FileInfo
		isNewer bool
	}

	aDate := FileTime{2017, 1, 1, 0, 0, 0, 0}
	bDate := FileTime{2018, 1, 1, 0, 0, 0, 0}

	testCases := []testCase{
		{
			FileInfo{},
			FileInfo{},
			false,
		},
		{
			FileInfo{"noname", aDate, bDate},
			FileInfo{"noname", aDate, aDate},
			true,
		},
		{
			FileInfo{"noname", aDate, bDate},
			FileInfo{"noname", aDate, aDate},
			true,
		},
		{
			FileInfo{"noname", bDate, aDate},
			FileInfo{"noname", aDate, aDate},
			true,
		},
	}

	for _, tt := range testCases {
		if got := tt.a.IsNewer(tt.b); got != tt.isNewer {
			t.Errorf("%v.IsNewer(%v):%v, WANT:%v", tt.a, tt.b, got, tt.isNewer)
		}
	}
}

