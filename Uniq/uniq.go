package main

import (
 "bufio"
 "os"
 "fmt"
)

func main() {
 in := bufio.NewScanner(os.Stdin)
 var prev string

 for in.Scan()  {
  txt := in.Text()

  if txt == prev {
   continue
  }

  fmt.Println(txt)
 }
}
