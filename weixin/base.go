package weixin

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/rosefinchlj/easy-go/cache"
    "github.com/rosefinchlj/easy-go/gen"
    "github.com/rosefinchlj/easy-go/log"
    "gopkg.in/go-with/wxpay.v1"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/bitly/go-simplejson"

    "github.com/ddliu/go-httpclient"
)

const (
    _getTicket      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
    _getJsurl       = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
    _getToken       = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="
    _getSubscribe   = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="
    _getTempMsg     = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
    _createMenu     = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
    _deleteMenu     = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
    _sendCustom     = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
    _setGuideConfig = "https://api.weixin.qq.com/cgi-bin/guide/setguideconfig?access_token="
    _cacheToken     = "wx_access_token"
    _cacheTicket    = "weixin_card_ticket"
)

var localCache = cache.NewLocalCache()

type wxTools struct {
    client     *wxpay.Client
    wxInfo     WxInfo
    certFile   string // 微信支付商户平台证书路径
    keyFile    string
    rootcaFile string
}

// GetAccessToken 获取微信accesstoken
// 获取登录凭证
func (_wx *wxTools) GetAccessToken() (accessToken string, err error) {
    // 先从缓存中获取 access_token
    value, found := localCache.Get(_cacheToken)
    if found {
        accessToken = value.(string)
        return
    }

    var url = _getToken + _wx.wxInfo.AppID + "&secret=" + _wx.wxInfo.AppSecret

    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    js, err := simplejson.NewJson(body)
    if err == nil {
        accessToken, _ = js.Get("access_token").String()
        if len(accessToken) == 0 {
            log.Error(js)
            return
        }
        // 保存缓存
        localCache.Set(_cacheToken, &accessToken, time.Duration(7000)*time.Second)
    }

    return
}

// clearAccessTokenCache 清除accesstoken缓存
func (_wx *wxTools) clearAccessTokenCache() {
    // 先从缓存中获取 access_token
    localCache.Delete(_cacheToken)
}

// GetAPITicket 获取微信卡券ticket
func (_wx *wxTools) GetAPITicket() (ticket string, err error) {
    //先从缓存中获取
    value, found := localCache.Get(_cacheTicket)
    if found {
        ticket = value.(string)
        return
    }

    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        log.Error(err)
        return
    }
    var url = _getTicket + accessToken

    resp, err := http.Get(url)
    if err != nil {
        log.Error(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return
    }
    var result APITicket
    json.Unmarshal(body, &result)
    ticket = result.Ticket
    //保存缓存
    localCache.Set(_cacheTicket, ticket, 7000*time.Second)

    return
}

// GetJsTicket 获取微信js ticket
func (_wx *wxTools) GetJsTicket() (ticket string, err error) {
    //先从缓存中获取
    value, found := localCache.Get("base")
    if found {
        ticket = value.(string)
        return
    }

    accessToken, e := _wx.GetAccessToken()
    if e != nil {
        log.Error(e)
        err = e
        return
    }
    var url = _getJsurl + accessToken

    resp, e1 := http.Get(url)
    if e1 != nil {
        log.Error(e1)
        err = e1
        return
    }
    defer resp.Body.Close()
    body, e2 := ioutil.ReadAll(resp.Body)
    if e2 != nil {
        log.Error(e2)
        err = e2
        return
    }
    var result APITicket
    json.Unmarshal(body, &result)
    ticket = result.Ticket
    //保存缓存
    localCache.Set("base", ticket, 7000*time.Second)

    return
}

// SendTemplateMsg 发送订阅消息
func (_wx *wxTools) SendTemplateMsg(msg TempMsg) bool {
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        log.Error(err)
        return false
    }

    resp, err := httpclient.PostJson(_getSubscribe+accessToken, msg)
    if err != nil {
        log.Error(err)
        return false
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return false
    }

    var res ResTempMsg
    json.Unmarshal(readAll, &res)
    return res.Errcode == 0
}

// SendWebTemplateMsg 发送订阅消息
func (_wx *wxTools) SendWebTemplateMsg(msg TempWebMsg) error {
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        log.Errorf("SendWebTemplateMsg error: openid:%v,err:%v", msg.Touser, err)
        return err
    }

    resp, err := httpclient.PostJson(_getTempMsg+accessToken, msg)
    if err != nil {
        log.Error(err)
        return err
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return err
    }
    var res ResTempMsg
    _ = json.Unmarshal(readAll, &res)

    if res.Errcode != 0 { // try again
        _wx.clearAccessTokenCache()
        accessToken, err = _wx.GetAccessToken()
        if err != nil {
            log.Error(err)
            return err
        }
        resp, err = httpclient.PostJson(_getTempMsg+accessToken, msg)
        if err != nil {
            log.Error(err)
            return err
        }

        readAll, err = resp.ReadAll()
        if err != nil {
            log.Error(err)
            return err
        }
        var res ResTempMsg
        _ = json.Unmarshal(readAll, &res)

        if res.Errcode == 0 {
            if res.Errcode == 43004 {
                return errors.New(gen.Unfollow.String())
            }
            log.Errorf("SendWebTemplateMsg error: openid:%v,res:%v", msg.Touser, res)
        }
    }
    return nil
}

// CreateMenu 创建自定义菜单
func (_wx *wxTools) CreateMenu(menu WxMenu) error { // 创建自定义菜单
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        return err
    }

    resp, err := httpclient.PostJson(_createMenu+accessToken, menu)
    if err != nil {
        log.Error(err)
        return err
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return err
    }

    var res ResTempMsg
    json.Unmarshal(readAll, &res)
    b := res.Errcode == 0
    if !b {
        return fmt.Errorf("SendWebTemplateMsg error: res:%v", res)
    }

    return nil
}

// DeleteMenu 删除自定义菜单
func (_wx *wxTools) DeleteMenu() error { // 创建自定义菜单
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        return err
    }

    resp, err := httpclient.PostJson(_deleteMenu+accessToken, nil)
    if err != nil {
        log.Error(err)
        return err
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return err
    }

    var res ResTempMsg
    json.Unmarshal(readAll, &res)
    b := res.Errcode == 0
    if !b {
        return fmt.Errorf("SendWebTemplateMsg error: res:%v", res)
    }

    return nil
}

// SendCustomMsg 发送客服消息
func (_wx *wxTools) SendCustomMsg(msg CustomMsg) error {
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        return err
    }

    resp, err := httpclient.PostJson(_sendCustom+accessToken, msg)
    if err != nil {
        log.Error(err)
        return err
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return err
    }

    var res ResTempMsg
    json.Unmarshal(readAll, &res)
    b := res.Errcode == 0
    if !b {
        return fmt.Errorf("SendWebTemplateMsg error: res:%v", res)
    }

    return nil
}

// SetGuideConfig 快捷回复与关注自动回复
func (_wx *wxTools) SetGuideConfig(guideConfig GuideConfig) error {
    accessToken, err := _wx.GetAccessToken()
    if err != nil {
        return err
    }

    resp, err := httpclient.PostJson(_setGuideConfig+accessToken, guideConfig)
    if err != nil {
        log.Error(err)
        return err
    }

    readAll, err := resp.ReadAll()
    if err != nil {
        log.Error(err)
        return err
    }

    var res ResTempMsg
    json.Unmarshal(readAll, &res)
    b := res.Errcode == 0
    if !b {
        return fmt.Errorf("SetGuideConfig error: res:%v", res)
    }

    return nil
}
