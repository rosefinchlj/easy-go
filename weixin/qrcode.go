package weixin

import (
    "encoding/json"
    "fmt"
    "github.com/ddliu/go-httpclient"
    "github.com/rosefinchlj/easy-go/log"
    "io/ioutil"
)

const (
    GETSHAREURL  = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="          // 小程序码
    GETQRCODEURL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=" // 小程序二维码
)

// GetShareQrcode 获取小程序码
// path：图片保存路径
// scene: 附带参数
// page:小程序页面头部
func (_wx *wxTools) GetShareQrcode(path string, scene, page string) (ret QrcodeRet) {
    accessToken, _ := _wx.GetAccessToken() // 获取access_token
    data := wxPostdata{Scene: scene, Page: page}

    resp, err := httpclient.PostJson(GETSHAREURL+accessToken, data)
    if err != nil {
        log.Error(err)
        return
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return
    }

    json.Unmarshal(readAll, &ret) //错误码45029 最大限制
    if ret.Errcode == 0 {
        ioutil.WriteFile(path, readAll, 0666) //写入文件(字节数组)
    }
    return
}

// GetWxQrcode 获取小程序二维码 （有限个）
// path：图片保存路径
// page: 小程序页面pages/index?query=1
// width: 二维码宽度

func (_wx *wxTools) GetWxQrcode(path, page string, width int) (ret QrcodeRet) {
    fmt.Println(path)
    //获取access_token
    accessToken, _ := _wx.GetAccessToken()
    data := wxQrcodedata{Path: page, Width: width}

    resp, err := httpclient.PostJson(GETQRCODEURL+accessToken, data)
    if err != nil {
        log.Error(err)
        return
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return
    }

    _ = json.Unmarshal(readAll, &ret) //错误码45029 最大限制
    if ret.Errcode == 0 {
        ioutil.WriteFile(path, readAll, 0666) //写入文件(字节数组)
    }
    return
}
