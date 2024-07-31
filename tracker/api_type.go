package tracker

var (
	Api GTApi

	Artists   []Artist
	Locations LocationST
	Dates     DateST
	Relations RelationST
	MoreInfos MoreInfo
)

type (
	GTApi struct {
		Artists   string `json:"artists"`
		Locations string `json:"locations"`
		Dates     string `json:"dates"`
		Relation  string `json:"relation"`
	}

	Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}

	LocationST struct {
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	}

	DateST struct {
		Dates []string `json:"dates"`
	}

	RelationST struct {
		DatesLocation map[string][]string `json:"datesLocations"`
	}

	MoreInfo struct {
		Artist
		LocationST
		DateST
		RelationST
	}
)
