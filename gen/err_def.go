package gen

const (
    ServerMaintenance Code = iota - 1 // 系统错误
    Ok                                // ok
    UnknownError                      // 未知错误
    InvalidArgument                   // 参数解析失败
    AppIdOverdue                      // appId过期

    UserExisted         Code = iota + 1000 // 用户已存在
    VerifyTimeError                        // 验证码请求过于频繁
    MailSendFailed                         // 邮箱发送失败
    SMSSendFailed                          // 手机发送失败
    PhoneParameterError                    // 手机号格式有问题
)
