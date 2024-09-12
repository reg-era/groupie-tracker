package webserver

import (
	"strconv"
	"strings"

	"GTapi/tracker"
)

func SearchProcess(key string) []int {
	res := []int{}
	value := strings.ToLower(key)
	for k, v := range tracker.Artists {
		if strings.HasPrefix(strings.ToLower(v.Name), value) || strings.ToLower(v.FirstAlbum) == value || strconv.Itoa(v.CreationDate) == value {
			res = append(res, k)
		}
	}
	for k, v := range tracker.Artists {
		for _, j := range v.Members {
			if strings.HasPrefix(strings.ToLower(j), value) {
				if !CheckVal(k, res) {
					res = append(res, k)
				}
			}
		}
	}
	for k, v := range tracker.Artists {
		for _, j := range v.LocationST.Locations {
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

func GetOptions(data []tracker.Artist) {
	for i, c := range data {
		Options[c.Name+" - artist/band"] = i
		Options[c.FirstAlbum+" - first album date"] = i
		Options[strconv.Itoa(c.CreationDate)+" - creation date"] = i
		for _, j := range c.Members {
			Options[j+" - members"] = i
		}
		for _, j := range c.LocationST.Locations {
			Options[j+" - locations"] = i
		}
	}
}
