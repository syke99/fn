package fn

import (
	"fmt"
	"reflect"
	"runtime"
)

type Result[Any any] struct {
	ok  bool
	val *Any
	err error
}

func (r *Result[Any]) Out() (*Any, error) {
	return r.val, r.err
}

func Try[In any, Out any](f func(in In) (Out, error), in any, wrapper error) *Result[Out] {
	fnName := funcName(f)

	if i, ok := in.(*Result[In]); ok {
		if !i.ok {
			if i.err != nil {
				return &Result[Out]{
					ok:  false,
					err: i.err,
				}
			}
		} else {
			if i.err != nil {

				return &Result[Out]{
					ok:  false,
					err: fmt.Errorf("%w: %w", argWithValueAndErrError(fnName), i.err),
				}
			}

			out, err := f(*i.val)
			if err != nil && wrapper != nil {
				err = fmt.Errorf("%w: %w", wrapper, err)
			}

			return newRes(&out, err)
		}
	} else if i, ok := in.(In); ok {
		out, err := f(i)
		if err != nil && wrapper != nil {
			err = fmt.Errorf("%w: %w", wrapper, err)
		}

		return newRes(&out, err)
	}

	return &Result[Out]{
		ok:  false,
		err: fmt.Errorf("%w: %w", wrapper, fmt.Errorf("failed to cast in as *Res[%[1]T] or %[1]T", in)),
	}
}

func newRes[Out any](out *Out, err error) *Result[Out] {
	if err != nil {
		return &Result[Out]{
			ok:  false,
			err: err,
		}
	}
	return &Result[Out]{
		ok:  true,
		val: out,
	}
}

func funcName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func argWithValueAndErrError(name string) error {
	return fmt.Errorf("okay Res with error found at:%s", name)
}
