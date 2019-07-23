package router

import (
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/global"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/object"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/repository"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

type base struct {
	app *iris.Application
	c   common
}

func NewRouterBase(app *iris.Application) *base {
	return &base{
		app: app,
		c:   common{},
	}
}

func (base *base) AddBase() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	v := base.app.Party("/", crs).AllowMethods(iris.MethodOptions)
	{
		v.Get("/version", base.version)
		v.Get("/data", base.getData)
	}
}

//获取版本
func (base *base) version(ctx iris.Context) {
	v := object.VersionResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Version: global.Version,
	}
	base.c.WriteResponse(ctx, v)
}

//获取门店数据
func (base *base) getData(ctx iris.Context) {
	rep, err := repository.NewRepOnLine()
	if err != nil {
		base.c.WriteResponse(ctx, object.NewDataResponse(-1, err.Error(), nil))
		return
	}
	if rep == nil {
		base.c.WriteResponse(ctx, object.NewDataResponse(-1, "rep online is nil", nil))
		return
	}
	d, err := rep.GetHeartBeatData()
	if err != nil {
		base.c.WriteResponse(ctx, object.NewDataResponse(-1, err.Error(), nil))
		return
	}
	base.c.WriteResponse(ctx,
		object.NewDataResponse(int(object.ErrTypeCodeNoError),
			string(object.ErrTypeMsgNoError),
			d))
	return
}
