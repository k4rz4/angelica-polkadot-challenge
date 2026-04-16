package main

import (
	_ "embed"
	"fmt"
)

//go:embed angelica.txt
var angelica string

type tildeRun struct {
	line, start, length int
}

type features struct {
	tildes   []tildeRun
	eyeLine  int // -1 if none
	pupils   int // ( + ) chars on eye line
	firstDot int // -1 if none
}

// computePolkadotScore is the entry point required by the challenge.
func computePolkadotScore() int {
	return scoreArt(angelica)
}

func scoreArt(art string) int {
	f := scan(art)
	lStart, lEnd := findLips(f)
	inside, outside := countDots(art, lStart, lEnd)
	return outside + inside*f.pupils
}

// scan walks the art once and collects the landmarks needed for scoring.
func scan(art string) features {
	f := features{eyeLine: -1, firstDot: -1}
	line, x := 0, 0
	runStart, runLen := 0, 0
	open, closes := 0, 0

	flushRun := func() {
		if runLen >= 6 {
			f.tildes = append(f.tildes, tildeRun{line, runStart, runLen})
		}
		runLen = 0
	}
	endLine := func() {
		flushRun()
		if f.eyeLine == -1 && open >= 2 && open == closes {
			f.eyeLine = line
			f.pupils = open + closes
		}
	}

	for _, ch := range art {
		if ch == '\n' {
			endLine()
			line++
			x, open, closes = 0, 0, 0
			continue
		}
		if ch == '~' {
			if runLen == 0 {
				runStart = x
			}
			runLen++
		} else {
			flushRun()
			switch ch {
			case 'O':
				if f.firstDot == -1 {
					f.firstDot = line
				}
			case '(':
				open++
			case ')':
				closes++
			}
		}
		x++
	}
	endLine()
	return f
}

func findLips(f features) (start, end int) {
	for _, r := range f.tildes {
		belowEyes := f.eyeLine == -1 || r.line > f.eyeLine
		aboveDots := f.firstDot == -1 || r.line < f.firstDot
		if belowEyes && aboveDots {
			return r.start, r.start + r.length - 1
		}
	}
	return 0, 0
}

func lipsRange(art string) (start, end int) {
	return findLips(scan(art))
}

func pupilCount(art string) int {
	return scan(art).pupils
}

func countDots(art string, colStart, colEnd int) (inside, outside int) {
	x := 0
	for _, ch := range art {
		if ch == '\n' {
			x = 0
			continue
		}
		if ch == 'O' {
			if x >= colStart && x <= colEnd {
				inside++
			} else {
				outside++
			}
		}
		x++
	}
	return
}

func main() {
	f := scan(angelica)
	lStart, lEnd := findLips(f)
	inside, outside := countDots(angelica, lStart, lEnd)
	score := outside + inside*f.pupils

	fmt.Printf("Lips range:         cols %d-%d\n", lStart, lEnd)
	fmt.Printf("Pupil chars:        %d\n", f.pupils)
	fmt.Printf("Outside lips range: %d\n", outside)
	fmt.Printf("Inside lips range:  %d\n", inside)
	fmt.Printf("Score: %d + (%d * %d) = %d\n", outside, inside, f.pupils, score)
}
