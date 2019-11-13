package main

import(
	"fmt"
	"time"
)

func numbers(){
	odd := 0
	for i:= 1; i <= 5; i++ {
		odd = 25
		time.Sleep(time.Duration(odd)*time.Millisecond)
		fmt.Printf("%d ",i)

	}
}

func alphabets(){
	for i:='a';i<='e';i++ {
		time.Sleep(40*time.Millisecond)
		fmt.Printf("%c ",i)
	}
}

func main(){
	go numbers()
	go alphabets()
	time.Sleep(time.Millisecond*500)
	fmt.Println("main terminated")
}
