package testing

import (
	"fmt"
	"runtime"
	"testing"
)

func TestStandard(t *testing.T) {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	trace()

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}

func trace() {

	pc := make([]uintptr, 10) // at least 1 entry needed

	runtime.Callers(2, pc)

	f := runtime.FuncForPC(pc[0])

	file, line := f.FileLine(pc[0])

	function_name := f.Name()

	fmt.Printf("file: %s:\nfunction name: %s\nline: %d\nfunction entry: %d\n", file, function_name, line, f.Entry())
}
