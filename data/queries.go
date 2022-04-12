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
