package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() (message string) {
	message = emoji.Sprint("Hello :world_map:!")
	return
}
