// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package tem provides methods and message types of the tem v1alpha1 API.
package tem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API: transactional Email API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type DomainStatus string

const (
	// Default unknown domain status
	DomainStatusUnknown = DomainStatus("unknown")
	// Domain is checked
	DomainStatusChecked = DomainStatus("checked")
	// Domain is unchecked
	DomainStatusUnchecked = DomainStatus("unchecked")
	// Domain is invalid
	DomainStatusInvalid = DomainStatus("invalid")
	// Domain is locked
	DomainStatusLocked = DomainStatus("locked")
	// Domain is revoked
	DomainStatusRevoked = DomainStatus("revoked")
	// Domain is pending for check
	DomainStatusPending = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type EmailRcptType string

const (
	// Default unknown recipient type
	EmailRcptTypeUnknownRcptType = EmailRcptType("unknown_rcpt_type")
	// Primary recipient
	EmailRcptTypeTo = EmailRcptType("to")
	// Carbon copy recipient
	EmailRcptTypeCc = EmailRcptType("cc")
	// Blind carbon copy recipient
	EmailRcptTypeBcc = EmailRcptType("bcc")
)

func (enum EmailRcptType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_rcpt_type"
	}
	return string(enum)
}

func (enum EmailRcptType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailRcptType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailRcptType(EmailRcptType(tmp).String())
	return nil
}

type EmailStatus string

const (
	// Default unknown email status
	EmailStatusUnknown = EmailStatus("unknown")
	// Email is new
	EmailStatusNew = EmailStatus("new")
	// Email is in the process of being sent
	EmailStatusSending = EmailStatus("sending")
	// Email is sent
	EmailStatusSent = EmailStatus("sent")
	// Email sending has failed
	EmailStatusFailed = EmailStatus("failed")
	// Email is canceled
	EmailStatusCanceled = EmailStatus("canceled")
)

func (enum EmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum EmailStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailStatus(EmailStatus(tmp).String())
	return nil
}

type ListEmailsRequestOrderBy string

const (
	// Order by creation date (descending chronological order)
	ListEmailsRequestOrderByCreatedAtDesc = ListEmailsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order)
	ListEmailsRequestOrderByCreatedAtAsc = ListEmailsRequestOrderBy("created_at_asc")
	// Order by last update date (descending chronological order)
	ListEmailsRequestOrderByUpdatedAtDesc = ListEmailsRequestOrderBy("updated_at_desc")
	// Order by last update date (ascending chronological order)
	ListEmailsRequestOrderByUpdatedAtAsc = ListEmailsRequestOrderBy("updated_at_asc")
	// Order by status (descending alphabetical order)
	ListEmailsRequestOrderByStatusDesc = ListEmailsRequestOrderBy("status_desc")
	// Order by status (ascending alphabetical order)
	ListEmailsRequestOrderByStatusAsc = ListEmailsRequestOrderBy("status_asc")
	// Order by mail_from (descending alphabetical order)
	ListEmailsRequestOrderByMailFromDesc = ListEmailsRequestOrderBy("mail_from_desc")
	// Order by mail_from (ascending alphabetical order)
	ListEmailsRequestOrderByMailFromAsc = ListEmailsRequestOrderBy("mail_from_asc")
	// Order by mail recipient (descending alphabetical order)
	ListEmailsRequestOrderByMailRcptDesc = ListEmailsRequestOrderBy("mail_rcpt_desc")
	// Order by mail recipient (ascending alphabetical order)
	ListEmailsRequestOrderByMailRcptAsc = ListEmailsRequestOrderBy("mail_rcpt_asc")
	// Order by subject (descending alphabetical order)
	ListEmailsRequestOrderBySubjectDesc = ListEmailsRequestOrderBy("subject_desc")
	// Order by subject (ascending alphabetical order)
	ListEmailsRequestOrderBySubjectAsc = ListEmailsRequestOrderBy("subject_asc")
)

func (enum ListEmailsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListEmailsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEmailsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEmailsRequestOrderBy(ListEmailsRequestOrderBy(tmp).String())
	return nil
}

// CreateEmailRequestAddress: create email request. address.
type CreateEmailRequestAddress struct {
	// Email: email address.
	Email string `json:"email"`
	// Name: (Optional) Name displayed.
	Name *string `json:"name"`
}

// CreateEmailRequestAttachment: create email request. attachment.
type CreateEmailRequestAttachment struct {
	// Name: filename of the attachment.
	Name string `json:"name"`
	// Type: mIME type of the attachment.
	Type string `json:"type"`
	// Content: content of the attachment encoded in base64.
	Content []byte `json:"content"`
}

// CreateEmailResponse: create email response.
type CreateEmailResponse struct {
	// Emails: single page of emails matching the requested criteria.
	Emails []*Email `json:"emails"`
}

// Domain: domain.
type Domain struct {
	// ID: ID of the domain.
	ID string `json:"id"`
	// OrganizationID: ID of the domain's Organization.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the domain's Project.
	ProjectID string `json:"project_id"`
	// Name: domain name (example.com).
	Name string `json:"name"`
	// Status: status of the domain.
	// Default value: unknown
	Status DomainStatus `json:"status"`
	// CreatedAt: date and time of domain creation.
	CreatedAt *time.Time `json:"created_at"`
	// NextCheckAt: date and time of the next scheduled check.
	NextCheckAt *time.Time `json:"next_check_at"`
	// LastValidAt: date and time the domain was last valid.
	LastValidAt *time.Time `json:"last_valid_at"`
	// RevokedAt: date and time of the domain's deletion.
	RevokedAt *time.Time `json:"revoked_at"`
	// LastError: error message returned if the last check failed.
	LastError *string `json:"last_error"`
	// SpfConfig: snippet of the SPF record to register in the DNS zone.
	SpfConfig string `json:"spf_config"`
	// DkimConfig: dKIM public key to record in the DNS zone.
	DkimConfig string `json:"dkim_config"`
	// Statistics: domain's statistics.
	Statistics *DomainStatistics `json:"statistics"`

	Region scw.Region `json:"region"`
}

type DomainStatistics struct {
	TotalCount uint32 `json:"total_count"`

	SentCount uint32 `json:"sent_count"`

	FailedCount uint32 `json:"failed_count"`

	CanceledCount uint32 `json:"canceled_count"`
}

// Email: email.
type Email struct {
	// ID: technical ID of the email.
	ID string `json:"id"`
	// MessageID: message ID of the email.
	MessageID string `json:"message_id"`
	// ProjectID: ID of the Project to which the email belongs.
	ProjectID string `json:"project_id"`
	// MailFrom: email address of the sender.
	MailFrom string `json:"mail_from"`
	// Deprecated: RcptTo: (Deprecated) Email address of the recipient.
	RcptTo *string `json:"rcpt_to,omitempty"`
	// MailRcpt: email address of the recipient.
	MailRcpt string `json:"mail_rcpt"`
	// RcptType: type of recipient.
	// Default value: unknown_rcpt_type
	RcptType EmailRcptType `json:"rcpt_type"`
	// Subject: subject of the email.
	Subject string `json:"subject"`
	// CreatedAt: creation date of the email object.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: last update of the email object.
	UpdatedAt *time.Time `json:"updated_at"`
	// Status: status of the email.
	// Default value: unknown
	Status EmailStatus `json:"status"`
	// StatusDetails: additional status information.
	StatusDetails *string `json:"status_details"`
	// TryCount: number of attempts to send the email.
	TryCount uint32 `json:"try_count"`
	// LastTries: information about the last three attempts to send the email.
	LastTries []*EmailTry `json:"last_tries"`
}

// EmailTry: email. try.
type EmailTry struct {
	// Rank: rank number of this attempt to send the email.
	Rank uint32 `json:"rank"`
	// TriedAt: date of the attempt to send the email.
	TriedAt *time.Time `json:"tried_at"`
	// Code: the SMTP status code received after the attempt. 0 if the attempt did not reach an SMTP server.
	Code int32 `json:"code"`
	// Message: the SMTP message received. If the attempt did not reach an SMTP server, the message returned explains what happened.
	Message string `json:"message"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	// TotalCount: number of domains that match the request (without pagination).
	TotalCount uint32 `json:"total_count"`

	Domains []*Domain `json:"domains"`
}

// ListEmailsResponse: list emails response.
type ListEmailsResponse struct {
	// TotalCount: number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
	// Emails: single page of emails matching the requested criteria.
	Emails []*Email `json:"emails"`
}

// Statistics: statistics.
type Statistics struct {
	// TotalCount: total number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
	// NewCount: number of emails still in the `new` transient state. This means emails received from the API but not yet processed.
	NewCount uint32 `json:"new_count"`
	// SendingCount: number of emails still in the `sending` transient state. This means emails received from the API but not yet in their final status.
	SendingCount uint32 `json:"sending_count"`
	// SentCount: number of emails in the final `sent` state. This means emails that have been delivered to the target mail system.
	SentCount uint32 `json:"sent_count"`
	// FailedCount: number of emails in the final `failed` state. This means emails that have been refused by the target mail system with a final error status.
	FailedCount uint32 `json:"failed_count"`
	// CanceledCount: number of emails in the final `canceled` state. This means emails that have been canceled upon request.
	CanceledCount uint32 `json:"canceled_count"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type CreateEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// From: sender information. Must be from a checked domain declared in the Project.
	From *CreateEmailRequestAddress `json:"from"`
	// To: array of recipient information (limited to 1 recipient).
	To []*CreateEmailRequestAddress `json:"to"`
	// Cc: array of recipient information (unimplemented).
	Cc []*CreateEmailRequestAddress `json:"cc"`
	// Bcc: array of recipient information (unimplemented).
	Bcc []*CreateEmailRequestAddress `json:"bcc"`
	// Subject: subject of the email.
	Subject string `json:"subject"`
	// Text: text content.
	Text string `json:"text"`
	// HTML: HTML content.
	HTML string `json:"html"`
	// ProjectID: ID of the Project in which to create the email.
	ProjectID string `json:"project_id"`
	// Attachments: array of attachments.
	Attachments []*CreateEmailRequestAttachment `json:"attachments"`
	// SendBefore: maximum date to deliver the email.
	SendBefore *time.Time `json:"send_before"`
}

// CreateEmail: send an email.
// You must specify the `region`, the sender and the recipient's information and the `project_id` to send an email from a checked domain. The subject of the email must contain at least 6 characters.
func (s *API) CreateEmail(req *CreateEmailRequest, opts ...scw.RequestOption) (*CreateEmailResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateEmailResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to retrieve.
	EmailID string `json:"-"`
}

// GetEmail: get an email.
// Retrieve information about a specific email using the `email_id` and `region` parameters.
func (s *API) GetEmail(req *GetEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "",
		Headers: http.Header{},
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListEmailsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// ProjectID: (Optional) ID of the Project in which to list the emails.
	ProjectID *string `json:"-"`
	// DomainID: (Optional) ID of the domain for which to list the emails.
	DomainID *string `json:"-"`
	// MessageID: (Optional) ID of the message for which to list the emails.
	MessageID *string `json:"-"`
	// Since: (Optional) List emails created after this date.
	Since *time.Time `json:"-"`
	// Until: (Optional) List emails created before this date.
	Until *time.Time `json:"-"`
	// MailFrom: (Optional) List emails sent with this sender's email address.
	MailFrom *string `json:"-"`
	// Deprecated: MailTo: (Deprecated) List emails sent to this recipient's email address.
	MailTo *string `json:"-"`
	// MailRcpt: (Optional) List emails sent to this recipient's email address.
	MailRcpt *string `json:"-"`
	// Statuses: (Optional) List emails with any of these statuses.
	Statuses []EmailStatus `json:"-"`
	// Subject: (Optional) List emails with this subject.
	Subject *string `json:"-"`
	// OrderBy: (Optional) List emails corresponding to specific criteria.
	// You can filter your emails in ascending or descending order using:
	//   - created_at
	//   - updated_at
	//   - status
	//   - mail_from
	//   - mail_rcpt
	//   - subject.
	// Default value: created_at_desc
	OrderBy ListEmailsRequestOrderBy `json:"-"`
}

// ListEmails: list emails.
// Retrieve the list of emails sent from a specific domain or for a specific Project or Organization. You must specify the `region`.
// You can filter your emails in ascending or descending order using:
//   - created_at
//   - updated_at
//   - status
//   - mail_from
//   - mail_rcpt
//   - subject.
func (s *API) ListEmails(req *ListEmailsRequest, opts ...scw.RequestOption) (*ListEmailsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "message_id", req.MessageID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)
	parameter.AddToQuery(query, "mail_to", req.MailTo)
	parameter.AddToQuery(query, "mail_rcpt", req.MailRcpt)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "subject", req.Subject)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetStatisticsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ProjectID: (Optional) Number of emails for this Project.
	ProjectID *string `json:"-"`
	// DomainID: (Optional) Number of emails sent from this domain (must be coherent with the `project_id` and the `organization_id`).
	DomainID *string `json:"-"`
	// Since: (Optional) Number of emails created after this date.
	Since *time.Time `json:"-"`
	// Until: (Optional) Number of emails created before this date.
	Until *time.Time `json:"-"`
	// MailFrom: (Optional) Number of emails sent with this sender's email address.
	MailFrom *string `json:"-"`
}

// GetStatistics: email statuses.
// Get information on your emails' statuses.
func (s *API) GetStatistics(req *GetStatisticsRequest, opts ...scw.RequestOption) (*Statistics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/statistics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Statistics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to cancel.
	EmailID string `json:"-"`
}

// CancelEmail: cancel an email.
// You can cancel the sending of an email if it has not been sent yet. You must specify the `region` and the `email_id` of the email you want to cancel.
func (s *API) CancelEmail(req *CancelEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "/cancel",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ProjectID: ID of the project to which the domain belongs.
	ProjectID string `json:"project_id"`
	// DomainName: fully qualified domain dame.
	DomainName string `json:"domain_name"`
	// AcceptTos: accept Scaleway's Terms of Service.
	AcceptTos bool `json:"accept_tos"`
}

// CreateDomain: register a domain in a project.
// You must specify the `region`, `project_id` and `domain_name` to register a domain in a specific Project.
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain.
	DomainID string `json:"-"`
}

// GetDomain: get information about a domain.
// Retrieve information about a specific domain using the `region` and `domain_id` parameters.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	Status []DomainStatus `json:"-"`

	OrganizationID *string `json:"-"`

	Name *string `json:"-"`
}

// ListDomains: list domains.
// Retrieve domains in a specific project or in a specific Organization using the `region` parameter.
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RevokeDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// RevokeDomain: delete a domain.
// You must specify the domain you want to delete by the `region` and `domain_id`. Deleting a domain is permanent and cannot be undone.
func (s *API) RevokeDomain(req *RevokeDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/revoke",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CheckDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to check.
	DomainID string `json:"-"`
}

// CheckDomain: domain DNS check.
// Perform an immediate DNS check of a domain using the `region` and `domain_id` parameters.
func (s *API) CheckDomain(req *CheckDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/check",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEmailsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Emails = append(r.Emails, results.Emails...)
	r.TotalCount += uint32(len(results.Emails))
	return uint32(len(results.Emails)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}
