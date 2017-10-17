// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// GetIdpGroupMappingRequest wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingRequest struct {

	// The OCID of the identity provider.
	IdentityProviderID string `mandatory:"true" contributesTo:"path" name:"identityProviderId"`

	// The OCID of the group mapping.
	MappingID string `mandatory:"true" contributesTo:"path" name:"mappingId"`
}

// GetIdpGroupMappingResponse wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingResponse struct {

	// The IdpGroupMapping instance
	IdpGroupMapping `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
