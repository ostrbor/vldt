package vldt

import "fmt"

type Array struct {
	Name    string
	Val     []any
	details []string
}

func Arr(n string, v []any) Array {
	return Array{Name: n, Val: v}
}

func (a Array) Req() Array {
	if a.Val == nil {
		d := requiredArr(a.Name)
		a.details = append(a.details, d)
	}
	return a
}

func (a Array) Min(v int) Array {
	if a.Val != nil && len(a.Val) < v {
		a.details = append(a.details, minItems(a.Name, len(a.Val), v))
	}
	return a
}

func (a Array) Max(v int) Array {
	if a.Val != nil && len(a.Val) > v {
		a.details = append(a.details, maxItems(a.Name, len(a.Val), v))
	}
	return a
}

// todo
// array of structures?
func (a Array) Unique() Array {
	if a.Val == nil {
		return a
	}
	m := make(map[any]bool)
	for _, v := range a.Val {
		if m[v] {
			d := fmt.Sprintf("%s: array must contain unique values, received duplicate value '%v'", a.Name, v)
			a.details = append(a.details, d)
			return a
		}
		m[v] = true
	}
	return a
}

func (a Array) Errs() (errs []Problem) {
	for _, d := range a.details {
		p := Problem{Detail: d}
		errs = append(errs, p)
	}
	return errs
}
