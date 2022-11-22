package coalesce

func String(args ...*string) *string {
	for _, arg := range args {
		if arg != nil {
			return arg
		}
	}
	return nil
}
