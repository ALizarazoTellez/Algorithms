package bf

import "os"

func ExampleInterpreter() {
	const prog = `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]
		      <+.+++++++..+++.>>++++++[<+++++++>-]<++.--
		      ----------.>++++++[<+++++++++>-]<+.<.+++.-
		      -----.--------.>>>++++[<++++++++>-]<+.`

	intrpr := NewInterpreter(os.Stdout, os.Stdin)
	if err := intrpr.Execute([]byte(prog)); err != nil {
		panic(err)
	}

	// Output: Hello, World!
}
