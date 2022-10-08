package main

var rules = []struct {
	character   string
	replacement string
}{
	{
		character:   "%",
		replacement: "",
	},
	{
		character:   "100",
		replacement: "1",
	},
	{
		character:   "90",
		replacement: "0.9",
	},
	{
		character:   "80",
		replacement: "0.8",
	},
	{
		character:   "70",
		replacement: "0.7",
	},
	{
		character:   "60",
		replacement: "0.6",
	},
	{
		character:   "50",
		replacement: "0.5",
	},
	{
		character:   "40",
		replacement: "0.4",
	},
	{
		character:   "30",
		replacement: "0.3",
	},
	{
		character:   "20",
		replacement: "0.2",
	},
	{
		character:   "10",
		replacement: "0.1",
	},
}
