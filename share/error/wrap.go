package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrDBConnTimeout = errors.New("timeout reaching db")

var DBErrorConnTimeout = &DBError{"connection timeout"}

func main() {
	fmt.Printf("1. f1() == ErrDBConnTimeout: %v \n", f1() == ErrDBConnTimeout)

	err := f2()
	if e, ok := err.(*DBError); ok {
		fmt.Printf("2. f2: %v \n", e)
	}

	f3Err := f3(err)
	if e, ok := f3Err.(*DBError); ok {
		fmt.Printf("3. f3: %v \n", e)
	}

	f4Err := f4(err)
	if e, ok := f4Err.(*HTTPError); ok && e.Err == DBErrorConnTimeout {
		fmt.Printf("4. f4: %v \n", e.Err)
	}

	/* go 1.13 */

	if errors.Is(f1(), ErrDBConnTimeout) {
		fmt.Printf("5. f1() is ErrDBConnTimeout\n")
	}

	var anDBError *DBError
	if errors.As(err, &anDBError) {
		fmt.Printf("6. anDBError: %v \n", anDBError.Description)
	}

	updateDBErr := updateDB(DBErrorConnTimeout)
	if errors.Is(updateDBErr, DBErrorConnTimeout) {
		fmt.Printf("7. [%v] is [%v] \n", updateDBErr, DBErrorConnTimeout)
	}

	wrappedErr := wrapErr(wrapErr(DBErrorConnTimeout))
	if errors.Is(wrappedErr, DBErrorConnTimeout) {
		fmt.Printf("8. [%v] is [%v]\n", wrappedErr, DBErrorConnTimeout)
	}

	if errors.As(wrapErr(err), &anDBError) {
		fmt.Printf("9. anDBError: %v \n", anDBError.Description)
	}

	fmt.Printf("10.wrapped: [%v], unwrapped: [%v] \n", wrapErr(err), errors.Unwrap(wrapErr(err)))
}

func f1() error {
	return ErrDBConnTimeout
}

type DBError struct {
	Description string
}

func (e *DBError) Error() string {
	return "db error: " + e.Description
}

func f2() error {
	return DBErrorConnTimeout
}

func f3(err error) error {
	return fmt.Errorf("f3: %v \n", err)
}

type HTTPError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *HTTPError) Error() string {
	return toJson(e)
	//return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

func f4(err error) error {
	return &HTTPError{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Err:        err,
	}
}

func updateDB(err error) error {
	return fmt.Errorf("err on updating data: %w", err)
}

func wrapErr(err error) error {
	return fmt.Errorf("wrapping error: %w", err)
}

func toJson(item interface{}) string {
	bytes, _ := json.Marshal(item)
	return string(bytes)
}
