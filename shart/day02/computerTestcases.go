package main

var testCasesRunComputer = []struct {
	description string
	input       []int
	expected    []int
}{
	{
		description: "1+1=2",
		input:       []int{1, 0, 0, 0, 99},
		expected:    []int{2, 0, 0, 0, 99},
	},
	{
		description: "3*2=6",
		input:       []int{2, 3, 0, 3, 99},
		expected:    []int{2, 3, 0, 6, 99},
	},
	{
		description: "99*99 = 9801",
		input:       []int{2, 4, 4, 5, 99, 0},
		expected:    []int{2, 4, 4, 5, 99, 9801},
	},
	{
		description: "more complex",
		input:       []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		expected:    []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
	{
		description: "challenge",
		input:       day2Input,
		expected:    []int{4714701, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 60, 1, 5, 19, 61, 2, 10, 23, 244, 1, 27, 5, 245, 2, 9, 31, 735, 1, 35, 5, 736, 2, 6, 39, 1472, 1, 43, 5, 1473, 2, 47, 10, 5892, 2, 51, 6, 11784, 1, 5, 55, 11785, 2, 10, 59, 47140, 1, 63, 6, 47142, 2, 67, 6, 94284, 1, 71, 5, 94285, 1, 13, 75, 94290, 1, 6, 79, 94292, 2, 83, 13, 471460, 1, 87, 6, 471462, 1, 10, 91, 471466, 1, 95, 9, 471469, 2, 99, 13, 2357345, 1, 103, 6, 2357347, 2, 107, 6, 4714694, 1, 111, 2, 4714696, 1, 115, 13, 0, 99, 2, 0, 14, 0},
	},
}

var day2Input = []int{
	1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 5, 19, 23, 2, 10, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 35, 5, 39, 2, 6, 39, 43, 1, 43, 5, 47, 2, 47, 10, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 1, 71, 5, 75, 1, 13, 75, 79, 1, 6, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 10, 91, 95, 1, 95, 9, 99, 2, 99, 13, 103, 1, 103, 6, 107, 2, 107, 6, 111, 1, 111, 2, 115, 1, 115, 13, 0, 99, 2, 0, 14, 0,
}
