package db

import "github.com/volatiletech/sqlboiler/v4/types"

func intArrToInt64Arr(intArr []int) types.Int64Array {
	res := make(types.Int64Array, 0, len(intArr))

	for _, arrEl := range intArr {
		res = append(res, int64(arrEl))
	}

	return res
}

func int64ArrToIntArr(int64Arr types.Int64Array) []int {
	res := make([]int, 0, len(int64Arr))

	for _, arrEl := range int64Arr {
		res = append(res, int(arrEl))
	}

	return res
}
