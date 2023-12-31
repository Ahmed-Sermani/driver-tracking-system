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

// RPCReader is a Reader for the RPC structure.
type RPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRPCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRPCOK creates a RPCOK with default headers values
func NewRPCOK() *RPCOK {
	return &RPCOK{}
}

/*
RPCOK describes a response with status code 200, with default header values.

A successful response.
*/
type RPCOK struct {
	Payload *models.CentrifugoProxyRPCResponse
}

// IsSuccess returns true when this rpc o k response has a 2xx status code
func (o *RPCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rpc o k response has a 3xx status code
func (o *RPCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rpc o k response has a 4xx status code
func (o *RPCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rpc o k response has a 5xx status code
func (o *RPCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rpc o k response a status code equal to that given
func (o *RPCOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rpc o k response
func (o *RPCOK) Code() int {
	return 200
}

func (o *RPCOK) Error() string {
	return fmt.Sprintf("[POST /api/proxy/rpc][%d] rpcOK  %+v", 200, o.Payload)
}

func (o *RPCOK) String() string {
	return fmt.Sprintf("[POST /api/proxy/rpc][%d] rpcOK  %+v", 200, o.Payload)
}

func (o *RPCOK) GetPayload() *models.CentrifugoProxyRPCResponse {
	return o.Payload
}

func (o *RPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CentrifugoProxyRPCResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRPCDefault creates a RPCDefault with default headers values
func NewRPCDefault(code int) *RPCDefault {
	return &RPCDefault{
		_statusCode: code,
	}
}

/*
RPCDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RPCDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this rpc default response has a 2xx status code
func (o *RPCDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rpc default response has a 3xx status code
func (o *RPCDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rpc default response has a 4xx status code
func (o *RPCDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rpc default response has a 5xx status code
func (o *RPCDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rpc default response a status code equal to that given
func (o *RPCDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rpc default response
func (o *RPCDefault) Code() int {
	return o._statusCode
}

func (o *RPCDefault) Error() string {
	return fmt.Sprintf("[POST /api/proxy/rpc][%d] rpc default  %+v", o._statusCode, o.Payload)
}

func (o *RPCDefault) String() string {
	return fmt.Sprintf("[POST /api/proxy/rpc][%d] rpc default  %+v", o._statusCode, o.Payload)
}

func (o *RPCDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RPCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
