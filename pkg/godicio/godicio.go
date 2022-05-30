package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

var wordList [][]string
func main() {

    argsWithoutProg := os.Args[1:]
    c := colly.NewCollector()
	var wordList []string
    c.OnHTML("p span" , func(e *colly.HTMLElement){
		wordList = append(wordList, e.Text)
		}) 
    c.Visit("https://www.dicio.com.br/"+argsWithoutProg[0])
	fmt.Println("Classe da palavra: ", wordList[0])
	fmt.Println("Significados: ", wordList[1], wordList[2])
}