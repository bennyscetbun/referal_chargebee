package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/bennyscetbun/referal_chargebee/chargebee"
	chargebeelib "github.com/chargebee/chargebee-go/v3"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	chargebeelib.Configure(os.Args[2], os.Args[1])
	r := gin.Default()
	r.POST("/chargebee", chargebee.WebhookHandler)
	tlsContext, tlsContextCancel := context.WithCancel(context.Background())
	defer tlsContextCancel()
	if err := autotls.RunWithContext(tlsContext, r, "lettresdalice.benny.ninja"); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen https:", err)
	}
}
