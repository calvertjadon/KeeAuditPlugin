package audit

type specDef struct {
	Code        string
	Description string
}

var SpecDefs = []specDef{
	{
		Code:        "entropy.min",
		Description: "specifies the minimum allowed entropy",
	},
	{
		Code:        "duplicates.max",
		Description: "specifies that a given password may only appear a specified number of times in a database",
	},
}
