package main

import (
	"flag"
	"fmt"
	"os"
)

// USAGE
// <program> <flag> <argument>
// ./2 -server https://127.0.0.1 -p 8081 -ssl

var banner = `
                                    
                                    
		lllllll   iiii  
		l:::::l  i::::i 
		l:::::l   iiii  
		l:::::l         
cccccccccccccccc l::::l iiiiiii 
cc:::::::::::::::c l::::l i:::::i 
c:::::::::::::::::c l::::l  i::::i 
c:::::::cccccc:::::c l::::l  i::::i 
c::::::c     ccccccc l::::l  i::::i 
c:::::c              l::::l  i::::i 
c:::::c              l::::l  i::::i 
c::::::c     ccccccc l::::l  i::::i 
c:::::::cccccc:::::cl::::::li::::::i
c:::::::::::::::::cl::::::li::::::i
cc:::::::::::::::cl::::::li::::::i
cccccccccccccccclllllllliiiiiiii
				
				
				
`

func main() {
	var (
		server string
		port   int
	)

	flag.StringVar(&server, "s", "http://127.0.0.1", "specify server address to use.  defaults to http://127.0.0.1.")
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000")
	ssl := flag.Bool("ssl", false, "specify ssl to use.  defaults to false")

	flag.Usage = func() {
		fmt.Println(banner)
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("./2 -s=https://127.0.0.1 -p=8080 -ssl true \n")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Printf("Starting Server \nListening %s:%d , SSL Mode: %v\n", server, port, *ssl)
}
