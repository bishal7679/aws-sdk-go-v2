// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type TrustAnchorType string

// Enum values for TrustAnchorType
const (
	TrustAnchorTypeAwsAcmPca            TrustAnchorType = "AWS_ACM_PCA"
	TrustAnchorTypeCertificateBundle    TrustAnchorType = "CERTIFICATE_BUNDLE"
	TrustAnchorTypeSelfSignedRepository TrustAnchorType = "SELF_SIGNED_REPOSITORY"
)

// Values returns all known values for TrustAnchorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TrustAnchorType) Values() []TrustAnchorType {
	return []TrustAnchorType{
		"AWS_ACM_PCA",
		"CERTIFICATE_BUNDLE",
		"SELF_SIGNED_REPOSITORY",
	}
}