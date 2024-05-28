package errs

import "fmt"

var (
	SearchNotAvailable = fmt.Errorf("搜索未启用")
)
