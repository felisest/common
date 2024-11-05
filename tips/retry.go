package main

import (
	"errors"
	"fmt"
	"time"
)

func DoRetryOne[T any](retry int, duration time.Duration, fn func(T) error, fnArgs ...any) error {
	for retryCount := 0; retryCount < retry; retryCount++ {

		time.Sleep(duration * time.Duration(retryCount))

		err := fn(fnArgs[0].(T))

		if err == nil {
			return nil
		}
	}

	return nil
}

func DoRetryTwo[T0 any, T1 any](retry int, duration time.Duration, fn func(T0, T1) error, fnArgs ...any) error {

	for retryCount := 0; retryCount < retry; retryCount++ {
		time.Sleep(duration * time.Duration(retryCount))

		err := fn(
			fnArgs[0].(T0),
			fnArgs[1].(T1),
		)

		if err == nil {
			return nil
		}
	}

	return nil
}

func testExecute(s string) error {
	fmt.Println(s)

	return errors.New("test error")
}

func main() {
	fmt.Println("Hello")

	DoRetryOne(3, time.Millisecond*1000, testExecute, "test Do")
}
