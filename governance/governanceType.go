package governance

type GovernanceStatus string
type EventType string

const (
	REGISTERED = 0
	APPROVED   = 1

	GovernanceRegisting    GovernanceStatus = "registering"
	GovernanceAvailable    GovernanceStatus = "available"
	GovernanceUnavailable  GovernanceStatus = "unavailable"
	GovernanceUpdating     GovernanceStatus = "updating"
	GovernanceFreezing     GovernanceStatus = "freezing"
	GovernanceActivating   GovernanceStatus = "activating"
	GovernanceFrozen       GovernanceStatus = "frozen"
	GovernanceLogouting    GovernanceStatus = "logouting"
	GovernanceBinding      GovernanceStatus = "binding"
	GovernanceUnbinding    GovernanceStatus = "unbinding"
	GovernanceBindable     GovernanceStatus = "bindable"
	GovernanceForbidden    GovernanceStatus = "forbidden"
	GovernanceTransferring GovernanceStatus = "transfering"
	GovernancePause        GovernanceStatus = "pause"

	EventRegister EventType = "register"
	EventUpdate   EventType = "update"
	EventFreeze   EventType = "freeze"
	EventActivate EventType = "activate"
	EventLogout   EventType = "logout"
	EventApprove  EventType = "approve"
	EventReject   EventType = "reject"
	EventBind     EventType = "bind"
	EventUnbind   EventType = "unbind"
	EventTransfer EventType = "transfer"
	EventPause    EventType = "pause"
)

type RegisterResult struct {
	IsRegistered bool   `json:"is_registered"`
	ID           string `json:"id"`
}

type GovernanceResult struct {
	ProposalID string `json:"proposal_id"`
	Extra      []byte `json:"extra"`
}

type InvokeRecord struct {
	Succeed int64 `json:"succeed"`
	Failure int64 `json:"failure"`
}

type EvaluationRecord struct {
	Addr       string  `json:"addr"`
	Score      float64 `json:"score"`
	Desc       string  `json:"desc"`
	CreateTime int64   `json:"create_time"`
}
