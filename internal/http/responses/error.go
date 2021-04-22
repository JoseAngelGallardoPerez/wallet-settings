package responses

import (
	"net/http"
)

const (
	TargetCommon = "common"

	CannotRetrieveSetting    = "CANNOT_RETRIEVE_SETTING"
	CanNotRetrieveCollection = "CANNOT_RETRIEVE_COLLECTION"
	CanNotUpdateCollection   = "CANNOT_UPDATE_COLLECTION"
	SettingsNotFound         = "SETTINGS_NOT_FOUND"
	UnprocessableEntity      = "UNPROCESSABLE_ENTITY"
	NotFound                 = "NOT_FOUND"
	CodeForbidden            = "FORBIDDEN"
)

var statusCodes = map[string]int{
	CannotRetrieveSetting:    http.StatusBadRequest,
	CanNotRetrieveCollection: http.StatusBadRequest,
	CanNotUpdateCollection:   http.StatusInternalServerError,
	SettingsNotFound:         http.StatusNotFound,
	UnprocessableEntity:      http.StatusUnprocessableEntity,
	CodeForbidden:            http.StatusForbidden,
}

// Error is the abstract error model
type Error struct {
	Title   string      `json:"title"`
	Details string      `json:"details"`
	Code    string      `json:"code"`
	Source  string      `json:"sourсe,omitempty"`
	Target  string      `json:"target"`
	Meta    interface{} `json:"meta,omitempty"`
}

// NewError creates a new Error instance.
func NewError() *Error {
	return new(Error)
}

// SetTitle sets the given title for the error object.
func (e *Error) SetTitle(title string) *Error {
	e.Title = title
	return e
}

// SetTitleByStatus sets the title for the error object by status code.
func (e *Error) SetTitleByStatus(status int) *Error {
	e.Title = http.StatusText(status)
	return e
}

// SetTitleByCode sets the title for the error object by code.
func (e *Error) SetTitleByCode(code string) *Error {
	if status, ok := statusCodes[code]; ok {
		e.SetTitleByStatus(status)
	}
	return e
}

// SetDetails sets the given details for the error object.
func (e *Error) SetDetails(details string) *Error {
	e.Details = details
	return e
}

// SetCode sets the code for the error object.
func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

// SetSourсe sets the sourсe for the error object.
func (e *Error) SetSource(source string) *Error {
	e.Source = source
	return e
}

// SetTarget sets the sourсe for the error object.
func (e *Error) SetTarget(target string) *Error {
	e.Target = target
	return e
}
