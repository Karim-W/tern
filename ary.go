package tern

func Ary[T any](
	condition bool,
	yes T,
	no T,
) T {
	if condition {
		return yes
	}
	return no
}
