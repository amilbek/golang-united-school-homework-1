package solution

import (
	"flag"
	"github.com/kyokomi/emoji"
)

func main() {
	emoji.Print(GetMessage())
}

func GetMessage() string {
	emojiName := flag.String("e", "Hello :world_map:!", "emoji name")
	flag.Parse()

	return *emojiName
}
