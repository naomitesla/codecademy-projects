package main

import (
    f "fmt"
)

func main(){
    var backTick string = "`"
    f.Printf(`      %[1]v.-::::::-.%[1]v
  .:-::::::::::::::-:.
  %[1]v_:::    ::    :::_%[1]v
   .:( ^   :: ^   ):.
   %[1]v:::   (..)   :::.
   %[1]v:::::::UU:::::::%[1]v
   .::::::::::::::::.
   O::::::::::::::::O
   -::::::::::::::::-
   %[1]v::::::::::::::::%[1]v
   .::::::::::::::.
     oO:::::::Oo\n`, backTick)

    f.Printf(`
         __      _
        o''}____//
         %[1]v_/      )
         (_(_/-(_/
    `, backTick)
}
