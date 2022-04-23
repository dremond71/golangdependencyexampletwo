package golangdependencyexampletwo

import (
	"github.com/dremond71/golangdependencyexampleone"
)

func GetSomeString() string {
	someString := "Obiwan Kenobi"
	someStringModified := golangdependencyexampleone.DoSomething(someString)
	return someStringModified
}
