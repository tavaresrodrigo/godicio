package dicio

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var wordList [][]string
func GetPTWord(arg[]string) {
    c := colly.NewCollector()
	var wordList []string
    c.OnHTML("p span" , func(e *colly.HTMLElement){
		wordList = append(wordList, e.Text)
		}) 
    c.Visit("https://www.dicio.com.br/"+arg[0])
	fmt.Println("Classe da palavra: ", wordList[0])
	fmt.Println("Significados: ", wordList[1], wordList[2])
}