package wxdata

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"path"
)

type DownloadItem struct {
	PartialPath string
	Url         string
}

func IsStatusCode200(url string) bool{

	response, err := http.Head(url)
	if err != nil {
		fmt.Println("Error while accessing", url, "-", err)
		return false
	}
	return response.StatusCode==200;
}

func Download(item DownloadItem, targetFolder string){

	filePath :=  path.Join(targetFolder,item.PartialPath)
	dirPath := path.Dir(filePath)

	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		fmt.Println("Could not create directory ", filePath)
	}

	//if file exists
	if _, err := os.Stat(filePath); err==nil {
		fmt.Println("File exists ", filePath)
		return;
	}


	//execute request
	response, err := http.Get(item.Url)
	if err != nil {
		fmt.Println("Error while downloading", item.Url, "-", err)
		return
	}
	defer response.Body.Close()

	if(response.StatusCode!=200){
		fmt.Println("Cancelling, got code", response.StatusCode, " for ", item.Url)
		return
	}

	//create file
	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", filePath, "-", err)
		return
	}
	defer output.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", item.Url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}
