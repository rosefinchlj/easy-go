package gen

type ResultMsg struct {
    Code    int
    Message string
    Data    interface{}
}

// GetResultMsg 通过code获取错误信息，同时支持输入自定义错误信息或者error
func GetResultMsg(errorCode ...interface{}) ResultMsg {
    resultMsg := ResultMsg{}

    if len(errorCode) == 0 {
        resultMsg.Code = int(UnknownError)
        resultMsg.Message = UnknownError.Message()
        return resultMsg
    }

    for _, e := range errorCode {
        switch v := e.(type) {
        case Code:
            resultMsg.Code = int(v)
            resultMsg.Message = v.Message()
        case string:
            resultMsg.Message = v
        case error:
            resultMsg.Message = v.Error()
        }
    }

    return resultMsg
}
