package basic_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Printf("Waktu sekarang: %v\n", now)
	fmt.Printf("Waktu yang akan datang: %v\n", now.AddDate(0, 0, 1))
	fmt.Println(now.Year())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.String())
}

func TestParsingTime(t *testing.T) {
	datetime := time.Now()
	//currentMilli := time.Now().UnixMilli()

	fmt.Println(datetime.Format(time.RFC850))

	//layoutFormat := "02/01/2006 MST"
	value := "24/12/2024 WIB"
	date, _ := time.Parse(time.RFC822, value)
	fmt.Println(value, "\t\t", date)
}
