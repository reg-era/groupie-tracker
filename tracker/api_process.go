package tracker

import (
	"strconv"
	"sync"
)

// Info tack the Url and Data struct for an Api
type Info struct {
	Url  string
	Data interface{}
}

// if all data fetched return true
var Stat bool = false

// processe the fetching from Api
func APiProcess(url string) {
	var wgp sync.WaitGroup

	// getting basic infos
	Get_Api_Data(Info{url, &Api})
	Get_Api_Data(Info{Api.Artists, &Artists})

	wgp.Add(1)
	go func(status *bool) {
		defer wgp.Done()
		wgp.Add(len(Artists) * 3)
		// for each artist fetching date location reletion
		for i := range Artists {
			artist := &Artists[i]
			go func(artist *Artist) {
				defer wgp.Done()
				Get_Api_Data(Info{Api.Dates + "/" + strconv.Itoa(artist.Id), &artist.DateST})
			}(artist)

			go func(artist *Artist) {
				defer wgp.Done()
				Get_Api_Data(Info{Api.Locations + "/" + strconv.Itoa(artist.Id), &artist.LocationST})
			}(artist)

			go func(artist *Artist) {
				defer wgp.Done()
				Get_Api_Data(Info{Api.Relation + "/" + strconv.Itoa(artist.Id), &artist.RelationST})
			}(artist)
		}
		*status = true
	}(&Stat)
	wgp.Wait()
}
