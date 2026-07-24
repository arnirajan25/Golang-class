package main
import "fmt"
func rectangle(l, b float64) (float64, float64) {
area:=l*b
perimeter:=2*(l+b)
return area,perimeter
}

func main(){
l:=3.50
b:=6.50
area,perimeter:= rectangle(l,b)//(_)underscore can be used if any one return value is not needed but the program needs to be executed[area,_:=]
fmt.Println("Area of square is = ",area)
fmt.Println("Perimeter of rectangle = ",perimeter)//("Perimeter of rectangle =")using underscore if perimeter is not needed
}
