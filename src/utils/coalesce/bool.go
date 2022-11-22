package coalesce

func Bool(args ...*bool) *bool {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
