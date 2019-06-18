package lshared

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-06-18

import (
	"testing"
)

func TestShared(t *testing.T) {

	sh := New(16)

	sh.Lock([]byte("12"))
	sh.UnLock([]byte("12"))

	sh.RLock([]byte("12"))
	sh.RUnLock([]byte("12"))
}
