package utils

import "strconv"

func ParseIDToInt(strId string) (uint, error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
