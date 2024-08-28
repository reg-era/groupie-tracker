package webserver

import (
	"strconv"
	"strings"

	"GTapi/tracker"
)

func SearchPro(key string) []int {
	res := []int{}
	value := strings.ToLower(key)
	for k, v := range tracker.Artists {
		i := 0
		if strings.Contains(strings.ToLower(v.Name), value) || strings.Contains(strings.ToLower(v.FirstAlbum), value) || strconv.Itoa(v.CreationDate) == value {
			i++
		}
		for _, j := range v.Members {
			if strings.Contains(strings.ToLower(j), value) {
				i++
			}
		}
		for _, j := range v.LocationST.Locations {
			if strings.Contains(strings.ToLower(j), value) {
				i++
			}
		}
		if i != 0 {
			res = append(res, k)
		}
	}
	return res
}

func GetOptions(data []tracker.Artist) {
	for i, c := range data {
		Options[c.Name] = i
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
