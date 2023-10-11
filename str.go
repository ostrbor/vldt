package vldt

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

type String struct {
	Name    string
	Value   *string // to prevent validation check in each validation func, serves as a flag that there is no need to validate value
	value   string  // derived from Value, otherwise in each validation call need to dereference
	details []string
}

// Str is used for required JSON string fields that do not allow empty string.
// Go string zero value signals that one of the following happened:
// - field was not set in JSON,
// - null was sent in JSON,
// - empty string was sent in JSON.
func Str(n, v string) String {
	if v == "" {
		d := n + ": required (must be set or null/empty string is not allowed)"
		return String{Name: n, Value: nil, details: []string{d}}
	}
	return String{Name: n, Value: &v, value: v}
}

// StrPtr is used for required JSON string fields that allow empty string
// or for optional fields.
func StrPtr(n string, v *string, required bool) String {
	if required && v == nil {
		d := n + ": required (must be set or null is not allowed)"
		return String{Name: n, Value: nil, details: []string{d}}
	}
	return String{Name: n, Value: v, value: *v}
}

func (s String) Min(length int) String {
	if s.Value == nil {
		return s
	}
	if len(s.value) < length {
		d := fmt.Sprintf("%s: minLength is %d, received length of %d in '%s'", s.Name, length, len(s.value), short(s.value))
		s.details = append(s.details, d)
	}
	return s
}

func (s String) Max(length int) String {
	if s.Value == nil {
		return s
	}
	if len(s.value) > length {
		d := fmt.Sprintf("%s: maxLength is %d, received length of %d in '%s'", s.Name, length, len(s.value), short(s.value))
		s.details = append(s.details, d)
	}
	return s
}

func (s String) Pattern(reg *regexp.Regexp) String {
	if s.Value == nil {
		return s
	}
	if reg.MatchString(s.value) {
		return s
	}
	d := fmt.Sprintf("%s: pattern %s is not matched by '%s'", s.Name, reg.String(), short(s.value))
	s.details = append(s.details, d)
	return s
}

func (s String) Enum(allowed []string) String {
	if s.Value == nil {
		return s
	}
	for _, a := range allowed {
		if s.value == a {
			return s
		}
	}
	d := fmt.Sprintf("%s: received '%s', allowed values {%s}", s.Name, short(s.value), short(strings.Join(allowed, ", ")))
	s.details = append(s.details, d)
	return s
}

func (s String) ASCII() String {
	if s.Value == nil {
		return s
	}
	if printableASCII(s.value) {
		return s
	}
	d := fmt.Sprintf("%s: must be printable US-ASCII, received '%s'", s.Name, short(s.value))
	s.details = append(s.details, d)
	return s
}

func (s String) Email() String {
	if s.Value == nil {
		return s
	}
	if _, err := mail.ParseAddress(s.value); err != nil {
		d := fmt.Sprintf("%s: must be valid email address, parsed '%s' with error '%s'", s.Name, short(s.value), err)
		s.details = append(s.details, d)
	}
	return s
}

func (s String) DateTime() String {
	if s.Value == nil {
		return s
	}
	if _, err := time.Parse(time.RFC3339, s.value); err != nil {
		d := fmt.Sprintf("%s: must be valid date-time in RFC3339 format, parsed '%s' with error '%s'", s.Name, short(s.value), err)
		s.details = append(s.details, d)
	}
	return s
}

func (s String) Errs() (errs []Problem) {
	for _, d := range s.details {
		p := Problem{Detail: d}
		if IncludeTitle {
			p.Title = InvalidValue
		}
		errs = append(errs, p)
	}
	return errs
}
