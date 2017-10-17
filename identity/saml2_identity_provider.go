// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"time"
)

// Saml2IdentityProvider. A special type of IdentityProvider that
// supports the SAML 2.0 protocol. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
type Saml2IdentityProvider struct {

	// The OCID of the `IdentityProvider`.
	ID string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentID string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation. The name
	// must be unique across all `IdentityProvider` objects in the tenancy and
	// cannot be changed. This is the name federated users see when choosing
	// which identity provider to use when signing in to the Oracle Bare Metal Cloud
	// Services Console.
	Name string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation. Does
	// not have to be unique, and it's changeable.
	Description string `mandatory:"true" json:"description,omitempty"`

	// The identity provider service or product (e.g., Oracle Identity Cloud Service).
	// Allowed value: `IDCS`.
	// Example: `IDCS`
	ProductType string `mandatory:"true" json:"productType,omitempty"`

	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `mandatory:"true" json:"timeCreated,omitempty"`

	// The current state. After creating an `IdentityProvider`, make sure its
	// `lifecycleState` changes from CREATING to ACTIVE before using it.
	LifecycleState string `mandatory:"true" json:"lifecycleState,omitempty"`

	// The protocol used for federation. Allowed value: `SAML2`.
	// Example: `SAML2`
	Protocol string `mandatory:"true" json:"protocol,omitempty"`

	// The URL for retrieving the identity provider's metadata, which
	// contains information required for federating.
	MetadataURL string `mandatory:"true" json:"metadataUrl,omitempty"`

	// The identity provider's signing certificate used by the IAM Service
	// to validate the SAML2 token.
	SigningCertificate string `mandatory:"true" json:"signingCertificate,omitempty"`

	// The URL to redirect federated users to for authentication with the
	// identity provider.
	RedirectURL string `mandatory:"true" json:"redirectUrl,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `mandatory:"false" json:"inactiveStatus,omitempty"`
}