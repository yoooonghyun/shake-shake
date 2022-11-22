package coalesce

func Int(args ...*int) *int {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
