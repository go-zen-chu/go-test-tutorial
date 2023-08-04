package bug

import (
	"fmt"
	"time"
)

var (
	now time.Time
)

func init() {
	now = time.Now()
}

func Bug3() error {
	sec := now.Second()
	if sec%3 == 0 {
		return fmt.Errorf("sorry we got bug3...: %d", sec)
	}
	return nil
}

func Bug4() error {
	sec := now.Second()
	if sec%4 == 0 {
		return fmt.Errorf("sorry we got bug4...: %d", sec)
	}
	return nil
}
