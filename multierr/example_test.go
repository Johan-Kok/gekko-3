
package multierr

import (
	"errors"
	"fmt"
)

func ExampleCombine() {
	err := Combine(
		errors.New("call 1 failed"),
		nil, // successful request
		errors.New("call 3 failed"),
		nil, // successful request
		errors.New("call 5 failed"),
	)
	fmt.Printf("%+v", err)
	// Output:
	// the following errors occurred:
	//  -  call 1 failed
	//  -  call 3 failed
	//  -  call 5 failed
}

func ExampleAppend() {
	var err error
	err = Append(err, errors.New("call 1 failed"))
	err = Append(err, errors.New("call 2 failed"))
	fmt.Println(err)
	// Output:
	// call 1 failed; call 2 failed
}

func ExampleErrors() {
	err := Combine(
		nil, // successful request
		errors.New("call 2 failed"),
		errors.New("call 3 failed"),
	)
	err = Append(err, nil) // successful request
	err = Append(err, errors.New("call 5 failed"))

	errors := Errors(err)
	for _, err := range errors {
		fmt.Println(err)
	}
	// Output:
	// call 2 failed
	// call 3 failed
	// call 5 failed
}

func ExampleAppendInto() {
	var err error

	if AppendInto(&err, errors.New("foo")) {
		fmt.Println("call 1 failed")
	}

	if AppendInto(&err, nil) {
		fmt.Println("call 2 failed")
	}

	if AppendInto(&err, errors.New("baz")) {
		fmt.Println("call 3 failed")
	}

	fmt.Println(err)
	// Output:
	// call 1 failed
	// call 3 failed
	// foo; baz
}
