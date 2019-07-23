package object

import (
	"encoding/json"
	"fmt"
	"strings"
)

import log "github.com/Deansquirrel/goToolLog"

//系统配置（Server|Client）
type SystemConfig struct {
	Total          systemConfigTotal    `toml:"total"`
	OnLineDbConfig systemOnLineDbConfig `toml:"onLineDbConfig"`
	Web            systemConfigWeb      `toml:"web"`
	Service        systemConfigService  `toml:"service"`
	Iris           systemConfigIris     `toml:"iris"`
}

func (sc *SystemConfig) FormatConfig() {
	sc.Total.FormatConfig()
	sc.OnLineDbConfig.FormatConfig()
	sc.Web.FormatConfig()
	sc.Service.FormatConfig()
	sc.Iris.FormatConfig()
}

func (sc *SystemConfig) ToString() string {
	d, err := json.Marshal(sc)
	if err != nil {
		log.Warn(fmt.Sprintf("SystemConfig转换为字符串时遇到错误：%s", err.Error()))
		return ""
	}
	return string(d)
}

//通用配置
type systemConfigTotal struct {
	StdOut   bool   `toml:"stdOut"`
	LogLevel string `toml:"logLevel"`
}

func (t *systemConfigTotal) FormatConfig() {
	//去除首尾空格
	t.LogLevel = strings.Trim(t.LogLevel, " ")
	//设置默认日志级别
	if t.LogLevel == "" {
		t.LogLevel = "warn"
	}
	//设置字符串转换为小写
	t.LogLevel = strings.ToLower(t.LogLevel)
	t.LogLevel = t.checkLogLevel(t.LogLevel)
}

//校验SysConfig中iris日志级别设置
func (t *systemConfigTotal) checkLogLevel(level string) string {
	switch level {
	case "debug", "info", "warn", "error":
		return level
	default:
		return "warn"
	}
}

//门店库配置库
type systemOnLineDbConfig struct {
	Address string
}

func (c *systemOnLineDbConfig) FormatConfig() {
}

//服务配置
type systemConfigService struct {
	Name        string `toml:"name"`
	DisplayName string `toml:"displayName"`
	Description string `toml:"description"`
}

//格式化
func (sc *systemConfigService) FormatConfig() {
	sc.Name = strings.Trim(sc.Name, " ")
	sc.DisplayName = strings.Trim(sc.DisplayName, " ")
	sc.Description = strings.Trim(sc.Description, " ")
	if sc.Name == "" {
		sc.Name = "Z9MdDataTransHeartBeatWatcher_ws"
	}
	if sc.DisplayName == "" {
		sc.DisplayName = "Z9MdDataTransHeartBeatWatcher_ws"
	}
	if sc.Description == "" {
		sc.Description = sc.Name
	}
}

//Web
type systemConfigWeb struct {
	OffLine int `toml:"offLine"`
}

//格式化
func (sc *systemConfigWeb) FormatConfig() {
	if sc.OffLine < 30 {
		sc.OffLine = 30
	}
}

//Iris
type systemConfigIris struct {
	Port     int    `toml:"port"`
	LogLevel string `toml:"logLevel"`
}

//格式化
func (i *systemConfigIris) FormatConfig() {
	//设置默认端口 8000
	if i.Port == 0 {
		i.Port = 8000
	}
	//去除首尾空格
	i.LogLevel = strings.Trim(i.LogLevel, " ")
	//设置Iris默认日志级别
	if i.LogLevel == "" {
		i.LogLevel = "warn"
	}
	//设置字符串转换为小写
	i.LogLevel = strings.ToLower(i.LogLevel)
}
