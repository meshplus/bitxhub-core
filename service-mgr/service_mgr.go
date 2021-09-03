package service_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/looplab/fsm"
	"github.com/meshplus/bitxhub-core/governance"
	"github.com/sirupsen/logrus"
)

const (
	SERVICE_PREFIX          = "service"
	SERVICE_APPCHAIN_PREFIX = "appchain"
)

type ServiceManager struct {
	governance.Persister
}

type ServiceType string

var (
	ServiceCallContract       ServiceType = "CallContract"
	ServiceDepositCertificate ServiceType = "DepositCertificate"
	ServiceDataMigration      ServiceType = "DataMigration"
)

type Service struct {
	ChainID    string              `json:"chain_id"`    // aoochain id
	ServiceID  string              `json:"service_id"`  // service id, contract addr(ServiceCallContract), service name(others)
	Name       string              `json:"name"`        // service name
	Type       ServiceType         `json:"type"`        // service type
	Intro      string              `json:"intro"`       // service introduction
	Ordered    bool                `json:"ordered"`     // service should be in order or not
	Permission map[string]struct{} `json:"permission"`  // counter party services which are allowed to call the service
	Details    string              `json:"details"`     // Detailed description of the service
	CreateTime int64               `json:"create_time"` // service create time

	Score             float64                                 `json:"score"`
	EvaluationRecords map[string]*governance.EvaluationRecord `json:"evaluation_records"`

	InvokeCount       uint64                              `json:"invoke_count"`
	InvokeSuccessRate float64                             `json:"invoke_success_rate"`
	InvokeRecords     map[string]*governance.InvokeRecord `json:"transfer_records"`

	Status governance.GovernanceStatus `json:"status"`
	FSM    *fsm.FSM                    `json:"fsm"`
}

var serviceStateMap = map[governance.EventType][]governance.GovernanceStatus{
	governance.EventRegister: {governance.GovernanceUnavailable},
	governance.EventUpdate:   {governance.GovernanceAvailable, governance.GovernanceFrozen},
	governance.EventFreeze:   {governance.GovernanceAvailable, governance.GovernanceUpdating, governance.GovernanceActivating},
	governance.EventActivate: {governance.GovernanceFrozen},
	governance.EventPause:    {governance.GovernanceAvailable, governance.GovernanceUpdating, governance.GovernanceFreezing, governance.GovernanceActivating, governance.GovernanceLogouting},
	governance.EventUnpause:  {governance.GovernancePause},
	governance.EventLogout:   {governance.GovernanceAvailable, governance.GovernanceUpdating, governance.GovernanceFreezing, governance.GovernanceActivating, governance.GovernanceFrozen, governance.GovernancePause},
}

var serviceAvailableMap = map[governance.GovernanceStatus]struct{}{
	governance.GovernanceAvailable: {},
	governance.GovernanceFreezing:  {},
}

func New(persister governance.Persister) ServiceMgr {
	return &ServiceManager{persister}
}

func (s *Service) IsAvailable() bool {
	if _, ok := serviceAvailableMap[s.Status]; ok {
		return true
	} else {
		return false
	}
}

func (s *Service) CheckPermission(serviceId string) bool {
	_, ok := s.Permission[serviceId]

	if ok {
		return false
	} else {
		return true
	}
}

func (s *Service) setFSM(lastStatus governance.GovernanceStatus) {
	s.FSM = fsm.NewFSM(
		string(s.Status),
		fsm.Events{
			// register 3
			{Name: string(governance.EventRegister), Src: []string{string(governance.GovernanceUnavailable)}, Dst: string(governance.GovernanceRegisting)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceRegisting)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceRegisting)}, Dst: string(lastStatus)},

			// update 1
			{Name: string(governance.EventUpdate), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceFrozen), string(governance.GovernanceFreezing), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceUpdating)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceUpdating)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceUpdating)}, Dst: string(governance.GovernanceFrozen)},

			// freeze 2
			{Name: string(governance.EventFreeze), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceUpdating), string(governance.GovernanceActivating), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceFreezing)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceFreezing)}, Dst: string(governance.GovernanceFrozen)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceFreezing)}, Dst: string(lastStatus)},

			// activate 1
			{Name: string(governance.EventActivate), Src: []string{string(governance.GovernanceFrozen), string(governance.GovernanceFreezing), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceActivating)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceActivating)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceActivating)}, Dst: string(lastStatus)},

			// pause
			{Name: string(governance.EventPause), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceUpdating), string(governance.GovernanceFreezing), string(governance.GovernanceActivating)}, Dst: string(governance.GovernancePause)},

			// unpause
			{Name: string(governance.EventUnpause), Src: []string{string(governance.GovernancePause)}, Dst: string(governance.GovernanceAvailable)},

			// logout 3
			{Name: string(governance.EventLogout), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceUpdating), string(governance.GovernanceFreezing), string(governance.GovernanceFrozen), string(governance.GovernanceActivating), string(governance.GovernancePause)}, Dst: string(governance.GovernanceLogouting)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceForbidden)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceLogouting)}, Dst: string(lastStatus)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) {
				s.Status = governance.GovernanceStatus(s.FSM.Current())
			},
		},
	)
}

func (sm *ServiceManager) GovernancePre(id string, event governance.EventType, _ []byte) (interface{}, error) {
	service := &Service{}
	if ok := sm.GetObject(ServiceKey(id), service); !ok {
		if event == governance.EventRegister {
			return service, nil
		} else {
			return service, fmt.Errorf("the service does not exist")
		}
	}

	for _, s := range serviceStateMap[event] {
		if service.Status == s {
			return service, nil
		}
	}

	return service, fmt.Errorf("the service (%s) can not be %s", string(service.Status), string(event))
}

func (sm *ServiceManager) ChangeStatus(id, trigger, lastStatus string, _ []byte) (bool, []byte) {
	service := &Service{}
	if ok := sm.GetObject(ServiceKey(id), service); !ok {
		return false, []byte(fmt.Sprintf("this service does not exist: %s", id))
	}

	service.setFSM(governance.GovernanceStatus(lastStatus))
	err := service.FSM.Event(trigger)
	if err != nil {
		return false, []byte(fmt.Sprintf("change status error: %v", err))
	}

	sm.SetObject(ServiceKey(id), *service)
	return true, nil
}

func (sm *ServiceManager) CountAvailable(_ []byte) (bool, []byte) {
	ok, value := sm.Query(SERVICE_PREFIX)
	if !ok {
		return true, []byte("0")
	}

	count := 0
	for _, v := range value {
		service := &Service{}
		if err := json.Unmarshal(v, service); err != nil {
			return false, []byte(fmt.Sprintf("unmarshal json error: %v", err))
		}

		if service.IsAvailable() {
			count++
		}

	}
	return true, []byte(strconv.Itoa(count))
}

func (sm *ServiceManager) CountAll(_ []byte) (bool, []byte) {
	ok, value := sm.Query(SERVICE_PREFIX)
	if !ok {
		return true, []byte("0")
	}
	return true, []byte(strconv.Itoa(len(value)))
}

func (sm *ServiceManager) All(_ []byte) (interface{}, error) {
	ret := make([]*Service, 0)
	ok, value := sm.Query(SERVICE_PREFIX)
	if ok {
		for _, data := range value {
			service := &Service{}
			if err := json.Unmarshal(data, service); err != nil {
				return nil, err
			}
			ret = append(ret, service)
		}
	}

	return ret, nil
}

func (sm *ServiceManager) QueryById(id string, _ []byte) (interface{}, error) {
	var service Service
	ok := sm.GetObject(ServiceKey(id), &service)
	if !ok {
		return nil, fmt.Errorf("this service does not exist: %s", id)
	}

	return &service, nil
}

func (sm *ServiceManager) GetIDListByChainID(chainID string) (map[string]struct{}, error) {
	serviceMap := make(map[string]struct{})
	ok := sm.GetObject(AppchainServicesKey(chainID), &serviceMap)
	if !ok {
		return nil, fmt.Errorf("the service list does not exist: %s", chainID)
	}

	return serviceMap, nil
}

func (sm *ServiceManager) PackageServiceInfo(chainID, serviceID, name, typ, intro string, ordered bool, permits, details string, createTime int64, status governance.GovernanceStatus) (*Service, error) {
	permission := make(map[string]struct{}, 0)
	if permits != "" {
		for _, id := range strings.Split(permits, ",") {
			permission[id] = struct{}{}
		}
	}

	service := &Service{
		ChainID:           chainID,
		ServiceID:         serviceID,
		Name:              name,
		Type:              ServiceType(typ),
		Intro:             intro,
		Ordered:           ordered,
		Permission:        permission,
		Details:           details,
		CreateTime:        createTime,
		Score:             0,
		EvaluationRecords: make(map[string]*governance.EvaluationRecord),
		InvokeCount:       0,
		InvokeSuccessRate: 0,
		InvokeRecords:     make(map[string]*governance.InvokeRecord),
		Status:            status,
	}

	return service, nil
}

func (sm *ServiceManager) Register(info *Service) (bool, []byte) {
	chainServiceID := fmt.Sprintf("%s:%s", info.ChainID, info.ServiceID)
	sm.SetObject(ServiceKey(chainServiceID), *info)

	serviceMap := make(map[string]struct{})
	_ = sm.GetObject(AppchainServicesKey(info.ChainID), serviceMap)
	serviceMap[chainServiceID] = struct{}{}
	sm.SetObject(AppchainServicesKey(info.ChainID), serviceMap)

	sm.Logger().WithFields(logrus.Fields{
		"chainServiceID": chainServiceID,
	}).Info("service is registering")

	return true, nil
}

func (sm *ServiceManager) Update(updateInfo *Service) (bool, []byte) {
	chainServiceID := fmt.Sprintf("%s:%s", updateInfo.ChainID, updateInfo.ServiceID)
	service := &Service{}
	ok := sm.GetObject(ServiceKey(chainServiceID), service)
	if !ok {
		return false, []byte("the service is not exist")
	}

	service.Name = updateInfo.Name
	service.Intro = updateInfo.Intro
	service.Ordered = updateInfo.Ordered
	service.Details = updateInfo.Details
	service.Permission = updateInfo.Permission
	sm.SetObject(ServiceKey(chainServiceID), *service)
	sm.Logger().WithFields(logrus.Fields{
		"chainServiceId": chainServiceID,
	}).Info("service is updating")

	return true, nil
}

func ServiceKey(id string) string {
	return fmt.Sprintf("%s-%s", SERVICE_PREFIX, id)
}

func AppchainServicesKey(id string) string {
	return fmt.Sprintf("%s-%s", SERVICE_APPCHAIN_PREFIX, id)
}
