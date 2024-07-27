package tracker

var (
	Datapls []MRinfo
	Data    []Artist
	Api     GTApi
)

type GTApi struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type MRinfo struct {
	LocationsCONC []string            `json:"locations"`
	ConcertDates  []string            `json:"concertDates"`
	RelationsART  map[string][]string `json:"relations"`
}
