package freee

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// User is user.
type User struct {
	ID            int           `json:"id"`
	Email         string        `json:"email"`
	DisplayName   string        `json:"display_name"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	FirstNameKana string        `json:"first_name_kana"`
	LastNameKana  string        `json:"last_name_kana"`
	Companies     []UserCompany `json:"companies"`
}

// UserCompany is user company.
type UserCompany struct {
	ID            int    `json:"id"`
	DisplayName   string `json:"display_name"`
	Role          string `json:"role"`
	UseCustomRole bool   `json:"use_custom_role"`
}

// User returns login user.
func (c *Client) User(ctx context.Context) (*User, error) {
	q := url.Values{}
	q.Set("companies", "true")

	var u struct {
		User `json:"user"`
	}

	if err := c.do(ctx, http.MethodGet, "users/me", q, &u); err != nil {
		return nil, err
	}

	return &u.User, nil
}

// UserCapability is user capability.
type UserCapability struct {
	WalletTxns                   Capability `json:"wallet_txns"`
	Deals                        Capability `json:"deals"`
	Transfers                    Capability `json:"transfers"`
	Docs                         Capability `json:"docs"`
	DocPostings                  Capability `json:"doc_postings"`
	Receipts                     Capability `json:"receipts"`
	ReceiptStreamEditor          Capability `json:"receipt_stream_editor"`
	ExpenseApplications          Capability `json:"expense_applications"`
	Spreadsheets                 Capability `json:"spreadsheets"`
	PaymentRequests              Capability `json:"payment_requests"`
	RequestForms                 Capability `json:"request_forms"`
	ApprovalRequests             Capability `json:"approval_requests"`
	Reports                      Capability `json:"reports"`
	ReportsIncomeExpense         Capability `json:"reports_income_expense"`
	ReportsReceivables           Capability `json:"reports_receivables"`
	ReportsPayables              Capability `json:"reports_payables"`
	ReportsCash_balance          Capability `json:"reports_cash_balance"`
	ReportsCrosstabs             Capability `json:"reports_crosstabs"`
	ReportsGeneralLedgers        Capability `json:"reports_general_ledgers"`
	ReportsPl                    Capability `json:"reports_pl"`
	ReportsBs                    Capability `json:"reports_bs"`
	ReportsJournals              Capability `json:"reports_journals"`
	ReportsManagementsPlanning   Capability `json:"reports_managements_planning"`
	ReportsManagementsNavigation Capability `json:"reports_managements_navigation"`
	ManualJournals               Capability `json:"manual_journals"`
	FixedAssets                  Capability `json:"fixed_assets"`
	InventoryRefreshes           Capability `json:"inventory_refreshes"`
	BizAllocations               Capability `json:"biz_allocations"`
	PaymentRecords               Capability `json:"payment_records"`
	AnnualReports                Capability `json:"annual_reports"`
	TaxReports                   Capability `json:"tax_reports"`
	ConsumptionEntries           Capability `json:"consumption_entries"`
	TaxReturn                    Capability `json:"tax_return"`
	AccountItemStatements        Capability `json:"account_item_statements"`
	MonthEnd                     Capability `json:"month_end"`
	YearEnd                      Capability `json:"year_end"`
	Walletables                  Capability `json:"walletables"`
	Companies                    Capability `json:"companies"`
	Invitations                  Capability `json:"invitations"`
	SignInLogs                   Capability `json:"sign_in_logs"`
	Backups                      Capability `json:"backups"`
	OpeningBalances              Capability `json:"opening_balances"`
	SystemConversion             Capability `json:"system_conversion"`
	Resets                       Capability `json:"resets"`
	Partners                     Capability `json:"partners"`
	Items                        Capability `json:"items"`
	Sections                     Capability `json:"sections"`
	Tags                         Capability `json:"tags"`
	AccountItems                 Capability `json:"account_items"`
	Taxes                        Capability `json:"taxes"`
	PayrollItemSets              Capability `json:"payroll_item_sets"`
	UserMatchers                 Capability `json:"user_matchers"`
	DealTemplates                Capability `json:"deal_templates"`
	ManualJournalTemplates       Capability `json:"manual_journal_templates"`
	CostAllocations              Capability `json:"cost_allocations"`
	ApprovalFlowRoutes           Capability `json:"approval_flow_routes"`
	ExpenseApplicationTemplates  Capability `json:"expense_application_templates"`
	Workflows                    Capability `json:"workflows"`
	OauthApplications            Capability `json:"oauth_applications"`
	OauthAuthorizations          Capability `json:"oauth_authorizations"`
	DivisionTag1                 Capability `json:"division_tag_1"`
	DivisionTag2                 Capability `json:"division_tag_2"`
	DivisionTag3                 Capability `json:"division_tag_3"`
	BankAccountantStaffUsers     Capability `json:"bank_accountant_staff_users"`
	WorkflowValidationSettings   Capability `json:"workflow_validation_settings"`
}

// Capability is capability.
type Capability struct {
	Confirm       bool   `json:"confirm"`
	Read          bool   `json:"read"`
	Create        bool   `json:"create"`
	Update        bool   `json:"update"`
	Destroy       bool   `json:"destroy"`
	Sync          bool   `json:"sync"`
	AllowedTarget string `json:"allowed_target"`
}

// UserCapability returns login user capability.
func (c *Client) UserCapability(ctx context.Context, companyID int) (*UserCapability, error) {
	q := url.Values{}
	q.Set("company_id", strconv.Itoa(companyID))

	var uc UserCapability
	if err := c.do(ctx, http.MethodGet, "users/capabilities", q, &uc); err != nil {
		return nil, err
	}

	return &uc, nil
}
