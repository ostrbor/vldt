package vldt

type Integer struct {
	Name    string
	Val     int
	details []string
}

func Int(n string, v int) Integer {
	return Integer{Name: n, Val: v}
}

func (i Integer) Req() Integer {
	if i.Val == 0 {
		d := requiredInt(i.Name)
		i.details = append(i.details, d)
	}
	return i
}

func (i Integer) Min(v int) Integer {
	if i.Val != 0 && i.Val < v {
		i.details = append(i.details, minimum(i.Name, i.Val, v))
	}
	return i
}

func (i Integer) Max(v int) Integer {
	if i.Val != 0 && i.Val > v {
		i.details = append(i.details, maximum(i.Name, i.Val, v))
	}
	return i
}

func (i Integer) Errs() (errs []Problem) {
	for _, d := range i.details {
		p := Problem{Detail: d}
		errs = append(errs, p)
	}
	return errs
}
