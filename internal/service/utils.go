package service

func First[T, U any](val T, _ U) T {
	return val
}
