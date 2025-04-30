# GoRandomPassword

* A programme for generating passwords using a cryptographically secure random number generator, using the Fisher-Yates algorithm to shuffle characters. Passwords contain lowercase and uppercase letters, numbers and special characters.
* The programme also removes ambiguous characters such as ‘l’, ‘1’, “O”, ‘0’ to avoid confusion when entering the password.
* Generate 12 passwords that are 12 characters long and contain lowercase and uppercase letters, numbers and special characters.
* The programme uses a cryptographically secure random number generator to generate passwords and the Fisher-Yates algorithm to mix the characters.
* The programme is written in Go language and uses standard libraries for generating random numbers and working with strings.


Run without building:  
go run main.go  
Build for running with double click:  
go mod init  
go build -o PasswordRandow.exe  
(PasswordRandow.exe - for Windows, for Linux or MacOS the extension (.exe) is not needed)  

In all cases, you must first install the Go language compiler. You can read how to do this on the official language page. [Link here](https://go.dev/doc/install "official language page")

p.s. If desired, you can add code so that the program accepts parameters from the console for the number and length of passwords.  
