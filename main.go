package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bennyscetbun/referal_chargebee/chargebee"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	// Check if there are enough arguments provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <config.json>")
		os.Exit(1)
	}
	if err := chargebee.Setup(os.Args[1]); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.POST("/chargebee", chargebee.WebhookHandler)
	tlsContext, tlsContextCancel := context.WithCancel(context.Background())
	defer tlsContextCancel()
	if err := autotls.RunWithContext(tlsContext, r, "lettresdalice.benny.ninja"); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen https:", err)
	}
}
