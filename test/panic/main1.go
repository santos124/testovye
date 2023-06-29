package main

import (
	"log"
	"time"

)

func main() {
	log.Println("privet")
	go privet()
	time.Sleep(10 * time.Second)
	log.Println("privet2")

}

func privet() {
	defer log.Println("privet defer")
	defer func() {
		err := recover()
		log.Printf("%+v", err)
	}()
	a := map[int]int{}
	delete(a, 1)
	b := []int{}
	log.Println(maps.Values(a), a)
	time.Sleep(time.Second * 2)
	panic("privet panic")
}

// Принцип единственной ответственности
// Принцип открытости/закрытости
// Принцип подстановки Барбары Лисков
// Принцип разделения интерфейса
// Принцип инверсии зависимостей
