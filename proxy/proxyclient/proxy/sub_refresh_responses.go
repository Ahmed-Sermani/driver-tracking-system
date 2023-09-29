// Code generated by go-swagger; DO NOT EDIT.

package proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Ahmed-Sermani/driver-tracking-system/proxy/proxyclient/models"
)

// SubRefreshReader is a Reader for the SubRefresh structure.
type SubRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubRefreshDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubRefreshOK creates a SubRefreshOK with default headers values
func NewSubRefreshOK() *SubRefreshOK {
	return &SubRefreshOK{}
}

/*
SubRefreshOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubRefreshOK struct {
	Payload *models.CentrifugoProxySubRefreshResponse
}

// IsSuccess returns true when this sub refresh o k response has a 2xx status code
func (o *SubRefreshOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sub refresh o k response has a 3xx status code
func (o *SubRefreshOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sub refresh o k response has a 4xx status code
func (o *SubRefreshOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sub refresh o k response has a 5xx status code
func (o *SubRefreshOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sub refresh o k response a status code equal to that given
func (o *SubRefreshOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sub refresh o k response
func (o *SubRefreshOK) Code() int {
	return 200
}

func (o *SubRefreshOK) Error() string {
	return fmt.Sprintf("[POST /api/proxy/sub-refresh][%d] subRefreshOK  %+v", 200, o.Payload)
}

func (o *SubRefreshOK) String() string {
	return fmt.Sprintf("[POST /api/proxy/sub-refresh][%d] subRefreshOK  %+v", 200, o.Payload)
}

func (o *SubRefreshOK) GetPayload() *models.CentrifugoProxySubRefreshResponse {
	return o.Payload
}

func (o *SubRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CentrifugoProxySubRefreshResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubRefreshDefault creates a SubRefreshDefault with default headers values
func NewSubRefreshDefault(code int) *SubRefreshDefault {
	return &SubRefreshDefault{
		_statusCode: code,
	}
}

/*
SubRefreshDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SubRefreshDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this sub refresh default response has a 2xx status code
func (o *SubRefreshDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sub refresh default response has a 3xx status code
func (o *SubRefreshDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sub refresh default response has a 4xx status code
func (o *SubRefreshDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sub refresh default response has a 5xx status code
func (o *SubRefreshDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sub refresh default response a status code equal to that given
func (o *SubRefreshDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sub refresh default response
func (o *SubRefreshDefault) Code() int {
	return o._statusCode
}

func (o *SubRefreshDefault) Error() string {
	return fmt.Sprintf("[POST /api/proxy/sub-refresh][%d] sub-refresh default  %+v", o._statusCode, o.Payload)
}

func (o *SubRefreshDefault) String() string {
	return fmt.Sprintf("[POST /api/proxy/sub-refresh][%d] sub-refresh default  %+v", o._statusCode, o.Payload)
}

func (o *SubRefreshDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *SubRefreshDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}