package object

import (
	"github.com/Deansquirrel/goToolCommon"
	"strconv"
	"time"
)

type HeartBeat struct {
	MdId          int       `json:"mdId"`
	MdName        string    `json:"mdName"`
	ClientVersion string    `json:"clientName"`
	HeartBeat     time.Time `json:"heartBeat"`
	IsOffLine     bool      `json:"isOffLine"`
}

type HeartBeatStr struct {
	MdId          string `json:"mdId"`
	MdName        string `json:"mdName"`
	ClientVersion string `json:"clientName"`
	HeartBeat     string `json:"heartBeat"`
	IsOffLine     string `json:"isOffLine"`
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
