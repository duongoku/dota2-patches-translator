# dota2-patches-translator
Small program that can partly translate [*Dota2*](http://dota2.com/) Patches

[![License is MIT](https://img.shields.io/github/license/duongoku/dota2-patches-translator)](./LICENSE)

- [dota2-patches-translator](#dota2-patches-translator)
	- [Prerequisites](#prerequisites)
	- [How to run it](#how-to-run-it)
	- [Side notes](#side-notes)
	- [License](#license)

## Prerequisites
- [**Go**](https://golang.org/), you can search on google on how to install it

- [**Windows10**](https://www.microsoft.com/en-us/software-download/windows10), I used Windows10, idk if it would work or not on another OS

## How to run it
- Clone this repository

- Go to any [**Dota 2 Patch**](http://dota2.com/patches/) page and save it to you computer by press ```Ctrl + S``` or another way of your choice

- Move the source HTML file to this directory, name it *src.html*

- Open command prompt in this repository's directory by using ```cd dota2-patches-translator```

- Type ```go run main.go``` and press Enter

- Open *des.html* to see the result

##Side notes
- **This will only translate parts of the file so you'll have to do the rest yourself**

- **You may see some broken text, you will have to change the font yourself, [Google](http://google.com/) for more information**

## License

This project is licensed under the [**MIT License**](https://choosealicense.com/licenses/mit).