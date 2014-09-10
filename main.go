package main

import (
	"fmt"
	"runtime"
	"strings"
)

var loremIpsum string
var finalIpsum string
var letterSentChan chan string

func deliverToFinal(letter string, finalIpsum *string){
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIpsum *string){
	for *current < length{
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
}

func main(){
	runtime.GOMAXPROCS(2)

	index := new(int)
	*index = 0
	loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum venenatis magna eget libero tincidunt, accondimentum enim auctor. Integer mauris arcu, dignissim sit amet convallis vitae, ornare vel odio. Phasellus in lectus risus. Utsodales vehicula ligula eu ultricies. Fusce vulputate fringillaeros at congue. Nulla tempor neque enim, non malesuada arculaoreet quis. Aliquam eget magna metus. Vivamus laciniavenenatis dolor, blandit faucibus mi iaculis quis. Vestibulum sit amet feugiat ante, eu porta justo."

	letters := []byte(loremIpsum)
	length := len(letters)
	go capitalize(index, length, letters, &finalIpsum)
	go func(){
		go capitalize(index, length, letters, &finalIpsum)
	}()

	// Force the runtime to wait for the goroutines to complete
	runtime.Gosched()
	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)
}
