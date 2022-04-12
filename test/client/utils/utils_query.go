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

func GetTestQueryAllPhone() string {
	return `
	match
	  $p isa phone-number;
	get
		$p;
	`
}

func GetTestQueryPersonPhone() string {
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

func GetTestQueryPerson() string {
	return `
	match
		$x isa person;
	get
		$x;
	`
}
