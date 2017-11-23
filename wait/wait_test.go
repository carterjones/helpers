package wait_test

import (
	"testing"
	"time"

	"github.com/carterjones/helpers/wait"
)

func Test_UntilNextMinute(t *testing.T) {
	wait.UntilNextMinute(nil)
	now := time.Now()
	if now.Second() != 0 {
		t.Error("seconds value after waiting is not equal to 0")
	}
}
