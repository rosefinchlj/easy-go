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

    ParameterInvalid     Code = iota + 2000 // 微信相关
    TokenCheckError                         // token校验失败
    AppidNotFind                            // 应用id未找到
    HaveDeal                                // 已经处理
    ParseFilesError                         // 解析文件错误
    CacheException                          // 缓存异常
    TemplateExecuteError                    // 模板执行错误
    OpTimeError                             // 请不要平凡操作
    EmptyError                              // 数据为空
    Unfollow                                // 用户已取消关注
)
