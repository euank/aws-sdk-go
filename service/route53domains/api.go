package route53domains

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"time"
)

// CheckDomainAvailabilityRequest generates a request for the CheckDomainAvailability operation.
func (c *Route53Domains) CheckDomainAvailabilityRequest(input *CheckDomainAvailabilityInput) (req *aws.Request, output *CheckDomainAvailabilityOutput) {
	if opCheckDomainAvailability == nil {
		opCheckDomainAvailability = &aws.Operation{
			Name:       "CheckDomainAvailability",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCheckDomainAvailability, input, output)
	output = &CheckDomainAvailabilityOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) CheckDomainAvailability(input *CheckDomainAvailabilityInput) (output *CheckDomainAvailabilityOutput, err error) {
	req, out := c.CheckDomainAvailabilityRequest(input)
	output = out
	err = req.Send()
	return
}

var opCheckDomainAvailability *aws.Operation

// DeleteTagsForDomainRequest generates a request for the DeleteTagsForDomain operation.
func (c *Route53Domains) DeleteTagsForDomainRequest(input *DeleteTagsForDomainInput) (req *aws.Request, output *DeleteTagsForDomainOutput) {
	if opDeleteTagsForDomain == nil {
		opDeleteTagsForDomain = &aws.Operation{
			Name:       "DeleteTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteTagsForDomain, input, output)
	output = &DeleteTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DeleteTagsForDomain(input *DeleteTagsForDomainInput) (output *DeleteTagsForDomainOutput, err error) {
	req, out := c.DeleteTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTagsForDomain *aws.Operation

// DisableDomainAutoRenewRequest generates a request for the DisableDomainAutoRenew operation.
func (c *Route53Domains) DisableDomainAutoRenewRequest(input *DisableDomainAutoRenewInput) (req *aws.Request, output *DisableDomainAutoRenewOutput) {
	if opDisableDomainAutoRenew == nil {
		opDisableDomainAutoRenew = &aws.Operation{
			Name:       "DisableDomainAutoRenew",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableDomainAutoRenew, input, output)
	output = &DisableDomainAutoRenewOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DisableDomainAutoRenew(input *DisableDomainAutoRenewInput) (output *DisableDomainAutoRenewOutput, err error) {
	req, out := c.DisableDomainAutoRenewRequest(input)
	output = out
	err = req.Send()
	return
}

var opDisableDomainAutoRenew *aws.Operation

// DisableDomainTransferLockRequest generates a request for the DisableDomainTransferLock operation.
func (c *Route53Domains) DisableDomainTransferLockRequest(input *DisableDomainTransferLockInput) (req *aws.Request, output *DisableDomainTransferLockOutput) {
	if opDisableDomainTransferLock == nil {
		opDisableDomainTransferLock = &aws.Operation{
			Name:       "DisableDomainTransferLock",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableDomainTransferLock, input, output)
	output = &DisableDomainTransferLockOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DisableDomainTransferLock(input *DisableDomainTransferLockInput) (output *DisableDomainTransferLockOutput, err error) {
	req, out := c.DisableDomainTransferLockRequest(input)
	output = out
	err = req.Send()
	return
}

var opDisableDomainTransferLock *aws.Operation

// EnableDomainAutoRenewRequest generates a request for the EnableDomainAutoRenew operation.
func (c *Route53Domains) EnableDomainAutoRenewRequest(input *EnableDomainAutoRenewInput) (req *aws.Request, output *EnableDomainAutoRenewOutput) {
	if opEnableDomainAutoRenew == nil {
		opEnableDomainAutoRenew = &aws.Operation{
			Name:       "EnableDomainAutoRenew",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableDomainAutoRenew, input, output)
	output = &EnableDomainAutoRenewOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) EnableDomainAutoRenew(input *EnableDomainAutoRenewInput) (output *EnableDomainAutoRenewOutput, err error) {
	req, out := c.EnableDomainAutoRenewRequest(input)
	output = out
	err = req.Send()
	return
}

var opEnableDomainAutoRenew *aws.Operation

// EnableDomainTransferLockRequest generates a request for the EnableDomainTransferLock operation.
func (c *Route53Domains) EnableDomainTransferLockRequest(input *EnableDomainTransferLockInput) (req *aws.Request, output *EnableDomainTransferLockOutput) {
	if opEnableDomainTransferLock == nil {
		opEnableDomainTransferLock = &aws.Operation{
			Name:       "EnableDomainTransferLock",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableDomainTransferLock, input, output)
	output = &EnableDomainTransferLockOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) EnableDomainTransferLock(input *EnableDomainTransferLockInput) (output *EnableDomainTransferLockOutput, err error) {
	req, out := c.EnableDomainTransferLockRequest(input)
	output = out
	err = req.Send()
	return
}

var opEnableDomainTransferLock *aws.Operation

// GetDomainDetailRequest generates a request for the GetDomainDetail operation.
func (c *Route53Domains) GetDomainDetailRequest(input *GetDomainDetailInput) (req *aws.Request, output *GetDomainDetailOutput) {
	if opGetDomainDetail == nil {
		opGetDomainDetail = &aws.Operation{
			Name:       "GetDomainDetail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetDomainDetail, input, output)
	output = &GetDomainDetailOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) GetDomainDetail(input *GetDomainDetailInput) (output *GetDomainDetailOutput, err error) {
	req, out := c.GetDomainDetailRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetDomainDetail *aws.Operation

// GetOperationDetailRequest generates a request for the GetOperationDetail operation.
func (c *Route53Domains) GetOperationDetailRequest(input *GetOperationDetailInput) (req *aws.Request, output *GetOperationDetailOutput) {
	if opGetOperationDetail == nil {
		opGetOperationDetail = &aws.Operation{
			Name:       "GetOperationDetail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetOperationDetail, input, output)
	output = &GetOperationDetailOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) GetOperationDetail(input *GetOperationDetailInput) (output *GetOperationDetailOutput, err error) {
	req, out := c.GetOperationDetailRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetOperationDetail *aws.Operation

// ListDomainsRequest generates a request for the ListDomains operation.
func (c *Route53Domains) ListDomainsRequest(input *ListDomainsInput) (req *aws.Request, output *ListDomainsOutput) {
	if opListDomains == nil {
		opListDomains = &aws.Operation{
			Name:       "ListDomains",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListDomains, input, output)
	output = &ListDomainsOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListDomains(input *ListDomainsInput) (output *ListDomainsOutput, err error) {
	req, out := c.ListDomainsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDomains *aws.Operation

// ListOperationsRequest generates a request for the ListOperations operation.
func (c *Route53Domains) ListOperationsRequest(input *ListOperationsInput) (req *aws.Request, output *ListOperationsOutput) {
	if opListOperations == nil {
		opListOperations = &aws.Operation{
			Name:       "ListOperations",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListOperations, input, output)
	output = &ListOperationsOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListOperations(input *ListOperationsInput) (output *ListOperationsOutput, err error) {
	req, out := c.ListOperationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListOperations *aws.Operation

// ListTagsForDomainRequest generates a request for the ListTagsForDomain operation.
func (c *Route53Domains) ListTagsForDomainRequest(input *ListTagsForDomainInput) (req *aws.Request, output *ListTagsForDomainOutput) {
	if opListTagsForDomain == nil {
		opListTagsForDomain = &aws.Operation{
			Name:       "ListTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListTagsForDomain, input, output)
	output = &ListTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListTagsForDomain(input *ListTagsForDomainInput) (output *ListTagsForDomainOutput, err error) {
	req, out := c.ListTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTagsForDomain *aws.Operation

// RegisterDomainRequest generates a request for the RegisterDomain operation.
func (c *Route53Domains) RegisterDomainRequest(input *RegisterDomainInput) (req *aws.Request, output *RegisterDomainOutput) {
	if opRegisterDomain == nil {
		opRegisterDomain = &aws.Operation{
			Name:       "RegisterDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRegisterDomain, input, output)
	output = &RegisterDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) RegisterDomain(input *RegisterDomainInput) (output *RegisterDomainOutput, err error) {
	req, out := c.RegisterDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterDomain *aws.Operation

// RetrieveDomainAuthCodeRequest generates a request for the RetrieveDomainAuthCode operation.
func (c *Route53Domains) RetrieveDomainAuthCodeRequest(input *RetrieveDomainAuthCodeInput) (req *aws.Request, output *RetrieveDomainAuthCodeOutput) {
	if opRetrieveDomainAuthCode == nil {
		opRetrieveDomainAuthCode = &aws.Operation{
			Name:       "RetrieveDomainAuthCode",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRetrieveDomainAuthCode, input, output)
	output = &RetrieveDomainAuthCodeOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) RetrieveDomainAuthCode(input *RetrieveDomainAuthCodeInput) (output *RetrieveDomainAuthCodeOutput, err error) {
	req, out := c.RetrieveDomainAuthCodeRequest(input)
	output = out
	err = req.Send()
	return
}

var opRetrieveDomainAuthCode *aws.Operation

// TransferDomainRequest generates a request for the TransferDomain operation.
func (c *Route53Domains) TransferDomainRequest(input *TransferDomainInput) (req *aws.Request, output *TransferDomainOutput) {
	if opTransferDomain == nil {
		opTransferDomain = &aws.Operation{
			Name:       "TransferDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opTransferDomain, input, output)
	output = &TransferDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) TransferDomain(input *TransferDomainInput) (output *TransferDomainOutput, err error) {
	req, out := c.TransferDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opTransferDomain *aws.Operation

// UpdateDomainContactRequest generates a request for the UpdateDomainContact operation.
func (c *Route53Domains) UpdateDomainContactRequest(input *UpdateDomainContactInput) (req *aws.Request, output *UpdateDomainContactOutput) {
	if opUpdateDomainContact == nil {
		opUpdateDomainContact = &aws.Operation{
			Name:       "UpdateDomainContact",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainContact, input, output)
	output = &UpdateDomainContactOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainContact(input *UpdateDomainContactInput) (output *UpdateDomainContactOutput, err error) {
	req, out := c.UpdateDomainContactRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainContact *aws.Operation

// UpdateDomainContactPrivacyRequest generates a request for the UpdateDomainContactPrivacy operation.
func (c *Route53Domains) UpdateDomainContactPrivacyRequest(input *UpdateDomainContactPrivacyInput) (req *aws.Request, output *UpdateDomainContactPrivacyOutput) {
	if opUpdateDomainContactPrivacy == nil {
		opUpdateDomainContactPrivacy = &aws.Operation{
			Name:       "UpdateDomainContactPrivacy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainContactPrivacy, input, output)
	output = &UpdateDomainContactPrivacyOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainContactPrivacy(input *UpdateDomainContactPrivacyInput) (output *UpdateDomainContactPrivacyOutput, err error) {
	req, out := c.UpdateDomainContactPrivacyRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainContactPrivacy *aws.Operation

// UpdateDomainNameserversRequest generates a request for the UpdateDomainNameservers operation.
func (c *Route53Domains) UpdateDomainNameserversRequest(input *UpdateDomainNameserversInput) (req *aws.Request, output *UpdateDomainNameserversOutput) {
	if opUpdateDomainNameservers == nil {
		opUpdateDomainNameservers = &aws.Operation{
			Name:       "UpdateDomainNameservers",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainNameservers, input, output)
	output = &UpdateDomainNameserversOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainNameservers(input *UpdateDomainNameserversInput) (output *UpdateDomainNameserversOutput, err error) {
	req, out := c.UpdateDomainNameserversRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainNameservers *aws.Operation

// UpdateTagsForDomainRequest generates a request for the UpdateTagsForDomain operation.
func (c *Route53Domains) UpdateTagsForDomainRequest(input *UpdateTagsForDomainInput) (req *aws.Request, output *UpdateTagsForDomainOutput) {
	if opUpdateTagsForDomain == nil {
		opUpdateTagsForDomain = &aws.Operation{
			Name:       "UpdateTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateTagsForDomain, input, output)
	output = &UpdateTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateTagsForDomain(input *UpdateTagsForDomainInput) (output *UpdateTagsForDomainOutput, err error) {
	req, out := c.UpdateTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateTagsForDomain *aws.Operation

type CheckDomainAvailabilityInput struct {
	DomainName  *string `type:"string" json:",omitempty"`
	IDNLangCode *string `locationName:"IdnLangCode" type:"string" json:"IdnLangCode,omitempty"`

	metadataCheckDomainAvailabilityInput `json:"-", xml:"-"`
}

type metadataCheckDomainAvailabilityInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type CheckDomainAvailabilityOutput struct {
	Availability *string `type:"string" json:",omitempty"`

	metadataCheckDomainAvailabilityOutput `json:"-", xml:"-"`
}

type metadataCheckDomainAvailabilityOutput struct {
	SDKShapeTraits bool `type:"structure" required:"Availability" json:",omitempty"`
}

type ContactDetail struct {
	AddressLine1     *string       `type:"string" json:",omitempty"`
	AddressLine2     *string       `type:"string" json:",omitempty"`
	City             *string       `type:"string" json:",omitempty"`
	ContactType      *string       `type:"string" json:",omitempty"`
	CountryCode      *string       `type:"string" json:",omitempty"`
	Email            *string       `type:"string" json:",omitempty"`
	ExtraParams      []*ExtraParam `type:"list" json:",omitempty"`
	Fax              *string       `type:"string" json:",omitempty"`
	FirstName        *string       `type:"string" json:",omitempty"`
	LastName         *string       `type:"string" json:",omitempty"`
	OrganizationName *string       `type:"string" json:",omitempty"`
	PhoneNumber      *string       `type:"string" json:",omitempty"`
	State            *string       `type:"string" json:",omitempty"`
	ZipCode          *string       `type:"string" json:",omitempty"`

	metadataContactDetail `json:"-", xml:"-"`
}

type metadataContactDetail struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteTagsForDomainInput struct {
	DomainName   *string   `type:"string" json:",omitempty"`
	TagsToDelete []*string `type:"list" json:",omitempty"`

	metadataDeleteTagsForDomainInput `json:"-", xml:"-"`
}

type metadataDeleteTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName,TagsToDelete" json:",omitempty"`
}

type DeleteTagsForDomainOutput struct {
	metadataDeleteTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataDeleteTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DisableDomainAutoRenewInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataDisableDomainAutoRenewInput `json:"-", xml:"-"`
}

type metadataDisableDomainAutoRenewInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type DisableDomainAutoRenewOutput struct {
	metadataDisableDomainAutoRenewOutput `json:"-", xml:"-"`
}

type metadataDisableDomainAutoRenewOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DisableDomainTransferLockInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataDisableDomainTransferLockInput `json:"-", xml:"-"`
}

type metadataDisableDomainTransferLockInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type DisableDomainTransferLockOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataDisableDomainTransferLockOutput `json:"-", xml:"-"`
}

type metadataDisableDomainTransferLockOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type DomainLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataDomainLimitExceeded `json:"-", xml:"-"`
}

type metadataDomainLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DomainSummary struct {
	AutoRenew    *bool      `type:"boolean" json:",omitempty"`
	DomainName   *string    `type:"string" json:",omitempty"`
	Expiry       *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	TransferLock *bool      `type:"boolean" json:",omitempty"`

	metadataDomainSummary `json:"-", xml:"-"`
}

type metadataDomainSummary struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type DuplicateRequest struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataDuplicateRequest `json:"-", xml:"-"`
}

type metadataDuplicateRequest struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type EnableDomainAutoRenewInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataEnableDomainAutoRenewInput `json:"-", xml:"-"`
}

type metadataEnableDomainAutoRenewInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type EnableDomainAutoRenewOutput struct {
	metadataEnableDomainAutoRenewOutput `json:"-", xml:"-"`
}

type metadataEnableDomainAutoRenewOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type EnableDomainTransferLockInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataEnableDomainTransferLockInput `json:"-", xml:"-"`
}

type metadataEnableDomainTransferLockInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type EnableDomainTransferLockOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataEnableDomainTransferLockOutput `json:"-", xml:"-"`
}

type metadataEnableDomainTransferLockOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type ExtraParam struct {
	Name  *string `type:"string" json:",omitempty"`
	Value *string `type:"string" json:",omitempty"`

	metadataExtraParam `json:"-", xml:"-"`
}

type metadataExtraParam struct {
	SDKShapeTraits bool `type:"structure" required:"Name,Value" json:",omitempty"`
}

type GetDomainDetailInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataGetDomainDetailInput `json:"-", xml:"-"`
}

type metadataGetDomainDetailInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type GetDomainDetailOutput struct {
	AbuseContactEmail *string        `type:"string" json:",omitempty"`
	AbuseContactPhone *string        `type:"string" json:",omitempty"`
	AdminContact      *ContactDetail `type:"structure" json:",omitempty"`
	AdminPrivacy      *bool          `type:"boolean" json:",omitempty"`
	AutoRenew         *bool          `type:"boolean" json:",omitempty"`
	CreationDate      *time.Time     `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	DNSSec            *string        `locationName:"DnsSec" type:"string" json:"DnsSec,omitempty"`
	DomainName        *string        `type:"string" json:",omitempty"`
	ExpirationDate    *time.Time     `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Nameservers       []*Nameserver  `type:"list" json:",omitempty"`
	RegistrantContact *ContactDetail `type:"structure" json:",omitempty"`
	RegistrantPrivacy *bool          `type:"boolean" json:",omitempty"`
	RegistrarName     *string        `type:"string" json:",omitempty"`
	RegistrarURL      *string        `locationName:"RegistrarUrl" type:"string" json:"RegistrarUrl,omitempty"`
	RegistryDomainID  *string        `locationName:"RegistryDomainId" type:"string" json:"RegistryDomainId,omitempty"`
	Reseller          *string        `type:"string" json:",omitempty"`
	StatusList        []*string      `type:"list" json:",omitempty"`
	TechContact       *ContactDetail `type:"structure" json:",omitempty"`
	TechPrivacy       *bool          `type:"boolean" json:",omitempty"`
	UpdatedDate       *time.Time     `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	WhoIsServer       *string        `type:"string" json:",omitempty"`

	metadataGetDomainDetailOutput `json:"-", xml:"-"`
}

type metadataGetDomainDetailOutput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName,Nameservers,AdminContact,RegistrantContact,TechContact" json:",omitempty"`
}

type GetOperationDetailInput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataGetOperationDetailInput `json:"-", xml:"-"`
}

type metadataGetOperationDetailInput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type GetOperationDetailOutput struct {
	DomainName    *string    `type:"string" json:",omitempty"`
	Message       *string    `type:"string" json:",omitempty"`
	OperationID   *string    `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`
	Status        *string    `type:"string" json:",omitempty"`
	SubmittedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Type          *string    `type:"string" json:",omitempty"`

	metadataGetOperationDetailOutput `json:"-", xml:"-"`
}

type metadataGetOperationDetailOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidInput struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataInvalidInput `json:"-", xml:"-"`
}

type metadataInvalidInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListDomainsInput struct {
	Marker   *string `type:"string" json:",omitempty"`
	MaxItems *int    `type:"integer" json:",omitempty"`

	metadataListDomainsInput `json:"-", xml:"-"`
}

type metadataListDomainsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListDomainsOutput struct {
	Domains        []*DomainSummary `type:"list" json:",omitempty"`
	NextPageMarker *string          `type:"string" json:",omitempty"`

	metadataListDomainsOutput `json:"-", xml:"-"`
}

type metadataListDomainsOutput struct {
	SDKShapeTraits bool `type:"structure" required:"Domains" json:",omitempty"`
}

type ListOperationsInput struct {
	Marker   *string `type:"string" json:",omitempty"`
	MaxItems *int    `type:"integer" json:",omitempty"`

	metadataListOperationsInput `json:"-", xml:"-"`
}

type metadataListOperationsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListOperationsOutput struct {
	NextPageMarker *string             `type:"string" json:",omitempty"`
	Operations     []*OperationSummary `type:"list" json:",omitempty"`

	metadataListOperationsOutput `json:"-", xml:"-"`
}

type metadataListOperationsOutput struct {
	SDKShapeTraits bool `type:"structure" required:"Operations" json:",omitempty"`
}

type ListTagsForDomainInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataListTagsForDomainInput `json:"-", xml:"-"`
}

type metadataListTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type ListTagsForDomainOutput struct {
	TagList []*Tag `type:"list" json:",omitempty"`

	metadataListTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataListTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure" required:"TagList" json:",omitempty"`
}

type Nameserver struct {
	GlueIPs []*string `locationName:"GlueIps" type:"list" json:"GlueIps,omitempty"`
	Name    *string   `type:"string" json:",omitempty"`

	metadataNameserver `json:"-", xml:"-"`
}

type metadataNameserver struct {
	SDKShapeTraits bool `type:"structure" required:"Name" json:",omitempty"`
}

type OperationLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataOperationLimitExceeded `json:"-", xml:"-"`
}

type metadataOperationLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type OperationSummary struct {
	OperationID   *string    `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`
	Status        *string    `type:"string" json:",omitempty"`
	SubmittedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Type          *string    `type:"string" json:",omitempty"`

	metadataOperationSummary `json:"-", xml:"-"`
}

type metadataOperationSummary struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId,Status,Type,SubmittedDate" json:",omitempty"`
}

type RegisterDomainInput struct {
	AdminContact                    *ContactDetail `type:"structure" json:",omitempty"`
	AutoRenew                       *bool          `type:"boolean" json:",omitempty"`
	DomainName                      *string        `type:"string" json:",omitempty"`
	DurationInYears                 *int           `type:"integer" json:",omitempty"`
	IDNLangCode                     *string        `locationName:"IdnLangCode" type:"string" json:"IdnLangCode,omitempty"`
	PrivacyProtectAdminContact      *bool          `type:"boolean" json:",omitempty"`
	PrivacyProtectRegistrantContact *bool          `type:"boolean" json:",omitempty"`
	PrivacyProtectTechContact       *bool          `type:"boolean" json:",omitempty"`
	RegistrantContact               *ContactDetail `type:"structure" json:",omitempty"`
	TechContact                     *ContactDetail `type:"structure" json:",omitempty"`

	metadataRegisterDomainInput `json:"-", xml:"-"`
}

type metadataRegisterDomainInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName,DurationInYears,AdminContact,RegistrantContact,TechContact" json:",omitempty"`
}

type RegisterDomainOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataRegisterDomainOutput `json:"-", xml:"-"`
}

type metadataRegisterDomainOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type RetrieveDomainAuthCodeInput struct {
	DomainName *string `type:"string" json:",omitempty"`

	metadataRetrieveDomainAuthCodeInput `json:"-", xml:"-"`
}

type metadataRetrieveDomainAuthCodeInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type RetrieveDomainAuthCodeOutput struct {
	AuthCode *string `type:"string" json:",omitempty"`

	metadataRetrieveDomainAuthCodeOutput `json:"-", xml:"-"`
}

type metadataRetrieveDomainAuthCodeOutput struct {
	SDKShapeTraits bool `type:"structure" required:"AuthCode" json:",omitempty"`
}

type TLDRulesViolation struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataTLDRulesViolation `json:"-", xml:"-"`
}

type metadataTLDRulesViolation struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Tag struct {
	Key   *string `type:"string" json:",omitempty"`
	Value *string `type:"string" json:",omitempty"`

	metadataTag `json:"-", xml:"-"`
}

type metadataTag struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TransferDomainInput struct {
	AdminContact                    *ContactDetail `type:"structure" json:",omitempty"`
	AuthCode                        *string        `type:"string" json:",omitempty"`
	AutoRenew                       *bool          `type:"boolean" json:",omitempty"`
	DomainName                      *string        `type:"string" json:",omitempty"`
	DurationInYears                 *int           `type:"integer" json:",omitempty"`
	IDNLangCode                     *string        `locationName:"IdnLangCode" type:"string" json:"IdnLangCode,omitempty"`
	Nameservers                     []*Nameserver  `type:"list" json:",omitempty"`
	PrivacyProtectAdminContact      *bool          `type:"boolean" json:",omitempty"`
	PrivacyProtectRegistrantContact *bool          `type:"boolean" json:",omitempty"`
	PrivacyProtectTechContact       *bool          `type:"boolean" json:",omitempty"`
	RegistrantContact               *ContactDetail `type:"structure" json:",omitempty"`
	TechContact                     *ContactDetail `type:"structure" json:",omitempty"`

	metadataTransferDomainInput `json:"-", xml:"-"`
}

type metadataTransferDomainInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName,DurationInYears,AdminContact,RegistrantContact,TechContact" json:",omitempty"`
}

type TransferDomainOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataTransferDomainOutput `json:"-", xml:"-"`
}

type metadataTransferDomainOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type UnsupportedTLD struct {
	Message *string `locationName:"message" type:"string" json:"message,omitempty"`

	metadataUnsupportedTLD `json:"-", xml:"-"`
}

type metadataUnsupportedTLD struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateDomainContactInput struct {
	AdminContact      *ContactDetail `type:"structure" json:",omitempty"`
	DomainName        *string        `type:"string" json:",omitempty"`
	RegistrantContact *ContactDetail `type:"structure" json:",omitempty"`
	TechContact       *ContactDetail `type:"structure" json:",omitempty"`

	metadataUpdateDomainContactInput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type UpdateDomainContactOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataUpdateDomainContactOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type UpdateDomainContactPrivacyInput struct {
	AdminPrivacy      *bool   `type:"boolean" json:",omitempty"`
	DomainName        *string `type:"string" json:",omitempty"`
	RegistrantPrivacy *bool   `type:"boolean" json:",omitempty"`
	TechPrivacy       *bool   `type:"boolean" json:",omitempty"`

	metadataUpdateDomainContactPrivacyInput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactPrivacyInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type UpdateDomainContactPrivacyOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataUpdateDomainContactPrivacyOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactPrivacyOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type UpdateDomainNameserversInput struct {
	DomainName  *string       `type:"string" json:",omitempty"`
	FIAuthKey   *string       `type:"string" json:",omitempty"`
	Nameservers []*Nameserver `type:"list" json:",omitempty"`

	metadataUpdateDomainNameserversInput `json:"-", xml:"-"`
}

type metadataUpdateDomainNameserversInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName,Nameservers" json:",omitempty"`
}

type UpdateDomainNameserversOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" json:"OperationId,omitempty"`

	metadataUpdateDomainNameserversOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainNameserversOutput struct {
	SDKShapeTraits bool `type:"structure" required:"OperationId" json:",omitempty"`
}

type UpdateTagsForDomainInput struct {
	DomainName   *string `type:"string" json:",omitempty"`
	TagsToUpdate []*Tag  `type:"list" json:",omitempty"`

	metadataUpdateTagsForDomainInput `json:"-", xml:"-"`
}

type metadataUpdateTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure" required:"DomainName" json:",omitempty"`
}

type UpdateTagsForDomainOutput struct {
	metadataUpdateTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataUpdateTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}