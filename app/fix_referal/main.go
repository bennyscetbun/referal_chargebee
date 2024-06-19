package main

import (
	"fmt"
	"os"

	"github.com/bennyscetbun/referal_chargebee/chargebee"
	"github.com/chargebee/chargebee-go/v3/models/customer"
	subscriptionEnum "github.com/chargebee/chargebee-go/v3/models/subscription/enum"
	"github.com/ztrue/tracerr"
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
	subscriptionsInfo, err := chargebee.GetAllSubscriptionsInfo()
	if err != nil {
		tracerr.Print(err)
		return
	}
	customersByStatus := make(map[subscriptionEnum.Status]map[string]*customer.Customer)

	for _, subscriptionInfo := range subscriptionsInfo {
		customerByStatus := customersByStatus[subscriptionInfo.Subscription.Status]
		if customerByStatus == nil {
			customerByStatus = make(map[string]*customer.Customer)
			customersByStatus[subscriptionInfo.Subscription.Status] = customerByStatus
		}
		customerByStatus[subscriptionInfo.Customer.Id] = subscriptionInfo.Customer
	}

	mergedCustomers := make(map[string]*customer.Customer)
	for customerId, customer := range customersByStatus[subscriptionEnum.StatusActive] {
		mergedCustomers[customerId] = customer
	}
	for customerId, customer := range customersByStatus[subscriptionEnum.StatusPaused] {
		mergedCustomers[customerId] = customer
	}
	for customerId, customer := range customersByStatus[subscriptionEnum.StatusNonRenewing] {
		mergedCustomers[customerId] = customer
	}
	for customerId, customer := range customersByStatus[subscriptionEnum.StatusFuture] {
		mergedCustomers[customerId] = customer
	}
	SendReferals(mergedCustomers)
}

func SendReferals(customers map[string]*customer.Customer) {
	for customerId := range customers {
		if err := chargebee.CreateReferalCoupon(customerId); err != nil {
			tracerr.Print(err)
		}
		// chargebee
	}
}
