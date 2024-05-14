package chargebee

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/chargebee/chargebee-go/v3"
	couponAction "github.com/chargebee/chargebee-go/v3/actions/coupon"
	customerAction "github.com/chargebee/chargebee-go/v3/actions/customer"
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	couponEnum "github.com/chargebee/chargebee-go/v3/models/coupon/enum"
	"github.com/chargebee/chargebee-go/v3/models/customer"
	"github.com/gin-gonic/gin"
	"github.com/ztrue/tracerr"
)

func WebhookHandler(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	webhookData := &WebhookCallback{}
	if err := json.Unmarshal(jsonData, &webhookData); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	switch webhookData.EventType {
	case "subscription_created":
		if err := subcriptionCreatedHandler(webhookData); err != nil {
			tracerr.Print(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	ctx.Status(http.StatusOK)
}

const REFERAL_COUPON_PREFIX = "REFERAL"

func subcriptionCreatedHandler(webhookData *WebhookCallback) error {
	for _, couponInfo := range webhookData.Content.Subscription.Coupons {
		if !strings.HasPrefix(couponInfo.CouponId, REFERAL_COUPON_PREFIX) {
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
		return "", tracerr.Errorf("bad coupon referal format")
	}
	return splitted[1], nil
}

func GiveCreditToCustomer(customerID string) error {
	_, err := customerAction.AddPromotionalCredits(
		customerID, &customer.AddPromotionalCreditsRequestParams{
			Amount:       chargebee.Int64(100),
			Description:  "Credits de parainage",
			CreditType:   enum.CreditTypeReferralRewards,
			CurrencyCode: "EUR",
		}).Request()
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

func retrieveAllCoupon() ([]*coupon.Coupon, error) {
	ret := []*coupon.Coupon{}
	var offset string
	for {
		result, err := couponAction.List(&coupon.ListRequestParams{
			Limit:  chargebee.Int32(1),
			Offset: offset,
		}).ListRequest()
		if err != nil {
			return nil, tracerr.Wrap(err)
		}
		for idx := 0; idx < len(result.List); idx++ {
			ret = append(ret, result.List[idx].Coupon)
		}
		if result.NextOffset == "" {
			return ret, nil
		}
		offset = result.NextOffset
	}
}

func HasAlreadyReferalCoupon(customerID string) (bool, error) {
	var offset string
	for {
		result, err := couponAction.List(&coupon.ListRequestParams{
			Limit:  chargebee.Int32(100),
			Offset: offset,
		}).ListRequest()
		if err != nil {
			return false, tracerr.Wrap(err)
		}
		for idx := 0; idx < len(result.List); idx++ {
			if strings.HasPrefix(result.List[idx].Coupon.Id, REFERAL_COUPON_PREFIX+"_"+customerID+"_") {
				return true, nil
			}

		}
		if result.NextOffset == "" {
			return false, nil
		}
		offset = result.NextOffset
	}
}

func CreateReferalCoupon(customerID string) error {
	if alreadyDone, err := HasAlreadyReferalCoupon(customerID); err != nil {
		return err
	} else if alreadyDone {
		return nil
	}

	_, err := couponAction.Create(&coupon.CreateRequestParams{
		Id:                 makeCouponReferalForCustomer(customerID),
		Name:               "Coupon Parainage " + customerID,
		DiscountPercentage: chargebee.Float64(0.5),
		DiscountType:       couponEnum.DiscountTypePercentage,
		DurationType:       couponEnum.DurationTypeForever,
		ApplyOn:            couponEnum.ApplyOnEachSpecifiedItem,
		PlanConstraint:     couponEnum.PlanConstraintAll,
		AddonConstraint:    couponEnum.AddonConstraintAll,
	}).Request()
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}
