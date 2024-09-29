package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/leosm-07/mcscrape/internal/scraper"
)

func main() {
	codes := []string{
		"BIOC 212",
		"BIOC 220",
		"BIOL 200",
		"BIOL 202",
		"CHEM 204",
		"CHEM 212",
		"CHEM 222",
		"BIOL 205",
		"MIMM 211",
		"MIMM 214",
		"PHGY 209",
		"PHGY 210",
		"ANAT 262",
		"BIOC 311",
		"BIOC 312",
		"BIOC 320",
		"CHEM 214",
		"CHEM 302",
		"CHEM 362",
		"BIOL 309",
		"BIOL 373",
		"CHEM 267",
		"COMP 202",
		"COMP 204",
		"MATH 203",
		"MATH 222",
		"PSYC 204",
		"BIOC 450",
		"BIOC 454",
		"BIOC 458",
		"BIOC 470",
		"BIOC 491",
		"BIOC 503",
		"PSYT 455",
		"BIOL 300",
		"BIOL 303",
		"BIOL 304",
		"BIOL 313",
		"BIOL 314",
		"CHEM 267",
		"CHEM 482",
		"CHEM 502",
		"CHEM 532",
		"CHEM 552",
		"CHEM 572",
		"EXMD 502",
		"MIMM 324",
		"PHAR 300",
		"PHGY 311",
	}

	courses := scraper.ScrapeCourses(codes)

	file, err := os.Create("courses.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	enc.Encode(courses)
}
