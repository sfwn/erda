// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: guide.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ListGuideRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListGuideResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Guide)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GittarPushPayloadEvent)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Content)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Pusher)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateGuideResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ProcessGuideRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ProcessGuideResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteGuideResponse)(nil)

// ListGuideRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListGuideRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "kind":
				m.Kind = vals[0]
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			}
		}
	}
	return nil
}

// ListGuideResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListGuideResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Guide implement urlenc.URLValuesUnmarshaler.
func (m *Guide) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				m.ID = vals[0]
			case "status":
				m.Status = vals[0]
			case "creator":
				m.Creator = vals[0]
			case "kind":
				m.Kind = vals[0]
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrgID = val
			case "orgName":
				m.OrgName = vals[0]
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "appID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.AppID = val
			case "branch":
				m.Branch = vals[0]
			case "appName":
				m.AppName = vals[0]
			case "timeCreated":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
			case "timeCreated.seconds":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeCreated.Seconds = val
			case "timeCreated.nanos":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeCreated.Nanos = int32(val)
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			}
		}
	}
	return nil
}

// GittarPushPayloadEvent implement urlenc.URLValuesUnmarshaler.
func (m *GittarPushPayloadEvent) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "event":
				m.Event = vals[0]
			case "action":
				m.Action = vals[0]
			case "orgID":
				m.OrgID = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "applicationID":
				m.ApplicationID = vals[0]
			case "env":
				m.Env = vals[0]
			case "timeStamp":
				m.TimeStamp = vals[0]
			case "content":
				if m.Content == nil {
					m.Content = &Content{}
				}
			case "content.ref":
				if m.Content == nil {
					m.Content = &Content{}
				}
				m.Content.Ref = vals[0]
			case "content.after":
				if m.Content == nil {
					m.Content = &Content{}
				}
				m.Content.After = vals[0]
			case "content.before":
				if m.Content == nil {
					m.Content = &Content{}
				}
				m.Content.Before = vals[0]
			case "content.pusher":
				if m.Content == nil {
					m.Content = &Content{}
				}
				if m.Content.Pusher == nil {
					m.Content.Pusher = &Pusher{}
				}
			case "content.pusher.ID":
				if m.Content == nil {
					m.Content = &Content{}
				}
				if m.Content.Pusher == nil {
					m.Content.Pusher = &Pusher{}
				}
				m.Content.Pusher.ID = vals[0]
			case "content.pusher.name":
				if m.Content == nil {
					m.Content = &Content{}
				}
				if m.Content.Pusher == nil {
					m.Content.Pusher = &Pusher{}
				}
				m.Content.Pusher.Name = vals[0]
			case "content.pusher.nickName":
				if m.Content == nil {
					m.Content = &Content{}
				}
				if m.Content.Pusher == nil {
					m.Content.Pusher = &Pusher{}
				}
				m.Content.Pusher.NickName = vals[0]
			case "content.pusher.email":
				if m.Content == nil {
					m.Content = &Content{}
				}
				if m.Content.Pusher == nil {
					m.Content.Pusher = &Pusher{}
				}
				m.Content.Pusher.Email = vals[0]
			}
		}
	}
	return nil
}

// Content implement urlenc.URLValuesUnmarshaler.
func (m *Content) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ref":
				m.Ref = vals[0]
			case "after":
				m.After = vals[0]
			case "before":
				m.Before = vals[0]
			case "pusher":
				if m.Pusher == nil {
					m.Pusher = &Pusher{}
				}
			case "pusher.ID":
				if m.Pusher == nil {
					m.Pusher = &Pusher{}
				}
				m.Pusher.ID = vals[0]
			case "pusher.name":
				if m.Pusher == nil {
					m.Pusher = &Pusher{}
				}
				m.Pusher.Name = vals[0]
			case "pusher.nickName":
				if m.Pusher == nil {
					m.Pusher = &Pusher{}
				}
				m.Pusher.NickName = vals[0]
			case "pusher.email":
				if m.Pusher == nil {
					m.Pusher = &Pusher{}
				}
				m.Pusher.Email = vals[0]
			}
		}
	}
	return nil
}

// Pusher implement urlenc.URLValuesUnmarshaler.
func (m *Pusher) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				m.ID = vals[0]
			case "name":
				m.Name = vals[0]
			case "nickName":
				m.NickName = vals[0]
			case "email":
				m.Email = vals[0]
			}
		}
	}
	return nil
}

// CreateGuideResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateGuideResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Guide{}
				}
			case "data.ID":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.ID = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.Status = vals[0]
			case "data.creator":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.Creator = vals[0]
			case "data.kind":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.Kind = vals[0]
			case "data.orgID":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.OrgID = val
			case "data.orgName":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.OrgName = vals[0]
			case "data.projectID":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ProjectID = val
			case "data.appID":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.AppID = val
			case "data.branch":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.Branch = vals[0]
			case "data.appName":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				m.Data.AppName = vals[0]
			case "data.timeCreated":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
			case "data.timeCreated.seconds":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Seconds = val
			case "data.timeCreated.nanos":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeCreated == nil {
					m.Data.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeCreated.Nanos = int32(val)
			case "data.timeUpdated":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "data.timeUpdated.seconds":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Seconds = val
			case "data.timeUpdated.nanos":
				if m.Data == nil {
					m.Data = &Guide{}
				}
				if m.Data.TimeUpdated == nil {
					m.Data.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.TimeUpdated.Nanos = int32(val)
			}
		}
	}
	return nil
}

// ProcessGuideRequest implement urlenc.URLValuesUnmarshaler.
func (m *ProcessGuideRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "appID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.AppID = val
			case "branch":
				m.Branch = vals[0]
			case "kind":
				m.Kind = vals[0]
			}
		}
	}
	return nil
}

// ProcessGuideResponse implement urlenc.URLValuesUnmarshaler.
func (m *ProcessGuideResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DeleteGuideResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteGuideResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
