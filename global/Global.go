package global

import (
	"context"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/object"
)

const (
	//PreVersion = "1.0.0 Build20190724"
	//TestVersion = "0.0.0 Build20190101"
	Version = "1.0.1 Build20190724"

	SecretKey = "Z9MdDataTransHeartBeatWatcher"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
