package service

import (
	"ProjModules/utils/logging"
	"ProjModules/utils/validation"
)

func MarkErrors(errors []*validation.Error)  {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
