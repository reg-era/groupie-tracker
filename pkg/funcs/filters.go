package funcs

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var FoundArtists []Artist

func Filter(members []string, firstmin string, firstmax string, creatmin string, creatmax string, location string) {
	var found []Artist
	for i, v := range All.Artists {
		if CreationDateVal(creatmax, creatmin, v.CreationDate) && FirstAlbumVal(firstmax, firstmin, v.FirstAlbum) && LocationVal(location, v.Rel) && MembersVal(members, len(v.Members)) {
			found = append(found, All.Artists[i])
		}
	}
	FoundArtists = found
}

func CreationDateVal(creatmin string, creatmax string, creatDate int) bool {
	return creatDate <= ToNum(creatmin) && creatDate >= ToNum(creatmax)
}

func FirstAlbumVal(firstmin string, firstmax string, firstdate string) bool {
	year := firstdate[6:]
	return ToNum(year) <= ToNum(firstmin) && ToNum(year) >= ToNum(firstmax)
}

func LocationVal(location string, rel map[string][]string) bool {
	_, ok := rel[strings.ToLower(location)]
	if ok {
		return true
	}
	if !ok {
		for j := range rel {
			j = strings.ReplaceAll(strings.ReplaceAll(j, "_", " "), "-", " ")
			if strings.Contains(strings.ToLower(j), strings.ToLower(location)) {
				return true
			}
		}
	}
	return false
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

func ToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return num
}
