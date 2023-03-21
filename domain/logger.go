package domain

type ILogger interface {
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Errorln(args ...interface{})
	Println(args ...interface{})
}
