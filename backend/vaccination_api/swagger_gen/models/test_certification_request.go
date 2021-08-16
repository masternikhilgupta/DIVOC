// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestCertificationRequest test certification request
//
// swagger:model TestCertificationRequest
type TestCertificationRequest struct {

	// facility
	// Required: true
	Facility *TestCertificationRequestFacility `json:"facility"`

	// meta
	Meta interface{} `json:"meta,omitempty"`

	// pre enrollment code
	// Required: true
	PreEnrollmentCode *string `json:"preEnrollmentCode"`

	// program Id
	// Min Length: 1
	ProgramID string `json:"programId,omitempty"`

	// recipient
	// Required: true
	Recipient *TestCertificationRequestRecipient `json:"recipient"`

	// test details
	// Required: true
	TestDetails *TestCertificationRequestTestDetails `json:"testDetails"`

	// verifier
	Verifier *TestCertificationRequestVerifier `json:"verifier,omitempty"`
}

// Validate validates this test certification request
func (m *TestCertificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreEnrollmentCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequest) validateFacility(formats strfmt.Registry) error {

	if err := validate.Required("facility", "body", m.Facility); err != nil {
		return err
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility")
			}
			return err
		}
	}

	return nil
}

func (m *TestCertificationRequest) validatePreEnrollmentCode(formats strfmt.Registry) error {

	if err := validate.Required("preEnrollmentCode", "body", m.PreEnrollmentCode); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequest) validateProgramID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgramID) { // not required
		return nil
	}

	if err := validate.MinLength("programId", "body", string(m.ProgramID), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequest) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *TestCertificationRequest) validateTestDetails(formats strfmt.Registry) error {

	if err := validate.Required("testDetails", "body", m.TestDetails); err != nil {
		return err
	}

	if m.TestDetails != nil {
		if err := m.TestDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TestCertificationRequest) validateVerifier(formats strfmt.Registry) error {

	if swag.IsZero(m.Verifier) { // not required
		return nil
	}

	if m.Verifier != nil {
		if err := m.Verifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequest) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestFacility test certification request facility
//
// swagger:model TestCertificationRequestFacility
type TestCertificationRequestFacility struct {

	// address
	// Required: true
	Address *TestCertificationRequestFacilityAddress `json:"address"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this test certification request facility
func (m *TestCertificationRequestFacility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestFacility) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *TestCertificationRequestFacility) validateName(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestFacility) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestFacilityAddress test certification request facility address
//
// swagger:model TestCertificationRequestFacilityAddress
type TestCertificationRequestFacilityAddress struct {

	// address line1
	// Required: true
	AddressLine1 *string `json:"addressLine1"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	// Required: true
	// Min Length: 1
	District *string `json:"district"`

	// pincode
	// Required: true
	// Min Length: 1
	Pincode *string `json:"pincode"`

	// state
	// Required: true
	// Min Length: 1
	State *string `json:"state"`
}

// Validate validates this test certification request facility address
func (m *TestCertificationRequestFacilityAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistrict(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePincode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestFacilityAddress) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestFacilityAddress) validateDistrict(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"district", "body", m.District); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"district", "body", string(*m.District), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestFacilityAddress) validatePincode(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"pincode", "body", m.Pincode); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"pincode", "body", string(*m.Pincode), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestFacilityAddress) validateState(formats strfmt.Registry) error {

	if err := validate.Required("facility"+"."+"address"+"."+"state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinLength("facility"+"."+"address"+"."+"state", "body", string(*m.State), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestFacilityAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestFacilityAddress) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestFacilityAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestRecipient test certification request recipient
//
// swagger:model TestCertificationRequestRecipient
type TestCertificationRequestRecipient struct {

	// address
	Address *TestCertificationRequestRecipientAddress `json:"address,omitempty"`

	// contact
	// Required: true
	Contact []string `json:"contact"`

	// dob
	// Required: true
	// Format: date
	Dob *strfmt.Date `json:"dob"`

	// gender
	// Required: true
	// Min Length: 1
	Gender *string `json:"gender"`

	// identity
	// Required: true
	// Min Length: 1
	Identity *string `json:"identity"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// nationality
	// Min Length: 1
	Nationality string `json:"nationality,omitempty"`
}

// Validate validates this test certification request recipient
func (m *TestCertificationRequestRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationality(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestRecipient) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateContact(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"contact", "body", m.Contact); err != nil {
		return err
	}

	for i := 0; i < len(m.Contact); i++ {

		if err := validate.MinLength("recipient"+"."+"contact"+"."+strconv.Itoa(i), "body", string(m.Contact[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateDob(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"dob", "body", m.Dob); err != nil {
		return err
	}

	if err := validate.FormatOf("recipient"+"."+"dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"gender", "body", m.Gender); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"gender", "body", string(*m.Gender), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateIdentity(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"identity", "body", m.Identity); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"identity", "body", string(*m.Identity), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateName(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipient) validateNationality(formats strfmt.Registry) error {

	if swag.IsZero(m.Nationality) { // not required
		return nil
	}

	if err := validate.MinLength("recipient"+"."+"nationality", "body", string(m.Nationality), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestRecipient) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestRecipientAddress test certification request recipient address
//
// swagger:model TestCertificationRequestRecipientAddress
type TestCertificationRequestRecipientAddress struct {

	// address line1
	// Required: true
	AddressLine1 *string `json:"addressLine1"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	// Required: true
	// Min Length: 1
	District *string `json:"district"`

	// pincode
	// Required: true
	Pincode *string `json:"pincode"`

	// state
	// Required: true
	// Min Length: 1
	State *string `json:"state"`
}

// Validate validates this test certification request recipient address
func (m *TestCertificationRequestRecipientAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistrict(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePincode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestRecipientAddress) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipientAddress) validateDistrict(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"district", "body", m.District); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"district", "body", string(*m.District), 1); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipientAddress) validatePincode(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"pincode", "body", m.Pincode); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestRecipientAddress) validateState(formats strfmt.Registry) error {

	if err := validate.Required("recipient"+"."+"address"+"."+"state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinLength("recipient"+"."+"address"+"."+"state", "body", string(*m.State), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestRecipientAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestRecipientAddress) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestRecipientAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestTestDetails test certification request test details
//
// swagger:model TestCertificationRequestTestDetails
type TestCertificationRequestTestDetails struct {

	// batch
	Batch string `json:"batch,omitempty"`

	// disease
	// Required: true
	Disease *string `json:"disease"`

	// manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// result
	// Required: true
	// Enum: [Positive Negative Inconclusive Void]
	Result *string `json:"result"`

	// result timestamp
	// Required: true
	// Format: date-time
	ResultTimestamp *strfmt.DateTime `json:"resultTimestamp"`

	// sample collection timestamp
	// Required: true
	// Format: date-time
	SampleCollectionTimestamp *strfmt.DateTime `json:"sampleCollectionTimestamp"`

	// sample origin
	SampleOrigin string `json:"sampleOrigin,omitempty"`

	// test name
	TestName string `json:"testName,omitempty"`

	// test type
	// Required: true
	TestType *string `json:"testType"`
}

// Validate validates this test certification request test details
func (m *TestCertificationRequestTestDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSampleCollectionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestTestDetails) validateDisease(formats strfmt.Registry) error {

	if err := validate.Required("testDetails"+"."+"disease", "body", m.Disease); err != nil {
		return err
	}

	return nil
}

var testCertificationRequestTestDetailsTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Positive","Negative","Inconclusive","Void"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		testCertificationRequestTestDetailsTypeResultPropEnum = append(testCertificationRequestTestDetailsTypeResultPropEnum, v)
	}
}

const (

	// TestCertificationRequestTestDetailsResultPositive captures enum value "Positive"
	TestCertificationRequestTestDetailsResultPositive string = "Positive"

	// TestCertificationRequestTestDetailsResultNegative captures enum value "Negative"
	TestCertificationRequestTestDetailsResultNegative string = "Negative"

	// TestCertificationRequestTestDetailsResultInconclusive captures enum value "Inconclusive"
	TestCertificationRequestTestDetailsResultInconclusive string = "Inconclusive"

	// TestCertificationRequestTestDetailsResultVoid captures enum value "Void"
	TestCertificationRequestTestDetailsResultVoid string = "Void"
)

// prop value enum
func (m *TestCertificationRequestTestDetails) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, testCertificationRequestTestDetailsTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TestCertificationRequestTestDetails) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("testDetails"+"."+"result", "body", m.Result); err != nil {
		return err
	}

	// value enum
	if err := m.validateResultEnum("testDetails"+"."+"result", "body", *m.Result); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestTestDetails) validateResultTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("testDetails"+"."+"resultTimestamp", "body", m.ResultTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("testDetails"+"."+"resultTimestamp", "body", "date-time", m.ResultTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestTestDetails) validateSampleCollectionTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("testDetails"+"."+"sampleCollectionTimestamp", "body", m.SampleCollectionTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("testDetails"+"."+"sampleCollectionTimestamp", "body", "date-time", m.SampleCollectionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TestCertificationRequestTestDetails) validateTestType(formats strfmt.Registry) error {

	if err := validate.Required("testDetails"+"."+"testType", "body", m.TestType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestTestDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestTestDetails) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestTestDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TestCertificationRequestVerifier test certification request verifier
//
// swagger:model TestCertificationRequestVerifier
type TestCertificationRequestVerifier struct {

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this test certification request verifier
func (m *TestCertificationRequestVerifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestCertificationRequestVerifier) validateName(formats strfmt.Registry) error {

	if err := validate.Required("verifier"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("verifier"+"."+"name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCertificationRequestVerifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCertificationRequestVerifier) UnmarshalBinary(b []byte) error {
	var res TestCertificationRequestVerifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
