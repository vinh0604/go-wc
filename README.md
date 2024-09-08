# wc implementation in Go

This is a simple implementation of the `wc` command in Go, for the coding challenge at [https://codingchallenges.fyi/challenges/challenge-wc](https://codingchallenges.fyi/challenges/challenge-wc).

## Usage

```bash
go-wc -c data/test.txt

go-wc -l data/test.txt

go-wc -w data/test.txt

go-wc -m data/test.txt

go-wc data/test.txt

cat data/test.txt | go-wc -l
```