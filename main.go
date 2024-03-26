package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction uint8

const (
	Ausgabe direction = iota
	Einnahme
)

type Entry struct {
	year     string
	month    string
	day      string
	category string
	amount   float64
	desc     string
	dir      direction
}

func main() {
	var FILE string
	flag.StringVar(&FILE, "file", "", "CSV file to process")
	flag.Parse()
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.Comma = ';'
	content, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	parseEntries(content)
}

func parseEntries(c [][]string) {
	for line, e := range c {
		if line == 0 {
			continue
		}
		//fmt.Printf("line:%d, content:%s\n", line, e)
		var t Entry
		var err error
		// year
		t.year = e[0]
		// month
		if len(e[1]) < 2 {
			t.month = "0" + e[1]
		} else {
			t.month = e[1]
		}
		// day
		if len(e[2]) < 2 {
			t.day = "0" + e[2]
		} else {
			t.day = e[2]
		}
		// category
		t.category = e[3]
		// amount
		tmpAmount := strings.Replace(e[4], ",", ".", 1)
		t.amount, err = strconv.ParseFloat(tmpAmount, 8)
		if err != nil {
			log.Fatal(err)
		}
		// description
		if e[5] == "" {
			t.desc = "?"
		} else {
			t.desc = e[5]
		}
		// direction
		if e[6] == "Einnahme" {
			t.dir = Einnahme
		} else {
			t.dir = Ausgabe
		}

		date := fmt.Sprintf("%s-%s-%s * \"%s\"\n", t.year, t.month, t.day, t.desc)
		var rec string
		if t.dir == Einnahme {
			rec = fmt.Sprintf("  Income:%s\n  Assets:                %.2f EUR\n\n", t.category, t.amount)
		} else {
			rec = fmt.Sprintf("  Expenses:%s              %.2f EUR\n  Assets:\n\n", t.category, t.amount)
		}
		fmt.Printf(date)
		fmt.Printf(rec)
	}
}
