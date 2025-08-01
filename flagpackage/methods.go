package flagpackage

import "strings"

func (u *URLList) String() string {
	return strings.Join(*u, ",")
}

func (u *URLList) Set(value string) error {
	*u = append(*u, strings.Split(value, ",")...)
	return nil
}
