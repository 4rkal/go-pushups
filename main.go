package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"os"
	"time"
)

func main() {
	var reps uint
	var increase uint
	var rest uint
	var confirm string
	var amount_done uint
	amount_done = 0

	for {
		fmt.Println("Welcome to go-pushups")
		fmt.Printf("Enter amount of reps: ")
		fmt.Scan(&reps)
		fmt.Printf("Percentage increase per round eg 10%% first round 10 then 11 etc (leave 0 for none): ")
		fmt.Scan(&increase)
		fmt.Printf("Amount of rest (in seconds): ")
		fmt.Scan(&rest)
		fmt.Printf("You will be doing %v reps with a %v%% increase per round with %v sec rest [y/n]: ", reps, increase, rest)
		fmt.Scan(&confirm)
		if confirm == "y" {
			break
		} else if confirm == "n" {
			fmt.Println("ok, redoing")
		} else {
			fmt.Printf("Didn't get that")
		}
	}

	fmt.Println("Starting")
	time.Sleep(time.Duration(rest) * time.Second)
	response := fmt.Sprintf("Do %v push-ups", reps)
	beeep.Alert("go-pushups", response, "logo.png")
	fmt.Printf("Do %v push-ups \n", reps)
	fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
	fmt.Scan(&confirm)
	if confirm == "y" {
		fmt.Println("Ok")
	} else if confirm == "q" {
		fmt.Printf("%d", amount_done)
		os.Exit(0)
	}

	time.Sleep(time.Duration(rest) * time.Second)

	if increase == 0 {
		for {
			response := fmt.Sprintf("Do %v push-ups", reps)
			beeep.Alert("go-pushups", response, "logo.png")
			time.Sleep(time.Duration(rest) * time.Second)
			fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
			fmt.Scan(&confirm)
			if confirm == "y" {
				fmt.Println("Ok")
			} else if confirm == "q" {
				fmt.Printf("%d", amount_done)
				os.Exit(0)
			}
			amount_done = amount_done + reps
			time.Sleep(time.Duration(rest) * time.Second)
		}
	} else {
		for {
			amount := reps + reps*increase/100
			response := fmt.Sprintf("Do %v push-ups", amount)
			beeep.Alert("go-pushups", response, "logo.png")
			fmt.Println(response)
			fmt.Printf("Did you do them (yes or quit)? [y/q]: ")
			fmt.Scan(&confirm)
			if confirm == "y" {
				fmt.Println("Ok")
			} else if confirm == "q" {
				fmt.Printf("%d", amount_done)
				os.Exit(0)
			}
			time.Sleep(time.Duration(rest) * time.Second)
			reps = reps + reps*increase/100
			amount_done = amount_done + reps
		}
	}
}
