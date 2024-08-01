package tracker

// Info tack the Url and Data struct for an Api
type Info struct {
	Url  string
	Data interface{}
}

// initialize the map that contant the Url and Data strict
var URLS = map[string]interface{}{
	Api.Locations: &Locations,
	Api.Dates:     &Dates,
	Api.Relation:  &Relations,
}

// processe the fetching from Api
func APiProcess(url string) {
	// get the Api info in first
	Get_Api_Data(Info{url, &Api})
	// fetch the necessary info for Artist
	Get_Api_Data(Info{Api.Artists, &Artists})
	// make the map and store supplement info for moreinfo function
	URLS = map[string]interface{}{
		Api.Locations: &Locations,
		Api.Dates:     &Dates,
		Api.Relation:  &Relations,
	}
}
