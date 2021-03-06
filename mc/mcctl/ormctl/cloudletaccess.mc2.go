// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudletaccess.proto

package ormctl

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var IssueCertRequestRequiredArgs = []string{}
var IssueCertRequestOptionalArgs = []string{
	"commonname",
}
var IssueCertRequestAliasArgs = []string{
	"commonname=issuecertrequest.commonname",
}
var IssueCertRequestComments = map[string]string{
	"commonname": "Certificate common name",
}
var IssueCertRequestSpecialArgs = map[string]string{}
var GetCasRequestRequiredArgs = []string{}
var GetCasRequestOptionalArgs = []string{
	"issuer",
}
var GetCasRequestAliasArgs = []string{
	"issuer=getcasrequest.issuer",
}
var GetCasRequestComments = map[string]string{
	"issuer": "Issuer",
}
var GetCasRequestSpecialArgs = map[string]string{}
var UpgradeAccessKeyClientMsgRequiredArgs = []string{}
var UpgradeAccessKeyClientMsgOptionalArgs = []string{
	"msg",
	"verifyonly",
}
var UpgradeAccessKeyClientMsgAliasArgs = []string{
	"msg=upgradeaccesskeyclientmsg.msg",
	"verifyonly=upgradeaccesskeyclientmsg.verifyonly",
}
var UpgradeAccessKeyClientMsgComments = map[string]string{
	"msg":        "Message type",
	"verifyonly": "Client node type",
}
var UpgradeAccessKeyClientMsgSpecialArgs = map[string]string{}
var AccessDataRequestRequiredArgs = []string{}
var AccessDataRequestOptionalArgs = []string{
	"type",
	"data",
}
var AccessDataRequestAliasArgs = []string{
	"type=accessdatarequest.type",
	"data=accessdatarequest.data",
}
var AccessDataRequestComments = map[string]string{
	"type": "Data type",
	"data": "Any request data (type specific)",
}
var AccessDataRequestSpecialArgs = map[string]string{}
