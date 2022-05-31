package pack

import "strconv"

func StrSlice2IntSlice(strSlice []string) (res []int64, err error) {
	for _, v := range strSlice {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, int64(val))
	}
	return res, nil
}
