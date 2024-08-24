package tracker

// declar all DataStruct variable
var (
	Api GTApi

	Artists []Artist
)

// make DataStruct type to use in fetching
type (
	// for fetching the Great Api
	GTApi struct {
		Artists   string `json:"artists"`
		Locations string `json:"locations"`
		Dates     string `json:"dates"`
		Relation  string `json:"relation"`
	}
	// for fetching the basic info from the artist
	Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		// fetching location
		LocationST struct {
			Locations []string `json:"locations"`
			Dates     string   `json:"dates"`
		}
		// fetching date
		DateST struct {
			Dates []string `json:"dates"`
		}
		// fetching the relation
		RelationST struct {
			DatesLocation map[string][]string `json:"datesLocations"`
		}
	}
)
