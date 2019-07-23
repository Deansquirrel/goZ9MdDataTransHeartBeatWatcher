package object

import (
	"github.com/Deansquirrel/goToolCommon"
	"strconv"
	"time"
)

type HeartBeat struct {
	MdId          int       `toml:"mdId"`
	MdName        string    `toml:"mdName"`
	ClientVersion string    `toml:"clientName"`
	HeartBeat     time.Time `toml:"heartBeat"`
	IsOffLine     bool      `toml:"isOffLine"`
}

type HeartBeatStr struct {
	MdId          string `toml:"mdId"`
	MdName        string `toml:"mdName"`
	ClientVersion string `toml:"clientName"`
	HeartBeat     string `toml:"heartBeat"`
	IsOffLine     string `toml:"isOffLine"`
}

func NewHeartBeatStr(d *HeartBeat) *HeartBeatStr {
	return &HeartBeatStr{
		MdId:          strconv.Itoa(d.MdId),
		MdName:        d.MdName,
		ClientVersion: d.ClientVersion,
		HeartBeat:     goToolCommon.GetDateTimeStrWithMillisecond(d.HeartBeat),
		IsOffLine:     strconv.FormatBool(d.IsOffLine),
	}
}
