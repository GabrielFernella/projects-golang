package main

import "sync"

func main() {
	var wg sync.WaitGroup
	i := 0
	go ChangeNumber(&i, 5)
	go ChangeNumber(&i, 5)
	go ChangeNumber(&i, 20)

	wg.Wait()
}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		go showMessage(strconv.Itoa(i))
// 	}
// }

// func showMessage(message string) {
// 	fmt.Println(message)
// }

// func main() {

// 	var wg sync.WaitGroup // async method waiting
// 	wg.Add(3)
// 	go callDatabase(&wg)
// 	go callAPI(&wg)
// 	go processInternal(&wg)

// 	wg.Wait()

// }

// func callDatabase(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("callDatabase end")
// 	wg.Done()
// }

// func callAPI(wg *sync.WaitGroup) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("callAPI end")
// 	wg.Done()
// }

// func processInternal(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("processInternal end")
// 	wg.Done()
// }
