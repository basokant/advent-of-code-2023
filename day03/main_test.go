package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestPart2(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  int
// 	}{
// 		{
// 			name: "example",
// 			input: `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..`,
// 			want: 467835,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := part2(tt.input); got != tt.want {
// 				t.Errorf("part2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
