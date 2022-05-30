# GoDicio

Godicio is a CLI tool to search for words definitions in online dictionaries, it helps CLI people to use a dictionary from the Terminal in different lanaguages as English, Spanish and Portuguese.
	
## Example
	
Searching for the Portuguese word "verdade".

```
$godicio pt verdade
```

Searching for the English word "truth".

```
$godicio en truth
```

Searching for the Spanish word "verdad".

```
$godicio es verdad
```

## Installing Godicio

This project uses [GoReleaser](https://goreleaser.com/) to build multi-platform bynaries that are available on the [releases tab](https://github.com/tavaresrodrigo/godicio/releases) of this repository. 

If you prefer you can also clone this repository and build your own bynary, you just need to make sure you have your GOPATH variable set. 

```
$ git clone https://github.com/tavaresrodrigo/godicio
$ git install 
$ ls $GOPATH/bin
$ godicio pt obrigado
Classe da palavra:  interjeição
Significados:  Expressão de gratidão, usada pelo gênero masculino: já recebi o pagamento, obrigado. adjetivo
```

## Contributing

Feel free to raise an issue if you have questions or if you want to contribute, I will be happy to collaborate. 