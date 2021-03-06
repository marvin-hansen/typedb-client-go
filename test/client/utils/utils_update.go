package utils

// GetPersonTelUpdate returns update query for testing
// https://docs.vaticle.com/docs/query/update-query
func GetPersonTelUpdate() string {
	return `
	match $p isa person, has phone-number $tel;
	$tel contains "+44 091 xxx"; 
	delete $tel isa phone-number;
	insert $p has phone-number "+44-091-81726-474";
`
}
