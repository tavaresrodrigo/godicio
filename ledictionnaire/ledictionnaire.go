package ledictionnaire

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func GetFRWord(arg[]string) {
    c := colly.NewCollector()
	fmt.Println(arg[0])

	// Fetching the word definition
    c.OnHTML("#maincontent > div:nth-child(4) > ul > li:nth-child(1)", func(e *colly.HTMLElement){
		wordMeaning := e.Text
		fmt.Println("Meaning: ", wordMeaning)
		}) 
    c.Visit("https://www.le-dictionnaire.com/definition/"+arg[0])

}
