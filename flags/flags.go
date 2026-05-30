package main

import "flag"

func main() {
	host := flag.String("host", "localhost", "Database host to connect to") // Define a string flag with a default value and an instruction
	username := flag.String("username", "root", "Database username")        // Define a string flag for the username
	password := flag.String("password", "root", "Password")                 // Define a string flag for the password
	//localhost and root are default values for host, username and password respectively.
	flag.Parse()                    // Parse the command-line flags
	println("Host:", *host)         // Print the value of the host flag
	println("Username:", *username) // Print the value of the username flag
	println("Password:", *password) // Print the value of the password flag
}

var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname") // Define an integer flag with a default value and a help message
}
