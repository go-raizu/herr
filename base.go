// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
// Copyright(c) 2022-2023 individual contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

package herr

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	code int
}

func (b HTTPError) StatusCode() int {
	return b.code
}

func (b HTTPError) Error() string {
	return fmt.Sprintf("%d (%s)", b.code, http.StatusText(b.code))
}

var (
	ErrBadRequest                   = HTTPError{http.StatusBadRequest}
	ErrUnauthorized                 = HTTPError{http.StatusUnauthorized}
	ErrPaymentRequired              = HTTPError{http.StatusPaymentRequired}
	ErrForbidden                    = HTTPError{http.StatusForbidden}
	ErrNotFound                     = HTTPError{http.StatusNotFound}
	ErrMethodNotAllowed             = HTTPError{http.StatusMethodNotAllowed}
	ErrNotAcceptable                = HTTPError{http.StatusNotAcceptable}
	ErrProxyAuthRequired            = HTTPError{http.StatusProxyAuthRequired}
	ErrRequestTimeout               = HTTPError{http.StatusRequestTimeout}
	ErrConflict                     = HTTPError{http.StatusConflict}
	ErrGone                         = HTTPError{http.StatusGone}
	ErrLengthRequired               = HTTPError{http.StatusLengthRequired}
	ErrPreconditionFailed           = HTTPError{http.StatusPreconditionFailed}
	ErrTooLarge                     = HTTPError{http.StatusRequestEntityTooLarge}
	ErrRequestURITooLong            = HTTPError{http.StatusRequestURITooLong}
	ErrUnsupportedMediaType         = HTTPError{http.StatusUnsupportedMediaType}
	ErrRequestedRangeNotSatisfiable = HTTPError{http.StatusRequestedRangeNotSatisfiable}
	ErrExpectationFailed            = HTTPError{http.StatusExpectationFailed}
	ErrTeapot                       = HTTPError{http.StatusTeapot}
	ErrMisdirectedRequest           = HTTPError{http.StatusMisdirectedRequest}
	ErrUnprocessableEntity          = HTTPError{http.StatusUnprocessableEntity}
	ErrLocked                       = HTTPError{http.StatusLocked}
	ErrFailedDependency             = HTTPError{http.StatusFailedDependency}
	ErrTooEarly                     = HTTPError{http.StatusTooEarly}
	ErrUpgradeRequired              = HTTPError{http.StatusUpgradeRequired}
	ErrPreconditionRequired         = HTTPError{http.StatusPreconditionRequired}
	ErrTooManyRequests              = HTTPError{http.StatusTooManyRequests}
	ErrRequestHeaderFieldsTooLarge  = HTTPError{http.StatusRequestHeaderFieldsTooLarge}
	ErrUnavailableForLegalReasons   = HTTPError{http.StatusUnavailableForLegalReasons}

	ErrInternalServerError           = HTTPError{http.StatusInternalServerError}
	ErrNotImplemented                = HTTPError{http.StatusNotImplemented}
	ErrBadGateway                    = HTTPError{http.StatusBadGateway}
	ErrServiceUnavailable            = HTTPError{http.StatusServiceUnavailable}
	ErrGatewayTimeout                = HTTPError{http.StatusGatewayTimeout}
	ErrHTTPVersionNotSupported       = HTTPError{http.StatusHTTPVersionNotSupported}
	ErrVariantAlsoNegotiates         = HTTPError{http.StatusVariantAlsoNegotiates}
	ErrInsufficientStorage           = HTTPError{http.StatusInsufficientStorage}
	ErrLoopDetected                  = HTTPError{http.StatusLoopDetected}
	ErrNotExtended                   = HTTPError{http.StatusNotExtended}
	ErrNetworkAuthenticationRequired = HTTPError{http.StatusNetworkAuthenticationRequired}
)
