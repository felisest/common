package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Произвольное число аргументов на входе, ошибка на выходе
func DoRetryReflect(retry int, duration int, fn interface{}, args ...interface{}) error {
	fnValue := reflect.ValueOf(fn)
	argsValue := make([]reflect.Value, len(args))
	for i, arg := range args {
		argsValue[i] = reflect.ValueOf(arg)
	}

	var err error
	for retryCount := 0; retryCount < retry; retryCount++ {

		time.Sleep(time.Millisecond * time.Duration(duration*retryCount))

		result := fnValue.Call(argsValue)

		err = (result[0].Interface()).(error)

		if err == nil {
			return nil
		}
	}

	return err
}

// Один аргумент на входе, ошибка на выходе
func DoRetryOne[T any](retry int, duration int, fn func(T) error, fnArgs ...any) error {

	var err error
	for retryCount := 0; retryCount < retry; retryCount++ {

		time.Sleep(time.Millisecond * time.Duration(duration*retryCount))

		err = fn(fnArgs[0].(T))

		if err == nil {
			return nil
		}
	}

	return err
}

// Два аргумента на входе, ошибка на выходе
func DoRetryTwo[T0 any, T1 any](retry int, duration int, fn func(T0, T1) error, fnArgs ...any) error {

	var err error
	for retryCount := 0; retryCount < retry; retryCount++ {
		time.Sleep(time.Millisecond * time.Duration(duration*retryCount))

		err = fn(
			fnArgs[0].(T0),
			fnArgs[1].(T1),
		)

		if err == nil {
			return nil
		}
	}

	return err
}

func testExecute(s string) error {
	fmt.Println(s)

	return errors.New("test error")
}

func main() {

	fmt.Println("WO reflection")
	err := DoRetryOne(3, 3000, testExecute, "test Do [T]")

	fmt.Println("With reflection")
	err = DoRetryReflect(3, 3000, testExecute, "test Do reflect") // Если передадим аргументы неверного типа - упадем во время выполнения !!!

	_ = err
}
