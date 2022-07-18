package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
)

// MemberInput 入力メンバー情報
type MemberInput struct {
	ID         string `csv:"id"`
	FamilyName string `csv:"family_name"`
	FirstName  string `csv:"first_name"`
}

// MemberOutput 出力用メンバー情報
type MemberOutput struct {
	ID          string `csv:"id"`
	FamilyName  string `csv:"family_name"`
	FirstName   string `csv:"first_name"`
	StampedDate string `csv:"stamped_date"`
	StampDate   string `csv:"stamp_date"`
	StampTime   string `csv:"stamp_time"`
	StampType   string `csv:"stamp_type"`
}

func main() {
	fmt.Println("start")

	if err := TransformCSV(); err != nil {
		panic(err)
	}
}

// TransformCSV csvファイルのデータを変換する
func TransformCSV() error {
	inputFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Printf("open input csv file failed (err=%+v)\n", err)
		return err
	}
	defer inputFile.Close()

	inputs := []*MemberInput{}

	if err = gocsv.UnmarshalFile(inputFile, &inputs); err != nil {
		fmt.Printf("unmarshal input csv file failed (err=%+v)\n", err)
		return err
	}

	startDate := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2021, time.June, 30, 0, 0, 0, 0, time.Local)

	dateStrs, err := CreateDateStrs(startDate, endDate)
	if err != nil {
		fmt.Printf("create date string array failed (err=%+v)\n", err)
		return err
	}

	outputs := []*MemberOutput{}
	for _, input := range inputs {
		for _, ds := range dateStrs {
			outputs = append(outputs, &MemberOutput{
				ID:          input.ID,
				FamilyName:  input.FamilyName,
				FirstName:   input.FirstName,
				StampedDate: *ds,
				StampDate:   *ds,
				StampTime:   "09:00",
				StampType:   "出勤",
			})
		}
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Printf("create output csv file failed (err=%+v)\n", err)
		return err
	}

	err = gocsv.MarshalFile(outputs, outputFile)
	if err != nil {
		fmt.Printf("marshal output csv file failed (err=%+v)\n", err)
		return err
	}

	return nil
}

// CreateDateStrs 開始日から終了日までの日付文字列の配列を作成する
func CreateDateStrs(startDate time.Time, endDate time.Time) ([]*string, error) {
	diffUnix := endDate.Unix() - startDate.Unix()
	diffDateStr := fmt.Sprintf("%d", (diffUnix / (60 * 60 * 24)))
	diffDate, err := strconv.Atoi(diffDateStr)
	if err != nil {
		fmt.Printf("convert string to int failed (err=%+v)\n", err)
		return nil, err
	}

	dateStrs := []*string{}
	for i := 0; i <= diffDate; i++ {
		date := startDate.Add(time.Hour * 24 * time.Duration(i))
		year := date.Year()
		month := int(date.Month())
		day := date.Day()
		dateStr := fmt.Sprintf("%d", year) + "/" + fmt.Sprintf("%d", month) + "/" + fmt.Sprintf("%d", day)
		dateStrs = append(dateStrs, &dateStr)
	}

	return dateStrs, nil
}
