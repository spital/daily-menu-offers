package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

/* no global vars, no init functions !!
var (
    mutex   sync.Mutex
    balance int
)

func init() {
    balance = 1000
}
*/

// checks if values are equal
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func deposit(value int, wg *sync.WaitGroup, strchan chan string, balance *int, mutex *sync.Mutex) {
	rndval := rand.Intn(2) + 1
	time.Sleep(time.Duration(rndval*500) * time.Millisecond)
	mutex.Lock()
	s := fmt.Sprintf("Depositing %d to account with balance: %d (sleep %d)", value, *balance, rndval)
	fmt.Printf("++ %s\n", s)
	strchan <- s
	*balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup, strchan chan string, balance *int, mutex *sync.Mutex) {
	rndval := rand.Intn(2) + 1
	time.Sleep(time.Duration(rndval*500) * time.Millisecond)
	mutex.Lock()
	s := fmt.Sprintf("Withdrawing %d from account with balance: %d (sleep %d)", value, *balance, rndval)
	fmt.Printf("-- %s\n", s)
	strchan <- s
	*balance -= value
	mutex.Unlock()
	wg.Done()
}

func TestMutextWaitgroupChanOfString(t *testing.T) {
	var (
		mutex   sync.Mutex
		balance int
	)

	balance = 1000
	fmt.Println("Go example of Mutex / Waitgroup / Channel\nInitial balance =", balance)
	strchan := make(chan string, 40)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		/*	type Fn_arg struct  {fn func(); arg int}
			rangevar := []Fn_arg{{deposit,500},{withdraw,700}}
			for j,fnx := range rangevar{
			  wg.Add(1)
			  go fnx.fn(fnx.arg, &wg, strchan,&balance,&mutex)
			  }
		*/
		wg.Add(2)
		go withdraw(700, &wg, strchan, &balance, &mutex)
		go deposit(500, &wg, strchan, &balance, &mutex)
		wg.Wait()

	}
	close(strchan)
	fmt.Printf("New Balance %d\n", balance)
	assertEqual(t, balance, -1000)
	fmt.Println("====final strchan===")
	for s := range strchan {
		fmt.Println(s)
	}
}
