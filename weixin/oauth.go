package weixin

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/rosefinchlj/easy-go/log"
)

// SmallAppOauth 小程序授权
func (_wx *wxTools) SmallAppOauth(jscode string) string {
    var url = "https://api.weixin.qq.com/sns/jscode2session?appid=" + _wx.wxInfo.AppID + "&secret=" +
        _wx.wxInfo.AppSecret + "&js_code=" + jscode + "&grant_type=authorization_code&trade_type=JSAPI"

    resp, e := http.Get(url)
    if e != nil {
        log.Error(e)
        return ""
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(e)
        return ""
    }
    return string(body)
}

// GetWebOauth 网页授权
func (_wx *wxTools) GetWebOauth(code string) (*AccessToken, error) {
    var url = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + _wx.wxInfo.AppID + "&secret=" +
        _wx.wxInfo.AppSecret + "&code=" + code + "&grant_type=authorization_code"

    resp, err := http.Get(url)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    var res AccessToken
    json.Unmarshal(body, &res)
    return &res, nil
}

// GetWebUserinfo 获取用户信息
func (_wx *wxTools) GetWebUserinfo(openid, accessToken string) (*WxUserinfo, error) {
    var url = "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN"

    resp, err := http.Get(url)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    var res WxUserinfo
    json.Unmarshal(body, &res)
    return &res, nil
}
