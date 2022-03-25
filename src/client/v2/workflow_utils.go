// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

func getSchemaQuery() string {
	return `
	match
		$x sub thing;
	get
		$x;
	`
}
