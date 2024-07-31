package tracker

var URLS = map[string]interface{}{
	Api.Locations: &Locations,
	Api.Dates:     &Dates,
	Api.Relation:  &Relations,
}

func APiProcess(url string) {
	Get_Api_Data(&Api, url)
	Get_Api_Data(&Artists, Api.Artists)
	
	URLS = map[string]interface{}{
		Api.Locations: &Locations,
		Api.Dates:     &Dates,
		Api.Relation:  &Relations,
	}
}
