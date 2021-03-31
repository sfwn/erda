package dbclient

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
	"github.com/erda-project/erda/modules/pipeline/spec"
	"github.com/erda-project/erda/pkg/strutil"
)

const (
	queueLabelKeyID               string = "__queue_id"
	queueLabelKeyName             string = "__queue_name"
	queueLabelKeyMaxCPU           string = "__queue_max_cpu"
	queueLabelKeyMaxMemoryMB      string = "__queue_max_memory_MB"
	queueLabelKeyClusterName      string = "__queue_cluster_name"
	queueLabelKeyScheduleStrategy string = "__queue_schedule_strategy"
	queueLabelKeyPriority         string = "__queue_priority"
)

// CreatePipelineQueue
func (client *Client) CreatePipelineQueue(req apistructs.PipelineQueueCreateRequest, ops ...SessionOption) (*apistructs.PipelineQueue, error) {
	session := client.NewSession(ops...)
	defer session.Close()

	// create label which represents queue for queue_id
	queueIDLabel := spec.PipelineLabel{
		Type:           apistructs.PipelineLabelTypeQueue,
		TargetID:       0,
		PipelineSource: req.PipelineSource,
		Key:            queueLabelKeyID,
		Value:          "see db id",
	}
	_, err := session.InsertOne(&queueIDLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to insert queue to db, pipelineSource: %s, name: %s, err: %v", req.PipelineSource, req.Name, err)
	}
	// update after insert
	queueIDLabel.TargetID = queueIDLabel.ID
	queueIDLabel.Value = strutil.String(queueIDLabel.ID)
	if _, err := session.ID(queueIDLabel.ID).Update(queueIDLabel); err != nil {
		return nil, fmt.Errorf("failed to update queue id label after insert, queueID: %d, err: %v", queueIDLabel.ID, err)
	}

	// create queue other fields
	queue, err := client.createPipelineQueueFields(req, queueIDLabel, ops...)
	if err != nil {
		return nil, err
	}

	return queue, nil
}

// createPipelineQueueFields create queue's other fields after queue id label created.
func (client *Client) createPipelineQueueFields(req apistructs.PipelineQueueCreateRequest, queueIDLabel spec.PipelineLabel, ops ...SessionOption) (*apistructs.PipelineQueue, error) {
	session := client.NewSession(ops...)
	defer session.Close()

	// store all fields as labels who's target_id is queue
	queueID := queueIDLabel.ID
	nameLabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyName, req.Name)
	clusterNameLabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyClusterName, req.ClusterName)
	scheduleStrategyLabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyScheduleStrategy, req.ScheduleStrategy.String())
	priorityLabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyPriority, strutil.String(req.Priority))
	maxCPULabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyMaxCPU, strutil.String(req.MaxCPU))
	maxMemoryMBLabel := genMetaLabelFunc(queueID, req.PipelineSource, queueLabelKeyMaxMemoryMB, strutil.String(req.MaxMemoryMB))
	queueMetaLabels := []spec.PipelineLabel{
		nameLabel,
		clusterNameLabel,
		scheduleStrategyLabel,
		priorityLabel,
		maxCPULabel,
		maxMemoryMBLabel,
	}
	for k, v := range req.Labels {
		queueMetaLabels = append(queueMetaLabels, genMetaLabelFunc(queueID, req.PipelineSource, k, v))
	}
	if _, err := session.InsertMulti(queueMetaLabels); err != nil {
		return nil, fmt.Errorf("failed to insert queue fields to db, queueID: %d, err: %v", queueID, err)
	}

	// construct queue
	return constructQueueByLabels(append(queueMetaLabels, queueIDLabel))
}

var genMetaLabelFunc = func(queueID uint64, source apistructs.PipelineSource, key, value string) spec.PipelineLabel {
	return spec.PipelineLabel{
		Type:           apistructs.PipelineLabelTypeQueue,
		TargetID:       queueID,
		PipelineSource: source,
		Key:            key,
		Value:          value,
	}
}

func constructQueueByLabels(labels []spec.PipelineLabel) (*apistructs.PipelineQueue, error) {
	var q apistructs.PipelineQueue

	// parse id first
	for _, label := range labels {
		if label.Key != queueLabelKeyID {
			continue
		}
		q.ID = label.ID
		q.PipelineSource = label.PipelineSource
		q.TimeCreated = &label.TimeCreated
	}
	if q.ID == 0 {
		return nil, fmt.Errorf("failed to construct queue, not found key id")
	}

	// parse other fields
	for _, label := range labels {
		switch label.Key {
		case queueLabelKeyID:
			continue
		case queueLabelKeyName:
			q.Name = label.Value
		case queueLabelKeyClusterName:
			q.ClusterName = label.Value
		case queueLabelKeyScheduleStrategy:
			q.ScheduleStrategy = apistructs.ScheduleStrategyInsidePipelineQueue(label.Value)
		case queueLabelKeyPriority:
			priority, err := strconv.ParseInt(label.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to construct queue for priority, queueID: %d, value: %s, err: %v", q.ID, label.Value, err)
			}
			q.Priority = priority
		case queueLabelKeyMaxCPU:
			maxCPU, err := strconv.ParseFloat(label.Value, 10)
			if err != nil {
				return nil, fmt.Errorf("failed to construct queue for maxCPU, queueID: %d, value: %s, err: %v", q.ID, label.Value, err)
			}
			q.MaxCPU = maxCPU
		case queueLabelKeyMaxMemoryMB:
			maxMemoryMB, err := strconv.ParseFloat(label.Value, 10)
			if err != nil {
				return nil, fmt.Errorf("failed to construct queue for maxMemoryMB, queueID: %d, value: %s, err: %v", q.ID, label.Value, err)
			}
			q.MaxMemoryMB = maxMemoryMB

		default:
			// other labels
			if q.Labels == nil {
				q.Labels = make(map[string]string)
			}
			q.Labels[label.Key] = label.Value
		}
	}

	// get time updated
	var timeUpdated time.Time
	for _, label := range labels {
		if timeUpdated.IsZero() {
			timeUpdated = label.TimeUpdated
			continue
		}
		if label.TimeUpdated.After(timeUpdated) {
			timeUpdated = label.TimeUpdated
		}
	}
	q.TimeUpdated = &timeUpdated

	return &q, nil
}

// GetPipelineQueue
func (client *Client) GetPipelineQueue(queueID uint64, ops ...SessionOption) (*apistructs.PipelineQueue, bool, error) {
	session := client.NewSession(ops...)
	defer session.Close()

	// queue id
	var idLabel spec.PipelineLabel
	exist, err := session.ID(queueID).Get(&idLabel)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get idLabel of queue, id: %d, err: %v", queueID, err)
	}
	if !exist {
		return nil, false, nil
	}

	// queue other fields
	var fieldLabels []spec.PipelineLabel
	cond := spec.PipelineLabel{
		Type:     apistructs.PipelineLabelTypeQueue,
		TargetID: idLabel.ID,
	}
	if err := session.Find(&fieldLabels, &cond); err != nil {
		return nil, false, fmt.Errorf("failed to get fieldlabels of queue, id: %d, err: %v", queueID, err)
	}

	queue, err := constructQueueByLabels(append(fieldLabels, idLabel))
	return queue, true, err
}

// transferQueuePagingRequestToMustMatchLabels transfer field to label key, just like: name => __queue_name
func transferQueuePagingRequestToMustMatchLabels(req *apistructs.PipelineQueuePagingRequest) {
	genMustMatchLabelFunc := func(k, v string) string { return fmt.Sprintf("%s=%s", k, v) }
	// name
	if req.Name != "" {
		req.MustMatchLabels = append(req.MustMatchLabels, genMustMatchLabelFunc(queueLabelKeyName, req.Name))
	}
	// clusterName
	if req.ClusterName != "" {
		req.MustMatchLabels = append(req.MustMatchLabels, genMustMatchLabelFunc(queueLabelKeyClusterName, req.ClusterName))
	}
	// scheduleStrategy
	if req.ScheduleStrategy != "" {
		req.MustMatchLabels = append(req.MustMatchLabels, genMustMatchLabelFunc(queueLabelKeyScheduleStrategy, req.ScheduleStrategy.String()))
	}
	// priority
	if req.Priority != 0 {
		req.MustMatchLabels = append(req.MustMatchLabels, genMustMatchLabelFunc(queueLabelKeyPriority, strutil.String(req.Priority)))
	}
}

func (client *Client) PagingPipelineQueues(req apistructs.PipelineQueuePagingRequest, ops ...SessionOption) (*apistructs.PipelineQueuePagingData, error) {
	session := client.NewSession(ops...)
	defer session.Close()

	// set default
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	// transfer req meta to labels
	transferQueuePagingRequestToMustMatchLabels(&req)

	// validate
	if !req.AllowNoPipelineSources && len(req.PipelineSources) == 0 {
		return nil, apierrors.ErrPagingPipelineQueues.InvalidParameter("missing pipeline sources")
	}
	mustMatchLabelMap, err := transferMustMatchLabelsToMap(req.MustMatchLabels)
	if err != nil {
		return nil, apierrors.ErrPagingPipelineQueues.InvalidParameter(fmt.Errorf("invalid mustMatchLabels, err: %v", err))
	}
	anyMatchLabelMap, err := transferMustMatchLabelsToMap(req.AnyMatchLabels)
	if err != nil {
		return nil, apierrors.ErrPagingPipelineQueues.InvalidParameter(fmt.Errorf("invalid anyMatchLabels, err: %v", err))
	}

	// get all queueIDs
	labelRequest := apistructs.TargetIDSelectByLabelRequest{
		Type:                   apistructs.PipelineLabelTypeQueue,
		PipelineSources:        req.PipelineSources,
		AllowNoMatchLabels:     true,
		MustMatchLabels:        mustMatchLabelMap,
		AnyMatchLabels:         anyMatchLabelMap,
		AllowNoPipelineSources: req.AllowNoPipelineSources,
		OrderByTargetIDAsc:     req.OrderByTargetIDAsc,
	}
	queueIDs, err := client.SelectTargetIDsByLabels(labelRequest)
	if err != nil {
		return nil, apierrors.ErrPagingPipelineQueues.InternalError(err)
	}

	// paging queueIDs
	queueIDs = filterAndOrder(queueIDs)
	pagingQueueIDs := paging(queueIDs, req.PageNo, req.PageSize)
	total := int64(len(queueIDs))

	// list queue details
	labelMap, err := client.ListPipelineLabelsByTypeAndTargetIDs(apistructs.PipelineLabelTypeQueue, pagingQueueIDs)
	if err != nil {
		return nil, apierrors.ErrPagingPipelineQueues.InternalError(fmt.Errorf("failed to list queue details by ids, err: %v", err))
	}
	var queues []*apistructs.PipelineQueue
	for _, queueID := range pagingQueueIDs {
		queue, err := constructQueueByLabels(labelMap[queueID])
		if err != nil {
			return nil, apierrors.ErrPagingPipelineQueues.InternalError(fmt.Errorf("failed to construct queue, err: %v", err))
		}
		queues = append(queues, queue)
	}

	pagingResult := apistructs.PipelineQueuePagingData{
		Queues: queues,
		Total:  total,
	}

	return &pagingResult, nil
}

func transferMustMatchLabelsToMap(ss []string) (map[string][]string, error) {
	result := make(map[string][]string)
	for _, s := range ss {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid label format(need: k=v): %s", s)
		}
		result[kv[0]] = strutil.DedupSlice(append(result[kv[0]], kv[1]))
	}
	return result, nil
}

// UpdatePipelineQueue
func (client *Client) UpdatePipelineQueue(req apistructs.PipelineQueueUpdateRequest, ops ...SessionOption) (*apistructs.PipelineQueue, error) {
	session := client.NewSession(ops...)
	defer session.Close()

	// query queue id label
	var queueIDLabel spec.PipelineLabel
	exist, err := session.ID(req.ID).Get(&queueIDLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue id label, err: %v", err)
	}
	if !exist {
		return nil, fmt.Errorf("queue not found")
	}

	// delete old fields except __queue_id
	_, err = session.Where("`target_id` = ?", req.ID).Where("`key` != ?", queueLabelKeyID).Delete(&spec.PipelineLabel{})
	if err != nil {
		return nil, fmt.Errorf("failed to delete queue fields, queueID: %d, err: %v", req.ID, err)
	}

	// insert new one
	updatedQueue, err := client.createPipelineQueueFields(req.PipelineQueueCreateRequest, queueIDLabel, ops...)
	if err != nil {
		return nil, fmt.Errorf("failed to update queue fields, queueID: %d, err: %v", req.ID, err)
	}

	return updatedQueue, nil
}

// DeletePipelineQueue
func (client *Client) DeletePipelineQueue(queueID uint64, ops ...SessionOption) error {
	session := client.NewSession(ops...)
	defer session.Close()

	_, err := session.ID(queueID).Delete(&spec.PipelineLabel{})
	if err != nil {
		return fmt.Errorf("failed to delete queue, queueID: %d, err: %v", queueID, err)
	}

	return nil
}
