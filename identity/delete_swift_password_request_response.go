// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// DeleteSwiftPasswordRequest wrapper for the DeleteSwiftPassword operation
type DeleteSwiftPasswordRequest struct {

	// The OCID of the user.
	UserID *string `mandatory:"true" contributesTo:"path" name:"userId"`

	// The OCID of the Swift password.
	SwiftPasswordID *string `mandatory:"true" contributesTo:"path" name:"swiftPasswordId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

// DeleteSwiftPasswordResponse wrapper for the DeleteSwiftPassword operation
type DeleteSwiftPasswordResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}
