package main

import "fmt"

// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Println(a)
// }
// func reverse(b *[4]int) {
// 	for index, value := range *b {
// 		(*b)[len(b)-1-index] = value
// 	}
// }
//

/* Создать приложение, которое сначала выдаёт меню:
- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход
При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}

// package main

// import "fmt"

// func main() {
// 	m := map[string]string{}
// 	// 2

// 	for {

// 		fmt.Println("1.Show\n2.Add\n3.Delete\n4.Exit")

// 		var a int
// 		fmt.Scan(&a)
// 		switch {
// 		case a == 1:
// 			printUser(m)
// 		case a == 2:
// 			addUser(m)
// 		case a == 3:
// 			removeUser(m)
// 		case a == 4:
// 			break
// 		default:
// 			fmt.Println("Error")
// 		}
// 	}

// 	// 1

// 	// m := map[string]string{
// 	// 	"Beka": "boxing",
// 	// }
// 	// fmt.Println(m)
// 	// m["Beka"] = "IT"
// 	// m["Ali"] = "IT"
// 	// fmt.Println(m)
// 	// delete(m, "Ali")
// 	// fmt.Println(m)

// 	// for key, value := range m {
// 	// 	fmt.Println(key, value)
// 	// }

// }
// func printUser(m map[string]string) {
// 	for key, value := range m {
// 		fmt.Println(key, " : ", value)
// 	}
// }
// func addUser(m map[string]string) {
// 	var name string
// 	var fullName string
// 	fmt.Print("Name: ")
// 	fmt.Scan(&name)
// 	fmt.Print("FullName: ")
// 	fmt.Scan(&fullName)
// 	m[name] = fullName

// }
// func removeUser(m map[string]string) {
// 	var name string
// 	fmt.Print("Write Name")
// 	fmt.Scan(&name)
// 	delete(m, name)
// }
