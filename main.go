package main

import "fmt"

func hawa(a *int) {
if a==nil {
fmt.Println("The pointer is empty.")
} else {
fmt.Println(a)
}}

func Namaskar(name *string) {
*name ="Hello user " + *name
}
func togglebool(a *bool) {
*a=!*a
}
func reset_to_zero(a *int) {
*a = 0
}
func swap(a, b *int) {
        temp:= *a
        *a = *b
        *b = temp
}
func main() {
x:= 10
y:= 91
fmt.Println("x = ",x)
fmt.Println("y = ",y)
swap(&x,&y)
fmt.Println("After swap:")
fmt.Println("x = ",x)
fmt.Println("y = ",y)
z:= 1
fmt.Println("z = ",z)
reset_to_zero(&z)
fmt.Println("z = ",z)
d:= true
fmt.Println("Boolean =",d)
togglebool(&d)
fmt.Println("Boolean =",d)
name:="Nirajan"
Namaskar(&name)
fmt.Println(name)
var g *int
hawa(g)
h:=8
hawa(&h)
fmt.Println(g)
fmt.Println(h)
}
