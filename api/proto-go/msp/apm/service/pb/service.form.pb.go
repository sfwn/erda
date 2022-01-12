// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: service.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetServiceLanguageResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceLanguageRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceOverviewTopRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceOverviewTopResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceCountRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceCountResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceAnalyzerOverviewRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServiceAnalyzerOverviewResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServicesView)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServicesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetServicesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TopTable)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Chart)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServiceTop)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServiceChart)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Service)(nil)
var _ urlenc.URLValuesUnmarshaler = (*AggregateMetric)(nil)

// GetServiceLanguageResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceLanguageResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "language":
				m.Language = vals[0]
			}
		}
	}
	return nil
}

// GetServiceLanguageRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceLanguageRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantId":
				m.TenantId = vals[0]
			case "serviceId":
				m.ServiceId = vals[0]
			}
		}
	}
	return nil
}

// GetServiceOverviewTopRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceOverviewTopRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantId":
				m.TenantId = vals[0]
			}
		}
	}
	return nil
}

// GetServiceOverviewTopResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceOverviewTopResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetServiceCountRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceCountRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantId":
				m.TenantId = vals[0]
			case "status":
				m.Status = vals[0]
			}
		}
	}
	return nil
}

// GetServiceCountResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceCountResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "totalCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TotalCount = val
			case "hasErrorCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.HasErrorCount = val
			case "withoutRequestCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.WithoutRequestCount = val
			}
		}
	}
	return nil
}

// GetServiceAnalyzerOverviewRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceAnalyzerOverviewRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantId":
				m.TenantId = vals[0]
			case "serviceIds":
				m.ServiceIds = vals
			case "view":
				m.View = vals[0]
			case "startTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTime = val
			case "endTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.EndTime = val
			}
		}
	}
	return nil
}

// GetServiceAnalyzerOverviewResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetServiceAnalyzerOverviewResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ServicesView implement urlenc.URLValuesUnmarshaler.
func (m *ServicesView) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "serviceId":
				m.ServiceId = vals[0]
			}
		}
	}
	return nil
}

// GetServicesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetServicesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "tenantId":
				m.TenantId = vals[0]
			case "serviceName":
				m.ServiceName = vals[0]
			case "serviceStatus":
				m.ServiceStatus = vals[0]
			}
		}
	}
	return nil
}

// GetServicesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetServicesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Total = val
			}
		}
	}
	return nil
}

// TopTable implement urlenc.URLValuesUnmarshaler.
func (m *TopTable) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "serviceId":
				m.ServiceId = vals[0]
			case "serviceName":
				m.ServiceName = vals[0]
			case "value":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Value = val
			}
		}
	}
	return nil
}

// Chart implement urlenc.URLValuesUnmarshaler.
func (m *Chart) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "timestamp":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Timestamp = val
			case "value":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Value = val
			case "dimension":
				m.Dimension = vals[0]
			case "extraValues":
				list := make([]float64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseFloat(text, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.ExtraValues = list
			}
		}
	}
	return nil
}

// ServiceTop implement urlenc.URLValuesUnmarshaler.
func (m *ServiceTop) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// ServiceChart implement urlenc.URLValuesUnmarshaler.
func (m *ServiceChart) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			case "maxValue":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.MaxValue = val
			}
		}
	}
	return nil
}

// Service implement urlenc.URLValuesUnmarshaler.
func (m *Service) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			case "language":
			case "lastHeartbeat":
				m.LastHeartbeat = vals[0]
			case "aggregateMetric":
				if m.AggregateMetric == nil {
					m.AggregateMetric = &AggregateMetric{}
				}
			case "aggregateMetric.avgRps":
				if m.AggregateMetric == nil {
					m.AggregateMetric = &AggregateMetric{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.AggregateMetric.AvgRps = val
			case "aggregateMetric.avgDuration":
				if m.AggregateMetric == nil {
					m.AggregateMetric = &AggregateMetric{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.AggregateMetric.AvgDuration = val
			case "aggregateMetric.errorRate":
				if m.AggregateMetric == nil {
					m.AggregateMetric = &AggregateMetric{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.AggregateMetric.ErrorRate = val
			}
		}
	}
	return nil
}

// AggregateMetric implement urlenc.URLValuesUnmarshaler.
func (m *AggregateMetric) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "avgRps":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.AvgRps = val
			case "avgDuration":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.AvgDuration = val
			case "errorRate":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.ErrorRate = val
			}
		}
	}
	return nil
}
