/************************
Created by: Mitchell Noun
Date created: 2/20/22
Class: COMP415 Emerging Languages
Assignment: Project 2
*************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//create channel and spawn the fortune function with the channel as a param
	message := make(chan string)
	fmt.Println("Your fortune is: ")
	go fortune(message)

	//for loop until user responds no
	for true {
		fmt.Println("Do you want another fortune? (yes/no)")
		var response string
		fmt.Scanln(&response) //gets input from user
		response = strings.ToLower(response)
		//if yes, send another fortune
		if response == "yes" {
			go fortune(message)
		}
		//if no, break out of loop
		if response == "no" {
			break
		}
	}
	fmt.Println("Program has ended")
	return
}

func fortune(message chan string) {
	//open and read fortunes.txt
	content, err := ioutil.ReadFile(Fortunes.txt)
	if err != nil {
		log.Fatal(err)
	}
	//split the contents of file
}
