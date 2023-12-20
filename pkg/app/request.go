package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/moshushu/gin-blog/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(err.Key, err.Message)
	}
}
