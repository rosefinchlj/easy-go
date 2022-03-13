package util

import (
    "fmt"
    "os"
    "regexp"
    "strings"
)

// 参数校验工具

// CheckNullParam 检测参数
func CheckNullParam(params ...string) bool {
    for _, value := range params {
        if len(value) == 0 {
            return true
        }
    }
    return false
}

// IsPhone 判断是否是手机号
func IsPhone(mobileNum string) bool {
    pattern := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
    reg := regexp.MustCompile(pattern)
    return reg.MatchString(mobileNum)
}

// IsMail 判断用户是否是邮件用户
func IsMail(email string) bool {
    pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
    reg := regexp.MustCompile(pattern)
    return reg.MatchString(email)
}

// IsRunTesting 判断是否在测试环境下使用
func IsRunTesting() bool {
    if len(os.Args) > 1 {
        fmt.Println(os.Args[1])
        return strings.HasPrefix(os.Args[1], "-test")
    }
    return false
}

// IsIDCard 判断是否是18或15位身份证
func IsIDCard(cardNo string) bool {
    // 18位身份证 ^(\d{17})([0-9]|X)$
    if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
        return false
    }
    return true
}
