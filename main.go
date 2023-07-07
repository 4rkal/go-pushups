package main

import (
	"fmt"
	"time"
	"github.com/gen2brain/beeep"
	"os"
)


func main() {
	var reps uint
	var increase uint
	var rest uint
	var confirm string
	for {
		fmt.Println("Welcome to go-pushups")
		fmt.Printf("Enter amount of reps: ")
		fmt.Scan(&reps)
		fmt.Printf("Percentage increase per round eg 10%% first round 10 then 11 etc (leave 0 for none): ")
		fmt.Scan(&increase)
		fmt.Printf("Amount of rest (in seconds): ")
		fmt.Scan(&rest)
		fmt.Printf("You will be doing %v reps with a %v %% increase per round with %v sec rest [y/n]: ",reps,increase,rest)
		fmt.Scan(&confirm)
		if confirm == "y" {
			break
		} else if confirm == "n" {
			fmt.Println("ok, redoing")
		} else {
			fmt.Printf("Didnt get that")
		}
	}
	fmt.Println("Starting")
	time.Sleep(time.Duration(rest)* time.Second)
	response := fmt.Sprintf("Do %v push-ups",reps)
	beeep.Alert("go-pushups", response,"logo.png")
	fmt.Printf("Do %v push-ups \n",reps)
	fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
	fmt.Scan(&confirm)
	if confirm == "y" {
		fmt.Println("Ok")
	} else if confirm == "q" {
		os.Exit(0)
	}
	time.Sleep(time.Duration(rest)* time.Second)
	if increase == 0{
		for {
			response := fmt.Sprintf("Do %v push-ups",reps)
			beeep.Alert("go-pushups", response,"logo.png")
			time.Sleep(time.Duration(rest) * time.Second)
			fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
			fmt.Scan(&confirm)
			if confirm == "y" {
				fmt.Println("Ok")
			} else if confirm == "q" {
				os.Exit(0)
			}
			time.Sleep(time.Duration(rest) * time.Second)
		}
	} else {
		for {
			amount := reps + reps * increase /100
			response := fmt.Sprintf("Do %v push-ups",amount)
			beeep.Alert("go-pushups", response,"logo.png")
			fmt.Println(response)
			fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
			fmt.Scan(&confirm)
			if confirm == "y" {
				println("Ok")
			} else if confirm == "q" {
				os.Exit(0)
			}
			time.Sleep(time.Duration(rest) * time.Second)
			reps = reps + reps * increase /100
		}
	}
}
