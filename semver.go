package random

import "fmt"

// SemVer returns random semver value
func SemVer() string {
	switch Element([]string{"x.x", "x.x.x", "x.x.x.x", "x.x.x-x.x"}) {
	case "x.x":
		return fmt.Sprintf("%v.%v", Value([]int{0, 20}), Value([]int{0, 100}))
	case "x.x.x":
		return fmt.Sprintf("%v.%v.%v", Value([]int{0, 20}), Value([]int{0, 50}), Value([]int{0, 200}))
	case "x.x.x.x":
		return fmt.Sprintf("%v.%v.%v.%v", Value([]int{0, 50}), Value([]int{0, 100}), Value([]int{0, 200}), Value([]int{0, 100000}))
	case "x.x.x-x.x":
		return fmt.Sprintf("%v.%v.%v-%v.%v", Value([]int{0, 20}), Value([]int{0, 50}), Value([]int{0, 300}), Element([]string{"alpha", "beta", "gamma"}), Value([]int{0, 1000}))
	}

	return ""
}
