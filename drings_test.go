package drings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLen(t *testing.T) {
	testdata := []struct {
		array    []string
		expected int
	}{
		{
			array: []string{
				"apple",
				"banana",
				"carrot",
				"duck",
				"elephant",
			},
			expected: 8,
		}, {
			array: []string{
				"apple",
			},
			expected: 5,
		}, {
			array:    []string{},
			expected: 0,
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("maxLen #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, MaxLen(test.array))
		})
	}
}

func TestDedup(t *testing.T) {
	testdata := []struct {
		in       []string
		expected []string
	}{
		{
			in: []string{
				"Around the world, around the world",
				"Around the world, around the world",
				"Around the world, around the world",
				"Around the world, around the world",
			},
			expected: []string{
				"Around the world, around the world",
			},
		}, {
			in: []string{
				"apple",
				"banana",
				"carrot",
				"duck",
				"elephant",
			},
			expected: []string{
				"apple",
				"banana",
				"carrot",
				"duck",
				"elephant",
			},
		}, {
			in: DeepCopy(technologicLines),
			expected: []string{
				"Buy it, use it, break it, fix it,",
				"Charge it, point it, zoom it, press it,",
				"Cross it, crack it, twitch, update it,",
				"Drag and drop it, zip, unzip it,",
				"Load it, check it, quick, rewrite it",
				"Lock it, fill it, curl it, find it,",
				"Name it, read it, tune it, print it,",
				"Plug it, play it, burn it, rip it,",
				"Scan it, send it, fax, rename it,",
				"Snap it, work it, quick, erase it,",
				"Surf it, scroll it, pose it, click it",
				"Technologic, technologic, technologic, technologic",
				"Technologic, technologic, technologic, technologic.",
				"Touch it, bring it, pay it, watch it,",
				"Touch it, scroll it, pose it, click it",
				"Trash it, change it, mail, upgrade it,",
				"Turn it, leave it, stop, format it.",
				"View it, code it, jam, unlock it",
				"Write it, cut it, paste it, save it,",
			},
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("dedup #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, Dedup(test.in))
		})
	}
}

func ExampleSplitAndTrimSpace() {
	testdata := `
Around the world, around the world
Around the world, around the world
Around the world, around the world
Around the world, around the world

			Around the world, around the world
			Around the world, around the world
			Around the world, around the world
			Around the world, around the world

Around the world, around the world
	Around the world, around the world
			Around the world, around the world
Around the world, around the world

Around the world, around the world
Around the world, around the world
		Around the world, around the world
Around the world, around the world
`
	// break up the lyrics for the Daft Punk song Around the World

	lyrics := SplitAndTrimSpace(testdata, "\n")

	deduped := Dedup(lyrics)

	fmt.Println(deduped)

	//Output:
	// [Around the world, around the world]

}

func TestSplitAndTrimSpace(t *testing.T) {
	// formattedLines, _ := json.MarshalIndent(technologicLines, "", "\t")
	// fmt.Println(string(formattedLines))

	// formattedSong, _ := json.MarshalIndent(SplitAndTrimSpace(technologicSong, "\n"), "", "\t")
	// fmt.Println(string(formattedSong))

	assert.Equal(t, technologicLines, SplitAndTrimSpace(technologicSong, "\n"))
}

func TestPadRight(t *testing.T) {
	testdata := []struct {
		in       string
		pad      string
		size     int
		expected string
	}{
		{
			in:       "apple",
			pad:      " ",
			size:     10,
			expected: "apple     ",
		}, {
			in:       "apple",
			pad:      " ",
			size:     5,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      " ",
			size:     4,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      " ",
			size:     0,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      "+",
			size:     8,
			expected: "apple+++",
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("PadRight #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, PadRight(test.in, test.pad, test.size))
		})
	}
}

func TestPadLeft(t *testing.T) {
	testdata := []struct {
		in       string
		pad      string
		size     int
		expected string
	}{
		{
			in:       "apple",
			pad:      " ",
			size:     10,
			expected: "     apple",
		}, {
			in:       "apple",
			pad:      " ",
			size:     5,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      " ",
			size:     4,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      " ",
			size:     0,
			expected: "apple",
		}, {
			in:       "apple",
			pad:      "+",
			size:     8,
			expected: "+++apple",
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("PadLeft #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, PadLeft(test.in, test.pad, test.size))
		})
	}
}

func TestPadAllRight(t *testing.T) {
	testdata := []struct {
		in       []string
		pad      string
		size     int
		expected []string
	}{
		{
			in: []string{
				"apple",
				"banana",
				"carrot cake",
				"dog",
			},
			pad:  "+",
			size: 10,
			expected: []string{
				"apple+++++",
				"banana++++",
				"carrot cake",
				"dog+++++++",
			},
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("PadAllRight #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, PadAllRight(test.in, test.pad, test.size))
		})
	}

}

func TestPadAllLeft(t *testing.T) {
	testdata := []struct {
		in       []string
		pad      string
		size     int
		expected []string
	}{
		{
			in: []string{
				"apple",
				"banana",
				"carrot cake",
				"dog",
			},
			pad:  "+",
			size: 10,
			expected: []string{
				"+++++apple",
				"++++banana",
				"carrot cake",
				"+++++++dog",
			},
		},
	}

	for id, test := range testdata {
		t.Run(fmt.Sprintf("PadAllLeft #%d", id), func(t *testing.T) {
			assert.Equal(t, test.expected, PadAllLeft(test.in, test.pad, test.size))
		})
	}

}
