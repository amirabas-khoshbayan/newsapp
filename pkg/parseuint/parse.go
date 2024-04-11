package parseuint

import "strconv"

func ConvertStrToUint64(id string) (uint64, error) {
	parseUintID, err := strconv.ParseUint(id, 10, 64)

	return parseUintID, err
}
