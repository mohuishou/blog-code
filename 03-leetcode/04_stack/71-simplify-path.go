package stack

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var res []string
	for _, p := range paths {
		switch p {
		case "..":
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		case "", ".":
		default:
			res = append(res, p)
		}
	}
	return "/" + strings.Join(res, "/")
}
