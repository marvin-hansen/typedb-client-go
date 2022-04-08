package utils

func GetPersonDelete() string {
	return `
	match $p isa person, has phone-number "+44 091 xxx";
	delete $p isa person;
`
}
