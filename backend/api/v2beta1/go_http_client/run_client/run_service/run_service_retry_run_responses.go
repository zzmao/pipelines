// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/run_model"
)

// RunServiceRetryRunReader is a Reader for the RunServiceRetryRun structure.
type RunServiceRetryRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunServiceRetryRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunServiceRetryRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRunServiceRetryRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunServiceRetryRunOK creates a RunServiceRetryRunOK with default headers values
func NewRunServiceRetryRunOK() *RunServiceRetryRunOK {
	return &RunServiceRetryRunOK{}
}

/*RunServiceRetryRunOK handles this case with default header values.

A successful response.
*/
type RunServiceRetryRunOK struct {
	Payload interface{}
}

func (o *RunServiceRetryRunOK) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/runs/{run_id}:retry][%d] runServiceRetryRunOK  %+v", 200, o.Payload)
}

func (o *RunServiceRetryRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunServiceRetryRunDefault creates a RunServiceRetryRunDefault with default headers values
func NewRunServiceRetryRunDefault(code int) *RunServiceRetryRunDefault {
	return &RunServiceRetryRunDefault{
		_statusCode: code,
	}
}

/*RunServiceRetryRunDefault handles this case with default header values.

An unexpected error response.
*/
type RunServiceRetryRunDefault struct {
	_statusCode int

	Payload *run_model.RuntimeError
}

// Code gets the status code for the run service retry run default response
func (o *RunServiceRetryRunDefault) Code() int {
	return o._statusCode
}

func (o *RunServiceRetryRunDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/runs/{run_id}:retry][%d] RunService_RetryRun default  %+v", o._statusCode, o.Payload)
}

func (o *RunServiceRetryRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}