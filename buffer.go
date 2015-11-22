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
	RESET   = "\033[00m"
	BLACK   = "\033[30m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[37m"
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
			color = CYAN
		case PENDING:
			color = YELLOW
		case SUCCESS:
			color = GREEN
		case FAILURE:
			color = RED
		case INFO:
			color = BLUE
		case PLAIN:
			color = WHITE
		}

		fmt.Print(color, record.String(), RESET)
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
