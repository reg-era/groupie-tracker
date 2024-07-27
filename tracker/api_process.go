package tracker

func APiProcess(url string) {
	Get_Api_Data(url)
	Get_Artist_Data(Api.Artists)
	urls := []string{Api.Locations, Api.Dates, Api.Relation}
	Get_Artist_MoreData(urls)
}
