package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// array of first names to pick from
	firstName := [] string {"Lightning", "Blazing", "Piercing", "Sharp", "Mega", "Phantom", "Ice", "Unbreakable", "Enduring", "Bullseye", "Screaming", "Elusive", "Magical", "Power", "Fierce", "Eternal"}
	// array of last names to pick from
	lastName := [] string {"Thunderbolt", "Arrow", "Shard", "Music", "Weapon", "Bow", "String", "Fury", "Hit", "Razor", "Ice", "Fire", "Aim", "Flame", "Power", "Magic"}	
	
	rand.Seed(time.Now().UnixNano())
		
		// generate random number between 0 and length of array
		randomln := rand.Intn(len(lastName))
		// use time for the seed
		// rand.Seed(time.Now().UnixNano())
		// generate random number between 0 and length of array
		randomfn := rand.Intn(len(firstName))
		// use random number to select last name from array
		pickln := lastName[randomln]
		// use random number to select first name from array
		pickfn := firstName[randomfn]
		// print random first name and last name
		fmt.Println(pickfn, pickln)
	
}