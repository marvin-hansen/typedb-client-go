// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package utils

func GetTestQuery() string {
	return `
	match
		$x sub thing;
	get
		$x;
	`

}

func GetTestQueryPersonPhone() string {
	return `
	match
	  $customer isa person, has phone-number $phone-number;
	get
		$phone-number;
	`
}

func GetTestQueryPerson() string {
	return `
	match
		$x isa person;
	get
		$x;
	`
}
