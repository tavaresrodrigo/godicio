package dicio

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func GetENWord(arg[]string) {
    c := colly.NewCollector()
	fmt.Println(arg[0])

	// Fetching the word classification
	c.OnHTML("#father_topg_2 > div > span.pos" , func(e *colly.HTMLElement){
		wordClass := e.Text
		fmt.Println("Class: ", wordClass)
		}) 

	// Fetching the word definition
    c.OnHTML("#father_sng_1 > span.def" , func(e *colly.HTMLElement){
		wordMeaning := e.Text
		fmt.Println("Meaning: ", wordMeaning)
		}) 
		
    c.Visit("https://www.oxfordlearnersdictionaries.com/definition/english/"+arg[0]+"_1?q="+arg[0])

}