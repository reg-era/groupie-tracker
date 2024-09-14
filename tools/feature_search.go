package tools

import (
	"strconv"
	"strings"
)

var Options = make(map[string]int)

func SearchProcess(key string) []int {
	res := []int{}
	value := strings.ToLower(key)
	for k, v := range Data.Cards {
		if strings.HasPrefix(strings.ToLower(v.Name), value) || strings.ToLower(v.FirstAlbum) == value || strconv.Itoa(v.CreationDate) == value {
			res = append(res, k)
		}
		for _, j := range v.Members {
			if strings.HasPrefix(strings.ToLower(j), value) {
				if !CheckVal(k, res) {
					res = append(res, k)
				}
			}
		}
		for _, j := range v.Locations {
			if strings.Contains(strings.ToLower(j), value) {
				if !CheckVal(k, res) {
					res = append(res, k)
				}
			}
		}
	}
	return res
}

func CheckVal(n int, tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if n == tab[i] {
			return true
		}
	}
	return false
}

func GetOptions(data PageData) {
	for i, c := range data.Cards {
		Options[c.Name] = i
		Options[c.FirstAlbum+" - first album date"] = i
		Options[strconv.Itoa(c.CreationDate)+" - creation date"] = i
		for _, j := range c.Members {
			Options[j+" - members"] = i
		}
		for _, j := range c.Locations {
			Options[j+" - locations"] = i
		}
	}
}


