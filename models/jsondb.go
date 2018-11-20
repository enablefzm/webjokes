package models

type JsonDb struct {
	Cmd  string
	Info interface{}
}

func CreateJsonDb(cmd string, info interface{}) *JsonDb {
	return &JsonDb{
		Cmd:  cmd,
		Info: info,
	}
}
