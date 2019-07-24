package repository

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/global"
	"github.com/Deansquirrel/goZ9MdDataTransHeartBeatWatcher/object"
	"github.com/kataras/iris/core/errors"
	"time"
)

const (
	sqlGetHeartBeatData = "" +
		"SELECT [mdid],[mdname],[clientversion],[heartbeat] " +
		"FROM [heartbeat] " +
		"Order By [mdid]"
)

type repOnLine struct {
	dbConfig *goToolMSSql.MSSqlConfig
}

func NewRepOnLine() (*repOnLine, error) {
	comm := NewCommon()
	dbConfig, err := comm.GetOnLineDbConfig()
	if err != nil {
		return nil, err
	} else {
		return &repOnLine{
			dbConfig: dbConfig,
		}, nil
	}
}

func (r *repOnLine) GetHeartBeatData() ([]*object.HeartBeat, error) {
	if r.dbConfig == nil {
		return nil, errors.New("rep online dbConfig is nil")
	}
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL(r.dbConfig, sqlGetHeartBeatData)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	list := make([]*object.HeartBeat, 0)
	var mdId int
	var mdName, clientVersion string
	var heartBeat time.Time
	offTimeLine := time.Now().Add(-goToolCommon.GetDurationBySecond(global.SysConfig.Web.OffLine))
	for rows.Next() {
		err := rows.Scan(&mdId, &mdName, &clientVersion, &heartBeat)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("read data err: %s", err.Error()))
		}
		list = append(list, &object.HeartBeat{
			MdId:          mdId,
			MdName:        mdName,
			ClientVersion: clientVersion,
			HeartBeat:     heartBeat,
			IsOffLine:     goToolCommon.GetDateTimeStr(offTimeLine) > goToolCommon.GetDateTimeStr(heartBeat),
		})
	}
	if rows.Err() != nil {
		return nil, errors.New(fmt.Sprintf("read data err: %s", rows.Err().Error()))
	}
	return list, nil
}
