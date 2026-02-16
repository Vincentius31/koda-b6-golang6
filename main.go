package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
	Wait int
}

func processOrder(p Person, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(p.Wait) * time.Second)

	if p.Name != "" {
		fmt.Println("Menunggu antrian...")
		fmt.Printf("Pesanan \"%s\" sudah siap\n", p.Name)
	} else {
		fmt.Println("Nama tidak boleh kosong")
	}
}

func main() {
	
}

