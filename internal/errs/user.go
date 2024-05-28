package errs

import "errors"

var (
	EmptyUsername      = errors.New("用户名为空")
	EmptyPassword      = errors.New("密码是空的")
	WrongPassword      = errors.New("密码不正确")
	DeleteAdminOrGuest = errors.New("无法删除管理员或客人")
	ErrRecordNotFound  = errors.New("找不到记录")
)
