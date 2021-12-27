package errors

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

var ExErr = errors.New("connect err")

type ExampleError struct {
	s int
}

func (receiver *ExampleError) Error() string {
	spew.Dump() // like print but better
	return strconv.Itoa(receiver.s)
}

func createTypedErrors() error {
	return &ExampleError{s: 1}
}

func createNewErrors() error {
	return errors.New("ошибка")
}

func CreateManyArgWithErrors() (int, error) {
	return 1, errors.New("ошибка")
}

func checkError() error {

	err := createNewErrors()
	if err != nil {
		return errors.Wrapf(err, "default")
	}

	// err alwas last arg
	arg, argErr := CreateManyArgWithErrors()
	if argErr != nil {
		return errors.Wrapf(argErr, "default")
	}
	log.Println(arg)

	// with ignore args
	_, argErr1 := CreateManyArgWithErrors()
	if argErr1 != nil {
		return errors.Wrapf(argErr1, "default")
	}

	newErr := createNewErrors()
	switch newErr {
	case nil:
	case ExErr:
	case errors.Wrap(newErr, "создана"):
	default:
		if err != nil {
			return errors.Wrapf(newErr, "создана")
		}
	}

	typedErr := createTypedErrors()
	switch typedErr.(type) {
	case nil:
	case *ExampleError:
		return errors.Wrapf(typedErr, "создана")
	default:
		return errors.Wrapf(typedErr, "default")
	}

	// errors.Is
	log.Printf("is = %s", errors.Is(ExErr, errors.Wrap(ExErr, "wrapped")))

	// errors.As
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	return nil
}
