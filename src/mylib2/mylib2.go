package mylib2

import "mylib"

func Bar() string {
	return mylib.Foo()
}
