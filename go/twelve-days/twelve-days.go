package twelve

import "fmt"

const testVersion = 1

type verse struct {
	dayCountWord string
	verseEnd     string
}

var versesData = []verse{
	{"first", " a Partridge in a Pear Tree."},
	{"second", ", two Turtle Doves"},
	{"third", ", three French Hens"},
	{"fourth", ", four Calling Birds"},
	{"fifth", ", five Gold Rings"},
	{"sixth", ", six Geese-a-Laying"},
	{"seventh", ", seven Swans-a-Swimming"},
	{"eighth", ", eight Maids-a-Milking"},
	{"ninth", ", nine Ladies Dancing"},
	{"tenth", ", ten Lords-a-Leaping"},
	{"eleventh", ", eleven Pipers Piping"},
	{"twelfth", ", twelve Drummers Drumming"},
}

func Song() string {
	song := ""
	for i, _ := range versesData {
		song += Verse(i+1) + "\n"
	}
	return song
}

func Verse(lineNb int) string {
	endVerse := ""
	if lineNb > 1 {
		verses := versesData[1:lineNb]
		max := len(verses)
		for i := max - 1; i >= 0; i-- {
			endVerse += verses[i].verseEnd
		}
	}
	if lineNb != 1 {
		endVerse += ", and"
	} else {
		endVerse += ","
	}
	endVerse += versesData[0].verseEnd
	return fmt.Sprintf("On the %[1]s day of Christmas my true love gave to me%[2]s", versesData[lineNb-1].dayCountWord, endVerse)
}
