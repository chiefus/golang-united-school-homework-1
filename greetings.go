package chiefuslecture00

import (
	"github.com/kyokomi/emoji/v2"
)

func GetGreetings() (greetings string) {
	greetings = emoji.Sprint("Hello :world_map:!")
	return
}
