package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

// Simulate deadlock for mutex
// The 1st Transfer call will lock user Eko
// The 2nd Transfer call will lock user Budi
// The 1st Transfer call need to lock user Budi,
  // but cannot do,
  // because 2nd Transfer call already locked user Budi,
  // it will wait for 2nd Transfer call to unlock user Budi.
// The 2nd Transfer call need to lock user Eko
  // but cannot do,
  // because 1st Transfer call already locked user Eko,
  // it will wait for 1st Transfer call to unlock user Eko.
// Both Transfer calls will wait for each other to unlock the users forever
// This situation is named deadlock
func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	// Both users' balance is already deducted
	// But since we exceed the timeout we set with time.Sleep()
	// We proceed the code, and printed out this inconsistent state
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
