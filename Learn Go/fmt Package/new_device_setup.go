package main

import f "fmt"

func main(){
    var name string
    var age int8
    f.Print("What's your nameeeee? ")
    f.Scan(&name)
    f.Printf("hiiiiii %v!!!" + "\n", name)
    f.Print("What's your agee!?!? ")
    f.Scan(&age)
    var newMessage = f.Sprintf("\n" + "umm i have terrible news for uu %v, you're %d years old :c", name, age)
    f.Println(newMessage)
}
