package tools

import (
	"fmt"
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

func CheckRange(firstmin string, firstmax string, firstdate string) bool {
	real, err := strconv.Atoi(firstdate)
	if err != nil {
		return false
	}
	max, err := strconv.Atoi(firstdate)
	if err != nil {
		return false
	}
	min, err := strconv.Atoi(firstdate)
	fmt.Println(max, min, real)
	if err != nil {
		return false
	} else {
		if real >= min && real <= max {
			return true
		} else {
			return false
		}
	}
}

func MembersVal(nmembers []string, num int) bool {
	if nmembers == nil {
		return true
	}
	for _, v := range nmembers {
		if strconv.Itoa(num) == v {
			return true
		}
	}
	return false
}

func Checkloca(n string, tab []string) bool {
	for i := 0; i < len(tab); i++ {
		if n == tab[i] {
			return true
		}
	}
	return false
}
