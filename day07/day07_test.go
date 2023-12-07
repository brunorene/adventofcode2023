package day07_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brunorene/adventofcode2023/day07"
)

func TestHandGetHandType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		h    day07.Hand
		want day07.HandType
	}{
		{
			name: "KTJJT == two pair",
			h:    day07.Hand("KTJJT"),
			want: day07.TwoPair,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, day07.GetHandType(tt.h), "GetHandType()")
		})
	}
}

func TestSortPlays(t *testing.T) {
	t.Parallel()

	type args struct {
		input []day07.Play
	}

	tests := []struct {
		name string
		args args
		want []day07.Play
	}{
		{
			name: "sort 8967A 24JT3",
			args: args{[]day07.Play{{day07.Hand("8967A"), 1}, {day07.Hand("24JT3"), 1}}},
			want: []day07.Play{{day07.Hand("24JT3"), 1}, {day07.Hand("8967A"), 1}},
		},
	}

	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, test.want, day07.SortPlays(test.args.input, day07.Cards1, day07.GetHandType),
				"SortPlays(%v)", test.args.input)
		})
	}
}
