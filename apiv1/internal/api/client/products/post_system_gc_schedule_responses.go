// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSystemGcScheduleReader is a Reader for the PostSystemGcSchedule structure.
type PostSystemGcScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSystemGcScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSystemGcScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSystemGcScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSystemGcScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSystemGcScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSystemGcScheduleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSystemGcScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSystemGcScheduleOK creates a PostSystemGcScheduleOK with default headers values
func NewPostSystemGcScheduleOK() *PostSystemGcScheduleOK {
	return &PostSystemGcScheduleOK{}
}

/*PostSystemGcScheduleOK handles this case with default header values.

GC schedule successfully.
*/
type PostSystemGcScheduleOK struct {
}

func (o *PostSystemGcScheduleOK) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleOK ", 200)
}

func (o *PostSystemGcScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemGcScheduleBadRequest creates a PostSystemGcScheduleBadRequest with default headers values
func NewPostSystemGcScheduleBadRequest() *PostSystemGcScheduleBadRequest {
	return &PostSystemGcScheduleBadRequest{}
}

/*PostSystemGcScheduleBadRequest handles this case with default header values.

Invalid schedule type.
*/
type PostSystemGcScheduleBadRequest struct {
}

func (o *PostSystemGcScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleBadRequest ", 400)
}

func (o *PostSystemGcScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemGcScheduleUnauthorized creates a PostSystemGcScheduleUnauthorized with default headers values
func NewPostSystemGcScheduleUnauthorized() *PostSystemGcScheduleUnauthorized {
	return &PostSystemGcScheduleUnauthorized{}
}

/*PostSystemGcScheduleUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostSystemGcScheduleUnauthorized struct {
}

func (o *PostSystemGcScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleUnauthorized ", 401)
}

func (o *PostSystemGcScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemGcScheduleForbidden creates a PostSystemGcScheduleForbidden with default headers values
func NewPostSystemGcScheduleForbidden() *PostSystemGcScheduleForbidden {
	return &PostSystemGcScheduleForbidden{}
}

/*PostSystemGcScheduleForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type PostSystemGcScheduleForbidden struct {
}

func (o *PostSystemGcScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleForbidden ", 403)
}

func (o *PostSystemGcScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemGcScheduleConflict creates a PostSystemGcScheduleConflict with default headers values
func NewPostSystemGcScheduleConflict() *PostSystemGcScheduleConflict {
	return &PostSystemGcScheduleConflict{}
}

/*PostSystemGcScheduleConflict handles this case with default header values.

There is a "gc" job in progress, so the request cannot be served.
*/
type PostSystemGcScheduleConflict struct {
}

func (o *PostSystemGcScheduleConflict) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleConflict ", 409)
}

func (o *PostSystemGcScheduleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemGcScheduleInternalServerError creates a PostSystemGcScheduleInternalServerError with default headers values
func NewPostSystemGcScheduleInternalServerError() *PostSystemGcScheduleInternalServerError {
	return &PostSystemGcScheduleInternalServerError{}
}

/*PostSystemGcScheduleInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PostSystemGcScheduleInternalServerError struct {
}

func (o *PostSystemGcScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/gc/schedule][%d] postSystemGcScheduleInternalServerError ", 500)
}

func (o *PostSystemGcScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
