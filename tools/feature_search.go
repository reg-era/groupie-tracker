package tools

import (
	"strconv"
	"strings"
)

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

func GetOptions(data PageData) map[string]int {
	options := make(map[string]int)
	for i, c := range data.Cards {
		options[c.Name] = i
		options[c.FirstAlbum+" - first album date"] = i
		options[strconv.Itoa(c.CreationDate)+" - creation date"] = i
		for _, j := range c.Members {
			options[j+" - members"] = i
		}
		for _, j := range c.Locations {
			options[j+" - locations"] = i
		}
	}
	return options
}
