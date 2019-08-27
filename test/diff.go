package test

import "github.com/gcla/deep"

// Diff in test structs
func Diff(a, b interface{}) []string {
	return deep.DefaultComparer.Equal(a, b)
}
