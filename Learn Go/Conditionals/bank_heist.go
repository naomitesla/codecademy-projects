package main

import (
    f "fmt"
    crypto_rand "crypto/rand"
    "encoding/binary"
    "math/rand"
)

func main() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed :C")
    }

    var cryptoSeed = int64(binary.LittleEndian.Uint64(b[:]))
    rand.Seed(cryptoSeed)

    var eludedGuards = rand.Intn(100)
    var isHeistOn = true

    if eludedGuards >= 50 {
        f.Println("yayyyy!! dumb guards >:3")
    } else {
        isHeistOn = false
        f.Println("omg, u guys r soooo dumb :C")
    }

    f.Print("\n")
    var openedVault = rand.Intn(100)

    if isHeistOn && openedVault >= 70 {
            f.Println("yayayayayyyy!! freee moniesss!!!! :D" + "\n")
    } else if isHeistOn {
        isHeistOn = false
        f.Println("ughghgh uu guys didn't open the safe?")
        f.Println("i should've just done it myself ;-;" + "\n")
    }
    
    var leftSafely = rand.Intn(5)

    if isHeistOn {
        f.Println("NOW, LETS LEAVEEEE >:3")
    } else {
        f.Println("umm less just go homeee :c")
    }

    f.Println("...")
    var amtStolen int
    switch leftSafely {
    case 0:
        f.Println("ooo nyooo!! D:")
        f.Println(" we forgot to turn off the cameras :c")
        isHeistOn = false
    case 1:
        f.Println("i got sleepyy and fell asleep :c")
        f.Println(" i sorryyyyy :c")
        isHeistOn = false
    case 2:
        f.Println("aaaaahhhhh!!!! aliensss!!!!")
        isHeistOn = false
    case 3:
        f.Println("my frennn wanted to hang out and i had to gooo sorryyyyy!!!")
        isHeistOn = false
    default:
        if isHeistOn {
            f.Println("YAYYYYYYYY!!! WE GOT OUT AND IT WASN'T EVEN UMM *THAT* SCARYY!! :D")
            amtStolen = 10000 + rand.Intn(1000000)
        } else {
            f.Println("well at least we made it out okieee c:")
        }
    }
    
    f.Print("\n")

    if isHeistOn {
        f.Printf("we managed to stealllllll.... $%d MONIESSSS!!!!!!!!!!" + "\n", amtStolen)
    } else if leftSafely <= 3 {
        f.Println("we didn't get the moniess and still got caught :c")
    } else {
        f.Println("we didn't make any monies but we made it out okiee c:")
    }
}
