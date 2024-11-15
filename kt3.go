package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

type Lion struct {
	HighSpeed         int
	RecognizeDiseases bool
}

type Giraffe struct {
	Size      int
	ClimbTree bool
}

type Snake struct {
	Length   int
	Venomous bool
}

func processLion(wg *sync.WaitGroup) {
	defer wg.Done()
	lion := Lion{HighSpeed: 80, RecognizeDiseases: true}
	fmt.Printf("Лев - Скорость: %d км/ч, Распознавание болезней: %v\n", lion.HighSpeed, lion.RecognizeDiseases)

	go func() {
		beeep.Notify("Лев", "Обработана информация о льве", "")
	}()
	time.Sleep(1 * time.Second)
}

func processGiraffe(wg *sync.WaitGroup) {
	defer wg.Done()
	giraffe := Giraffe{Size: 500, ClimbTree: false}
	fmt.Printf("Жираф - Размер: %d кг, Может залезать на деревья: %v\n", giraffe.Size, giraffe.ClimbTree)

	go func() {
		beeep.Notify("Жираф", "Обработана информация о жирафе", "")
	}()
	time.Sleep(1 * time.Second)
}

func processSnake(wg *sync.WaitGroup) {
	defer wg.Done()
	snake := Snake{Length: 2, Venomous: true}
	fmt.Printf("Змея - Длина: %d м, Ядовитая: %v\n", snake.Length, snake.Venomous)

	go func() {
		beeep.Notify("Змея", "Обработана информация о змее", "")
	}()
	time.Sleep(1 * time.Second)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go processLion(&wg)
	go processGiraffe(&wg)
	go processSnake(&wg)

	wg.Wait()
	fmt.Println("Все животные обработаны.")
}
