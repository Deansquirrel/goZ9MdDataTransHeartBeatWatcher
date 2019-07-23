package object

type VersionResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Version string `json:"version"`
}

type DataResponse struct {
	ErrCode int            `json:"errcode"`
	ErrMsg  string         `json:"errmsg"`
	Data    []HeartBeatStr `json:"data"`
}

func NewDataResponse(errCode int, errMsg string, dList []*HeartBeat) DataResponse {
	resultD := make([]HeartBeatStr, 0)
	if dList != nil && len(dList) > 0 {
		for _, d := range dList {
			resultD = append(resultD, *NewHeartBeatStr(d))
		}
	}
	return DataResponse{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    resultD,
	}
}
