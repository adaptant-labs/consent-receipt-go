package category

func CompareMatch(s1, s2 []DataCategory) []DataCategory {
	match := make([]DataCategory, 0)

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				match = append(match, s2[j])
			}
		}
	}

	return match
}

func calculateDifference(s1, s2 []DataCategory) []DataCategory {
	diff := make([]DataCategory, 0)

	for i := 0; i < len(s1); i++ {
		matched := false

		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				matched = true
				break
			}
		}

		if !matched {
			diff = append(diff, s1[i])
		}
	}

	return diff
}

func CompareDifference(s1, s2 []DataCategory) []DataCategory {
	diff := calculateDifference(s1, s2)
	diff = append(diff, calculateDifference(s2, s1)...)
	return diff
}
