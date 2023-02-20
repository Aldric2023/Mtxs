// mutexes
package main

import (
	"fmt"
	"sync"
)

func withdraw(balance *int, wg *sync.WaitGroup, mx *sync.Mutex, amount int){
	defer wg.Done()
	defer mx.Unlock()

	//get exclusive acceess to the thing we want to change :: balance variable
	mx.Lock()

	//critical region
	*balance = *balance - amount;

	mx.Unlock()
}

func deposit(balance *int, wg *sync.WaitGroup, mx *sync.Mutex, amount int){
	defer wg.Done()
	defer mx.Unlock()

	//get exclusive acceess to the thing we want to change :: balance variable
	mx.Lock()
	
	//critical region
	*balance = *balance + amount;

	mx.Unlock()
	
}

func main(){

	balance := 1000;

	//create waitgroup
	var wg sync.WaitGroup
	//create mutex
	var mx sync.Mutex
	//create my go routines
	wg.Add(2)

	go deposit(&balance, &wg, &mx,100)

	go withdraw(&balance, &wg, &mx, 40)


	//wait on the go routines to finish

	wg.Wait()

	fmt.Println("Balance: ",balance)

}