// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

type StatusType uint8 // 255 possible values

const OK = 0

type DBStatusType StatusType

const (
	ReadAllDBError DBStatusType = iota + 1
	CreateError
	CheckExistsError
	DBNotExists
	DeleteError
)

type SchemaStatusType StatusType

const (
	SchemaReadError = iota + 1
	ErrorCreateTransaction
)

type SessionStatusType StatusType

const (
	SessionOpenError = iota + 1
	SessionCloseError
)
