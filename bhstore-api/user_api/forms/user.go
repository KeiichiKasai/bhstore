package forms

type RegisterForm struct {
	Mobile   string `form:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" binding:"required,min=3,max=20"`
	Nickname string `form:"nickname" binging:"required"`
}

type LoginForm struct {
	Mobile   string `form:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" binding:"required,min=3,max=20"`
}
