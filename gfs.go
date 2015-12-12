package wxdata

import (
	"time"
	"strings"
	"github.com/bjartn/go-wx/easy"
)

func GetGfsCandidateAnatimes() []time.Time {
	date := easy.UtcDate()
	startHour := easy.UtcNow().Hour() - easy.UtcNow().Hour() % 6;
	return []time.Time{
		date.Add(time.Hour * time.Duration(startHour)),
		date.Add(time.Hour * time.Duration(startHour) - time.Hour * 6),
	}
}

func GetFirstTimestep(anaTime time.Time) string{
	url := "http://nomads.ncep.noaa.gov/pub/data/nccf/com/gfs/prod/gfs.{anaTime}/gfs.t{anaHour}z.pgrb2.1p00.f000"
	url = strings.Replace(url,"{anaHour}",anaTime.Format(easy.Date_HH),1)
	url = strings.Replace(url,"{anaTime}", anaTime.Format(easy.Date_yyyyMMddHH),1)
	return url;
}

func GetGfsDownloadItems(anaTime time.Time) []DownloadItem{

	urlTemplate := "http://nomads.ncep.noaa.gov/pub/data/nccf/com/gfs/prod/"
	fcTimes := []string { "000", "003", "006", "009", "012", "015", "018", "021", "024", "027", "030", "036", "039", "042", "048", "051", "054", "057", "060", "063", "066", "069", "072" }
	l:=make([]DownloadItem,len(fcTimes))


	for i,fcTime := range fcTimes {
		uniqueSubPath := "gfs.{anaTime}/gfs.t{anaHour}z.pgrb2.1p00.f{fcHour}"
		uniqueSubPath = strings.Replace(uniqueSubPath,"{anaHour}",anaTime.Format(easy.Date_HH),1)
		uniqueSubPath =strings.Replace(uniqueSubPath,"{fcHour}",fcTime,1)
		uniqueSubPath = strings.Replace(uniqueSubPath,"{anaTime}", anaTime.Format(easy.Date_yyyyMMddHH),1)

		url := urlTemplate + uniqueSubPath

		l[i] = DownloadItem{uniqueSubPath, url}
	}

	return l
}