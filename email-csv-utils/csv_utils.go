package csv_utils

import (
	"encoding/csv"
    "strconv"
	"fmt"
	"log"
	"os"
)

type EmailRecord struct {
	ID string
	Email string
}

func CheckEmailPresence(emailRecords []EmailRecord, emailToCheck string) bool{
	for _, emailRecord := range emailRecords {
		if emailRecord.Email == emailToCheck {
			return true
		}
	}
	return false
}

func WriteEmailRecordToFile(pathCSV string, email string) int{
	file, _ := os.OpenFile(pathCSV, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	emailRecords := ReadCSV(pathCSV)

	if CheckEmailPresence(emailRecords, email) {
		return 1
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	lastRecordIDString := emailRecords[len(emailRecords)-1].ID
	lastRecordIDInteger, _ := strconv.Atoi(lastRecordIDString)
	newIDInteger := lastRecordIDInteger + 1
	newIDString := strconv.Itoa(newIDInteger)

	newRecord := EmailRecord {ID: newIDString, Email: email}

	row := []string{newRecord.ID, newRecord.Email}
	if err := w.Write(row); err != nil {
		log.Fatalln("Error writing record to file", err)
		return 2
	}
	return 0
}

func ReadCSV(pathCSV string) []EmailRecord {
	file, err := os.Open(pathCSV)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var emailRecords []EmailRecord
	for _, record := range records{
		data := EmailRecord {
			ID: record[0],
			Email: record[1],
		}
		emailRecords = append(emailRecords, data)
		// fmt.Println(data.Email)
		// fmt.Println(data.ID)
	}

	// fmt.Println(emailRecords)
	return emailRecords
}

// func main(){
// 	pathCSV := "example.csv"
// 	// sampleRecord := EmailRecord{ID: "4", Email: "jojo@jotaro.com"}
// 	WriteEmailRecordToFile(pathCSV, "jojo@jotaro.com")
// 	emailRecords := readCSV(pathCSV)
// 	fmt.Println(emailRecords)
// }