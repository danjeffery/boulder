// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "fmt"

const _FeatureFlag_name = "unusedAllowAccountDeactivationAllowKeyRolloverResubmitMissingSCTsOnlyUseAIAIssuerURLAllowTLS02ChallengesGenerateOCSPEarlyReusePendingAuthzCountCertificatesExactRandomDirectoryEntryIPv6FirstDirectoryMetaAllowRenewalFirstRLRecheckCAALegacyCAAUDPDNSROCACheckWildcardDomains"

var _FeatureFlag_index = [...]uint16{0, 6, 30, 46, 69, 84, 104, 121, 138, 160, 180, 189, 202, 221, 231, 240, 246, 255, 270}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return fmt.Sprintf("FeatureFlag(%d)", i)
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
