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
