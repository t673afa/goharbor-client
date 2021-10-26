// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// GetGCScheduleReader is a Reader for the GetGCSchedule structure.
type GetGCScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGCScheduleOK creates a GetGCScheduleOK with default headers values
func NewGetGCScheduleOK() *GetGCScheduleOK {
	return &GetGCScheduleOK{}
}

/*GetGCScheduleOK handles this case with default header values.

Get gc's schedule.
*/
type GetGCScheduleOK struct {
	Payload *model.GCHistory
}

func (o *GetGCScheduleOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGCScheduleOK  %+v", 200, o.Payload)
}

func (o *GetGCScheduleOK) GetPayload() *model.GCHistory {
	return o.Payload
}

func (o *GetGCScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GCHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCScheduleUnauthorized creates a GetGCScheduleUnauthorized with default headers values
func NewGetGCScheduleUnauthorized() *GetGCScheduleUnauthorized {
	return &GetGCScheduleUnauthorized{}
}

/*GetGCScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type GetGCScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGCScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCScheduleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCScheduleForbidden creates a GetGCScheduleForbidden with default headers values
func NewGetGCScheduleForbidden() *GetGCScheduleForbidden {
	return &GetGCScheduleForbidden{}
}

/*GetGCScheduleForbidden handles this case with default header values.

Forbidden
*/
type GetGCScheduleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGCScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetGCScheduleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCScheduleInternalServerError creates a GetGCScheduleInternalServerError with default headers values
func NewGetGCScheduleInternalServerError() *GetGCScheduleInternalServerError {
	return &GetGCScheduleInternalServerError{}
}

/*GetGCScheduleInternalServerError handles this case with default header values.

Internal server error
*/
type GetGCScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/schedule][%d] getGCScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCScheduleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
