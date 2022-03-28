// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

type StatusType uint8 // 255 possible values

const (
	OK StatusType = iota
	DBReadAllError
	DBCreateError
	DBCheckExistsError
	DBNotExists
	DBDeleteError

	ErrorSessionOpen
	ErrorSessionClose

	ErrorCreateTransaction
	ErrorOpenTransaction

	ErrorReadSchema
	ErrorQuerySchema
	ErrorWriteSchema
	ErrorCommitSchemaTransaction
	ErrorRollbackSchemaTransaction
	ErrorCloseSchemaTransaction
)
