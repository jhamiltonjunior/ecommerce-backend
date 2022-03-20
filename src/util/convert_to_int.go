package util

import "strconv"

func ConvertToInt(params map[string]string) (int64, error) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return int64(id), err
	}
	return int64(id), nil
}
