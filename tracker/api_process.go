package tracker

var URLS = map[string]interface{}{
	Api.Locations: &Locations,
	Api.Dates:     &Dates,
	Api.Relation:  &Relations,
}

func APiProcess(url string) {
	Get_Api_Data(url)
	Get_Artist_Data(Api.Artists)
	URLS = map[string]interface{}{
		Api.Locations: &Locations,
		Api.Dates:     &Dates,
		Api.Relation:  &Relations,
	}
}
