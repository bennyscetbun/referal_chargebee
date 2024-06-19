package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bennyscetbun/referal_chargebee/chargebee"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
	subscriptionEnum "github.com/chargebee/chargebee-go/v3/models/subscription/enum"
	"github.com/ztrue/tracerr"
)

func getNextMonth() time.Time {
	now := time.Now().UTC()
	y, m, _ := now.Date()
	forth := time.Date(y, m, 4, 0, 0, 0, 0, time.UTC)
	if forth.Before(now) {
		forth = forth.AddDate(0, 1, 0)
	}
	return forth
}

func main() {
	// Check if there are enough arguments provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <config.json>")
		os.Exit(1)
	}
	if err := chargebee.Setup(os.Args[1]); err != nil {
		tracerr.Print(err)
		return
	}
	subscriptionsInfo, err := chargebee.GetAllSubscriptionsInfo()
	if err != nil {
		tracerr.Print(err)
		return
	}
	byStatus := make(map[subscriptionEnum.Status][]*subscription.Subscription)

	for _, subscriptionInfo := range subscriptionsInfo {
		byStatus[subscriptionInfo.Subscription.Status] = append(byStatus[subscriptionInfo.Subscription.Status], subscriptionInfo.Subscription)
	}
	nextMonth := getNextMonth()
	nextNextMonth := nextMonth.AddDate(0, 1, 0)
	renewedNextMonth := 0
	renewedNextNextMonth := 0

	nextMoney := float64(0)
	nextNextMoney := float64(0)
	for _, sub := range byStatus[subscriptionEnum.StatusActive] {
		currentTermEnd := time.Unix(sub.CurrentTermEnd, 0)
		if currentTermEnd.Before(nextMonth) && currentTermEnd.Add(time.Hour*24).After(nextMonth) {
			renewedNextMonth += 1
			nextMoney += float64(sub.PlanUnitPrice) * float64(sub.PlanQuantity)
			if sub.BillingPeriodUnit == "month" && sub.BillingPeriod == 1 {
				renewedNextNextMonth += 1
				nextNextMoney += float64(sub.PlanUnitPrice) * float64(sub.PlanQuantity)
			}
		} else if currentTermEnd.Before(nextNextMonth) && currentTermEnd.Add(time.Hour*24).After(nextNextMonth) {
			renewedNextNextMonth += 1
			nextNextMoney += float64(sub.PlanUnitPrice) * float64(sub.PlanQuantity)
		}
	}
	fmt.Println("The", nextMonth.Format(time.DateOnly), "there will be", renewedNextMonth, "for a value of", nextMoney/100)
	fmt.Println("The", nextNextMonth.Format(time.DateOnly), "there will be", renewedNextNextMonth, "for a value of", nextNextMoney/100)
}
