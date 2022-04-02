// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cicdcms.proto

package pb

import (
	fmt "fmt"
	math "math"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CICDCmsUpdateRequest) Validate() error {
	for _, item := range this.Configs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Configs", err)
			}
		}
	}
	return nil
}
func (this *CICDCmsUpdateResponse) Validate() error {
	return nil
}
func (this *CICDCmsCreateRequest) Validate() error {
	for _, item := range this.Configs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Configs", err)
			}
		}
	}
	return nil
}
func (this *CICDCmsCreateResponse) Validate() error {
	return nil
}
func (this *CICDCmsDeleteRequest) Validate() error {
	return nil
}
func (this *CICDCmsDeleteResponse) Validate() error {
	return nil
}
func (this *Config) Validate() error {
	return nil
}
