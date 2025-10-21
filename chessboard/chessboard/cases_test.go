package chessboard

var testCases = []struct {
	description string
	width       int
	length      int
	want        string
	expectError bool
}{
	{
		description: "empty chessboard",
		width:       0,
		length:      0,
		want:        "",
		expectError: false,
	},
	{
		description: "1x1 chessboard",
		width:       1,
		length:      1,
		want:        "#",
		expectError: false,
	},
	{
		description: "2x2",
		width:       2,
		length:      2,
		want:        "# \n #",
		expectError: false,
	},
	{
		description: "3x4",
		width:       3,
		length:      4,
		want:        "# #\n # \n# #\n # ",
		expectError: false,
	},
	{
		description: "7x7",
		width:       7,
		length:      7,
		want:        "# # # #\n # # # \n# # # #\n # # # \n# # # #\n # # # \n# # # #",
		expectError: false,
	},
	{
		description: "10x10",
		width:       10,
		length:      10,
		want:        "# # # # # \n # # # # #\n# # # # # \n # # # # #\n# # # # # \n # # # # #\n# # # # # \n # # # # #\n# # # # # \n # # # # #",
		expectError: false,
	},
	{
		description: "1001x1",
		width:       1001,
		length:      1,
		want:        "",
		expectError: true,
	},
}
