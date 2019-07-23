package service

import (
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/global"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/webServer"
)

//启动服务内容
func StartService() error {
	log.Debug("Start Service")
	defer log.Debug("Start Service Complete")

	go func() {
		ws := webServer.NewWebServer(global.SysConfig.Iris.Port)
		ws.StartWebService()
	}()

	return nil
}
