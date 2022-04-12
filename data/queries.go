package data

func GetQueryAllPeople() string {
	return `
	match
		$x isa person;
	get
		$x;
	`
}

func GetQueryAllPhoneNumbers() string {
	return `
	match
	  $p isa phone-number;
	get
		$p;
	`
}

func GetQueryPersonWithPhone() string {
	return `
	match
	  $p isa person,
	  has first-name $fName,	
	  has last-name $lName,
	  has phone-number $number;
	get
		$fName,
		$lName,
		$number;
	`
}

func GetQuerySinglePerson() string {
	return `
	match $p isa person, has phone-number "+54 398 559 0423";
`
}

func GetQueryDeletePerson() string {
	return `
	match $p isa person, has phone-number "+54 398 559 0423";
	delete $p isa person;
`
}
