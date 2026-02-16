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
	persons := []Person{
		{Name: "Koda", Wait: 5},
		{Name: "Mamet", Wait: 4},
		{Name: "Cath", Wait: 7},
		{Name: "Kevin", Wait: 6},
		{Name: "Mark", Wait: 8},
	}

	queueChan := make(chan Person)
	var wg sync.WaitGroup

	go func() {
		for p := range queueChan {
			processOrder(p, &wg)
		}
	}()

	fmt.Println("Memulai antrian pemesanan...")
	for _, p := range persons {
		wg.Add(1)
		queueChan <- p
	}

	wg.Wait()
	close(queueChan)
	fmt.Println("Selesai")
}

