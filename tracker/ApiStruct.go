package box


type DataValu struct {
	Image string
	Name  string
}

type Api_attits_locatons struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Api_artists_conertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Api_artists_relations struct {
	Id             int
	DatesLocations map[string][]string
}

type Api_artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
