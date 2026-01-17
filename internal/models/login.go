package models

// Login ShouldBind 会按顺序合并 Query + Form + JSON + Header + Uri，一份代码多端通用
type Login struct {
	Username string `form:"username" json:"username" binding:"required" header:"username" uri:"user"`
	Password string `form:"password" json:"password"  binding:"required" header:"password" uri:"user"`
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
