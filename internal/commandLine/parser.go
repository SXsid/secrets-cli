package cli

type CliFlags struct {
	Key   string
	Value string
}

func CliParser(args []string) *CliFlags {
	var flags CliFlags
	i := 2
	for i+1 < len(args) {
		switch args[i] {
		case "-k":
			i++
			flags.Key = args[i]
		case "-v":
			i++
			flags.Value = args[i]
		default:
			i++
		}
	}
	return &flags
}
