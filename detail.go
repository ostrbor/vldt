package vldt

import (
	"fmt"
	"strings"
)

func requiredStr(n string) string {
	return fmt.Sprintf("%s: required (must be set or null/empty string are not allowed)", n)
}

func requiredStrPtr(n string) string {
	return fmt.Sprintf("%s: required (must be set or null is not allowed)", n)
}

func requiredInt(n string) string {
	return fmt.Sprintf("%s: required (must be set or null/0 are not allowed)", n)
}

func requiredIntPtr(n string) string {
	return fmt.Sprintf("%s: required (must be set or null is not allowed)", n)
}

func requiredBool(n string) string {
	return fmt.Sprintf("%s: required (must be set or null is not allowed)", n)
}

func minLength(n, v string, min int) string {
	return fmt.Sprintf("%s: minLength is %d, received a length of %d in the string '%s'", n, min, len(v), short(v))
}

func maxLength(n, v string, max int) string {
	return fmt.Sprintf("%s: maxLength is %d, received a length of %d in the string '%s'", n, max, len(v), short(v))
}

func ascii(n, v string) string {
	return fmt.Sprintf("%s: must be printable US-ASCII, received '%s'", n, v)
}

func pattern(n, v, p string) string {
	return fmt.Sprintf("%s: the pattern %s is not matched by the string '%s'", n, p, short(v))
}

func enum(n, v string, allowed []string) string {
	return fmt.Sprintf("%s: allowed values {%s}, received '%s'", n, strings.Join(allowed, ", "), short(v))
}

func minimum(n string, v, min int) string {
	return fmt.Sprintf("%s: minimum is %d, received %d", n, min, v)
}

func maximum(n string, v, max int) string {
	return fmt.Sprintf("%s: maximum is %d, received %d", n, max, v)
}

func minItems(n string, l, min int) string {
	return fmt.Sprintf("%s: minItems is %d, received array with length %d", n, min, l)
}

func maxItems(n string, l, max int) string {
	return fmt.Sprintf("%s: maxItems is %d, received array with length %d", n, max, l)
}
