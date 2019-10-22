package nesting

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNestingFunc(t *testing.T) {
	convey.Convey("Test nesttingFunc", t, func() {
		nestingFunc()
	})
}
