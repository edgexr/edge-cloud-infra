// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trustpolicy.proto

package ormctl

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
	"github.com/mobiledgex/edge-cloud/cli"
	edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	math "math"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var CreateTrustPolicyCmd = &cli.Command{
	Use:                  "CreateTrustPolicy",
	RequiredArgs:         "region " + strings.Join(TrustPolicyRequiredArgs, " "),
	OptionalArgs:         strings.Join(TrustPolicyOptionalArgs, " "),
	AliasArgs:            strings.Join(TrustPolicyAliasArgs, " "),
	SpecialArgs:          &TrustPolicySpecialArgs,
	Comments:             addRegionComment(TrustPolicyComments),
	ReqData:              &ormapi.RegionTrustPolicy{},
	ReplyData:            &edgeproto.Result{},
	Run:                  runRest("/auth/ctrl/CreateTrustPolicy"),
	StreamOut:            true,
	StreamOutIncremental: true,
}

var DeleteTrustPolicyCmd = &cli.Command{
	Use:                  "DeleteTrustPolicy",
	RequiredArgs:         "region " + strings.Join(TrustPolicyRequiredArgs, " "),
	OptionalArgs:         strings.Join(TrustPolicyOptionalArgs, " "),
	AliasArgs:            strings.Join(TrustPolicyAliasArgs, " "),
	SpecialArgs:          &TrustPolicySpecialArgs,
	Comments:             addRegionComment(TrustPolicyComments),
	ReqData:              &ormapi.RegionTrustPolicy{},
	ReplyData:            &edgeproto.Result{},
	Run:                  runRest("/auth/ctrl/DeleteTrustPolicy"),
	StreamOut:            true,
	StreamOutIncremental: true,
}

var UpdateTrustPolicyCmd = &cli.Command{
	Use:          "UpdateTrustPolicy",
	RequiredArgs: "region " + strings.Join(TrustPolicyRequiredArgs, " "),
	OptionalArgs: strings.Join(TrustPolicyOptionalArgs, " "),
	AliasArgs:    strings.Join(TrustPolicyAliasArgs, " "),
	SpecialArgs:  &TrustPolicySpecialArgs,
	Comments:     addRegionComment(TrustPolicyComments),
	ReqData:      &ormapi.RegionTrustPolicy{},
	ReplyData:    &edgeproto.Result{},
	Run: runRest("/auth/ctrl/UpdateTrustPolicy",
		withSetFieldsFunc(setUpdateTrustPolicyFields),
	),
	StreamOut:            true,
	StreamOutIncremental: true,
}

func setUpdateTrustPolicyFields(in map[string]interface{}) {
	// get map for edgeproto object in region struct
	obj := in[strings.ToLower("TrustPolicy")]
	if obj == nil {
		return
	}
	objmap, ok := obj.(map[string]interface{})
	if !ok {
		return
	}
	fields := cli.GetSpecifiedFields(objmap, &edgeproto.TrustPolicy{}, cli.JsonNamespace)
	// include fields already specified
	if inFields, found := objmap["fields"]; found {
		if fieldsArr, ok := inFields.([]string); ok {
			fields = append(fields, fieldsArr...)
		}
	}
	objmap["fields"] = fields
}

var ShowTrustPolicyCmd = &cli.Command{
	Use:          "ShowTrustPolicy",
	RequiredArgs: "region",
	OptionalArgs: strings.Join(append(TrustPolicyRequiredArgs, TrustPolicyOptionalArgs...), " "),
	AliasArgs:    strings.Join(TrustPolicyAliasArgs, " "),
	SpecialArgs:  &TrustPolicySpecialArgs,
	Comments:     addRegionComment(TrustPolicyComments),
	ReqData:      &ormapi.RegionTrustPolicy{},
	ReplyData:    &edgeproto.TrustPolicy{},
	Run:          runRest("/auth/ctrl/ShowTrustPolicy"),
	StreamOut:    true,
}

var TrustPolicyApiCmds = []*cli.Command{
	CreateTrustPolicyCmd,
	DeleteTrustPolicyCmd,
	UpdateTrustPolicyCmd,
	ShowTrustPolicyCmd,
}

var SecurityRuleRequiredArgs = []string{}
var SecurityRuleOptionalArgs = []string{
	"protocol",
	"portrangemin",
	"portrangemax",
	"remotecidr",
}
var SecurityRuleAliasArgs = []string{
	"protocol=securityrule.protocol",
	"portrangemin=securityrule.portrangemin",
	"portrangemax=securityrule.portrangemax",
	"remotecidr=securityrule.remotecidr",
}
var SecurityRuleComments = map[string]string{
	"protocol":     "tcp, udp, icmp",
	"portrangemin": "TCP or UDP port range start",
	"portrangemax": "TCP or UDP port range end",
	"remotecidr":   "remote CIDR X.X.X.X/X",
}
var SecurityRuleSpecialArgs = map[string]string{}
var TrustPolicyRequiredArgs = []string{
	"cloudlet-org",
	"name",
}
var TrustPolicyOptionalArgs = []string{
	"outboundsecurityrules:#.protocol",
	"outboundsecurityrules:#.portrangemin",
	"outboundsecurityrules:#.portrangemax",
	"outboundsecurityrules:#.remotecidr",
}
var TrustPolicyAliasArgs = []string{
	"fields=trustpolicy.fields",
	"cloudlet-org=trustpolicy.key.organization",
	"name=trustpolicy.key.name",
	"outboundsecurityrules:#.protocol=trustpolicy.outboundsecurityrules:#.protocol",
	"outboundsecurityrules:#.portrangemin=trustpolicy.outboundsecurityrules:#.portrangemin",
	"outboundsecurityrules:#.portrangemax=trustpolicy.outboundsecurityrules:#.portrangemax",
	"outboundsecurityrules:#.remotecidr=trustpolicy.outboundsecurityrules:#.remotecidr",
}
var TrustPolicyComments = map[string]string{
	"fields":                               "Fields are used for the Update API to specify which fields to apply",
	"cloudlet-org":                         "Name of the organization for the cluster that this policy will apply to",
	"name":                                 "Policy name",
	"outboundsecurityrules:#.protocol":     "tcp, udp, icmp",
	"outboundsecurityrules:#.portrangemin": "TCP or UDP port range start",
	"outboundsecurityrules:#.portrangemax": "TCP or UDP port range end",
	"outboundsecurityrules:#.remotecidr":   "remote CIDR X.X.X.X/X",
}
var TrustPolicySpecialArgs = map[string]string{
	"trustpolicy.fields": "StringArray",
}