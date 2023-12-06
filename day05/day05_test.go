package day05_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brunorene/adventofcode2023/day05"
)

//nolint:funlen // it is a test func so it is ok to be long
func TestSplitRanges(t *testing.T) {
	t.Parallel()

	type args struct {
		start  int64
		end    int64
		ranges []day05.ValRange
	}

	tests := []struct {
		name       string
		args       args
		wantResult []day05.ValRange
	}{
		{
			name: "case 1",
			args: args{
				start: 2,
				end:   100,
				ranges: []day05.ValRange{
					{SourceStart: 5, SourceEnd: 10},
					{SourceStart: 25, SourceEnd: 30},
					{SourceStart: 15, SourceEnd: 20},
					{SourceStart: 50, SourceEnd: 60},
					{SourceStart: 61, SourceEnd: 80},
				},
			},
			wantResult: []day05.ValRange{
				{SourceStart: 2, SourceEnd: 4},
				{SourceStart: 5, SourceEnd: 10},
				{SourceStart: 11, SourceEnd: 14},
				{SourceStart: 15, SourceEnd: 20},
				{SourceStart: 21, SourceEnd: 24},
				{SourceStart: 25, SourceEnd: 30},
				{SourceStart: 31, SourceEnd: 49},
				{SourceStart: 50, SourceEnd: 60},
				{SourceStart: 61, SourceEnd: 80},
				{SourceStart: 81, SourceEnd: 100},
			},
		},
		{
			name: "case 2",
			args: args{
				start: 2,
				end:   45,
				ranges: []day05.ValRange{
					{SourceStart: 5, SourceEnd: 10},
					{SourceStart: 25, SourceEnd: 30},
					{SourceStart: 15, SourceEnd: 20},
					{SourceStart: 50, SourceEnd: 60},
					{SourceStart: 61, SourceEnd: 80},
				},
			},
			wantResult: []day05.ValRange{
				{SourceStart: 2, SourceEnd: 4},
				{SourceStart: 5, SourceEnd: 10},
				{SourceStart: 11, SourceEnd: 14},
				{SourceStart: 15, SourceEnd: 20},
				{SourceStart: 21, SourceEnd: 24},
				{SourceStart: 25, SourceEnd: 30},
				{SourceStart: 31, SourceEnd: 45},
			},
		},
		{
			name: "case 3",
			args: args{
				start: 2,
				end:   55,
				ranges: []day05.ValRange{
					{SourceStart: 5, SourceEnd: 10},
					{SourceStart: 25, SourceEnd: 30},
					{SourceStart: 15, SourceEnd: 20},
					{SourceStart: 50, SourceEnd: 60},
					{SourceStart: 61, SourceEnd: 80},
				},
			},
			wantResult: []day05.ValRange{
				{SourceStart: 2, SourceEnd: 4},
				{SourceStart: 5, SourceEnd: 10},
				{SourceStart: 11, SourceEnd: 14},
				{SourceStart: 15, SourceEnd: 20},
				{SourceStart: 21, SourceEnd: 24},
				{SourceStart: 25, SourceEnd: 30},
				{SourceStart: 31, SourceEnd: 49},
				{SourceStart: 50, SourceEnd: 55},
			},
		},
		{
			name: "case 4",
			args: args{
				start: 2,
				end:   5,
				ranges: []day05.ValRange{
					{SourceStart: 5, SourceEnd: 10},
					{SourceStart: 25, SourceEnd: 30},
					{SourceStart: 15, SourceEnd: 20},
					{SourceStart: 50, SourceEnd: 60},
					{SourceStart: 61, SourceEnd: 80},
				},
			},
			wantResult: []day05.ValRange{
				{SourceStart: 2, SourceEnd: 4},
				{SourceStart: 5, SourceEnd: 5},
			},
		},
		{
			name: "case 5",
			args: args{
				start: 2,
				end:   4,
				ranges: []day05.ValRange{
					{SourceStart: 5, SourceEnd: 10},
					{SourceStart: 25, SourceEnd: 30},
					{SourceStart: 15, SourceEnd: 20},
					{SourceStart: 50, SourceEnd: 60},
					{SourceStart: 61, SourceEnd: 80},
				},
			},
			wantResult: []day05.ValRange{
				{SourceStart: 2, SourceEnd: 4},
			},
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, testCase.wantResult,
				day05.SplitRanges(testCase.args.start, testCase.args.end, testCase.args.ranges),
				"SplitRanges(%v, %v, %v)", testCase.args.start, testCase.args.end, testCase.args.ranges)
		})
	}
}

func TestMergeRanges(t *testing.T) {
	t.Parallel()

	type args struct {
		in []day05.ValRange
	}

	tests := []struct {
		name    string
		args    args
		wantOut []day05.ValRange
	}{
		{
			name: "case 1",
			args: args{[]day05.ValRange{
				{SourceStart: 1, SourceEnd: 5},
				{SourceStart: 6, SourceEnd: 15},
				{SourceStart: 17, SourceEnd: 34},
				{SourceStart: 35, SourceEnd: 60},
				{SourceStart: 61, SourceEnd: 63},
				{SourceStart: 150, SourceEnd: 150},
				{SourceStart: 151, SourceEnd: 170},
				{SourceStart: 80, SourceEnd: 90},
			}},
			wantOut: []day05.ValRange{
				{SourceStart: 1, SourceEnd: 15},
				{SourceStart: 17, SourceEnd: 63},
				{SourceStart: 80, SourceEnd: 90},
				{SourceStart: 150, SourceEnd: 170},
			},
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, testCase.wantOut, day05.MergeRanges(testCase.args.in), "MergeRanges(%v)", testCase.args.in)
		})
	}
}
