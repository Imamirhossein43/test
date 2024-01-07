package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	
)
type User struct {
    Username       string
    Email          string
    Password       string
    TrafficInGB    int // مقدار ترافیک بر حسب گیگابایت
    TimeInMonths   int     // مقدار زمان بر حسب ماه
}

type UserManager struct {
    Users []User
}


func (um *UserManager) Register(username, email, password string , traffic ,time int) {
    newUser := User{Username: username, Email: email, Password: password , TrafficInGB: traffic, TimeInMonths: time}
    um.Users = append(um.Users, newUser)
    fmt.Println("User registered successfully!")
}

func (um *UserManager) Login(username, password string) bool {
    for _, user := range um.Users {
        if user.Username == username && user.Password == password {
            fmt.Println("Login successful!")
            return true
        }
    }
    fmt.Println("Invalid username or password.")
    return false
}

func (um *UserManager) DisplayUserInfo(username string) {
    for _, user := range um.Users {
        if user.Username == username {
            fmt.Printf("Username: %s \n", user.Username)
            fmt.Printf("Email: %s \n", user.Email)
            fmt.Printf("traficc: %d \n", user.TrafficInGB)
            fmt.Printf("time: %d \n",user.TimeInMonths)
            return
        }
    }
    fmt.Println("User not found.")
}

func (um *UserManager) DeleteUser(username string) {
    for i, user := range um.Users {
        if user.Username == username {
            um.Users = append(um.Users[:i], um.Users[i+1:]...)
            fmt.Println("User deleted successfully!")
            return
        }
    }
    fmt.Println("User not found.")
}

func main() {
    userManager := UserManager{}
    for {
    var option string
    fmt.Println("  please choice :")
    fmt.Println("1: create")
    fmt.Println("2: display")
    fmt.Println("3: delete")
    fmt.Println("4: exit")
    fmt.Scanln(&option)

    switch option {
    case "1":
        var username, email, password string
        var traffic int
        var time int
        fmt.Println("Enter username:")
        fmt.Scanln(&username)
        fmt.Println("Enter email:")
        fmt.Scanln(&email)
        fmt.Println("Enter password:")
        fmt.Scanln(&password)
        fmt.Println("Enter traffic in gigabytes:")
        fmt.Scanln(&traffic)
        fmt.Println("Enter time in months:")
        fmt.Scanln(&time)
        userManager.Register(username, email, password, traffic, time)
    case "2":
        var username string
        fmt.Println("Enter username:")
        fmt.Scanln(&username)
        userManager.DisplayUserInfo(username)
    case "3":
        var username string
        fmt.Println("Enter username:")
        fmt.Scanln(&username)
        userManager.DeleteUser(username)
   
        case "4":
            fmt.Println("exiting...")
            return 
    default:
        fmt.Println("Invalid option.")
    }


url := "http://87.248.150.177:9011"
data := userManager
jsonData,err := json.Marshal(data)
if err != nil{
    fmt.Println("خطا در تبدیل درخواست به json",err)
    return
}
response,err:= http.Post(url ,"application/json",bytes.NewBuffer(jsonData))
if err != nil{
    fmt.Println("خطا در ارسال درخواست",err)
    return
}
defer response.Body.Close()
fmt.Println("وضعیت درخواست :", response.Status)

    }
}