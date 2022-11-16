package main

func SmartPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
