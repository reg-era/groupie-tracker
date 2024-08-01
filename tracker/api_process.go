package tracker

var URLS = map[string]interface{}{
	Api.Locations: &Locations,
	Api.Dates:     &Dates,
	Api.Relation:  &Relations,
}

type Info struct {
	Url  string
	Data interface{}
}

func APiProcess(url string) {
	Get_Api_Data(Info{url, &Api})

	Get_Api_Data(Info{Api.Artists, &Artists})

	URLS = map[string]interface{}{
		Api.Locations: &Locations,
		Api.Dates:     &Dates,
		Api.Relation:  &Relations,
	}
}
