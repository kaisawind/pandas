//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

// Package api contains implementation of lbs service HTTP API.

package api

import (
	"context"
	"fmt"
	"time"

	"github.com/cloustone/pandas/lbs"
	lbp "github.com/cloustone/pandas/lbs/proxy"
	log "github.com/cloustone/pandas/pkg/logger"
)

var _ lbs.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger log.Logger
	svc    lbs.Service
}

// LoggingMiddleware adds logging facilities to the core service.
func LoggingMiddleware(svc lbs.Service, logger log.Logger) lbs.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) ListCollections(ctx context.Context, token string) (products []string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method list_collections for token %s took %s to complete", token, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.ListCollections(ctx, token)
}

// Geofence
func (lm *loggingMiddleware) CreateCircleGeofence(ctx context.Context, token string, projectId string, fence *lbs.CircleGeofence) (fenceId string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method create_circle_geofence for token %s and project %s and fanceName %s took %s to complete", token, projectId, fence.Name, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.CreateCircleGeofence(ctx, token, projectId, fence)
}

func (lm *loggingMiddleware) UpdateCircleGeofence(ctx context.Context, token string, projectId string, fence *lbs.CircleGeofence) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method update_circle_geofence for token %s and project %s and fanceName %s took %s to complete", token, projectId, fence.Name, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.UpdateCircleGeofence(ctx, token, projectId, fence)
}

func (lm *loggingMiddleware) DeleteGeofence(ctx context.Context, token string, projectId string, fenceIds []string, objects []string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method delete_geofence for token %s and project %s and fanceIds %s took %s to complete", token, projectId, fenceIds, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.DeleteGeofence(ctx, token, projectId, fenceIds, objects)
}

func (lm *loggingMiddleware) ListGeofences(ctx context.Context, token string, projectId string, fenceIds []string, objects []string) (fenceList []*lbs.Geofence, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method list_geofences for token %s and project %s and fanceIds %s took %s to complete", token, projectId, fenceIds, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.ListGeofences(ctx, token, projectId, fenceIds, objects)
}

func (lm *loggingMiddleware) AddMonitoredObject(ctx context.Context, token string, projectId string, fenceId string, objects []string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method list_geofences for token %s and project %s and fanceId %s took %s to complete", token, projectId, fenceId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.AddMonitoredObject(ctx, token, projectId, fenceId, objects)
}

func (lm *loggingMiddleware) RemoveMonitoredObject(ctx context.Context, token string, projectId string, fenceId string, objects []string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method remove_monitored_object for token %s and project %s and fanceId %s took %s to complete", token, projectId, fenceId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.RemoveMonitoredObject(ctx, token, projectId, fenceId, objects)
}

func (lm *loggingMiddleware) ListMonitoredObjects(ctx context.Context, token string, projectId string, fenceId string, pageIndex int32, pageSize int32) (total int32, objects []string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method list_monitored_objects for token %s and project %s and fanceId %s and pageIndex %d and pageSize %d took %s to complete", token, projectId, fenceId, pageIndex, pageSize, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.ListMonitoredObjects(ctx, token, projectId, fenceId, pageIndex, pageSize)
}

func (lm *loggingMiddleware) CreatePolyGeofence(ctx context.Context, token string, projectId string, fence *lbs.PolyGeofence) (fenceId string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method create_poly_geofence for token %s and project %s and fanceName %s took %s to complete", token, projectId, fence.Name, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.CreatePolyGeofence(ctx, token, projectId, fence)
}

func (lm *loggingMiddleware) UpdatePolyGeofence(ctx context.Context, token string, projectId string, fence *lbs.PolyGeofence) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method update_poly_geofence for token %s and project %s and fanceName %s took %s to complete", token, projectId, fence.Name, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.UpdatePolyGeofence(ctx, token, projectId, fence)
}

func (lm *loggingMiddleware) GetFenceIds(ctx context.Context, token string, projectId string) (fenceIds []string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method get_fenceids for token %s and project %s took %s to complete", token, projectId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.GetFenceIds(ctx, token, projectId)
}

// Alarm
func (lm *loggingMiddleware) QueryStatus(ctx context.Context, token string, projectId string, monitoredPerson string, fenceIds []string) (status *lbp.QueryStatus, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method query_status for token %s and project %s monitoredpersion %s fenceIds %s took %s to complete", token, projectId, monitoredPerson, fenceIds, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.QueryStatus(ctx, token, projectId, monitoredPerson, fenceIds)
}

func (lm *loggingMiddleware) GetHistoryAlarms(ctx context.Context, token string, projectId string, monitoredPerson string, fenceIds []string) (alarms *lbp.HistoryAlarms, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method get_history_alarms for token %s and project %s monitoredpersion %s fenceIds %s took %s to complete", token, projectId, monitoredPerson, fenceIds, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.GetHistoryAlarms(ctx, token, projectId, monitoredPerson, fenceIds)
}

func (lm *loggingMiddleware) BatchGetHistoryAlarms(ctx context.Context, token string, projectId string, input *lbp.BatchGetHistoryAlarmsRequest) (alarms *lbp.HistoryAlarms, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method batchget_history_alarms for token %s and project %s took %s to complete", token, projectId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.BatchGetHistoryAlarms(ctx, token, projectId, input)
}

func (lm *loggingMiddleware) GetStayPoints(ctx context.Context, token string, projectId string, input *lbp.GetStayPointsRequest) (points *lbp.StayPoints, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method get_stay_points for token %s and project %s took %s to complete", token, projectId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.GetStayPoints(ctx, token, projectId, input)
}

// NotifyAlarms is used by apiserver to provide asynchrous notication
func (lm *loggingMiddleware) NotifyAlarms(ctx context.Context, token string, projectId string, content []byte) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method notify_alarms for token %s and project %s took %s to complete", token, projectId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.NotifyAlarms(ctx, token, projectId, content)
}

func (lm *loggingMiddleware) GetFenceUserId(ctx context.Context, token string, fenceId string) (userId string, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method get_fence_userids for token %s and fenceId %s took %s to complete", token, fenceId, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.GetFenceUserId(ctx, token, fenceId)
}

//Entity
func (lm *loggingMiddleware) AddEntity(ctx context.Context, token string, projectId string, entityName string, entityDesc string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method add_entiry for token %s and entityName %s took %s to complete", token, entityName, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.AddEntity(ctx, token, projectId, entityName, entityDesc)
}

func (lm *loggingMiddleware) UpdateEntity(ctx context.Context, token string, projectId string, entityName string, entityDesc string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method update_entity for token %s and entityName %s took %s to complete", token, entityName, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.UpdateEntity(ctx, token, projectId, entityName, entityDesc)
}

func (lm *loggingMiddleware) DeleteEntity(ctx context.Context, token string, projectId string, entityName string) (err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method delete_entity for token %s and entityName %s took %s to complete", token, entityName, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.DeleteEntity(ctx, token, projectId, entityName)
}

func (lm *loggingMiddleware) ListEntity(ctx context.Context, token string, projectId string, coordTypeOutput string, pageIndex int32, pageSize int32) (total int32, infos []*lbs.EntityInfo, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method list_entity for token %s and coordTypeOutput %s and pageIndex %d and pageSize %d took %s to complete", token, coordTypeOutput, pageIndex, pageSize, time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.ListEntity(ctx, token, projectId, coordTypeOutput, pageIndex, pageSize)
}