//Read 2 variables, named A and B and make the sum of these two variables,
//assigning its result to the variable X. Print X as shown below.
//Print endline after the result otherwise you will get “Presentation Error”.

package ativs
 
import "fmt"
    
 

func Somando() {

var A int
var B int
fmt.Println("")
fmt.Scanln(&A)
fmt.Scanln(&B)
fmt.Println("\nX =", A + B)
}

