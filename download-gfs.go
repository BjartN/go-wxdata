package wxdata

import (
	"fmt"
	"sync"
	"time"
)

func DownloadGfs(targetFolder string) {

	for _,t:= range GetGfsCandidateAnatimes(){
		downloadGfs(targetFolder, t)
	}

}

func downloadGfs(targetFolder string, date time.Time) {

	items := GetGfsDownloadItems(date)

	var wg sync.WaitGroup
	wg.Add(len(items))

	for _,item := range items {

		go (func(item DownloadItem){
			defer wg.Done()
			Download(item, targetFolder)
		})(item)

		fmt.Println(item.Url)
	}

	wg.Wait()
}
