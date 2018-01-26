package main

var (
	initInterval = 1000
	interval     = initInterval
	maxInterval  = 100000
)

func main() {}

// What does this function do?
// In what kind of scenario such function can be useful?
func computeInterval(err error) int {
	if err != nil {
		if interval < maxInterval {
			interval = interval * 2
		} else {
			interval = maxInterval
		}
	} else {
		interval = initInterval
	}
	return interval
}
