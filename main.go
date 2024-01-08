package main

import (
	"flag"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan1(path string) {
	print("scanned path", path)
}

// stats generates a nice graph of your Git contributions
func stats1(email string) {
	print("stats email ", email)
}

func main() {
	var folder string
	var email string

	/*
		build or run time flags from command line and pass the path with forward slash '/'
		to add folder go run main.go scan.go stats.go -add /folder/path
		go run main.go scan.go stats.go -email email
	*/
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()
	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
