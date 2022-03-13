package weixin

import (
    "crypto/md5"
    "fmt"
    "github.com/rosefinchlj/easy-go/gen"
    "github.com/rosefinchlj/easy-go/log"

    "github.com/rosefinchlj/easy-go/util"
    "strings"
    "time"

    "gopkg.in/go-with/wxpay.v1"
)

const (
    unifiedOrderURL  = "https://api.mch.weixin.qq.com/pay/unifiedorder"                      // 统一下单请求URL
    queryOrderURL    = "https://api.mch.weixin.qq.com/pay/orderquery"                        // 统一查询请求URL
    refundURL        = "https://api.mch.weixin.qq.com/secapi/pay/refund"                     // 退款请求URL
    enterprisePayURL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers" // 查询企业付款接口请求URL
)

const (
    PAY_SUCCESS    = 1  // 支付成功
    PAY_REFUND     = 2  // 转入退款
    PAY_CLOSED     = 3  // 已关闭
    PAY_NOTPAY     = 4  // 未支付
    PAY_REVOKED    = 5  // 已撤销
    PAY_USERPAYING = 6  // 支付中
    PAY_ERROR      = -1 // 支付失败
)

// UnifiedOrder 小程序统一下单接口
//	小程序统一下单接口
//	openID:用户唯一标识
//	price : 预支付价钱(单位分)
//	priceBody ： 支付描述
//	orderID ： 商户订单号
func (_wx *wxTools) UnifiedOrder(openID string, price int64, priceBody, orderID, clientIP string) gen.ResultMsg {
    if util.CheckNullParam(openID, orderID) || price <= 0 { //参数检测
        return gen.GetResultMsg(gen.ParameterInvalid)
    }

    params := make(wxpay.Params)
    // 查询企业付款接口请求参数
    params.SetString("appid", _wx.client.AppId)
    params.SetString("mch_id", _wx.client.MchId)
    params.SetString("body", priceBody)
    params.SetInt64("total_fee", price)
    params.SetString("spbill_create_ip", clientIP)
    params.SetString("notify_url", _wx.wxInfo.NotifyURL)
    params.SetString("trade_type", "JSAPI")
    params.SetString("openid", openID)
    params.SetString("nonce_str", util.GetRandomString(32)) // 随机字符串
    params.SetString("out_trade_no", orderID)               // 商户订单号
    params.SetString("sign", _wx.client.Sign(params))       // 签名 c.Sign(params)

    // 发送查询企业付款请求
    ret, err := _wx.client.Post(unifiedOrderURL, params, true)
    if err != nil {
        log.Error(err)
        msg := gen.GetResultMsg(err)
        return msg
    }

    //ret["order_id"] = order_tbl.Order_id
    log.Info(ret)

    if strings.EqualFold(ret["result_code"], "SUCCESS") { //再次签名
        dd := make(map[string]string)
        dd["timeStamp"] = fmt.Sprintf("%v", time.Now())
        dd["nonceStr"] = util.GetRandomString(32)
        dd["package"] = "prepay_id=" + ret["prepay_id"]
        dd["signType"] = "MD5"
        dd["paySign"] = "MD5"
        //appId=wxd678efh567hg6787&nonceStr=5K8264ILTKCH16CQ2502SI8ZNMTM67VS&package=prepay_id=&signType=MD5&timeStamp=1490840662&key=qazwsxedcrfvtgbyhnujmikolp111111
        str := "appId=" + _wx.wxInfo.AppID + "&nonceStr=" + dd["nonceStr"] + "&package=" +
            dd["package"] + "&signType=MD5&timeStamp=" + dd["timeStamp"] + "&key=" + _wx.wxInfo.APIKey
        by := md5.Sum([]byte(str))
        dd["paySign"] = strings.ToUpper(fmt.Sprintf("%x", by))
        dd["orderId"] = orderID

        msg := gen.GetResultMsg()
        msg.Data = dd
        return msg
    }

    msg := gen.GetResultMsg(ret["err_code_des"])
    msg.Data = ret
    return msg
}

// SelectOrder 统一查询接口
/*
	统一查询接口
	open_id:用户唯一标识
	order_id ： 商户订单号
*/
func (_wx *wxTools) SelectOrder(openID, orderID string) (int, gen.ResultMsg) {
    if util.CheckNullParam(openID, orderID) { //参数检测
        return 0, gen.GetResultMsg(gen.ParameterInvalid)
    }

    code := 0

    params := make(wxpay.Params)
    // 查询企业付款接口请求参数
    params.SetString("appid", _wx.client.AppId)
    params.SetString("mch_id", _wx.client.MchId)
    params.SetString("out_trade_no", orderID)               //商户订单号
    params.SetString("nonce_str", util.GetRandomString(32)) // 随机字符串
    params.SetString("sign", _wx.client.Sign(params))       // 签名 c.Sign(params)

    // 发送查询企业付款请求
    ret := make(wxpay.Params)
    var err error
    ret, err = _wx.client.Post(queryOrderURL, params, true)
    if err != nil { //做再次确认
        time.Sleep(time.Second * 1)
        ret, err = _wx.client.Post(queryOrderURL, params, true)
        if err != nil {
            log.Error(err)
            msg := gen.GetResultMsg(gen.Ok)
            return code, msg
        }
    }

    msg := gen.GetResultMsg(gen.Ok)

    /*
    	SUCCESS—支付成功
    	REFUND—转入退款
    	NOTPAY—未支付
    	CLOSED—已关闭
    	REVOKED—已撤销（刷卡支付）
    	USERPAYING--用户支付中
    	PAYERROR--支付失败(其他原因，如银行返回失败)
    */
    if ret["trade_state"] == "SUCCESS" {
        code = PAY_SUCCESS
    } else if ret["trade_state"] == "REFUND" {
        code = PAY_REFUND
    } else if ret["trade_state"] == "CLOSED" {
        code = PAY_CLOSED
    } else if ret["trade_state"] == "NOTPAY" {
        code = PAY_NOTPAY
    } else if ret["trade_state"] == "REVOKED" {
        code = PAY_REVOKED
    } else if ret["trade_state"] == "USERPAYING" {
        code = PAY_USERPAYING
    } else {
        code = PAY_ERROR
    }

    if ret["trade_state"] == "SUCCESS" { //支付成功
        msg.Success = true
    } else {
        msg.Success = false
    }

    msg.Data = ret
    return code, msg
}

// RefundPay 申请退款
/*
	申请退款
	openID:用户唯一标识
	orderID ： 商户订单号
	rrefundNO ：商户退款单号
	totalFee: 订单总金额 分
	refundFee: 退款总金额 分
*/
func (_wx *wxTools) RefundPay(openID, orderID, refundNO string, totalFee, refundFee int64) (bool, gen.ResultMsg) {
    if util.CheckNullParam(openID, orderID) { //参数检测
        return false, gen.GetResultMsg(gen.ParameterInvalid)
    }
    code := false
    params := make(wxpay.Params)
    // 退款请求参数
    params.SetString("appid", _wx.client.AppId)
    params.SetString("mch_id", _wx.client.MchId)
    params.SetString("out_trade_no", orderID)               //商户订单号
    params.SetString("out_refund_no", refundNO)             //商户退款单号
    params.SetInt64("total_fee", totalFee)                  // 订单总金额（分）
    params.SetInt64("refund_fee", refundFee)                // 退款金额（分）
    params.SetString("nonce_str", util.GetRandomString(32)) // 随机字符串
    params.SetString("sign", _wx.client.Sign(params))       // 签名 c.Sign(params)

    // 发送申请退款请求
    ret, err := _wx.client.Post(refundURL, params, true)
    if err != nil {
        log.Error(err)
        msg := gen.GetResultMsg(err)
        return code, msg
    }

    msg := gen.GetResultMsg(gen.Ok)

    if len(ret["err_code_des"]) > 0 {
        msg = gen.GetResultMsg(ret["err_code_des"])
    } else {
        if ret["result_code"] == "SUCCESS" { //申请成功
            msg.Success = true
            code = true
        }
    }
    msg.Data = ret
    return code, msg
}

// WxEnterprisePay 企业付款
/*
企业付款
open_id:用户唯一标识
trade_no : 商户订单号
desc ： 操作说明
amount：付款金额 分
*/
func (_wx *wxTools) WxEnterprisePay(openID, tradeNO, desc, ipAddr string, amount int) bool {
    // c := wxpay.NewClient(_wx.wxInfo.AppID, _wx.wxInfo.MchID, _wx.wxInfo.APIKey)

    // // 附着商户证书
    // err := c.WithCert(_wx.certFile, _wx.keyFile, _wx.rootcaFile)
    // if err != nil {
    // 	log.Fatal(err)
    // }

    params := make(wxpay.Params)
    nonceStr := util.GetRandomString(16)
    // 查询企业付款接口请求参数
    params.SetString("mch_appid", _wx.client.AppId) // 商户账号appid
    params.SetString("mchid", _wx.client.MchId)     // 商户号
    params.SetString("nonce_str", nonceStr)         // 随机字符串
    params.SetString("partner_trade_no", tradeNO)   // 商户订单号
    params.SetString("openid", openID)              // 用户openid
    params.SetString("check_name", "NO_CHECK")      // 校验用户姓名选项
    params.SetInt64("amount", int64(amount))        // 企业付款金额，单位为分
    params.SetString("desc", desc)                  // 企业付款操作说明信息。必填。
    params.SetString("spbill_create_ip", ipAddr)

    params.SetString("sign", _wx.client.Sign(params)) // 签名

    // 发送查询企业付款请求
    ret, err := _wx.client.Post(enterprisePayURL, params, true)
    if err != nil {
        log.Fatal(err)
    }
    log.Debug(ret)
    returnCode := ret.GetString("return_code")
    resultCode := ret.GetString("result_code")
    if returnCode == "SUCCESS" && resultCode == "SUCCESS" {
        return true
    }

    return false
}
