package emoji

import "testing"

var notEmojiTest = []string{
	"a",
	"😀b",
	"b😀",
	"\n",
	"test",
	".",
	"😀👍",
	"🇧",
}

var emojiTest = []string{
	"😀",
	"👍",
	"⛰️",
	"🏕️",
	"🏥",
	"🛢️",
	"💈",
	"⛱️",
	"🪁",
	"☎️",
	"💡",
	"💳",
	"🖌️",
	"🔓",
	"⛓️",
	"🧴",
	"🍌",
	"🍓",
	"🥔",
	"🧅",
	"🥐",
	"🥞",
	"🍳",
	"🍠",
	"🥃",
	"🍽️",
	"😆",
	"😍",
	"😚",
	"😬",
	"🤯",
	"💀",
	"😫",
	"✌️",
	"🧠",
	"🧙",
	"💇",
	"🎒",
	"⛑️",
	"🥺",
	"😂",
	"😊",
	"🔥",

	"🙍‍♂️",
	"👩‍🎓",
	"🧑‍🏫",
	"🧑‍⚕️",
	"🧑‍🍳",
	"👨‍🍳",
	"👮‍♀️",
	"🧙‍♂️",
	"🧟‍♂️",
	"👩‍🦯",
	"👨‍👩‍👦‍👦",
	"👩‍👩‍👧‍👧",
	"👨‍👧‍👦",

	"👋🏼",
	"🖖🏿",
	"🦻🏻",
	"🧏🏼‍♀️",
	"🧜🏼‍♀️",
	"🧝🏼‍♀️",
	"👯🏼‍♀️",
	"👩‍❤️‍👩",
	"👩🏾‍👨🏾‍👦🏾",

	"🏳️",
	"🎌",
	"🇧🇳",
	"🏳️‍⚧️",
	"🇧🇸",
	"🇪🇨",
	"🇭🇲",
	"🇱🇮",
	"🇳🇬",
	"🇸🇬",
	"🇺🇬",
	"🏴󠁵󠁳󠁴󠁸󠁿",
	"🏴󠁧󠁢󠁳󠁣󠁴󠁿",
}

func Test_PossibleGlyph(t *testing.T) {
	for _, s := range notEmojiTest {
		if PossibleGlyph(s) {
			t.Errorf("%q returned positive", s)
		}
	}

	for _, s := range emojiTest {
		if !PossibleGlyph(s) {
			t.Errorf("%q returned negative", s)
		}
	}
}
