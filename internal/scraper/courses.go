package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Course struct {
	Title       string
	Code        string
	Credits     int
	Overview    string
	Meta        string
	Terms       []string
	Instructors string
	Notes       []string
}

func ScrapeCourses(codes []string) []Course {
	course := Course{}
	courses := []Course{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error while scraping", e.Error())
	})

	// Course Title
	c.OnHTML("h1#page-title", func(h *colly.HTMLElement) {
		course.Title = cleanText(h.Text)
		credits, err := extractCredits(course.Title)
		if err == nil {
			course.Credits = credits
		}
	})

	// Terms
	c.OnHTML("p.catalog-terms", func(e *colly.HTMLElement) {
		trimmed := strings.TrimPrefix(cleanText(e.Text), "Terms:      ")
		terms := strings.Split(trimmed, ", ")
		course.Terms = terms
	})

	// Instructors
	c.OnHTML("p.catalog-instructors", func(e *colly.HTMLElement) {
		trimmed := strings.TrimPrefix(cleanText(e.Text), "Instructors:      ")
		course.Instructors = trimmed
	})

	// Meta
	c.OnHTML("div.meta", func(e *colly.HTMLElement) {
		course.Meta = cleanText(e.Text)
		// fmt.Println(cleanText(e.Text))
	})

	// Overview Section
	c.OnHTML("div.content > h3:first-child", func(e *colly.HTMLElement) {
		overview := e.DOM.Next()

		course.Overview = cleanText(overview.Text())
	})

	// Notes Section
	c.OnHTML("ul.catalog-notes", func(e *colly.HTMLElement) {
		notes := e.DOM.Find("li")
		notes.Each(func(_ int, s *goquery.Selection) {
			course.Notes = append(course.Notes, cleanText(s.Text()))
		})
	})

	c.OnScraped(func(r *colly.Response) {
		courses = append(courses, course)
		course = Course{}
	})

	for _, code := range codes {
		course.Code = code
		c.Visit(courseUrl(code))
	}

	return courses
}

func cleanText(s string) string {
	return strings.TrimSpace(s)
}

func extractCredits(input string) (int, error) {
	re := regexp.MustCompile(`\((\d+) credit[s]?\)`)

	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		credits, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		return credits, nil
	}

	return 0, fmt.Errorf("no match found for credits")
}

func courseUrl(code string) string {
	return fmt.Sprint(
		"https://www.mcgill.ca/study/2024-2025/courses/",
		strings.ReplaceAll(strings.ToLower(code), " ", "-"),
	)
}
