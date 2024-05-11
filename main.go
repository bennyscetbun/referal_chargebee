package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bennyscetbun/referal_chargebee/chargebee"
	chargebeelib "github.com/chargebee/chargebee-go/v3"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func readArgumentsFromFile() []string {
	// Check if there are enough arguments provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: program_name <file_path>")
		os.Exit(1)
	}

	// Get the file path from command line arguments
	filePath := os.Args[1]

	// Read the entire file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	splitted := strings.Split((string)(data), "\n")
	for i := range splitted {
		splitted[i] = strings.TrimSpace(splitted[i])
	}
	return splitted
}

func main() {
	myargs := readArgumentsFromFile()
	chargebeelib.Configure(myargs[0], myargs[1])
	r := gin.Default()
	r.POST("/chargebee", chargebee.WebhookHandler)
	tlsContext, tlsContextCancel := context.WithCancel(context.Background())
	defer tlsContextCancel()
	if err := autotls.RunWithContext(tlsContext, r, "lettresdalice.benny.ninja"); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen https:", err)
	}
}
