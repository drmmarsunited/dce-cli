// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Lease Lease Details
// swagger:model lease
type Lease struct {

	// accountId of the AWS account
	AccountID string `json:"accountId,omitempty"`

	// budget amount
	BudgetAmount float64 `json:"budgetAmount,omitempty"`

	// budget currency
	BudgetCurrency string `json:"budgetCurrency,omitempty"`

	// budget notification emails
	BudgetNotificationEmails []string `json:"budgetNotificationEmails"`

	// creation date in epoch seconds
	CreatedOn float64 `json:"createdOn,omitempty"`

	// date lease should expire in epoch seconds
	ExpiresOn float64 `json:"expiresOn,omitempty"`

	// Lease ID
	ID string `json:"id,omitempty"`

	// date last modified in epoch seconds
	LastModifiedOn float64 `json:"lastModifiedOn,omitempty"`

	// Status of the Lease.
	// "Active": The principal is leased and has access to the account
	// "Inactive": The lease has become inactive, either through expiring, exceeding budget, or by request.
	//
	// Enum: [Active Inactive]
	LeaseStatus string `json:"leaseStatus,omitempty"`

	// date lease status was last modified in epoch seconds
	LeaseStatusModifiedOn float64 `json:"leaseStatusModifiedOn,omitempty"`

	// A reason behind the lease status.
	// "LeaseExpired": The lease exceeded its expiration time ("expiresOn") and
	// the associated account was reset and returned to the account pool.
	// "LeaseOverBudget": The lease exceeded its budgeted amount and the
	// associated account was reset and returned to the account pool.
	// "LeaseDestroyed": The lease was adminstratively ended, which can be done
	// via the leases API.
	// "LeaseActive": The lease is active.
	// "LeaseRolledBack": A system error occurred while provisioning the lease.
	// and it was rolled back.
	//
	// Enum: [LeaseExpired LeaseOverBudget LeaseDestroyed LeaseActive LeaseRolledBack]
	LeaseStatusReason string `json:"leaseStatusReason,omitempty"`

	// principalId of the lease to get
	PrincipalID string `json:"principalId,omitempty"`
}

// Validate validates this lease
func (m *Lease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeaseStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaseStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var leaseTypeLeaseStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		leaseTypeLeaseStatusPropEnum = append(leaseTypeLeaseStatusPropEnum, v)
	}
}

const (

	// LeaseLeaseStatusActive captures enum value "Active"
	LeaseLeaseStatusActive string = "Active"

	// LeaseLeaseStatusInactive captures enum value "Inactive"
	LeaseLeaseStatusInactive string = "Inactive"
)

// prop value enum
func (m *Lease) validateLeaseStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, leaseTypeLeaseStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Lease) validateLeaseStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.LeaseStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLeaseStatusEnum("leaseStatus", "body", m.LeaseStatus); err != nil {
		return err
	}

	return nil
}

var leaseTypeLeaseStatusReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LeaseExpired","LeaseOverBudget","LeaseDestroyed","LeaseActive","LeaseRolledBack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		leaseTypeLeaseStatusReasonPropEnum = append(leaseTypeLeaseStatusReasonPropEnum, v)
	}
}

const (

	// LeaseLeaseStatusReasonLeaseExpired captures enum value "LeaseExpired"
	LeaseLeaseStatusReasonLeaseExpired string = "LeaseExpired"

	// LeaseLeaseStatusReasonLeaseOverBudget captures enum value "LeaseOverBudget"
	LeaseLeaseStatusReasonLeaseOverBudget string = "LeaseOverBudget"

	// LeaseLeaseStatusReasonLeaseDestroyed captures enum value "LeaseDestroyed"
	LeaseLeaseStatusReasonLeaseDestroyed string = "LeaseDestroyed"

	// LeaseLeaseStatusReasonLeaseActive captures enum value "LeaseActive"
	LeaseLeaseStatusReasonLeaseActive string = "LeaseActive"

	// LeaseLeaseStatusReasonLeaseRolledBack captures enum value "LeaseRolledBack"
	LeaseLeaseStatusReasonLeaseRolledBack string = "LeaseRolledBack"
)

// prop value enum
func (m *Lease) validateLeaseStatusReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, leaseTypeLeaseStatusReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Lease) validateLeaseStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(m.LeaseStatusReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateLeaseStatusReasonEnum("leaseStatusReason", "body", m.LeaseStatusReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Lease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Lease) UnmarshalBinary(b []byte) error {
	var res Lease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
