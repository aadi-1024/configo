package argsparser

type ArgsParser struct {
	Prefix rune
}

func (a *ArgsParser) GenerateConfig(argv []string) (map[string][]string, error) {
	n := len(argv)
	m := make(map[string][]string)
	flaginds := make([]int, 0)

	for i := 0; i < n; i++ {
		if argv[i][0] == byte(a.Prefix) {
			flaginds = append(flaginds, i)
		}
	}

	for _, i := range flaginds {
		x := make([]string, 0)
		for j := i + 1;j<n;j++ {
			if (argv[j][0] == byte(a.Prefix)) {
				break
			}
			x = append(x, argv[j])
		}
		m[argv[i]] = x
	}

	return m, nil
}