package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type AlpineCsvData struct {
	BugId              string
	Project            string
	Tracker            string
	ParentTask         string
	Status             string
	Priority           string
	Subject            string
	Author             string
	Assignee           string
	Updated            string
	Category           string
	TargetVersion      string
	StartDate          string
	DueDate            string
	EstimatedTime      string
	TotalEstimatedTime string
	PercentageDone     string
	Created            string
	Closed             string
	RelatedIssues      string
	AffectedVersions   string
}

func csvFileParser(path string) ([]string, error) {
	results := []string{}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	results = strings.Split(string(dat), ",")
	return results, nil
}
func getDataFromAlpineIssuesUrl(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	strings := strings.Split(string(contents), ",")
	return strings, nil
}
func detectNumberOfCves(list []string) []string {
	results := []string{}
	re := regexp.MustCompile(`CVE-[0-9]{4}-[0-9]{4}`)
	for _, elem := range list {
		if strings.Contains(elem, "CVE") {
			miniSlice := re.FindStringSubmatch(elem)
			results = append(results, miniSlice...)
		}

	}
	return results
}
func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func addAlpineDataFromResponse(data []string) []AlpineCsvData {
	results := []AlpineCsvData{}
	for i := 0; i < len(data); {
		fmt.Printf("Entry: %s \n", data[i])
		alpineDataEntry := AlpineCsvData{
			BugId:              data[i],
			Project:            data[i+1],
			Tracker:            data[i+2],
			ParentTask:         data[i+3],
			Status:             data[i+4],
			Priority:           data[i+5],
			Subject:            data[i+6],
			Author:             data[i+7],
			Assignee:           data[i+8],
			Updated:            data[i+9],
			Category:           data[i+10],
			TargetVersion:      data[i+11],
			StartDate:          data[i+12],
			DueDate:            data[i+13],
			EstimatedTime:      data[i+14],
			TotalEstimatedTime: data[i+15],
			PercentageDone:     data[i+16],
			Created:            data[i+17],
			Closed:             data[i+18],
			RelatedIssues:      data[i+19],
			AffectedVersions:   data[i+20],
		}
		results = append(results, alpineDataEntry)
		i = i + 21
	}
	return results
}
func main() {
	/*fileData, err := csvFileParser("./data.csv")
	if err != nil {
		fmt.Printf(string(err.Error()))
	}*/
	urlData, err := getDataFromAlpineIssuesUrl("http://bugs.alpinelinux.org/projects/alpine/issues.csv?c%5B%5D=project&c%5B%5D=tracker&c%5B%5D=status&c%5B%5D=priority&c%5B%5D=subject&c%5B%5D=assigned_to&c%5B%5D=updated_on&f%5B%5D=&group_by=&set_filter=1&utf8=%E2%9C%93")
	if err != nil {
		fmt.Printf(string(err.Error()))
	}

	fmt.Printf("Got %d entries from alpine URL. Processing...\n", len(urlData))
	//alpineData := addAlpineDataFromResponse(fileData)
	/*for _, entry := range alpineData {
		fmt.Printf("Entry %v\n", entry)
	}*/
	//fmt.Printf("Got %d processed alpine data objects. Processing...\n", len(alpineData))
	cves := detectNumberOfCves(urlData)
	fmt.Printf("Found %d CVEs in entries\n", len(cves))
	RemoveDuplicates(&cves)
	fmt.Printf("Found %d CVEs after removing duplicate entries\n", len(cves))
	for _, cve := range cves {
		fmt.Printf("Found: %s \n", cve)
	}

}
