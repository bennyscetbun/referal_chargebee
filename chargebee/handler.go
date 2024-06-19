package chargebee

import (
	"encoding/json"
	"io"
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
		tracerr.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	webhookData := &WebhookCallback{}
	if err := json.Unmarshal(jsonData, &webhookData); err != nil {
		tracerr.Print(err)
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

const REFERRAL_COUPON_PREFIX = "REF"

func subcriptionCreatedHandler(webhookData *WebhookCallback) error {
	for _, couponInfo := range webhookData.Content.Subscription.Coupons {
		if !strings.HasPrefix(couponInfo.CouponId, REFERRAL_COUPON_PREFIX) {
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
	/*
		sum := 0
		for _, c := range customerID {
			sum += int(c)
		}
		sum *= 100 // so we are sure we ve got at least 3 digits
	*/
	return REFERRAL_COUPON_PREFIX + "_" + customerID
}

func extractCustomerFromReferalCoupon(couponID string) (string, error) {
	splitted := strings.Split(couponID, "_")
	if len(splitted) != 2 {
		return "", tracerr.Errorf("bad coupon referral format")
	}
	return splitted[1], nil
}

func GiveCreditToCustomer(customerID string) error {
	customerInfo, err := GetCustomer(customerID)
	if err != nil {
		return err
	}

	if _, err := customerAction.AddPromotionalCredits(
		customerID, &customer.AddPromotionalCreditsRequestParams{
			Amount:       &cfg.CreditOffertEnCentime,
			Description:  "Credits de parainage",
			CreditType:   enum.CreditTypeReferralRewards,
			CurrencyCode: "EUR",
		}).Request(); err != nil {
		return tracerr.Wrap(err)
	}
	if err := sendCreditAddedEmail(customerInfo.Email, customerInfo.PromotionalCredits+cfg.CreditOffertEnCentime); err != nil {
		return err
	}
	return nil
}

func HasAlreadyReferalCoupon(couponID string) (bool, error) {
	_, err := couponAction.Retrieve(couponID).Request()
	if err != nil {
		switch v := err.(type) {
		case *chargebee.Error:
			if v.HTTPStatusCode == 404 {
				return false, nil
			}
		}
		return false, tracerr.Wrap(err)
	}
	return true, nil
}

func CreateReferalCoupon(customerID string) error {
	couponID := makeCouponReferalForCustomer(customerID)
	if alreadyDone, err := HasAlreadyReferalCoupon(couponID); err != nil {
		return err
	} else if alreadyDone {
		return nil
	}
	customerInfo, err := GetCustomer(customerID)
	if err != nil {
		return err
	}

	if _, err := couponAction.Create(&coupon.CreateRequestParams{
		Id:                 couponID,
		Name:               "Coupon Parainage " + couponID,
		DiscountPercentage: &cfg.ReductionEnPourcent,
		DiscountType:       couponEnum.DiscountTypePercentage,
		DurationType:       couponEnum.DurationTypeOneTime,
		ApplyOn:            couponEnum.ApplyOnEachSpecifiedItem,
		PlanConstraint:     couponEnum.PlanConstraintAll,
		AddonConstraint:    couponEnum.AddonConstraintAll,
	}).Request(); err != nil {
		return tracerr.Wrap(err)
	}
	if !strings.Contains(customerInfo.Email, "@") {
		return nil
	}
	if err := sendReferalEmail(customerInfo.Email, couponID); err != nil {
		return err
	}
	return nil
}

func GetCustomer(customerID string) (*customer.Customer, error) {
	resp, err := customerAction.Retrieve(customerID).Request()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	return resp.Customer, nil
}
