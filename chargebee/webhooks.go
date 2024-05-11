package chargebee

type WebhookCallback struct {
	ApiVersion string `json:"api_version"`
	Content    struct {
		Customer struct {
			AllowDirectDebit bool   `json:"allow_direct_debit"`
			AutoCollection   string `json:"auto_collection"`
			Balances         []struct {
				BalanceCurrencyCode string `json:"balance_currency_code"`
				CurrencyCode        string `json:"currency_code"`
				ExcessPayments      int64  `json:"excess_payments"`
				Object              string `json:"object"`
				PromotionalCredits  int64  `json:"promotional_credits"`
				RefundableCredits   int64  `json:"refundable_credits"`
				UnbilledCharges     int64  `json:"unbilled_charges"`
			} `json:"balances"`
			CardStatus            string `json:"card_status"`
			Channel               string `json:"channel"`
			CreatedAt             int64  `json:"created_at"`
			CreatedFromIp         string `json:"created_from_ip"`
			Deleted               bool   `json:"deleted"`
			ExcessPayments        int64  `json:"excess_payments"`
			FirstName             string `json:"first_name"`
			Id                    string `json:"id"`
			LastName              string `json:"last_name"`
			Mrr                   int64  `json:"mrr"`
			NetTermDays           int64  `json:"net_term_days"`
			Object                string `json:"object"`
			PiiCleared            string `json:"pii_cleared"`
			PreferredCurrencyCode string `json:"preferred_currency_code"`
			PromotionalCredits    int64  `json:"promotional_credits"`
			RefundableCredits     int64  `json:"refundable_credits"`
			ResourceVersion       int64  `json:"resource_version"`
			Taxability            string `json:"taxability"`
			UnbilledCharges       int64  `json:"unbilled_charges"`
			UpdatedAt             int64  `json:"updated_at"`
		} `json:"customer"`
		Subscription struct {
			ActivatedAt             int64  `json:"activated_at"`
			BillingPeriod           int64  `json:"billing_period"`
			BillingPeriodUnit       string `json:"billing_period_unit"`
			CancelScheduleCreatedAt int64  `json:"cancel_schedule_created_at"`
			CancelledAt             int64  `json:"cancelled_at"`
			Channel                 string `json:"channel"`
			Coupon                  string `json:"coupon"`
			Coupons                 []struct {
				AppliedCount int64  `json:"applied_count"`
				CouponId     string `json:"coupon_id"`
				Object       string `json:"object"`
			} `json:"coupons"`
			CreatedAt                   int64  `json:"created_at"`
			CreatedFromIp               string `json:"created_from_ip"`
			CurrencyCode                string `json:"currency_code"`
			CurrentTermEnd              int64  `json:"current_term_end"`
			CurrentTermStart            int64  `json:"current_term_start"`
			CustomerId                  string `json:"customer_id"`
			Deleted                     bool   `json:"deleted"`
			DueInvoicesCount            int64  `json:"due_invoices_count"`
			HasScheduledAdvanceInvoices bool   `json:"has_scheduled_advance_invoices"`
			HasScheduledChanges         bool   `json:"has_scheduled_changes"`
			Id                          string `json:"id"`
			Mrr                         int64  `json:"mrr"`
			Object                      string `json:"object"`
			PlanAmount                  int64  `json:"plan_amount"`
			PlanFreeQuantity            int64  `json:"plan_free_quantity"`
			PlanId                      string `json:"plan_id"`
			PlanQuantity                int64  `json:"plan_quantity"`
			PlanUnitPrice               int64  `json:"plan_unit_price"`
			RemainingBillingCycles      int64  `json:"remaining_billing_cycles"`
			ResourceVersion             int64  `json:"resource_version"`
			StartedAt                   int64  `json:"started_at"`
			Status                      string `json:"status"`
			UpdatedAt                   int64  `json:"updated_at"`
		} `json:"subscription"`
		UnbilledCharges []struct {
			Amount         int64  `json:"amount"`
			CurrencyCode   string `json:"currency_code"`
			CustomerId     string `json:"customer_id"`
			DateFrom       int64  `json:"date_from"`
			DateTo         int64  `json:"date_to"`
			Deleted        bool   `json:"deleted"`
			Description    string `json:"description"`
			DiscountAmount int64  `json:"discount_amount"`
			EntityId       string `json:"entity_id"`
			EntityType     string `json:"entity_type"`
			Id             string `json:"id"`
			IsVoided       bool   `json:"is_voided"`
			Object         string `json:"object"`
			PricingModel   string `json:"pricing_model"`
			Quantity       int64  `json:"quantity"`
			SubscriptionId string `json:"subscription_id"`
			UnitAmount     int64  `json:"unit_amount"`
			UpdatedAt      int64  `json:"updated_at"`
		} `json:"unbilled_charges"`
	} `json:"content"`
	EventType     string `json:"event_type"`
	Id            string `json:"id"`
	Object        string `json:"object"`
	OccurredAt    int64  `json:"occurred_at"`
	Source        string `json:"source"`
	User          string `json:"user"`
	WebhookStatus string `json:"webhook_status"`
	Webhooks      []struct {
		Id            string `json:"id"`
		Object        string `json:"object"`
		WebhookStatus string `json:"webhook_status"`
	} `json:"webhooks"`
}
