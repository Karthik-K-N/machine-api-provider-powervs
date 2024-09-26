// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1HostGroupsIDPutReader is a Reader for the V1HostGroupsIDPut structure.
type V1HostGroupsIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1HostGroupsIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1HostGroupsIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1HostGroupsIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1HostGroupsIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1HostGroupsIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1HostGroupsIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1HostGroupsIDPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1HostGroupsIDPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewV1HostGroupsIDPutGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/host-groups/{host_group_id}] v1.hostGroups.id.put", response, response.Code())
	}
}

// NewV1HostGroupsIDPutOK creates a V1HostGroupsIDPutOK with default headers values
func NewV1HostGroupsIDPutOK() *V1HostGroupsIDPutOK {
	return &V1HostGroupsIDPutOK{}
}

/*
V1HostGroupsIDPutOK describes a response with status code 200, with default header values.

OK
*/
type V1HostGroupsIDPutOK struct {
	Payload *models.HostGroup
}

// IsSuccess returns true when this v1 host groups Id put o k response has a 2xx status code
func (o *V1HostGroupsIDPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 host groups Id put o k response has a 3xx status code
func (o *V1HostGroupsIDPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put o k response has a 4xx status code
func (o *V1HostGroupsIDPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups Id put o k response has a 5xx status code
func (o *V1HostGroupsIDPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put o k response a status code equal to that given
func (o *V1HostGroupsIDPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 host groups Id put o k response
func (o *V1HostGroupsIDPutOK) Code() int {
	return 200
}

func (o *V1HostGroupsIDPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutOK %s", 200, payload)
}

func (o *V1HostGroupsIDPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutOK %s", 200, payload)
}

func (o *V1HostGroupsIDPutOK) GetPayload() *models.HostGroup {
	return o.Payload
}

func (o *V1HostGroupsIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutBadRequest creates a V1HostGroupsIDPutBadRequest with default headers values
func NewV1HostGroupsIDPutBadRequest() *V1HostGroupsIDPutBadRequest {
	return &V1HostGroupsIDPutBadRequest{}
}

/*
V1HostGroupsIDPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1HostGroupsIDPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put bad request response has a 2xx status code
func (o *V1HostGroupsIDPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put bad request response has a 3xx status code
func (o *V1HostGroupsIDPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put bad request response has a 4xx status code
func (o *V1HostGroupsIDPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups Id put bad request response has a 5xx status code
func (o *V1HostGroupsIDPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put bad request response a status code equal to that given
func (o *V1HostGroupsIDPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 host groups Id put bad request response
func (o *V1HostGroupsIDPutBadRequest) Code() int {
	return 400
}

func (o *V1HostGroupsIDPutBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutBadRequest %s", 400, payload)
}

func (o *V1HostGroupsIDPutBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutBadRequest %s", 400, payload)
}

func (o *V1HostGroupsIDPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutUnauthorized creates a V1HostGroupsIDPutUnauthorized with default headers values
func NewV1HostGroupsIDPutUnauthorized() *V1HostGroupsIDPutUnauthorized {
	return &V1HostGroupsIDPutUnauthorized{}
}

/*
V1HostGroupsIDPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1HostGroupsIDPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put unauthorized response has a 2xx status code
func (o *V1HostGroupsIDPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put unauthorized response has a 3xx status code
func (o *V1HostGroupsIDPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put unauthorized response has a 4xx status code
func (o *V1HostGroupsIDPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups Id put unauthorized response has a 5xx status code
func (o *V1HostGroupsIDPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put unauthorized response a status code equal to that given
func (o *V1HostGroupsIDPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 host groups Id put unauthorized response
func (o *V1HostGroupsIDPutUnauthorized) Code() int {
	return 401
}

func (o *V1HostGroupsIDPutUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutUnauthorized %s", 401, payload)
}

func (o *V1HostGroupsIDPutUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutUnauthorized %s", 401, payload)
}

func (o *V1HostGroupsIDPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutForbidden creates a V1HostGroupsIDPutForbidden with default headers values
func NewV1HostGroupsIDPutForbidden() *V1HostGroupsIDPutForbidden {
	return &V1HostGroupsIDPutForbidden{}
}

/*
V1HostGroupsIDPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1HostGroupsIDPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put forbidden response has a 2xx status code
func (o *V1HostGroupsIDPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put forbidden response has a 3xx status code
func (o *V1HostGroupsIDPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put forbidden response has a 4xx status code
func (o *V1HostGroupsIDPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups Id put forbidden response has a 5xx status code
func (o *V1HostGroupsIDPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put forbidden response a status code equal to that given
func (o *V1HostGroupsIDPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 host groups Id put forbidden response
func (o *V1HostGroupsIDPutForbidden) Code() int {
	return 403
}

func (o *V1HostGroupsIDPutForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutForbidden %s", 403, payload)
}

func (o *V1HostGroupsIDPutForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutForbidden %s", 403, payload)
}

func (o *V1HostGroupsIDPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutNotFound creates a V1HostGroupsIDPutNotFound with default headers values
func NewV1HostGroupsIDPutNotFound() *V1HostGroupsIDPutNotFound {
	return &V1HostGroupsIDPutNotFound{}
}

/*
V1HostGroupsIDPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1HostGroupsIDPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put not found response has a 2xx status code
func (o *V1HostGroupsIDPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put not found response has a 3xx status code
func (o *V1HostGroupsIDPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put not found response has a 4xx status code
func (o *V1HostGroupsIDPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups Id put not found response has a 5xx status code
func (o *V1HostGroupsIDPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put not found response a status code equal to that given
func (o *V1HostGroupsIDPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 host groups Id put not found response
func (o *V1HostGroupsIDPutNotFound) Code() int {
	return 404
}

func (o *V1HostGroupsIDPutNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutNotFound %s", 404, payload)
}

func (o *V1HostGroupsIDPutNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutNotFound %s", 404, payload)
}

func (o *V1HostGroupsIDPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutConflict creates a V1HostGroupsIDPutConflict with default headers values
func NewV1HostGroupsIDPutConflict() *V1HostGroupsIDPutConflict {
	return &V1HostGroupsIDPutConflict{}
}

/*
V1HostGroupsIDPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1HostGroupsIDPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put conflict response has a 2xx status code
func (o *V1HostGroupsIDPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put conflict response has a 3xx status code
func (o *V1HostGroupsIDPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put conflict response has a 4xx status code
func (o *V1HostGroupsIDPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups Id put conflict response has a 5xx status code
func (o *V1HostGroupsIDPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups Id put conflict response a status code equal to that given
func (o *V1HostGroupsIDPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 host groups Id put conflict response
func (o *V1HostGroupsIDPutConflict) Code() int {
	return 409
}

func (o *V1HostGroupsIDPutConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutConflict %s", 409, payload)
}

func (o *V1HostGroupsIDPutConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutConflict %s", 409, payload)
}

func (o *V1HostGroupsIDPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutInternalServerError creates a V1HostGroupsIDPutInternalServerError with default headers values
func NewV1HostGroupsIDPutInternalServerError() *V1HostGroupsIDPutInternalServerError {
	return &V1HostGroupsIDPutInternalServerError{}
}

/*
V1HostGroupsIDPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1HostGroupsIDPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put internal server error response has a 2xx status code
func (o *V1HostGroupsIDPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put internal server error response has a 3xx status code
func (o *V1HostGroupsIDPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put internal server error response has a 4xx status code
func (o *V1HostGroupsIDPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups Id put internal server error response has a 5xx status code
func (o *V1HostGroupsIDPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 host groups Id put internal server error response a status code equal to that given
func (o *V1HostGroupsIDPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 host groups Id put internal server error response
func (o *V1HostGroupsIDPutInternalServerError) Code() int {
	return 500
}

func (o *V1HostGroupsIDPutInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutInternalServerError %s", 500, payload)
}

func (o *V1HostGroupsIDPutInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutInternalServerError %s", 500, payload)
}

func (o *V1HostGroupsIDPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsIDPutGatewayTimeout creates a V1HostGroupsIDPutGatewayTimeout with default headers values
func NewV1HostGroupsIDPutGatewayTimeout() *V1HostGroupsIDPutGatewayTimeout {
	return &V1HostGroupsIDPutGatewayTimeout{}
}

/*
V1HostGroupsIDPutGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type V1HostGroupsIDPutGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups Id put gateway timeout response has a 2xx status code
func (o *V1HostGroupsIDPutGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups Id put gateway timeout response has a 3xx status code
func (o *V1HostGroupsIDPutGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups Id put gateway timeout response has a 4xx status code
func (o *V1HostGroupsIDPutGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups Id put gateway timeout response has a 5xx status code
func (o *V1HostGroupsIDPutGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 host groups Id put gateway timeout response a status code equal to that given
func (o *V1HostGroupsIDPutGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the v1 host groups Id put gateway timeout response
func (o *V1HostGroupsIDPutGatewayTimeout) Code() int {
	return 504
}

func (o *V1HostGroupsIDPutGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutGatewayTimeout %s", 504, payload)
}

func (o *V1HostGroupsIDPutGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/host-groups/{host_group_id}][%d] v1HostGroupsIdPutGatewayTimeout %s", 504, payload)
}

func (o *V1HostGroupsIDPutGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsIDPutGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}