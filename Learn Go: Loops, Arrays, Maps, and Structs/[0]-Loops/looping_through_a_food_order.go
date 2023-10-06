/*
        i decided to stray outside of the bounds of the assignment for this
                cuz i was having fun c:
*/

package main

import (
    f "fmt"
    "bufio"
    "os"
    crypto_rand "crypto/rand"
    "encoding/binary"
    "math/rand"
)

func askOrder() (string) {
    var input string
    f.Print("whaddyaa wantttt?? c: > ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input = scanner.Text()
    return input
}

func contains(menu []string, order string) bool {
    for _, menuItem := range menu {
        if order == menuItem {
            return true
        }
    }
    return false
}

func genRandom(limit int) int{
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed :C")
    }

    var cryptoSeed = int64(binary.LittleEndian.Uint64(b[:]))
    rand.Seed(cryptoSeed)
    var randomInt = rand.Intn(limit)

    return randomInt
}

func printMenu(menu []string) {
    f.Print("the foodies we have areee: ")
    for index, item := range menu{
        if index == 0 {
            f.Print(item)
        } else {
            f.Printf(", %s", item)
        }
    }
    f.Println("s" + "\n" + " > (type 'quit' to quit c:)" + "\n") // the 's' may seem silly but it's to make lox bowls plural
                                                                 //  and still a singular form input for the user c:
}

func orderPrompt(menu []string) int {
      var order string
      var total int
      for order != "quit" {
        order = askOrder()
        if contains(menu, order) {
            f.Println("\n" + " > ooooo yummyyy!!!" + "\n")
            total += genRandom(800)                             // silly lil jokee cuz we're just charging the consumers arbitrary
        } else if order == "quit" {                             //   amounts each time cuz we're an illegitimate business >:3
            f.Println("\n" + " > ummmm wellll... byeeee!!! :D")
        } else {
            f.Println("\n" + " > ummm we don't have that :c" + "\n")
        }
    }
    return total
}


func main() {
    fastfoodMenu := []string{"sushi", "pizza", "chicken", "ice cream", "ramen", "lox bowl"}
    var total = 10
  
    printMenu(fastfoodMenu)
    total += orderPrompt(fastfoodMenu)

    f.Printf("\n\n" + " [*] sorry but u owe us: $%d" + "\n\n", total)
    f.Println("\t" + ":c" + "\n")
}


