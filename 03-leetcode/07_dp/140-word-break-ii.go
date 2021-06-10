package dp

func wordBreak(s string, wordDict []string) []string {
	words := map[string]struct{}{}
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	data := map[int][]string{0: {""}}
	for {
		if len(data) == 0 {
			return nil
		}
		_, ok := data[len(s)]
		if len(data) == 1 && ok {
			return data[len(s)]
		}
		nData := map[int][]string{}
		for begin, ss := range data {
			if begin == len(s) {
				nData[len(s)] = append(nData[len(s)], ss...)
				continue
			}
			for i := begin; i <= len(s); i++ {
				word := s[begin:i]
				if _, ok := words[word]; !ok {
					continue
				}
				for _, prev := range ss {
					if prev != "" {
						prev = prev + " "
					}
					nData[i] = append(nData[i], prev+word)
				}
			}
		}
		data = nData
	}
}
