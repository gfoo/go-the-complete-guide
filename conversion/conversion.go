package conversion

import "strconv"

func StringsToFloats(strings []string) ([]float64, error) {
	var err error
	res := make([]float64, len(strings))
	for i, line := range strings {
		res[i], err = strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
