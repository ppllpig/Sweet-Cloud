package errs

import (
	"errors"

	pkgerr "github.com/pkg/errors"
)

var (
	ObjectNotFound = errors.New("未找到对象")
	NotFolder      = errors.New("不是文件夹")
	NotFile        = errors.New("不是文件")
)

func IsObjectNotFound(err error) bool {
	return errors.Is(pkgerr.Cause(err), ObjectNotFound)
}
