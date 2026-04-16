# angelica polkadot challenge

Solution for the Angelica polkadot challenge. The answer is **71**.

## Problem

Count polkadots (`O`) on a dress in an ASCII art image. Polkadots whose column
falls within the lip range are weighted by the number of pupil characters.

```
score = outside + inside * pupils
```

See the [original challenge](https://gist.github.com/kmdupr33/8bab92e762d367de9455183abe04f38f) for the full problem statement and ASCII art.

## Usage

```bash
make run    # prints detected landmarks and final score
make test   # runs all tests
make build  # compiles to ./angelica binary
make clean  # removes compiled binary
```

## Approach

Two O(n) passes over the art:

**Pass 1 — `scan()`**: walks the art line by line, building a `features` struct with three landmarks:
- all tilde runs of length ≥ 6 (lips, hair, and hem candidates)
- the first line with ≥ 2 matched `()` pairs (the eye line) and its bracket count (pupil count)
- the first line containing an `O` (top of the polkadot region)

The lips are selected as the tilde run that falls **between** the eye line and the first polkadot row. (immune to horizontal shifts and feature reordering)

**Pass 2 — `countDots()`**: partitions every `O` into inside or outside the detected lip column range.

## Files

| File | Purpose |
|------|---------|
| `polkadot.go` | all logic |
| `polkadot_test.go` | unit + integration tests |
| `angelica.txt` | ASCII art, embedded at compile time via `//go:embed` |
| `Makefile` | build, test, run, clean |
