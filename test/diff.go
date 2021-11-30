package test

import "github.com/google/go-cmp/cmp"

// Diff in test structs
func Diff(a, b interface{}) string {
	return cmp.Diff(a, b)
}
