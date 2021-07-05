package solution

import ( "github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	x := emoji.Sprint("Hello from :flag_pl:!")
	return x
}
