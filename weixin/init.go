package weixin

import (
    "github.com/rosefinchlj/easy-go/file"
    "github.com/rosefinchlj/easy-go/gen"
    "github.com/rosefinchlj/easy-go/log"
    wxpay "gopkg.in/go-with/wxpay.v1"
)

// 微信支付商户平台证书路径

// CertFileLoc  cert.pem
var CertFileLoc = "/conf/cert/apiclient_cert.pem"

// KeyFileLoc key.pem
var KeyFileLoc = "/conf/cert/apiclient_key.pem"

// RootcaFileLoc rootca.pem
var RootcaFileLoc = "/conf/cert/rootca.pem"

// IWxTools 微信操作类型
type IWxTools interface {
    GetAccessToken() (accessToken string, err error)                                             // 获取登录凭证
    GetAPITicket() (ticket string, err error)                                                    // 获取微信卡券ticket
    GetJsTicket() (ticket string, err error)                                                     // 获取微信js ticket
    SendTemplateMsg(msg TempMsg) bool                                                            // 发送订阅消息
    SmallAppOauth(jscode string) string                                                          // 小程序授权
    UnifiedOrder(openID string, price int64, priceBody, orderID, clientIP string) gen.ResultMsg  // 小程序统一下单接口
    SelectOrder(openID, orderID string) (int, gen.ResultMsg)                                     // 统一查询接口
    RefundPay(openID, orderID, refundNO string, totalFee, refundFee int64) (bool, gen.ResultMsg) // 申请退款
    WxEnterprisePay(openID, tradeNO, desc, ipAddr string, amount int) bool                       // 企业付款
    GetShareQrcode(path string, scene, page string) (ret QrcodeRet)                              // 获取小程序码
    GetWxQrcode(path, page string, width int) (ret QrcodeRet)                                    // 获取小程序二维码 （有限个）

    GetWebOauth(code string) (*AccessToken, error)                  // 授权
    GetWebUserinfo(openid, accessToken string) (*WxUserinfo, error) // 获取用户信息
    SendWebTemplateMsg(msg TempWebMsg) error                        // 发送公众号模板消息
    CreateMenu(menu WxMenu) error                                   // 创建自定义菜单
    DeleteMenu() error                                              // 删除自定义菜单
    SetGuideConfig(guideConfig GuideConfig) error                   // 快捷回复与关注自动回复
    SendCustomMsg(msg CustomMsg) error                              // 发送客服消息
}

// New 新建及 初始化配置信息
func New(info WxInfo) (IWxTools, error) {
    t := &wxTools{
        wxInfo:     info,
        certFile:   file.GetCurrentDirectory() + CertFileLoc,
        keyFile:    file.GetCurrentDirectory() + KeyFileLoc,
        rootcaFile: file.GetCurrentDirectory() + RootcaFileLoc,
        client:     wxpay.NewClient(info.AppID, info.MchID, info.APIKey),
    }
    err := t.client.WithCert(t.certFile, t.keyFile, t.rootcaFile)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return t, nil
}
