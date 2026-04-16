package main

import (
	_ "embed"
	"fmt"
	"strings"
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

// scan walks the art line by line and records the landmarks.
func scan(art string) features {
	f := features{eyeLine: -1, firstDot: -1}
	for i, line := range strings.Split(art, "\n") {
		f.scanLine(i, line)
	}
	return f
}

func (f *features) scanLine(lineNum int, line string) {
	runStart, runLen := 0, 0
	opens, closes := 0, 0
	x := 0
	for _, ch := range line {
		if ch == '~' {
			if runLen == 0 {
				runStart = x
			}
			runLen++
			x++
			continue
		}
		f.collectRun(lineNum, runStart, runLen)
		runLen = 0
		switch ch {
		case 'O':
			if f.firstDot == -1 {
				f.firstDot = lineNum
			}
		case '(':
			opens++
		case ')':
			closes++
		}
		x++
	}
	f.collectRun(lineNum, runStart, runLen)
	if f.eyeLine == -1 && opens >= 2 && opens == closes {
		f.eyeLine = lineNum
		f.pupils = opens + closes
	}
}

func (f *features) collectRun(line, start, length int) {
	if length >= 6 {
		f.tildes = append(f.tildes, tildeRun{line, start, length})
	}
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
