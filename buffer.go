package stdres

import "fmt"

var colorDisabled bool

func DisableColor() {
	colorDisabled = true
}

func EnableColor() {
	colorDisabled = false
}

type Result uint

// Foreground colors
const (
	reset   = "\033[00m"
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

const (
	UNKNOWN Result = iota
	PENDING
	FAILURE
	SUCCESS
	INFO
	PLAIN
)

type Buffer struct {
	buffer []*Record
}

// Record struct keeps record of message to print and current result
type Record struct {
	Result
	Message string
}

func (out *Record) String() string {
	return string(out.Message)
}

// Flash iterates through all records in Buffer and flush it to stdout preceded by color matching current status.
// Each record will be followed by color reset if it was preceded by a color.
func (outBuffer *Buffer) Flush() {
	toFlush := outBuffer.buffer
	outBuffer.buffer = make([]*Record, 0, 500)

	// Print with colors disabled
	if colorDisabled {
		for _, out := range toFlush {
			message := out.String()
			fmt.Print(message)
		}

		return
	}

	// Print with colors enabled
	for _, record := range toFlush {
		var color string

		switch record.Result {
		case UNKNOWN:
			color = cyan
		case PENDING:
			color = yellow
		case SUCCESS:
			color = green
		case FAILURE:
			color = red
		case INFO:
			color = blue
		case PLAIN:
			color = white
		}

		fmt.Print(color, record.String(), reset)
	}
}

func (outBuffer *Buffer) printer(message string, newline bool) *Record {
	if outBuffer.buffer == nil {
		outBuffer.buffer = make([]*Record, 0, 500)
	}

	if newline {
		message = message + "\n"
	}

	out := &Record{UNKNOWN, message}
	outBuffer.buffer = append(outBuffer.buffer, out)

	return out
}

// Println records message to buffer to be printed later.
// Text not printed until Buffer.Flush has been called.
// It returns current status and text string as Record.
func (outBuffer *Buffer) Print(message string) *Record {
	return outBuffer.printer(message, false)
}

// Println records message followed by newline to buffer to be printed later
// Text not printed until Buffer.Flush has been called.
// It returns current status and text string as Record.
func (outBuffer *Buffer) Println(message string) *Record {
	return outBuffer.printer(message, true)
}
