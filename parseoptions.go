package ghostls

func IsFlag(s string) bool {
	return s[0] == '-'
}

func ParseFlags(args []string) []string {
	for _, argument := range args {
		if IsFlag(argument) {
			switch argument {
			case "-a":
				DisplayHidden = true
			case "-R":
				RecursiveSearch = true
			case "-l":
				LongFormat = true
			case "-r":
				ReverseOrder = true
			default:
				continue
			}
		}
	}
	return args
}
