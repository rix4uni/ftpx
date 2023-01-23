# FTPBruteForce


# Installation

```
https://github.com/rix4uni/FTPBruteForce.git
cd FTPBruteForce
```

# Usage
ftp login bruteforce for one username with multiple passwords
```
options:
  -ip string
        IP and port for FTP login (default "127.0.0.1:21")
  -p string
        file containing passwords to try (default "ftp-password.txt")
  -u string
        username for FTP login (default "anonymous")
examples:
  go run ftp-brute-force.go -u anonymous -ip 127.0.0.1:21 -p ftp-password.txt
```

# Usage
ftp login bruteforce for default credentails
```
options:
  -ip string
        IP and port for FTP login (default "127.0.0.1:21")
  -up string
        File containing usernames & passwords (default "ftp-username-password.txt")
examples:
  go run ftp-brute-force-default-credentails.go -ip 127.0.0.1:21 -up ftp-username-password.txt
```
