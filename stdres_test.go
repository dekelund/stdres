package stdres_test

import "github.com/dekelund/stdres"

func Example() {
	var buffer stdres.Buffer
	buffer = stdres.Buffer{}

	firstAction := buffer.Println("Executing first action")
	secondAction := buffer.Println("Executing second action")
	thirdAction := buffer.Println("Executing third action")
	fourthAction := buffer.Println("Executing fourth action")
	buffer.Print("Executing fifth action") // Never updated, printed with UNKNOWN state which result in Cyan color

	firstAction.Result = stdres.SUCCESS  // First action succeeded, line printed in green
	secondAction.Result = stdres.FAILURE // Second action failed, line printed in red
	thirdAction.Result = stdres.INFO     // Third action was just information, line printed in blue
	fourthAction.Result = stdres.PLAIN   // Fourth action printed with white text

	stdres.DisableColor() // Use this method to remove color output
	buffer.Flush()        // Only print if and when we flush, update result before flush.

	// Output:
	// Executing first action
	// Executing second action
	// Executing third action
	// Executing fourth action
	// Executing fifth action
}
