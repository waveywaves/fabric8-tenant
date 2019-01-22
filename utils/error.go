package utils

import "fmt"

func ListErrorsInMessage(errorChan chan error) string {
	index := 1
	msg := ""
	for er := range errorChan {
		if er != nil {
			msg += fmt.Sprintf("\n#%d: %s", index, er.Error())
			index++
		}
	}
	return msg
}
