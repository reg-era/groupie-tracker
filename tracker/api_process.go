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

// processe the fetching from Api
func APiProcess(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	var wgp sync.WaitGroup

	wgp.Add(1)
	go func() {
		defer wgp.Done()
		Get_Api_Data(Info{url, &Api})
	}()
	wgp.Wait()

	wgp.Add(1)
	go func() {
		defer wgp.Done()
		Get_Api_Data(Info{Api.Artists, &Artists})
	}()
	wgp.Wait()

	wgp.Add(1)
	if len(Artists) != 0 {
		var nestedWgp sync.WaitGroup
		nestedWgp.Add(len(Artists) * 3)

		for i := range Artists {
			artist := &Artists[i]

			go func(artist *Artist) {
				defer nestedWgp.Done()
				Get_Api_Data(Info{Api.Dates + "/" + strconv.Itoa(artist.Id), &artist.DateST})
			}(artist)

			go func(artist *Artist) {
				defer nestedWgp.Done()
				Get_Api_Data(Info{Api.Locations + "/" + strconv.Itoa(artist.Id), &artist.LocationST})
			}(artist)

			go func(artist *Artist) {
				defer nestedWgp.Done()
				Get_Api_Data(Info{Api.Relation + "/" + strconv.Itoa(artist.Id), &artist.RelationST})
			}(artist)
		}

		nestedWgp.Wait()
		wgp.Done()
	}
	wgp.Wait()
}
