package tracker

import "sync"

// Info tack the Url and Data struct for an Api
type Info struct {
	Url  string
	Data interface{}
}

// processe the fetching from Api
func APiProcess(url string) {
	var wgp sync.WaitGroup

	artists := make(chan []interface{})

	// get the Api info in first
	wgp.Add(1)
	go Get_Api_Data(Info{url, &Api}, artists, &wgp)

	go func() {
		wgp.Wait()
		close(artists)
	}()
}
