package models

// Login ShouldBind 会按顺序合并 Query + Form + JSON + Header + Uri，一份代码多端通用
type Login struct {
	Username string `form:"username" json:"username" binding:"required" header:"username" uri:"user"`
	Password string `form:"password" json:"password"  binding:"required" header:"password" uri:"user"`
}
