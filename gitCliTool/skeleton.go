package main

import "flag"

// Scan function to scan path folders and git repositories
func scan(path string) {
	print("scan")
}

// Stats function to get stats of git repositories
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "Add folder to scan")
	flag.StringVar(&email, "email", "your@email.com", "email to scan")
	flag.Parse()
	// Check if the folder flag is set and call the scan function
	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
