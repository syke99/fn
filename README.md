# fn
[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/fn.svg)](https://pkg.go.dev/github.com/syke99/fn)
[![go reportcard](https://goreportcard.com/badge/github.com/syke99/fn)](https://goreportcard.com/report/github.com/syke99/fn)
[![License](https://img.shields.io/github/license/syke99/fn)](https://github.com/syke99/fn/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/syke99/fn)</br>
an actual implementation of a Result type in Go


Why?
=====
As we all know, Go has a reputation for its error handling being very verbose. fn provides a clean
implementation of a Result type to cut down on this verbosity . While fn requires that functions 
used with it only take one argument (it leverages Generics to make its code clean), it still keeps 
function calls clear, as well as the ability to clearly wrap errors where they happen (and prevent 
further fubctions from being called).


How?
=====

### Installing
```bash
go get -u github.com/syke99/fn
```

Then you can import the package in any go file you'd like
```go
import "github.com/syke99/fn"
```

### Basic usage

Define your types for in and out:
```go
type myFirstType struct {
	name string
	id   int
}

type mySecondType struct {
	name  string
	id    int
	email string
}

type myThirdType struct {
	name  string
	id    int
	email string
	job   string
}
```

Define your functions:
```go
func addName(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: "jane doe",
	}, nil
}

func addId(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: v.name,
		id:   1234,
	}, nil
}

func addEmail(v myFirstType) (mySecondType, error) {
	return mySecondType{
		name:  v.name,
		id:    v.id,
		email: "jane_doe@work.com",
	}, nil
}

func addJob(v mySecondType) (myThirdType, error) {
	return myThirdType{
		name:  v.name,
		id:    v.id,
		email: v.email,
		job:   "sales",
	}, nil
}
```

Define any errors:
```go
var (
	testError     = errors.New("testing error")
	addNameError  = errors.New("add name error")
	addIdError    = errors.New("add id error")
	addEmailError = errors.New("add email error")
	addJobError   = errors.New("add job error")
)
```

Then simply call fn.Try with the function you want to call, the argument you want to pass the function you want to call,
and then any error you'd like to wrap the error returned from the function you want to call with:
```go
func main() {
    start := myFirstType{}
    
    withName := fn.Try(addName, start, addNameError)
    
    withId := fn.Try(addId, withName, addIdError)
    
    withEmail := fn.Try(addEmail, withId, addEmailError)
    
    final, err := fn.Try(addJob, withEmail, addJobError).Out()
	// err check
	
	// use final
}
```

All together:
```go
package main

import (
	"errors"
	
	"github.com/syke99/fn"
)

var (
	testError     = errors.New("testing error")
	addNameError  = errors.New("add name error")
	addIdError    = errors.New("add id error")
	addEmailError = errors.New("add email error")
	addJobError   = errors.New("add job error")
)

type myFirstType struct {
	name string
	id   int
}

type mySecondType struct {
	name  string
	id    int
	email string
}

type myThirdType struct {
	name  string
	id    int
	email string
	job   string
}

func addName(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: "jane doe",
	}, nil
}

func addId(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: v.name,
		id:   1234,
	}, nil
}

func addEmail(v myFirstType) (mySecondType, error) {
	return mySecondType{
		name:  v.name,
		id:    v.id,
		email: "jane_doe@work.com",
	}, nil
}

func addJob(v mySecondType) (myThirdType, error) {
	return myThirdType{
		name:  v.name,
		id:    v.id,
		email: v.email,
		job:   "sales",
	}, nil
}

func main() {
	start := myFirstType{}

	withName := fn.Try(addName, start, addNameError)

	withId := fn.Try(addId, withName, addIdError)

	withEmail := fn.Try(addEmail, withId, addEmailError)

	final, err := fn.Try(addJob, withEmail, addJobError).Out()
	// err check

	// use final
}
```

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the MIT license, see [LICENSE](../LICENSE) for details.
