package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kr/pretty"
)

func main() {
	f, err := excelize.OpenFile("Scholarship 2020 _RawData.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	uniFileTypeCounts := make(map[string]int)
	// // Get all the rows in the Sheet1.
	rows, err := f.GetRows("High School")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		name := row[3] + "_" + row[5]
		fmt.Println(name)
		os.MkdirAll("HighSchoolCounts/"+name, 0755)
		for j, col := range row {
			if j < 56 || j > 63 {
				continue
			}
			// r := strings.Split(col, "/")
			// fileName := r[len(r)-1]
			// fmt.Print(fileName, "\t")

			s := strings.Split(col, ".")
			fileType := s[len(s)-1]
			uniFileTypeCounts[fileType]++
			fmt.Print(fileType, "\t")

			// if fileName == "" {
			// 	continue
			// }
			// if err := DownloadFile(name+"/"+fileName, col); err != nil {
			// 	panic(err)
			// }
		}
		fmt.Println()
	}
	pretty.Println(uniFileTypeCounts)

	// rows, err := f.GetRows("High School")
	// for i, row := range rows {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	name := row[3] + "_" + row[5]
	// 	fmt.Println(name)
	// 	//os.Mkdir(name, 0755)
	// 	for j, col := range row {
	// 		if j < 56 || j > 63 {
	// 			continue
	// 		}
	// 		r := strings.Split(col, "/")
	// 		fileName := r[len(r)-1]
	// 		fmt.Print(fileName, "\t")
	// 		if fileName == "" {
	// 			continue
	// 		}
	// 		if err := DownloadFile(name+"/"+fileName, col); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
