// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// ListRegionsRequest wrapper for the ListRegions operation
type ListRegionsRequest struct {
}

// ListRegionsResponse wrapper for the ListRegions operation
type ListRegionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []Region instance
	Items []Region `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}
