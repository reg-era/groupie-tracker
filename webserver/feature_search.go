package webserver

import (
	"strconv"

	"GTapi/tracker"
)

func SearchPro(value string) []int {
	return []int{}
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

// add func FilterHandele(&tracker.Artists) Artists
