// BUG(Daniel Ekelund): Current version is not thread-safe

// Package stdres provides the means to select colors for STDOUT based on action results.
// By using this package you are able to print text to the buffer, and postpone
// the color selection until action has finished.
//
// Typical scenario:
// The tool decides what to do, and the tool informs the user what's going to happen.
// The tool starts to run the command and print more information.
// During execution of the action, the tool fails and sets the result to failure.
//
// Before the tool starts to exuecute the next action, it calls method Buffer.Flush
// and the existing records and flushed in corresponding colors, in this case red text
// due to failure.
package stdres
