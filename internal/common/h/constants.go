package h

const (
	BadRequest          string = "ERR_BAD_REQUEST"
	InternalServerError string = "ERR_INTERNAL_SERVER"
	MongoError          string = "ERR_MONGO"
	EncodeError         string = "ERR_ENCODE_RES"
	NotFoundError       string = "ERR_NOT_FOUND"
	InitError           string = "ERR_INIT"
	InitOk              string = "OK_INIT"
)

const (
	Limit int64 = 50
	Skip  int64 = 0
)
