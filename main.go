package main

import (
	"fmt"
	"os"

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

	todayStr := "2006/01/05"

	outputs := []*MemberOutput{}
	for _, input := range inputs {
		outputs = append(outputs, &MemberOutput{
			ID:          input.ID,
			FamilyName:  input.FamilyName,
			FirstName:   input.FirstName,
			StampedDate: todayStr,
			StampDate:   todayStr,
			StampTime:   "09:00",
			StampType:   "出勤",
		})
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Printf("create output csv file failed (err=%+v)\n", err)
		return err
	}

	err = gocsv.MarshalFile(outputs, outputFile); if err != nil {
		fmt.Printf("marshal output csv file failed (err=%+v)\n", err)
		return err
	}

	return nil
}
