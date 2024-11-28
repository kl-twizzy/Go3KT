package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

// Структура для животного
type Animal struct {
	Name              string
	HighSpeed         bool
	Size              string
	ClimbTree         bool
	RecognizeDiseases bool
}

// Функция для загрузки информации о животном
func loadAnimalInfo(animal Animal) {
	// Имитация загрузки информации
	time.Sleep(time.Second)
	fmt.Printf("Загружена информация о %s:\n", animal.Name)
	fmt.Printf("HighSpeed: %v\n", animal.HighSpeed)
	fmt.Printf("Size: %s\n", animal.Size)
	fmt.Printf("ClimbTree: %v\n", animal.ClimbTree)
	fmt.Printf("RecognizeDiseases: %v\n", animal.RecognizeDiseases)
}

// Функция для отправки уведомления
func sendNotification(animal Animal) {
	err := beeep.Notify("Информация о животном", animal.Name, "")
	if err != nil {
		panic(err)
	}
}

func main() {
	// Создаем WaitGroup для отслеживания горутин
	var wg sync.WaitGroup

	// Список животных
	animals := []Animal{
		{Name: "Лев", HighSpeed: true, Size: "Большой", ClimbTree: false, RecognizeDiseases: false},
		{Name: "Обезьяна", HighSpeed: false, Size: "Средний", ClimbTree: true, RecognizeDiseases: true},
		{Name: "Зебра", HighSpeed: true, Size: "Средний", ClimbTree: false, RecognizeDiseases: false},
		{Name: "Тигр", HighSpeed: true, Size: "Большой", ClimbTree: false, RecognizeDiseases: false},
		{Name: "Кот", HighSpeed: false, Size: "Маленький", ClimbTree: true, RecognizeDiseases: false},
	}

	// Выводим список животных для выбора
	fmt.Println("Доступные животные:")
	for i, animal := range animals {
		fmt.Printf("%d. %s\n", i+1, animal.Name)
	}

	// Считываем выбор пользователя
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Выберите животных (через запятую): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choices := strings.Split(input, ",")

	// Запускаем горутины для выбранных животных
	for _, choice := range choices {
		index, err := strconv.Atoi(strings.TrimSpace(choice))
		if err != nil || index < 1 || index > len(animals) {
			fmt.Printf("Неверный выбор: %s\n", choice)
			continue
		}

		animal := animals[index-1]
		wg.Add(1)
		go func(a Animal) {
			defer wg.Done()
			loadAnimalInfo(a)
			go sendNotification(a) // Асинхронный вызов уведомления
		}(animal)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Ожидание завершения всех звуков (в данном случае, уведомлений)
	time.Sleep(2 * time.Second)

	fmt.Println("Программа завершена.")
}
