# angelica polkadot challenge

Solution for the Angelica polkadot challenge. The answer is 71.

## Problem

Count polkadots on a dress in an ASCII art image. Polkadots whose column falls
between the lip columns are multiplied by the number of characters used to
represent the pupils.

```
score = outside + inside * pupils
```

See the original challenge for the full wording.

## Run

```bash
go run .
```

Prints the detected landmarks and the final score.

## Test

```bash
go test ./...
```

## Approach

Single pass over the art builds a `features` struct with three landmarks:

- all tilde runs of length 6 or more
- the first line that has matching `()` pairs (the eye line)
- the first line that contains an `O` (top of the dress)

The lips are the tilde run that sits between the eye line and the first
polkadot. Hair and hem runs fall outside that band so they are ignored.
The pupil count is the number of `(` and `)` characters on the eye line.

Once the lip range is known, a second pass partitions every `O` into inside
or outside the lip columns.

## Files

- `polkadot.go` main code
- `polkadot_test.go` tests
- `angelica.txt` the ASCII art, embedded at compile time via `//go:embed`
