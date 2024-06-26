package errs

import (
	"errors"
	pkgerr "github.com/pkg/errors"
	"testing"
)

func TestErrs(t *testing.T) {

	err1 := NewErr(StorageNotFound, "请先添加一个存储")
	t.Logf("err1: %s", err1)
	if !errors.Is(err1, StorageNotFound) {
		t.Errorf("failed, expect %s is %s", err1, StorageNotFound)
	}
	if !errors.Is(pkgerr.Cause(err1), StorageNotFound) {
		t.Errorf("failed, expect %s is %s", err1, StorageNotFound)
	}
	err2 := pkgerr.WithMessage(err1, "找不到存储")
	t.Logf("err2: %s", err2)
	if !errors.Is(err2, StorageNotFound) {
		t.Errorf("failed, expect %s is %s", err2, StorageNotFound)
	}
	if !errors.Is(pkgerr.Cause(err2), StorageNotFound) {
		t.Errorf("failed, expect %s is %s", err2, StorageNotFound)
	}
}
