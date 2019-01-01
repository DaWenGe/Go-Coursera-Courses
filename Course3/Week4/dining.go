package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number  int
	leftCS  *Chopstick
	rightCS *Chopstick
}

type Host struct {
	chopsticksMap map[*Chopstick]bool
	sync.Mutex
}

func (host *Host) getPermission(p *Philosopher) bool {
	inUse := false
	host.Lock()
	inUse = host.chopsticksMap[p.leftCS] || host.chopsticksMap[p.rightCS]
	if !inUse {
		host.chopsticksMap[p.leftCS] = true
		host.chopsticksMap[p.rightCS] = true
	}
	host.Unlock()
	return inUse
}

func (host *Host) doneEating(p *Philosopher) {
	host.Lock()
	host.chopsticksMap[p.leftCS] = false
	host.chopsticksMap[p.rightCS] = false
	host.Unlock()
}

func (p *Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println("Eating round: ", i+1)
		if ok := host.getPermission(p); !ok {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("starting to eat <%d>\n", p.number)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			fmt.Printf("finishing eating <%d>\n", p.number)
			host.doneEating(p)
		}
	}
	wg.Done()
}

func main() {
	host := Host{}
	host.chopsticksMap = make(map[*Chopstick]bool)

	Chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		Chopsticks[i] = new(Chopstick)
		host.chopsticksMap[Chopsticks[i]] = false
	}

	Philosophers := make([]*Philosopher, 5)
	for j := 0; j < 5; j++ {
		Philosophers[j] = &Philosopher{j + 1, Chopsticks[j], Chopsticks[(j+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go Philosophers[i].eat(&host, &wg)
	}

	wg.Wait()
	fmt.Println("Everyone is full!")
}
