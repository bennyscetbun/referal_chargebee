package chargebee

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/chargebee/chargebee-go/v3"
	couponAction "github.com/chargebee/chargebee-go/v3/actions/coupon"
	promotionalCreditAction "github.com/chargebee/chargebee-go/v3/actions/promotionalcredit"
	"github.com/chargebee/chargebee-go/v3/filter"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	couponEnum "github.com/chargebee/chargebee-go/v3/models/coupon/enum"
	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"
	"github.com/gin-gonic/gin"
)

func WebhookHandler(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	webhookData := &WebhookCallback{}
	if err := json.Unmarshal(jsonData, &webhookData); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	switch webhookData.EventType {
	case "subscription_created":
		subcriptionCreatedHandler(webhookData)
	}
	ctx.Status(http.StatusOK)
}

const REFERAL_COUPON_PREFIX = "REFERAL"

func subcriptionCreatedHandler(webhookData *WebhookCallback) error {
	if len(webhookData.Content.Subscription.Coupons) == 0 {
		return nil
	}
	for _, couponInfo := range webhookData.Content.Subscription.Coupons {
		if strings.HasPrefix(couponInfo.CouponId, REFERAL_COUPON_PREFIX) {
			return nil
		}
		referalCustomerID, err := extractCustomerFromReferalCoupon(couponInfo.CouponId)
		if err != nil {
			return err
		}
		if err := GiveCreditToCustomer(referalCustomerID); err != nil {
			return err
		}
	}
	if err := CreateReferalCoupon(webhookData.Content.Subscription.CustomerId); err != nil {
		return err
	}

	return nil
}

func makeCouponReferalForCustomer(customerID string) string {
	return REFERAL_COUPON_PREFIX + "_" + customerID + "_" + generateRandomString(3)
}

func extractCustomerFromReferalCoupon(couponID string) (string, error) {
	splitted := strings.Split(couponID, "_")
	if len(splitted) != 3 {
		return "", errors.New("bad coupon referal format")
	}
	return splitted[1], nil
}

func GiveCreditToCustomer(customerID string) error {
	_, err := promotionalCreditAction.Add(&promotionalcredit.AddRequestParams{
		CustomerId:  customerID,
		Amount:      chargebee.Int64(100),
		Description: "Credits de parainage",
	}).Request()
	if err != nil {
		return err
	}
	return nil
}

func hasAlreadyReferalCoupon(customerID string) (bool, error) {
	result, err := couponAction.List(&coupon.ListRequestParams{
		Id: &filter.StringFilter{
			StartsWith: REFERAL_COUPON_PREFIX + "_" + customerID,
		},
	}).ListRequest()
	if err != nil {
		return false, err
	}
	return len(result.List) > 0, nil
}

func CreateReferalCoupon(customerID string) error {
	if alreadyDone, err := hasAlreadyReferalCoupon(customerID); err != nil {
		return err
	} else if alreadyDone {
		return nil
	}

	_, err := couponAction.CreateForItems(&coupon.CreateForItemsRequestParams{
		Id:                 makeCouponReferalForCustomer(customerID),
		Name:               "Coupon Parainage",
		DiscountPercentage: chargebee.Float64(0.5),
		DiscountType:       couponEnum.DiscountTypePercentage,
		DurationType:       couponEnum.DurationTypeForever,
		ApplyOn:            couponEnum.ApplyOnEachSpecifiedItem,
		ItemConstraints: []*coupon.CreateForItemsItemConstraintParams{
			{
				Constraint: couponEnum.ItemConstraintConstraintAll,
				ItemType:   couponEnum.ItemConstraintItemTypePlan,
			},
		},
	}).Request()
	if err != nil {
		return err
	}
	return nil
}
