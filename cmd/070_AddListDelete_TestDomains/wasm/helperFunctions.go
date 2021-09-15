package main

import (
	"fmt"
	"runtime"
)

// Retrieve trace information regarding file, line and function of the caller
func (mt *MagicTable) trace(printToStandardOutput bool) string {
	var returnMessage string

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Sprintf("?", 0, "?")
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fmt.Sprintf(file, line, "?")
	}

	//fmt.Printf("%s:%d %s\n", file, line, fn.Name())
	returnMessage = fmt.Sprintf(file, line, fn.Name())

	if printToStandardOutput == true {
		fmt.Println(returnMessage)
	}

	return returnMessage
}
