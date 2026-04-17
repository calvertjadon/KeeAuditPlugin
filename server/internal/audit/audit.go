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
		Code:        "duplicates.none",
		Description: "specifies that a given password may only appear once in a given database",
	},
}
