package errs

import "errors"

var (
	PermissionDenied = errors.New("权限不够")
)
