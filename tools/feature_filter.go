package tools

import (
	"strconv"
)

func GetLocations(cards []Card) []string {
	res := []string{}
	for _, v := range cards {
		for _, loc := range v.Locations {
			if !Checkloca(loc, Rel) {
				res = append(res, loc)
			}
		}
	}
	return res
}

func CheckRange(firstmin string, firstmax string, date string) bool {
	if len(date) > 4 {
		date = date[len(date)-4:]
	}
	real, err := strconv.Atoi(date)
	if err != nil {
		return false
	}
	max, err := strconv.Atoi(firstmax)
	if err != nil {
		return false
	}
	min, err := strconv.Atoi(firstmin)
	if err != nil {
		return false
	}

	if real >= min && real <= max {
		return true
	} else {
		return false
	}
}

func MembersVal(nmembers []string, num int) bool {
	if len(nmembers) == 0 {
		return true
	}
	for _, v := range nmembers {
		nb, err := strconv.Atoi(v)
		if err != nil {
			return false
		} else {
			if nb == num {
				return true
			}
		}
	}
	return false
}

func Checkloca(n string, tab []string) bool {
	if n == "" {
		return true
	}
	for i := 0; i < len(tab); i++ {
		if n == tab[i] {
			return true
		}
	}
	return false
}
