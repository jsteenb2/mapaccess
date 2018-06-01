package mapaccess

// ok, not a real smooth API to work with as you have to deal
// with the inbound check
func AccessWithEmptyStructs(f []int) []int {
	validFeatures := map[int]struct{}{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	newF := make([]int, 0, len(validFeatures))
	for _, feature := range f {
		if _, ok := validFeatures[feature]; !ok {
			continue
		}
		newF = append(newF, feature)
	}
	return newF
}

// much better api to work with as you can check the zero value
// without the inbound check
func AccessWithBool(f []int) []int {
	validFeatures := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}

	newF := make([]int, 0, len(validFeatures))
	for _, feature := range f {
		if !validFeatures[feature] {
			continue
		}
		newF = append(newF, feature)
	}
	return newF
}
