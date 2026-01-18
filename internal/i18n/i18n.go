package i18n

// Language 语言类型
type Language string

const (
	// LanguageEn 英语
	LanguageEn Language = "en"
	// LanguageZh 中文
	LanguageZh Language = "zh"
)

// MessageKey 消息键类型
type MessageKey string

// 日志消息键（英文）
const (
	// 认证相关
	LogAuthFailedNoToken    MessageKey = "log.auth.failed.no_token"
	LogAuthFailedInvalidFmt MessageKey = "log.auth.failed.invalid_format"
	LogAuthFailedInvalid    MessageKey = "log.auth.failed.invalid"
	LogAuthSuccess          MessageKey = "log.auth.success"

	// 请求相关
	LogRequestCost    MessageKey = "log.request.cost"
	LogResponseBody   MessageKey = "log.response.body"
	LogPanicRecovered MessageKey = "log.panic.recovered"
	LogInternalError  MessageKey = "log.internal.error"
)

// 用户消息键（中文，用于API响应）
const (
	// 认证相关
	UserAuthNoToken    MessageKey = "user.auth.no_token"
	UserAuthInvalidFmt MessageKey = "user.auth.invalid_format"
	UserAuthInvalid    MessageKey = "user.auth.invalid"

	// 用户操作相关
	UserCreateSuccess       MessageKey = "user.create.success"
	UserRegisterSuccess     MessageKey = "user.register.success"
	UserGetSuccess          MessageKey = "user.get.success"
	UserGetAllSuccess       MessageKey = "user.get_all.success"
	UserUpdateSuccess       MessageKey = "user.update.success"
	UserDeleteSuccess       MessageKey = "user.delete.success"
	UserLoginSuccess        MessageKey = "user.login.success"
	UserRefreshTokenSuccess MessageKey = "user.refresh_token.success"

	// 错误相关
	UserErrorBadRequest MessageKey = "user.error.bad_request"
	UserErrorInvalidID  MessageKey = "user.error.invalid_id"
	UserErrorJSONFormat MessageKey = "user.error.json_format"
	UserErrorInternal   MessageKey = "user.error.internal"

	// 系统相关
	UserHealthCheckSuccess MessageKey = "user.health.check_success"
)

// messages 消息映射表
var messages = map[MessageKey]map[Language]string{
	// 日志消息（英文）
	LogAuthFailedNoToken: {
		LanguageEn: "Authentication failed: no token provided",
		LanguageZh: "认证失败：未提供认证令牌",
	},
	LogAuthFailedInvalidFmt: {
		LanguageEn: "Authentication failed: invalid token format",
		LanguageZh: "认证失败：令牌格式错误",
	},
	LogAuthFailedInvalid: {
		LanguageEn: "Authentication failed: token validation failed",
		LanguageZh: "认证失败：令牌验证失败",
	},
	LogAuthSuccess: {
		LanguageEn: "Authentication successful",
		LanguageZh: "认证成功",
	},
	LogRequestCost: {
		LanguageEn: "Request processing time",
		LanguageZh: "请求处理耗时",
	},
	LogResponseBody: {
		LanguageEn: "Response body",
		LanguageZh: "响应体",
	},
	LogPanicRecovered: {
		LanguageEn: "Panic recovered",
		LanguageZh: "Panic已恢复",
	},
	LogInternalError: {
		LanguageEn: "Internal server error",
		LanguageZh: "内部服务器错误",
	},

	// 用户消息（中文，用于API响应）
	UserAuthNoToken: {
		LanguageZh: "未提供认证令牌",
		LanguageEn: "No authentication token provided",
	},
	UserAuthInvalidFmt: {
		LanguageZh: "认证令牌格式错误",
		LanguageEn: "Invalid authentication token format",
	},
	UserAuthInvalid: {
		LanguageZh: "无效的认证令牌",
		LanguageEn: "Invalid authentication token",
	},
	UserCreateSuccess: {
		LanguageZh: "创建成功",
		LanguageEn: "Created successfully",
	},
	UserRegisterSuccess: {
		LanguageZh: "注册成功",
		LanguageEn: "Registration successful",
	},
	UserGetSuccess: {
		LanguageZh: "获取成功",
		LanguageEn: "Retrieved successfully",
	},
	UserGetAllSuccess: {
		LanguageZh: "获取成功",
		LanguageEn: "Retrieved successfully",
	},
	UserUpdateSuccess: {
		LanguageZh: "更新成功",
		LanguageEn: "Updated successfully",
	},
	UserDeleteSuccess: {
		LanguageZh: "删除成功",
		LanguageEn: "Deleted successfully",
	},
	UserLoginSuccess: {
		LanguageZh: "登录成功",
		LanguageEn: "Login successful",
	},
	UserErrorBadRequest: {
		LanguageZh: "请求参数错误",
		LanguageEn: "Bad request",
	},
	UserErrorInvalidID: {
		LanguageZh: "无效的用户ID: %s",
		LanguageEn: "Invalid user ID: %s",
	},
	UserErrorJSONFormat: {
		LanguageZh: "JSON格式错误",
		LanguageEn: "Invalid JSON format",
	},
	UserErrorInternal: {
		LanguageZh: "内部服务器错误",
		LanguageEn: "Internal server error",
	},
	UserHealthCheckSuccess: {
		LanguageZh: "服务运行正常",
		LanguageEn: "Service is running normally",
	},
}

// LogMessage 获取日志消息（始终返回英文）
func LogMessage(key MessageKey) string {
	if msg, ok := messages[key]; ok {
		if en, ok := msg[LanguageEn]; ok {
			return en
		}
	}
	return string(key) // 如果找不到，返回key本身
}

// UserMessage 获取用户消息（默认中文，可通过lang参数指定）
func UserMessage(key MessageKey, lang ...Language) string {
	language := LanguageZh // 默认中文
	if len(lang) > 0 {
		language = lang[0]
	}

	if msg, ok := messages[key]; ok {
		if text, ok := msg[language]; ok {
			return text
		}
		// 如果指定语言不存在，尝试返回中文
		if text, ok := msg[LanguageZh]; ok {
			return text
		}
	}
	return string(key) // 如果找不到，返回key本身
}

// UserMessagef 获取格式化的用户消息（类似fmt.Sprintf）
func UserMessagef(key MessageKey, args ...interface{}) string {
	msg := UserMessage(key)
	// 这里可以添加格式化逻辑，暂时先返回原始消息
	return msg
}
