package main

import (
	"fmt"

	documentstore "github.com/AzartArn/go-leng-course/lessons/lesson03/document_store"
)

func main() {
	fmt.Println("--- START ---")

	// створення колекції данних

	myDoc := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			// Обов'язкове поле key
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "user_555",
			},
			// Додаткове поле
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Dmytro",
			},
			"age": {
				Type:  documentstore.DocumentFieldTypeNumber,
				Value: 30,
			},
		},
	}

	// демонстрація функції Put()

	documentstore.Put(myDoc)
	fmt.Println("[Put] Документ збережено.")

	// демонстрація функції Get()

	doc, found := documentstore.Get("user_555")
	if found {
		nameVal := doc.Fields["name"].Value
		fmt.Printf("[Get] Знайдено! Ім'я: %v\n", nameVal)
	} else {
		fmt.Println("[Get] Помилка: Документ не знайдено.")
	}

	// демонстарція функції List()

	list := documentstore.List()

	// виводимо кількість документів у колекуії
	fmt.Printf("[List] Всього документів: %d\n", len(list))

	// виводемо кожен документ колекції
	for i, doc := range list {
		// %d - номер по порядку
		// %+v - формат для друку структур і мап
		fmt.Printf("  %d. Вміст: %+v\n", i+1, doc.Fields)
	}

	// демонстрація функції Delete()

	deleted := documentstore.Delete("user_555")
	if deleted {
		fmt.Println("[Delete] Документ успішно видалено.")
	}

	// перевірака що видалення відбулося

	_, foundAgain := documentstore.Get("user_555")
	if !foundAgain {
		fmt.Println("[Check] Документа більше немає в базі.")
	}
}
