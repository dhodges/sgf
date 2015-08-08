package parse

import (
	"testing"

  "github.com/stretchr/testify/assert"
)

var lexFixture = "" +
	"(;CA[UTF-8]SZ[19]EV[The Game of the Century]DT[1933-10-16]PB[Go Seigen]BR[5p]PW[Honinbo Shusai]WR[9p]PC[Tokyo, Japan]RE[W+2]KM[0]C[This match was sponsored by Yomiuri Newspaper. In the preliminary game, Go Seigen defeated the strong opponents Kitani Minoru 6p and Hashimoto Utaro 5p, thus earned the right to play this memorable game with Honinbo Shusai Meijin who hadn't played any game for almost ten years. This game started in Oct. 1933 and didn't finish until Feb. of next year, during the process the game was adjourned for more than 10 times, and it created a furore in this Newspaper and at that time it was also regarded as the game of the century. In the recent Go world, we never saw a game which had produced such a great impact.];B[qc]C[At that time, Go Seigen was just 20 years' old, in this match with Shusai Meijin, he started with the unprecedented opening of 3-3, star and tengen, which shocked the Go world. This special play was once criticized by some people: this is not polite to Meijin. As for today, this kind of criticism is just worth laughing. However, Mr. Go always holds firmly to his play, because Mr. Go just wanted to break away from all Shusaku's openings of 1, 3 and 5 and establish his own style, which is the \"New Fuseki\" he invented together with Kitani Minoru, so this game had another special meaning, it was a historical game between the new and old opening. In the game, Mr. Go had often made Meijin to cudgel his brain, and due to health cause, the game was adjourned for so many times (there was no sealing of moves, it meant that Shusai Meijin could adjourn the game at his will and continue to study the position leisurely at home with the whole Honinbo clique). In the end, although Mr. Go Seigen lost by two points, but it indeed let the Go world to acknowledge the \"New Fuseki\" and this meaningful match.];W[cd];B[dp];W[pq];B[jj])"

func TestLex(t *testing.T) {
	l := lex(lexFixture)
	for {
		item := l.nextItem()
		if item.typ == itemEOF {
			break
		}
	}
}

type example struct {
	example string
	message string
}

var invalidExamples = []example{
	{"(;CA)", "missing left bracket '['"},
	{"(;CA[UTF-8", "missing right bracket ']'"},
	{"(;CA[UTF-8];)", "missing property"},
	{"(CA[UTF-8])", "missing semi-colon"},
}

func TestLexFilterNewlines(t *testing.T) {
	l := lex("(;CA[UTF-8]\n\rSZ[19]\r\nEV[The Game of the Century]")

	assert.Equal(t, l.input, "(;CA[UTF-8]SZ[19]EV[The Game of the Century]", "expected lexer to strip newlines")
}

func TestLexErrors(t *testing.T) {
	for ndx := range invalidExamples {
		l := lex(invalidExamples[ndx].example)
		var i item
		for {
			i = l.nextItem()
			if i.typ == itemEOF || i.typ == itemError {
				break
			}
		}
		assert.Equal(t, i.typ, itemError, "expected an error")
	}
}
