package colors

import (
	"testing"
)

func TestGreen(t *testing.T) {
	str := "color me"
	expected := COLORS["GREEN"] + "color me" + COLORS["RESET"]
	actual := Green(str)

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s\n", expected, actual)
	}
}

func TestRed(t *testing.T) {
	str := "color me"
	expected := COLORS["RED"] + "color me" + COLORS["RESET"]
	actual := Red(str)

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s\n", expected, actual)
	}
}

func TestStrLen(t *testing.T) {
	str := Green("color me green. ") + Red("color me red")
	expected := 28
	actual := StrLen(str)

	if expected != actual {
		t.Fatalf("\nExpected:\t%d\nActual:\t\t%d\n", expected, actual)
	}
}

func TestLeft(t *testing.T) {
	width := 38
	unpadded := Green("color me green. ") + Red("color me red")
	expected := Green("color me green. ") + Red("color me red") + "          "
	actual := Left(unpadded, width)

	if expected != actual {
		t.Fatalf("\nExpected:\t'%s'\nActual:\t\t'%s'\n", expected, actual)
	}
}

func TestLocate(t *testing.T) {
	str := "normal " + Green("green") + " normal " + Red("red") + " normal"
	expected := []Location{
		{
			Start: 7,
			End:   11,
		},
		{
			Start: 17,
			End:   20,
		},
		{
			Start: 29,
			End:   33,
		},
		{
			Start: 37,
			End:   40,
		},
	}
	actual := LocateAll(str)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
		}
	}
}

func TestSkip(t *testing.T) {
	str := "no color " + Green("green") + " no color " + Red("red") + " no color"
	expected := "no color green no color red no color"
	var actual string

	locations := LocateAll(str)

	for i := 0; i < len(str); i++ {
		if Skip(&i, &locations) {
			continue
		}

		actual += string(str[i])
	}

	if expected != actual {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}
