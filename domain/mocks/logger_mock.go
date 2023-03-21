package mocks

type LoggerMock struct{}

func (o LoggerMock) Debugln(args ...interface{}) {}
func (o LoggerMock) Infoln(args ...interface{})  {}
func (o LoggerMock) Errorln(args ...interface{}) {}
func (o LoggerMock) Println(args ...interface{}) {}
