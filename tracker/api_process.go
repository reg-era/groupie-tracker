package tracker

import "fmt"

func APiProcess(url string) {
	if err := Get_Api_Data(&Api, url); err != nil{
		fmt.Printf("failed to get API data: %v", err)
	}
	if err := Get_Api_Data(&Artists, Api.Artists); err != nil{
		fmt.Printf("failed to get API data: %v", err)
	}
	
	URLS = map[string]interface{}{
		Api.Locations: &Locations,
		Api.Dates:     &Dates,
		Api.Relation:  &Relations,
	}
}
