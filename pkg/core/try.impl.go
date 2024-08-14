package core

import (
	"errors"
	"fmt"
)

const (
	_TRY_PND = 0
	_TRY_OK  = 1
	_TRY_ERR = 2
)

type try[A any, R any] struct {
	fn     func(A) R
	arg    A
	r      R
	e      interface{}
	status uint8
}

func (t try[A, R]) call() {
	defer func() {
		if e := recover(); e != nil {
			t.e = e
			t.status = _TRY_ERR
		} else {
			t.status = _TRY_OK
		}
	}()
	t.r = t.fn(t.arg)
}

func (t try[A, R]) test() {
	if t.status == _TRY_PND {
		t.call()
	}
}

func (t try[A, R]) IsSuccess() bool {
	t.test()
	return t.status == _TRY_OK
}

func (t try[A, R]) IsFailure() bool {
	t.test()
	return t.status == _TRY_ERR
}

func (t try[A, R]) Success() Option[R] {
	t.test()
	if t.status == _TRY_ERR {
		return None[R]()
	} else {
		return Some(t.r)
	}
}

func (t try[A, R]) Failure() Option[any] {
	t.test()
	if t.status == _TRY_ERR {
		return Some(t.e)
	} else {
		return None[any]()
	}
}

func (t try[A, R]) AsResult(errorFactory ...func(any) error) Result[R] {
	t.test()
	if t.status == _TRY_ERR {
		var err error
		if len(errorFactory) != 0 {
			err = errorFactory[0](t.e)
		} else {
			err = errors.New(fmt.Sprint(t.e))
		}
		return Err[R](err)
	} else {
		return Ok(t.r)
	}
}
