package main
import "fmt"
 
func main() {
    var num,temp,rem int
    var res int =0
    fmt.Print("Enter any three digit num : ")
    fmt.Scan(&num)
 
    temp = num
 
    for {
        rem = temp%10
        res += rem*rem*rem     
        temp /=10
         
        if(temp==0){
            break 
        }
    }
 
    if(res==num){
         fmt.Printf("%d is an Armstrong num.",num)
    }else{
        fmt.Printf("%d is not an Armstrong num.",num)
    }
}