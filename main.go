package main

import(
  "fmt"
  "./token"
)
func main(){
    fmt.Println(token.Wsp_Lexical("./Wsp/Test.wsp"))
    fmt.Println(len(token.Wsp_Lexical("./Wsp/Test.wsp")))
}