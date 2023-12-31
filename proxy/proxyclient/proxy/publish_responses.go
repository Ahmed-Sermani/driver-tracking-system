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

// PublishReader is a Reader for the Publish structure.
type PublishReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublishReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublishOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublishDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublishOK creates a PublishOK with default headers values
func NewPublishOK() *PublishOK {
	return &PublishOK{}
}

/*
PublishOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublishOK struct {
	Payload *models.CentrifugoProxyPublishResponse
}

// IsSuccess returns true when this publish o k response has a 2xx status code
func (o *PublishOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this publish o k response has a 3xx status code
func (o *PublishOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this publish o k response has a 4xx status code
func (o *PublishOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this publish o k response has a 5xx status code
func (o *PublishOK) IsServerError() bool {
	return false
}

// IsCode returns true when this publish o k response a status code equal to that given
func (o *PublishOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the publish o k response
func (o *PublishOK) Code() int {
	return 200
}

func (o *PublishOK) Error() string {
	return fmt.Sprintf("[POST /api/proxy/publish][%d] publishOK  %+v", 200, o.Payload)
}

func (o *PublishOK) String() string {
	return fmt.Sprintf("[POST /api/proxy/publish][%d] publishOK  %+v", 200, o.Payload)
}

func (o *PublishOK) GetPayload() *models.CentrifugoProxyPublishResponse {
	return o.Payload
}

func (o *PublishOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CentrifugoProxyPublishResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublishDefault creates a PublishDefault with default headers values
func NewPublishDefault(code int) *PublishDefault {
	return &PublishDefault{
		_statusCode: code,
	}
}

/*
PublishDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublishDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this publish default response has a 2xx status code
func (o *PublishDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this publish default response has a 3xx status code
func (o *PublishDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this publish default response has a 4xx status code
func (o *PublishDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this publish default response has a 5xx status code
func (o *PublishDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this publish default response a status code equal to that given
func (o *PublishDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the publish default response
func (o *PublishDefault) Code() int {
	return o._statusCode
}

func (o *PublishDefault) Error() string {
	return fmt.Sprintf("[POST /api/proxy/publish][%d] publish default  %+v", o._statusCode, o.Payload)
}

func (o *PublishDefault) String() string {
	return fmt.Sprintf("[POST /api/proxy/publish][%d] publish default  %+v", o._statusCode, o.Payload)
}

func (o *PublishDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublishDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
