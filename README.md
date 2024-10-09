## ftpx

ftpx - A faster & simpler way to bruteforce FTP server

## Installation
```
go install github.com/rix4uni/ftpx@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/ftpx/releases/download/v0.0.2/ftpx-linux-amd64-0.0.2.tgz
tar -xvzf ftpx-linux-amd64-0.0.2.tgz
rm -rf ftpx-linux-amd64-0.0.2.tgz
mv ftpx ~/go/bin/ftpx
```
Or download [binary release](https://github.com/rix4uni/ftpx/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/ftpx.git
cd ftpx; go install
```

## Usage
```
Usage of ftpx:
  -ip string
        IP and port for FTP login
  -mode string
        Mode of operation: 'su' for single-username or 'upc' for username-password-combination
  -username string
        Username for FTP login (required for 'su' mode)
  -version
        Print the version of the tool and exit.
  -wordlist string
        File containing passwords or usernames & passwords
```

## Usage Examples
#### bruteforce single ip with single username and multiple passwords wordlist
```
# Command:
ftpx -mode su -ip 127.0.0.1:21 -username anonymous -wordlist ftp-password.txt

# Output
[+] Trying anonymous:12hrs37
[+] Trying anonymous:rootpasswdb1uRR3
[+] Trying anonymous:admin
[+] Trying anonymous:localadmin
Password Not Found with ip:127.0.0.1:21
```

#### bruteforce single ip with default username and password wordlist
```
# Command:
ftpx -mode upc -ip 127.0.0.1:21 -wordlist ftp-username-password.txt

# Output
[+] Trying anonymous:anonymous
[+] Trying root:rootpasswd
[+] Trying root:12hrs37
[+] Trying ftp:b1uRR3
Successfully logged in with ip:127.0.0.1:21 username:admin password:default
```