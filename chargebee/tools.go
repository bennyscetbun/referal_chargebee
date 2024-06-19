package chargebee

import (
	"fmt"
	"math/rand"

	"github.com/chargebee/chargebee-go/v3"
	subscriptionAction "github.com/chargebee/chargebee-go/v3/actions/subscription"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
	"github.com/ztrue/tracerr"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(size int) string {
	var s string
	for i := 0; i < size; i++ {
		s += string(letters[rand.Intn(len(letters))])
	}
	return s
}

func GetAllSubscriptionsInfo() ([]*chargebee.Result, error) {
	var offset string
	nb := 0
	ret := make([]*chargebee.Result, 0, 5000)
	for {
		res, err := subscriptionAction.List(&subscription.ListRequestParams{
			Limit:  chargebee.Int32(50),
			Offset: offset,
		}).ListRequest()
		if err != nil {
			return nil, tracerr.Wrap(err)
		} else {
			ret = append(ret, res.List...)
			nb += len(res.List)
			fmt.Println(nb)
		}
		if len(res.NextOffset) != 0 {
			offset = res.NextOffset
		} else {
			break
		}
	}
	return ret, nil
}
