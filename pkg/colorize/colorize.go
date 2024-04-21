package colorize

var (
	reset   = "\033[0m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cWhite  = "\033[37m"
)

func C(color string, msg string) string {
	switch color {
	case "red":
		return cRed + msg + reset
	case "green":
		return cGreen + msg + reset
	case "yellow":
		return cYellow + msg + reset
	case "blue":
		return cBlue + msg + reset
	case "purple":
		return cPurple + msg + reset
	case "cyan":
		return cCyan + msg + reset
	default:
		return cWhite + msg + reset
	}
}
