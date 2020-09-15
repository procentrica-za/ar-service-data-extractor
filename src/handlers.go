package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (s *Server) handleextractassets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle extract Assets Has Been Called...")
		//Get Asset ID from URL
		assettypeid := r.URL.Query().Get("assettypeid")
		extractedFileName := r.URL.Query().Get("filename")

		//Check if Asset ID provided is null
		if assettypeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset Type ID not properly provided in URL")
			fmt.Println("Asset Type ID not proplery provided in URL")
			return
		}
		if extractedFileName == "" {
			extractedFileName = "AssetRegisterExtract.csv"
		} else {
			extractedFileName = extractedFileName + ".csv"
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/extract?assettypeid=" + assettypeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get asset data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get asset data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []AssetRegisterResponse{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		// Unmarshal JSON data
		var jsonData []Application
		err := json.Unmarshal([]byte(js), &jsonData)

		if err != nil {
			fmt.Println(err)
		}
		//generate UUID as a container name
		fileUUID, _ := newUUID()
		fileName := fileUUID + ".csv"

		//Create CSV file
		csvFile, err := os.Create(fileName)
		fmt.Println("Created file")
		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		//Add data into rows
		for _, usance := range jsonData {
			var row []string
			row = append(row, usance.Name)
			row = append(row, usance.Description)
			row = append(row, usance.SerialNo)
			row = append(row, usance.Size)
			row = append(row, usance.Type)
			row = append(row, usance.Class)
			row = append(row, usance.Dimension1Val)
			row = append(row, usance.Dimension2Val)
			row = append(row, usance.Dimension3Val)
			row = append(row, usance.Dimension4Val)
			row = append(row, usance.Dimension5Val)
			row = append(row, usance.Dimension6Val)
			row = append(row, usance.Extent)
			row = append(row, usance.ExtentConfidence)
			row = append(row, usance.DeRecognitionvalue)
			writer.Write(row)
		}
		fmt.Println("Populated CSV")

		// flush the writer
		writer.Flush()
		//Open file
		Openfile, err := os.Open(fileName)
		//Read 512 bytes of file data
		FileHeader := make([]byte, 512)
		Openfile.Read(FileHeader)
		//detect file type
		FileContentType := http.DetectContentType(FileHeader)
		FileStat, _ := Openfile.Stat()
		FileSize := strconv.FormatInt(FileStat.Size(), 10)

		//Make file downloadable
		w.Header().Set("Content-Disposition", "attachment; filename="+extractedFileName)
		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", FileSize)
		//return to beginning  of array
		Openfile.Seek(0, 0)

		var dlterror = os.Remove(fileName)
		if dlterror != nil {
			fmt.Println(dlterror)
		}
		fmt.Println("File has been deleted -->" + fileName)
		//Send the file
		io.Copy(w, Openfile)
	}
}

func populate() {
	// read data from file
	jsonDataFromFile, err := ioutil.ReadFile("./company.json")

	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal JSON data
	var jsonData []Application
	err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)

	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("./data.csv")
	fmt.Println("Created file")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, usance := range jsonData {
		var row []string
		row = append(row, usance.Name)
		row = append(row, usance.Description)
		row = append(row, usance.SerialNo)
		row = append(row, usance.Size)
		row = append(row, usance.Type)
		row = append(row, usance.Class)
		row = append(row, usance.Dimension1Val)
		row = append(row, usance.Dimension2Val)
		row = append(row, usance.Dimension3Val)
		row = append(row, usance.Dimension4Val)
		row = append(row, usance.Dimension5Val)
		row = append(row, usance.Dimension6Val)
		row = append(row, usance.Extent)
		row = append(row, usance.ExtentConfidence)
		row = append(row, usance.DeRecognitionvalue)
		writer.Write(row)
	}
	fmt.Println("Populated CSV")

	// flush the writer
	writer.Flush()
}
