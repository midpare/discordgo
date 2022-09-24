package language

import (
	. "golang.org/x/text/language"
	"golang.org/x/text/message"
)

func SetMessage(id string, data string) {
	message.SetString(Korean, id, data)
}
