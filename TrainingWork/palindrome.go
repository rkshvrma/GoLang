package main
import "fmt"
 
func main() {
    var num,rem,temp int
    var rev int = 0
 
    fmt.Print("Enter any positive integer : ")
    fmt.Scan(&num)
 
    temp=num
 
    for{
        rem = num%10
        rev = rev*10 + rem
        num /= 10
 
        if(num==0){
            break 
        }
    }
 
    if(temp==rev){
        fmt.Printf("%d is a Palindrome",temp)
    }else{
        fmt.Printf("%d is not a Palindrome",temp)
    }
 
}