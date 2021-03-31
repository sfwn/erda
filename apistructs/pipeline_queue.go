package apistructs

import (
	"fmt"
	"time"

	"github.com/erda-project/erda/pkg/strutil"
)

type PipelineQueue struct {
	ID uint64 `json:"id"`

	Name             string                              `json:"name"`
	PipelineSource   PipelineSource                      `json:"pipelineSource"`
	ClusterName      string                              `json:"clusterName"`
	ScheduleStrategy ScheduleStrategyInsidePipelineQueue `json:"scheduleStrategy"`
	Priority         int64                               `json:"priority"`
	MaxCPU           float64                             `json:"maxCPU"`
	MaxMemoryMB      float64                             `json:"maxMemoryMB"`

	Labels map[string]string `json:"labels,omitempty"`

	TimeCreated *time.Time `json:"timeCreated,omitempty"`
	TimeUpdated *time.Time `json:"timeUpdated,omitempty"`
}

// ScheduleStrategyInsidePipelineQueue represents the schedule strategy of workflows inside a queue.
type ScheduleStrategyInsidePipelineQueue string

var (
	ScheduleStrategyInsidePipelineQueueOfFIFO ScheduleStrategyInsidePipelineQueue = "FIFO"
)

func (strategy ScheduleStrategyInsidePipelineQueue) String() string {
	return string(strategy)
}

func (strategy ScheduleStrategyInsidePipelineQueue) IsValid() bool {
	switch strategy {
	case ScheduleStrategyInsidePipelineQueueOfFIFO:
		return true
	default:
		return false
	}
}

var (
	PipelineQueueDefaultPriority         int64 = 10
	PipelineQueueDefaultScheduleStrategy       = ScheduleStrategyInsidePipelineQueueOfFIFO
)

// PipelineQueueCreateRequest represents queue create request.
type PipelineQueueCreateRequest struct {

	// Name is the queue name.
	// +required
	Name string `json:"name,omitempty"`

	// PipelineSource group queues by source.
	// +required
	PipelineSource PipelineSource `json:"pipelineSource,omitempty"`

	// ClusterName represents which cluster this queue belongs to.
	// +required
	ClusterName string `json:"clusterName,omitempty"`

	// ScheduleStrategy defines schedule strategy.
	// If not present, will use default strategy.
	// +optional
	ScheduleStrategy ScheduleStrategyInsidePipelineQueue `json:"scheduleStrategy,omitempty"`

	// Priority defines priority between queues.
	// Higher number means higher priority.
	// If not present, will use default priority.
	// +optional
	Priority int64 `json:"priority,omitempty"`

	// MaxCPU is the cpu resource this queue holds.
	MaxCPU float64 `json:"maxCPU,omitempty"`

	// MaxMemoryMB is the memory resource this queue holds.
	MaxMemoryMB float64 `json:"maxMemoryMB,omitempty"`

	// Labels contains the other infos for this queue.
	// Labels can be used to query and filter queues.
	Labels map[string]string `json:"labels,omitempty"`

	IdentityInfo
}

// Validate validate and handle request.
func (req *PipelineQueueCreateRequest) Validate() error {
	// name
	if err := strutil.Validate(req.Name, strutil.MinLenValidator(1), strutil.MaxRuneCountValidator(191)); err != nil {
		return fmt.Errorf("invalid name: %v", err)
	}
	// source
	if !req.PipelineSource.Valid() {
		return fmt.Errorf("invalid source: %s", req.PipelineSource)
	}
	// clusterName
	if req.ClusterName == "" {
		return fmt.Errorf("empty clusterName: %s", req.ClusterName)
	}
	// strategy
	if req.ScheduleStrategy == "" {
		req.ScheduleStrategy = PipelineQueueDefaultScheduleStrategy
	}
	if !req.ScheduleStrategy.IsValid() {
		return fmt.Errorf("invalid schedule strategy: %s", req.ScheduleStrategy)
	}
	// priority
	if req.Priority == 0 {
		req.Priority = PipelineQueueDefaultPriority
	}
	if req.Priority < 0 {
		return fmt.Errorf("priority must > 0")
	}
	// max cpu
	if req.MaxCPU < 0 {
		return fmt.Errorf("max cpu must >= 0")
	}
	// max memoryMB
	if req.MaxMemoryMB < 0 {
		return fmt.Errorf("max memory(MB) must >= 0")
	}
	return nil
}

// PipelineQueuePagingRequest
type PipelineQueuePagingRequest struct {
	Name string `schema:"name"`

	PipelineSources []PipelineSource `schema:"pipelineSource"`

	ClusterName string `schema:"clusterName"`

	ScheduleStrategy ScheduleStrategyInsidePipelineQueue `schema:"scheduleStrategy"`

	Priority int64 `schema:"priority"`

	// MUST match
	MustMatchLabels []string `schema:"mustMatchLabel"`
	// ANY match
	AnyMatchLabels []string `schema:"anyMatchLabel"`

	// AllowNoPipelineSources, default is false.
	// 默认查询必须带上 pipeline source，增加区分度
	AllowNoPipelineSources bool `schema:"allowNoPipelineSources"`

	// OrderByTargetIDAsc 根据 target_id 升序，默认为 false，即降序
	OrderByTargetIDAsc bool `schema:"orderByTargetIDAsc"`

	PageNo   int `schema:"pageNo"`
	PageSize int `schema:"pageSize"`
}

// PipelineQueuePagingData
type PipelineQueuePagingData struct {
	Queues []*PipelineQueue `json:"queues"`
	Total  int64            `json:"total"`
}

// PipelineQueueUpdateRequest
type PipelineQueueUpdateRequest struct {
	ID uint64 `json:"-"` // get from path variable

	// create request include all fields can be updated
	PipelineQueueCreateRequest
}

// Validate request.
func (req *PipelineQueueUpdateRequest) Validate() error {
	// id
	if req.ID == 0 {
		return fmt.Errorf("missing queue id")
	}
	// pipeline source
	if req.PipelineSource != "" {
		return fmt.Errorf("cannot change queue's source")
	}

	return nil
}
