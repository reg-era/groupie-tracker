package webserver

import (
	"strconv"
	"strings"

	"GTapi/tracker"
)

func FilterProcess(key string) []int {
	res := []int{}
	value := strings.ToLower(key)
	type (
		newArtist struct {
			Id           int      `json:"id"`
			Members      []string `json:"members"`
			CreationDate int      `json:"creationDate"`
			FirstAlbum   string   `json:"firstAlbum"`
			// fetching location
			LocationST struct {
				Locations []string `json:"locations"`
			}
	}
	)

	for k, v := range tracker.Artists {
		if strings.Contains(v.Name, value) || strings.Contains(v.FirstAlbum, value) || strconv.Itoa(v.CreationDate) == value {
			res = append(res, k)
		}
	}
	for k, v := range tracker.Artists {
		for _, j := range v.Members {
			if strings.Contains(strings.ToLower(j), value) {
				if !CheckV(k, res) {
					res = append(res, k)
				}
			}
		}
	}
	for k, v := range tracker.Artists {
		for _, j := range v.LocationST.Locations {
			if strings.Contains(strings.ToLower(j), value) {
				if !CheckV(k, res) {
					res = append(res, k)
				}
			}
		}
	}
	return res
}

func CheckV(n int, tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if n == tab[i] {
			return true
		}
	}
	return false
}

// func GetOptions(data []tracker.Artist) {
// 	for i, c := range data {
// 		Options[c.Name] = i
// 		Options[c.FirstAlbum+" - first album date"] = i
// 		Options[strconv.Itoa(c.CreationDate)+" - creation date"] = i
// 		for _, j := range c.Members {
// 			Options[j+" - members"] = i
// 		}
// 		for _, j := range c.LocationST.Locations {
// 			Options[j+" - locations"] = i
// 		}
// 	}
// }
