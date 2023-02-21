package ncode

import "strconv"

type NationalID string

func (id NationalID) IsNationalIDvalid(nID NationalID) bool {
	if len(nID) == 10 {
		_, err := strconv.Atoi(string(nID))
		return err == nil

		// return true
	}
	return false
}
