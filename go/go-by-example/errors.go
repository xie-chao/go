package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("can't work with 2")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 2 {
		return -1, &argError{arg, "can't work with 2"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 2} {
		//在 if行内的错误检查代码，在 Go 中是一个普遍的用法
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 2} {
		//在 if行内的错误检查代码，在 Go 中是一个普遍的用法
		if r, e := f2(i); e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(2)
	fmt.Println(e.Error())
	//进行类型检查
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.Error())
	}
}
