package vldt

import (
	"fmt"
	"net/mail"
	"regexp"
	"slices"
	"time"
)

type String struct {
	Name    string
	Val     string
	details []string
}

func Str(n, v string) String {
	return String{Name: n, Val: v}
}

func (s String) Req() String {
	if s.Val == "" {
		d := requiredStr(s.Name)
		s.details = append(s.details, d)
	}
	return s
}

func (s String) Min(length int) String {
	if s.Val != "" && len(s.Val) < length {
		s.details = append(s.details, minLength(s.Name, s.Val, length))
	}
	return s
}

func (s String) Max(length int) String {
	if s.Val != "" && len(s.Val) > length {
		s.details = append(s.details, maxLength(s.Name, s.Val, length))
	}
	return s
}

func (s String) Pattern(reg *regexp.Regexp) String {
	if s.Val != "" && !reg.MatchString(s.Val) {
		s.details = append(s.details, pattern(s.Name, s.Val, reg.String()))
	}
	return s
}

func (s String) Enum(allowed []string) String {
	if s.Val != "" && !slices.Contains(allowed, s.Val) {
		s.details = append(s.details, enum(s.Name, s.Val, allowed))
	}
	return s
}

func (s String) ASCII() String {
	if s.Val != "" && !printableASCII(s.Val) {
		s.details = append(s.details, ascii(s.Name, s.Val))
	}
	return s
}

func (s String) Email() String {
	if s.Val == "" {
		return s
	}
	if _, err := mail.ParseAddress(s.Val); err != nil {
		d := fmt.Sprintf("%s: must be valid email, parsed '%s' with error '%s'", s.Name, short(s.Val), err)
		s.details = append(s.details, d)
	}
	return s
}

func (s String) DateTime() String {
	if s.Val == "" {
		return s
	}
	if _, err := time.Parse(time.RFC3339, s.Val); err != nil {
		d := fmt.Sprintf("%s: must be valid RFC3339 (example '%s'), parsed '%s' with error '%s'",
			s.Name, time.RFC3339, short(s.Val), err)
		s.details = append(s.details, d)
	}
	return s
}

func (s String) Errs() (errs []Problem) {
	for _, d := range s.details {
		p := Problem{Detail: d}
		errs = append(errs, p)
	}
	return errs
}
