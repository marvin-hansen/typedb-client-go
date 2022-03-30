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
