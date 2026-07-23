package main

import "fmt"
//func hawa(a *int) {
//if a==nil {
//fmt.Println("The pointer is empty.")
//} else {
//fmt.Println(a)
//}}

func cs(b *yippe, newval string){
b.name=newval
}

func Namaskar(name *string) {
*name ="Hello user " + *name
}


func togglebool(a *bool) {
*a=!*a
}

func pbv ( b yippe , newval string)  yippe{
b.name=newval
return  b
}


func reset_to_zero(a *int) {
*a = 0
}


func swap(a, b *int) {
       temp:= *a
       *a = *b
       *b = temp
}

type yippe struct{
name string
age int
}


func cv(n *yippe, newval string){//cv=change value
(*n).name= newval// (*)Is Not Required
}

func main() {
//x:= 10
//y:= 91
//fmt.Println("x = ",x)
//fmt.Println("y = ",y)
//swap(&x,&y)
//fmt.Println("After swap:")
//fmt.Println("x = ",x)
//fmt.Println("y = ",y)
//z:= 1
//fmt.Println("z = ",z)
//reset_to_zero(&z)
//fmt.Println("z = ",z)
//d:= true
//fmt.Println("Boolean =",d)
//togglebool(&d)
//fmt.Println("Boolean =",d)
//name:="Nirajan"
//Namaskar(&name)
//fmt.Println(name)
//var g *int
//hawa(g)
//h:=8
//hawa(&h)
//fmt.Println(g)
//fmt.Println(h)

s1:=yippe{
name:"Nirajan",
age:20,
}
var s2 yippe
s2.name="Suhag"
s2.age=19

fmt.Println(s1.name)
fmt.Println(s1.age)
fmt.Println(s2.name)
fmt.Println(s2.age)
fmt.Println("Value after Change")
cv(&s1,"Ayush")
fmt.Println(s1.name)
cs(&s2,"Mahim")
fmt.Println(s2.name)
ayush :=pbv(s2 , "Rohan")
fmt.Println(ayush.name)
}
