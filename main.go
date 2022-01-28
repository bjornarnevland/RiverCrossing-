package main


import (
	"bufio"
	"fmt"

	"github.com/bjornarnevland/RiverCrossing-/event"
	"os"

	//"rc1-verden/event"
)

/*var items = [] string{
"kylling",
"rev",
"korn",
}*/


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Hva vil du sende over? --> kylling, rev eller korn")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(event.FirstPut(text))
		fmt.Println(event.GetIn(text))
		fmt.Println(event.CrossRiver(text))
		fmt.Println(event.TakeOut(text))
		fmt.Println(event.GetOut(text))
		if(text == "kylling"){
			break
		}
	}
}