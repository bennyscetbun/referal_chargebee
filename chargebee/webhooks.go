package chargebee

type WebhookCallback struct {
	ApiVersion string `json:"api_version"`
	Content    struct {
		Invoice struct {
			//	AdjustmentCreditNotes                   Is of Type: TypeUndefined `json:"adjustment_credit_notes"`
			AmountAdjusted  float64 `json:"amount_adjusted"`
			AmountDue       float64 `json:"amount_due"`
			AmountPaid      float64 `json:"amount_paid"`
			AmountToCollect float64 `json:"amount_to_collect"`
			// AppliedCredits                  Is of Type: TypeUndefined `json:"applied_credits"`
			BaseCurrencyCode string `json:"base_currency_code"`
			BillingAddress   struct {
				FirstName        string `json:"first_name"`
				LastName         string `json:"last_name"`
				Object           string `json:"object"`
				ValidationStatus string `json:"validation_status"`
			} `json:"billing_address"`
			Channel        string  `json:"channel"`
			CreditsApplied float64 `json:"credits_applied"`
			CurrencyCode   string  `json:"currency_code"`
			CustomerId     string  `json:"customer_id"`
			Date           float64 `json:"date"`
			Deleted        bool    `json:"deleted"`
			Discounts      []struct {
				Amount             float64 `json:"amount"`
				Description        string  `json:"description"`
				DiscountPercentage float64 `json:"discount_percentage"`
				DiscountType       string  `json:"discount_type"`
				EntityId           string  `json:"entity_id"`
				EntityType         string  `json:"entity_type"`
				Object             string  `json:"object"`
			} `json:"discounts"`
			DueDate float64 `json:"due_date"`
			//	DunningAttempts                 Is of Type: TypeUndefined `json:"dunning_attempts"`
			ExchangeRate      float64 `json:"exchange_rate"`
			FirstInvoice      bool    `json:"first_invoice"`
			GeneratedAt       float64 `json:"generated_at"`
			HasAdvanceCharges bool    `json:"has_advance_charges"`
			Id                string  `json:"id"`
			IsGifted          bool    `json:"is_gifted"`
			//	IssuedCreditNotes                       Is of Type: TypeUndefined `json:"issued_credit_notes"`
			LineItemDiscounts []struct {
				CouponId       string  `json:"coupon_id"`
				DiscountAmount float64 `json:"discount_amount"`
				DiscountType   string  `json:"discount_type"`
				EntityId       string  `json:"entity_id"`
				LineItemId     string  `json:"line_item_id"`
				Object         string  `json:"object"`
			} `json:"line_item_discounts"`
			LineItems []struct {
				Amount                  float64 `json:"amount"`
				CustomerId              string  `json:"customer_id"`
				DateFrom                float64 `json:"date_from"`
				DateTo                  float64 `json:"date_to"`
				Description             string  `json:"description"`
				DiscountAmount          float64 `json:"discount_amount"`
				EntityId                string  `json:"entity_id"`
				EntityType              string  `json:"entity_type"`
				Id                      string  `json:"id"`
				IsTaxed                 bool    `json:"is_taxed"`
				ItemLevelDiscountAmount float64 `json:"item_level_discount_amount"`
				Object                  string  `json:"object"`
				PricingModel            string  `json:"pricing_model"`
				Quantity                float64 `json:"quantity"`
				SubscriptionId          string  `json:"subscription_id"`
				TaxAmount               float64 `json:"tax_amount"`
				TaxExemptReason         string  `json:"tax_exempt_reason"`
				UnitAmount              float64 `json:"unit_amount"`
			} `json:"line_items"`
			// LinkedOrders                    Is of Type: TypeUndefined `json:"linked_orders"`
			LinkedPayments []struct {
				AppliedAmount float64 `json:"applied_amount"`
				AppliedAt     float64 `json:"applied_at"`
				TxnAmount     float64 `json:"txn_amount"`
				TxnDate       float64 `json:"txn_date"`
				TxnId         string  `json:"txn_id"`
				TxnStatus     string  `json:"txn_status"`
			} `json:"linked_payments"`
			NetTermDays           float64 `json:"net_term_days"`
			NewSalesAmount        float64 `json:"new_sales_amount"`
			Object                string  `json:"object"`
			PaidAt                float64 `json:"paid_at"`
			PriceType             string  `json:"price_type"`
			Recurring             bool    `json:"recurring"`
			ResourceVersion       float64 `json:"resource_version"`
			RoundOffAmount        float64 `json:"round_off_amount"`
			SiteDetailsAtCreation struct {
				OrganizationAddress struct {
					City             string `json:"city"`
					CountryCode      string `json:"country_code"`
					Email            string `json:"email"`
					Line1            string `json:"line1"`
					OrganizationName string `json:"organization_name"`
					Phone            string `json:"phone"`
					State            string `json:"state"`
					Zip              string `json:"zip"`
				} `json:"organization_address"`
				Timezone string `json:"timezone"`
			} `json:"site_details_at_creation"`
			Status         string  `json:"status"`
			SubTotal       float64 `json:"sub_total"`
			SubscriptionId string  `json:"subscription_id"`
			Tax            float64 `json:"tax"`
			TaxOrigin      struct {
				Country string `json:"country"`
			} `json:"tax_origin"`
			TermFinalized  bool    `json:"term_finalized"`
			Total          float64 `json:"total"`
			UpdatedAt      float64 `json:"updated_at"`
			WriteOffAmount float64 `json:"write_off_amount"`
		} `json:"invoice"`
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
