package activedirectory

import (
	"github.com/lkarlslund/adalanche/modules/engine"
)

var (
	PwnForeignIdentity = engine.NewPwn("ForeignIdentity")

	DistinguishedName          = engine.NewAttribute("distinguishedName").Tag("AD")
	ObjectClass                = engine.NewAttribute("objectClass").Tag("AD")
	ObjectCategory             = engine.NewAttribute("objectCategory").Tag("AD")
	ObjectCategorySimple       = engine.NewAttribute("objectCategorySimple")
	StructuralObjectClass      = engine.NewAttribute("structuralObjectClass").Tag("AD")
	NTSecurityDescriptor       = engine.NewAttribute("nTSecurityDescriptor").Tag("AD")
	SAMAccountType             = engine.NewAttribute("sAMAccountType").Tag("AD")
	GroupType                  = engine.NewAttribute("groupType").Tag("AD")
	MemberOf                   = engine.NewAttribute("memberOf").Tag("AD")
	Member                     = engine.NewAttribute("member").Tag("AD")
	AccountExpires             = engine.NewAttribute("accountExpires").Tag("AD")
	RepsTo                     = engine.NewAttribute("repsTo").Tag("AD")
	InstanceType               = engine.NewAttribute("instanceType").Tag("AD")
	ModifiedCount              = engine.NewAttribute("modifiedCount").Tag("AD")
	MinPwdAge                  = engine.NewAttribute("minPwdAge").Tag("AD")
	MinPwdLength               = engine.NewAttribute("minPwdLength").Tag("AD")
	PwdProperties              = engine.NewAttribute("pwdProperties").Tag("AD")
	LockOutDuration            = engine.NewAttribute("lockoutDuration")
	PwdHistoryLength           = engine.NewAttribute("pwdHistoryLength")
	IsCriticalSystemObject     = engine.NewAttribute("isCriticalSystemObject").Tag("AD")
	FSMORoleOwner              = engine.NewAttribute("fSMORoleOwner")
	NTMixedDomain              = engine.NewAttribute("nTMixedDomain")
	SystemFlags                = engine.NewAttribute("systemFlags")
	PrimaryGroupID             = engine.NewAttribute("primaryGroupID").Tag("AD")
	LogonCount                 = engine.NewAttribute("logonCount")
	UserAccountControl         = engine.NewAttribute("userAccountControl").Tag("AD")
	LocalPolicyFlags           = engine.NewAttribute("localPolicyFlags")
	CodePage                   = engine.NewAttribute("codePage")
	CountryCode                = engine.NewAttribute("countryCode")
	OperatingSystem            = engine.NewAttribute("operatingSystem")
	OperatingSystemHotfix      = engine.NewAttribute("operatingSystemHotfix")
	OperatingSystemVersion     = engine.NewAttribute("operatingSystemVersion")
	OperatingSystemServicePack = engine.NewAttribute("operatingSystemServicePack")
	AdminCount                 = engine.NewAttribute("adminCount").Tag("AD")
	LogonHours                 = engine.NewAttribute("logonHours")
	BadPwdCount                = engine.NewAttribute("badPwdCount")
	GPCFileSysPath             = engine.NewAttribute("gPCFileSysPath").Tag("AD").Merge()
	SchemaIDGUID               = engine.NewAttribute("schemaIDGUID").Tag("AD")
	PossSuperiors              = engine.NewAttribute("possSuperiors")
	SystemMayContain           = engine.NewAttribute("systemMayContain")
	SystemMustContain          = engine.NewAttribute("systemMustContain")
	ServicePrincipalName       = engine.NewAttribute("servicePrincipalName").Tag("AD")
	Name                       = engine.NewAttribute("name").Tag("AD")
	DisplayName                = engine.NewAttribute("displayName").Tag("AD")
	LDAPDisplayName            = engine.NewAttribute("lDAPDisplayName").Tag("AD") // Attribute-Schema
	Description                = engine.NewAttribute("description").Tag("AD")
	SAMAccountName             = engine.NewAttribute("sAMAccountName").Tag("AD")
	ObjectSid                  = engine.NewAttribute("objectSid").Tag("AD").NonUnique().Merge()

	ObjectGUID                  = engine.NewAttribute("objectGUID").Tag("AD").Merge()
	PwdLastSet                  = engine.NewAttribute("pwdLastSet").Tag("AD")
	WhenCreated                 = engine.NewAttribute("whenCreated")
	WhenChanged                 = engine.NewAttribute("whenChanged")
	DsCorePropagationData       = engine.NewAttribute("dsCorePropagationData")
	MsExchLastUpdateTime        = engine.NewAttribute("msExchLastUpdateTime")
	GWARTLastModified           = engine.NewAttribute("gWARTLastModified")
	SpaceLastComputed           = engine.NewAttribute("spaceLastComputed")
	MsExchPolicyLastAppliedTime = engine.NewAttribute("msExchPolicyLastAppliedTime")
	MsExchWhenMailboxCreated    = engine.NewAttribute("msExchWhenMailboxCreated")
	SIDHistory                  = engine.NewAttribute("sIDHistory").Tag("AD")
	LastLogon                   = engine.NewAttribute("lastLogon")
	LastLogonTimestamp          = engine.NewAttribute("lastLogonTimestamp")
	MSDSGroupMSAMembership      = engine.NewAttribute("msDS-GroupMSAMembership").Tag("AD")
	MSDSHostServiceAccount      = engine.NewAttribute("msDS-HostServiceAccount").Tag("AD")
	MSDSHostServiceAccountBL    = engine.NewAttribute("msDS-HostServiceAccountBL").Tag("AD")
	MSmcsAdmPwdExpirationTime   = engine.NewAttribute("ms-mcs-AdmPwdExpirationTime").Tag("AD") // LAPS password timeout
	SecurityIdentifier          = engine.NewAttribute("securityIdentifier")
	TrustDirection              = engine.NewAttribute("trustDirection")
	TrustAttributes             = engine.NewAttribute("trustAttributes")
	TrustPartner                = engine.NewAttribute("trustPartner")
	DsHeuristics                = engine.NewAttribute("dsHeuristics").Tag("AD")
	AttributeSecurityGUID       = engine.NewAttribute("attributeSecurityGUID").Tag("AD")
	MSDSConsistencyGUID         = engine.NewAttribute("mS-DS-ConsistencyGuid")
	RightsGUID                  = engine.NewAttribute("rightsGUID").Tag("AD")
	GPLink                      = engine.NewAttribute("gPLink").Tag("AD")
	GPOptions                   = engine.NewAttribute("gPOptions").Tag("AD")
	ScriptPath                  = engine.NewAttribute("scriptPath").Tag("AD")
	MSPKICertificateNameFlag    = engine.NewAttribute("msPKI-Certificate-Name-Flag").Tag("AD")
	PKIExtendedUsage            = engine.NewAttribute("pKIExtendedKeyUsage").Tag("AD")
)

func init() {
	engine.AddMergeApprover("Don't merge differing distinguishedName", func(a, b *engine.Object) (result *engine.Object, err error) {
		if a.HasAttr(engine.DistinguishedName) && b.HasAttr(engine.DistinguishedName) && a.DN() != b.DN() {
			// Yes, it happens we have objects that have different DNs but the same merge attributes (objectSID, etc)
			return nil, engine.ErrDontMerge
		}
		return nil, nil
	})
}
