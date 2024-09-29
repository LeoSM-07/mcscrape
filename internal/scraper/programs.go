package scraper

import "fmt"

// import "github.com/gocolly/colly"

type Program struct {
	Title       string
	OfferedBy   string
	Degree      string
	Description string
	Credits     int
	Sections    []ProgramSection
}

type BasicCourse struct {
	Code    string
	Title   string
	Credits int
}

type ProgramSection struct {
	Title   string
	Credits string
	Courses []BasicCourse
}

func ScrapePrograms() {
	electricalEngineering := Program{
		Title:       "Electrical Engineering",
		OfferedBy:   "Electrical & Computer Engr",
		Degree:      "Bachelor of Engineering",
		Description: "...",
		Credits:     134,
		Sections: []ProgramSection{
			{
				Title:   "Required Non-Departmental Courses",
				Credits: "26",
				Courses: []BasicCourse{
					{
						Code:    "CIVE 281",
						Title:   "Analytical Mechanics",
						Credits: 3,
					},
					{
						Code:    "COMP 202",
						Title:   "Foundations of Programming",
						Credits: 3,
					},
					{
						Code:    "COMP 206",
						Title:   "Introduction to Software Systems",
						Credits: 3,
					},
				},
			},
		},
	}

	fmt.Println(electricalEngineering)

	// course := Course{}
	// courses := []Course{}
	//
	// c := colly.NewCollector()
}
