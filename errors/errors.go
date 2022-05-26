package errors

import (
	"github.com/hysios/utils/errors"

	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"gorm.io/gorm"
)

var (
	ErrNonimplement       = errors.New("nonimplement")
	ErrLoginUserNotExists = errors.New("login user not exists")
	ErrIncorrectPassword  = errors.New("password is incorrect")
	ErrAccessDenied       = oerrors.ErrAccessDenied
	ErrRecordNotFound     = gorm.ErrRecordNotFound
	ErrUserDisabled       = errors.New("user disabled")
)

func init() {
	errors.RegisterErrCode(ErrNonimplement, errors.ErrAuto)      // 未实现
	errors.RegisterErrCode(ErrLoginUserNotExists, 10001)         // 登陆用户名不存在
	errors.RegisterErrCode(ErrIncorrectPassword, 10002)          // 不正确的密码
	errors.RegisterErrCode(oerrors.ErrInvalidAccessToken, 10003) // 无效的 AccessToken
	errors.RegisterErrCode(ErrAccessDenied, 10004)               // 访问被拒绝
	errors.RegisterErrCode(ErrRecordNotFound, 10007)             // 记录没有找到
	errors.RegisterErrCode(ErrUserDisabled, 10013)               // 用户被禁
}
