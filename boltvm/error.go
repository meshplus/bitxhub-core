package boltvm

const (
	// other
	OtherInternalErrCode ErrorCode = "2000000"
	OtherInternalErrMsg  ErrorMsg  = "%s"

	// governance
	GovernanceInternalErrCode ErrorCode = "2010000"
	GovernanceInternalErrMsg  ErrorMsg  = "%s"

	GovernanceNoPermissionCode ErrorCode = "1010001"
	GovernanceNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	GovernanceNonexistentProposalCode ErrorCode = "1010002"
	GovernanceNonexistentProposalMsg  ErrorMsg  = "the proposal(%s) does not exist: %s"

	GovernanceEndEndedProposalCode ErrorCode = "1010003"
	GovernanceEndEndedProposalMsg  ErrorMsg  = "the proposal(%s) is %s, cannot be ended"

	GovernanceNotVoteAdminCode ErrorCode = "1010004"
	GovernanceNotVoteAdminMsg  ErrorMsg  = "the admin of the address(%s) has not voted"

	GovernanceIllegalProposalTypeCode ErrorCode = "1010005"
	GovernanceIllegalProposalTypeMsg  ErrorMsg  = "illegal proposal type(%s)"

	GovernanceIllegalProposalStatusCode ErrorCode = "1010006"
	GovernanceIllegalProposalStatusMsg  ErrorMsg  = "illegal proposal status(%s)"

	GovernanceUnavailableAdminVoteCode ErrorCode = "1010007"
	GovernanceUnavailableAdminVoteMsg  ErrorMsg  = "the admin(%s) is currently unavailable and can not vote"

	GovernanceVoteEndProposalCode ErrorCode = "1010008"
	GovernanceVoteEndProposalMsg  ErrorMsg  = "the current status of the proposal is %s and cannot be voted on"

	GovernanceAdminRepeatVoteCode ErrorCode = "1010009"
	GovernanceAdminRepeatVoteMsg  ErrorMsg  = "administrator of the address(%s) has voted"

	GovernanceIllegalVoteInfoCode ErrorCode = "1010010"
	GovernanceIllegalVoteInfoMsg  ErrorMsg  = "illegal vote info(%s), should be approve or reject"

	GovernanceAdminNoVotePermissonCode ErrorCode = "1010011"
	GovernanceAdminNoVotePermissonMsg  ErrorMsg  = "the admin(%s) can not vote to the proposal(%s)"

	GovernanceIllegalProposalStrategyInfoCode ErrorCode = "1010012"
	GovernanceIllegalProposalStrategyInfoMsg  ErrorMsg  = "illegal proposal strategy info: %s"

	GovernanceNonexistentProposalStrategyCode ErrorCode = "1010013"
	GovernanceNonexistentProposalStrategyMsg  ErrorMsg  = "the proposal strategy for the type(%s) does not exist"

	// appchain
	AppchainInternalErrCode ErrorCode = "2020000"
	AppchainInternalErrMsg  ErrorMsg  = "%s"

	AppchainNoPermissionCode ErrorCode = "1020001"
	AppchainNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	AppchainNilBrokerCode ErrorCode = "1020002"
	AppchainNilBrokerMsg  ErrorMsg  = "broker can not be nil"

	AppchainIllegalFabricBrokerCode ErrorCode = "1020003"
	AppchainIllegalFabricBrokerMsg  ErrorMsg  = "illegal fabric broker info(%s): %s"

	AppchainEmptyChainIDCode ErrorCode = "1020004"
	AppchainEmptyChainIDMsg  ErrorMsg  = "chain id can not be an empty string"

	AppchainDuplicateChainIDCode ErrorCode = "1020005"
	AppchainDuplicateChainIDMsg  ErrorMsg  = "the appchain id %s has been occupied"

	AppchainEmptyChainNameCode ErrorCode = "1020006"
	AppchainEmptyChainNameMsg  ErrorMsg  = "chain name can not be an empty string"

	AppchainDuplicateChainNameCode ErrorCode = "1020007"
	AppchainDuplicateChainNameMsg  ErrorMsg  = "the appchain name %s has been occupied by appchain %s"

	AppchainIncompleteAdminListCode ErrorCode = "1020008"
	AppchainIncompleteAdminListMsg  ErrorMsg  = "the admin list does not contain the current admin(%s)"

	AppchainIllegalAdminAddrCode ErrorCode = "1020009"
	AppchainIllegalAdminAddrMsg  ErrorMsg  = "illegal admin addr(%s): %s"

	AppchainDuplicateAdminCode ErrorCode = "1020010"
	AppchainDuplicateAdminMsg  ErrorMsg  = "the appchain admin %s has been occupied by %s"

	AppchainEmptyRuleUrlCode ErrorCode = "1020011"
	AppchainEmptyRuleUrlMsg  ErrorMsg  = "urls for custom rule cannot be empty string"

	AppchainStatusErrorCode ErrorCode = "1020012"
	AppchainStatusErrorMsg  ErrorMsg  = "the appchain(%s) is %s, can not do %s"

	AppchainRuleUpdatingCode ErrorCode = "1020013"
	AppchainRuleUpdatingMsg  ErrorMsg  = "chain master rule(%s) is updating, can not submit proposal to %s appchain(%s)"

	AppchainNonexistentChainCode ErrorCode = "1020014"
	AppchainNonexistentChainMsg  ErrorMsg  = "the appchain(%s) does not exist: %s"

	// rule
	RuleInternalErrCode ErrorCode = "2030000"
	RuleInternalErrMsg  ErrorMsg  = "%s"

	RuleNoPermissionCode ErrorCode = "1030001"
	RuleNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	RuleIllegalRuleAddrCode ErrorCode = "1030002"
	RuleIllegalRuleAddrMsg  ErrorMsg  = "illegal rule addr(%s): %s"

	RuleNonexistentRuleCode ErrorCode = "1030003"
	RuleNonexistentRuleMsg  ErrorMsg  = "the rule(%s) does not exist"

	RuleNonexistentChainCode ErrorCode = "1030004"
	RuleNonexistentChainMsg  ErrorMsg  = "the appchain(%s) does not exist: %s"

	RuleStatusErrorCode ErrorCode = "1030005"
	RuleStatusErrorMsg  ErrorMsg  = "the rule(%s) is %s, can not do %s"

	RuleMasterRuleUpdatingCode ErrorCode = "1030006"
	RuleMasterRuleUpdatingMsg  ErrorMsg  = "master rule(%s) is updating, can not update master rule"

	RuleRegisterDefaultCode ErrorCode = "1030007"
	RuleRegisterDefaultMsg  ErrorMsg  = "default rule(%s) can not be registered"

	RuleLogoutDefaultCode ErrorCode = "1030008"
	RuleLogoutDefaultMsg  ErrorMsg  = "default rule(%s) can not be logouted"

	RuleAppchainForbiddenCode ErrorCode = "1030009"
	RuleAppchainForbiddenMsg  ErrorMsg  = "appchain(%s) is forbidden, can not operate rule"

	RuleAppchainStatusErrorCode ErrorCode = "1030010"
	RuleAppchainStatusErrorMsg  ErrorMsg  = "appchain(%s) is %s, can not bind new master rule"

	// role
	RoleInternalErrCode ErrorCode = "2040000"
	RoleInternalErrMsg  ErrorMsg  = "%s"

	RoleNoPermissionCode ErrorCode = "1040001"
	RoleNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	RoleIllegalRoleIDCode ErrorCode = "1040002"
	RoleIllegalRoleIDMsg  ErrorMsg  = "illegal role id(%s): %s"

	RoleNonexistentNodeCode ErrorCode = "1040003"
	RoleNonexistentNodeMsg  ErrorMsg  = "the node(%s) does not exist: %s"

	RoleWrongNodeCode ErrorCode = "1040004"
	RoleWrongNodeMsg  ErrorMsg  = "the node(%s) bind to audit admin is not a nvp node"

	RoleWrongStatusNodeCode ErrorCode = "1040005"
	RoleWrongStatusNodeMsg  ErrorMsg  = "the status of node(%s) is %s, can not bind to audit admin"

	RoleIllegalRoleTypeCode ErrorCode = "1040006"
	RoleIllegalRoleTypeMsg  ErrorMsg  = "illegal role type(%s)"

	RoleNonexistentRoleCode ErrorCode = "1040007"
	RoleNonexistentRoleMsg  ErrorMsg  = "the role(%s) does not exist"

	RoleStatusErrorCode ErrorCode = "1040008"
	RoleStatusErrorMsg  ErrorMsg  = "the role(%s) is %s, can not do %s"

	RoleNonsupportSuperAdminCode ErrorCode = "1040009"
	RoleNonsupportSuperAdminMsg  ErrorMsg  = "the super admin role(%s) does not support %s"

	RoleNonsupportAppchainAdminCode ErrorCode = "1040010"
	RoleNonsupportAppchainAdminMsg  ErrorMsg  = "the appchain admin role(%s) does not support %s"

	RoleNonsupportAuditAdminCode ErrorCode = "1040011"
	RoleNonsupportAuditAdminMsg  ErrorMsg  = "the audit admin role(%s) does not support %s"

	RoleNonsupportGovernanceAdminCode ErrorCode = "1040012"
	RoleNonsupportGovernanceAdminMsg  ErrorMsg  = "the governance admin role(%s) does not support %s"

	RoleNoAppchainAdminCode ErrorCode = "1040013"
	RoleNoAppchainAdminMsg  ErrorMsg  = "there is no admin for the appchain(%s)"

	RoleNotGovernanceAdminCode ErrorCode = "1040014"
	RoleNotGovernanceAdminMsg  ErrorMsg  = "the role(%s) is not governane admin"

	// node
	NodeInternalErrCode ErrorCode = "2050000"
	NodeInternalErrMsg  ErrorMsg  = "%s"

	NodeNoPermissionCode ErrorCode = "1050001"
	NodeNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	NodeVPBeingGovernedCode ErrorCode = "1050002"
	NodeVPBeingGovernedMsg  ErrorMsg  = " a vp node is being governed"

	NodeIllegalAccountCode ErrorCode = "1050003"
	NodeIllegalAccountMsg  ErrorMsg  = "illegal node account(%s): %s"

	NodeDuplicateAccountCode ErrorCode = "1050004"
	NodeDuplicateAccountMsg  ErrorMsg  = "the node account %s has been occupied by %s"

	NodeIllegalNodeTypeCode ErrorCode = "1050005"
	NodeIllegalNodeTypeMsg  ErrorMsg  = "illegal node type(%s)"

	NodeIllegalVpIdCode ErrorCode = "1050006"
	NodeIllegalVpIdMsg  ErrorMsg  = "illegal vp node id(%d) (%s)"

	NodeEmptyPidCode ErrorCode = "1050007"
	NodeEmptyPidMsg  ErrorMsg  = "node pid can not be an empty string"

	NodeDuplicatePidCode ErrorCode = "1050008"
	NodeDuplicatePidMsg  ErrorMsg  = "the node pid %s has been occupied by node %s"

	NodeEmptyNameCode ErrorCode = "1050009"
	NodeEmptyNameMsg  ErrorMsg  = "node name can not be an empty string"

	NodeDuplicateNameCode ErrorCode = "1050010"
	NodeDuplicateNameMsg  ErrorMsg  = "the node name %s has been occupied by node %s"

	NodeEmptyPermissionCode ErrorCode = "1050011"
	NodeEmptyPermissionMsg  ErrorMsg  = "empty node permission"

	NodeIllegalPermissionCode ErrorCode = "1050012"
	NodeIllegalPermissionMsg  ErrorMsg  = "illegal node permission addr(%s): %s"

	NodeNonexistentNodeCode ErrorCode = "1050013"
	NodeNonexistentNodeMsg  ErrorMsg  = "the node(%s) does not exist"

	NodeStatusErrorCode ErrorCode = "1050014"
	NodeStatusErrorMsg  ErrorMsg  = "the node(account:%s) is %s, can not do %s"

	NodeLogoutPrimaryNodeCode ErrorCode = "1050015"
	NodeLogoutPrimaryNodeMsg  ErrorMsg  = "don't support logout primary vp node(account:%s)"

	NodeLogoutTooFewNodeCode ErrorCode = "1050016"
	NodeLogoutTooFewNodeMsg  ErrorMsg  = "don't support logout node when there're only %s vp nodes"

	NodeLogoutWrongIdNodeCode ErrorCode = "10500017"
	NodeLogoutWrongIdNodeMsg  ErrorMsg  = "only support logout last vp node(id:%s) currently"

	NodeUpdateVPNodeCode ErrorCode = "10500018"
	NodeUpdateVPNodeMsg  ErrorMsg  = "can not update vp node(account: %s)"

	NodeBindVPNodeCode ErrorCode = "10500019"
	NodeBindVPNodeMsg  ErrorMsg  = "can not bind vp node(account: %s)"

	NodeUnbindVPNodeCode ErrorCode = "10500020"
	NodeUnbindVPNodeMsg  ErrorMsg  = "can not unbind vp node(account: %s)"

	// service
	ServiceInternalErrCode ErrorCode = "2060000"
	ServiceInternalErrMsg  ErrorMsg  = "%s"

	ServiceNoPermissionCode ErrorCode = "1060001"
	ServiceNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	ServiceNonexistentServiceCode ErrorCode = "1060002"
	ServiceNonexistentServiceMsg  ErrorMsg  = "the service(%s) does not exist"

	ServiceStatusErrorCode ErrorCode = "1060003"
	ServiceStatusErrorMsg  ErrorMsg  = "the service(%s) is %s, can not do %s"

	ServiceUnavailableChainCode ErrorCode = "1060004"
	ServiceUnavailableChainMsg  ErrorMsg  = "the appchain(%s) is not available so that the service(%s) cannot be governed: %s"

	ServiceEmptyNameCode ErrorCode = "1060005"
	ServiceEmptyNameMsg  ErrorMsg  = "service name can not be an empty string"

	ServiceDuplicateNameCode ErrorCode = "1060006"
	ServiceDuplicateNameMsg  ErrorMsg  = "the service name %s has been occupied by service %s"

	ServiceIllegalServiceIDCode ErrorCode = "1060007"
	ServiceIllegalServiceIDMsg  ErrorMsg  = "illegal service id(%s): %s"

	ServiceIllegalTypeCode ErrorCode = "1060008"
	ServiceIllegalTypeMsg  ErrorMsg  = "illegal service type(%s)"

	ServiceIllegalPermissionFormatCode ErrorCode = "1060009"
	ServiceIllegalPermissionFormatMsg  ErrorMsg  = "illegal permission full service id(%s) format: %s"

	ServiceNonexistentPermissionServiceCode ErrorCode = "1060010"
	ServiceNonexistentPermissionServiceMsg  ErrorMsg  = "the permission service(%s) is not registered on this relay chain(%s): %s"

	ServiceLogoutedPermissionServiceCode ErrorCode = "1060011"
	ServiceLogoutedPermissionServiceMsg  ErrorMsg  = "the permission service(%s) is logouted"

	ServiceIllegalEvaluateScoreCode ErrorCode = "1060012"
	ServiceIllegalEvaluateScoreMsg  ErrorMsg  = "illegal evaluate score(%s), should be in the range [0,5]"

	ServiceRepeatEvaluateCode ErrorCode = "1060013"
	ServiceRepeatEvaluateMsg  ErrorMsg  = "the caller(%s) has evaluate the service(%s)"

	// dapp
	DappInternalErrCode ErrorCode = "2070000"
	DappInternalErrMsg  ErrorMsg  = "%s"

	DappNoPermissionCode ErrorCode = "1070001"
	DappNoPermissionMsg  ErrorMsg  = "regulatorAddr(%s) does not have the permission: %s"

	DappNonexistentDappCode ErrorCode = "1070002"
	DappNonexistentDappMsg  ErrorMsg  = "the dapp(%s) does not exist"

	DappStatusErrorCode ErrorCode = "1070003"
	DappStatusErrorMsg  ErrorMsg  = "the dapp(%s) is %s, can not do %s"

	DappEmptyNameCode ErrorCode = "1070004"
	DappEmptyNameMsg  ErrorMsg  = "dapp name can not be an empty string"

	DappDuplicateNameCode ErrorCode = "1070005"
	DappDuplicateNameMsg  ErrorMsg  = "the dapp name %s has been occupied by dapp %s"

	DappEmptyUrlCode ErrorCode = "1070006"
	DappEmptyUrlMsg  ErrorMsg  = "dapp url can not be an empty string"

	DappIllegalTypeCode ErrorCode = "1070007"
	DappIllegalTypeMsg  ErrorMsg  = "illegal dapp type(%s)"

	DappIllegalContractAddrCode ErrorCode = "1070008"
	DappIllegalContractAddrMsg  ErrorMsg  = "illegal dapp contract addr(%s)"

	DappDuplicateContractRegisterCode ErrorCode = "1070009"
	DappDuplicateContractRegisterMsg  ErrorMsg  = "the contract address %s belongs to dapp %s and cannot be registered repeatedly"

	DappDuplicateContractUpdateCode ErrorCode = "1070010"
	DappDuplicateContractUpdateMsg  ErrorMsg  = "the contract address %s belongs to dapp %s and cannot be update to others"

	DappNonexistentContractCode ErrorCode = "1070011"
	DappNonexistentContractMsg  ErrorMsg  = "the contract(%s) does not exist"

	DappIllegalPermissionCode ErrorCode = "10700012"
	DappIllegalPermissionMsg  ErrorMsg  = "illegal dapp permission addr(%s): %s"

	DappIllegalTransferAddrCode ErrorCode = "1070013"
	DappIllegalTransferAddrMsg  ErrorMsg  = "illegal new owner addr(%s): %s"

	DappTransferToSelfCode ErrorCode = "1070014"
	DappTransferToSelfMsg  ErrorMsg  = "can not transfer dapp to self(%s)"

	DappIllegalEvaluateScoreCode ErrorCode = "1070015"
	DappIllegalEvaluateScoreMsg  ErrorMsg  = "illegal evaluate score(%s), should be in the range [0,5]"

	DappRepeatEvaluateCode ErrorCode = "1070016"
	DappRepeatEvaluateMsg  ErrorMsg  = "the caller(%s) has evaluate the dapp(%s)"

	// interchain
	InterchainInternalErrCode ErrorCode = "2080000"
	InterchainInternalErrMsg  ErrorMsg  = "%s"

	InterchainNonexistentInterchainCode ErrorCode = "1080001"
	InterchainNonexistentInterchainMsg  ErrorMsg  = "the service %s interchain info does not exist"

	InterchainInvalidIBTPParseSourceErrorCode ErrorCode = "1080002"
	InterchainInvalidIBTPParseSourceErrorMsg  ErrorMsg  = "invalid ibtp, parse source chain service id: %s"

	InterchainInvalidIBTPParseDestErrorCode ErrorCode = "1080003"
	InterchainInvalidIBTPParseDestErrorMsg  ErrorMsg  = "invalid ibtp, parse dest chain service id: %s"

	InterchainInvalidIBTPNotInCurBXHCode ErrorCode = "1080004"
	InterchainInvalidIBTPNotInCurBXHMsg  ErrorMsg  = "invalid ibtp, neither source service nor dest service of IBTP %s is in current bitxhub"

	InterchainInvalidIBTPIllegalTypeCode ErrorCode = "1080005"
	InterchainInvalidIBTPIllegalTypeMsg  ErrorMsg  = "invalid ibtp, IBTP type %v is not expected"

	InterchainSourceAppchainNotAvailableCode ErrorCode = "1080006"
	InterchainSourceAppchainNotAvailableMsg  ErrorMsg  = "source appchain(%s) not available: %s"

	InterchainSourceServiceNotAvailableCode ErrorCode = "1080007"
	InterchainSourceServiceNotAvailableMsg  ErrorMsg  = "source service(%s) not available: %s"

	InterchainTargetAppchainNotAvailableCode ErrorCode = "1080008"
	InterchainTargetAppchainNotAvailableMsg  ErrorMsg  = "target appchain(%s) not available: %s"

	InterchainTargetServiceNotAvailableCode ErrorCode = "1080009"
	InterchainTargetServiceNotAvailableMsg  ErrorMsg  = "target service(%s) not available: %s"

	InterchainTargetServiceNoPermissionCode ErrorCode = "1080010"
	InterchainTargetServiceNoPermissionMsg  ErrorMsg  = "source service(%s) does not have permission call target service(%s)"

	InterchainSourceBitXHubNotAvailableCode ErrorCode = "1080011"
	InterchainSourceBitXHubNotAvailableMsg  ErrorMsg  = "source bitxhub(%s) not available: %v"

	InterchainTargetBitXHubNotAvailableCode ErrorCode = "1080012"
	InterchainTargetBitXHubNotAvailableMsg  ErrorMsg  = "target bitxhub(%s) not available: %v"

	InterchainIbtpIndexExistCode ErrorCode = "1080013"
	InterchainIbtpIndexExistMsg  ErrorMsg  = "index already exists, required %d, but %d"

	InterchainIbtpIndexWrongCode ErrorCode = "1080014"
	InterchainIbtpIndexWrongMsg  ErrorMsg  = "wrong index, required %d, but %d"

	InterchainWrongIBTPIDCode ErrorCode = "1080015"
	InterchainWrongIBTPIDMsg  ErrorMsg  = "wrong ibtp id %s"

	InterchainNonexistentIBTPCode ErrorCode = "1080016"
	InterchainNonexistentIBTPMsg  ErrorMsg  = "the ibtp(%s) does not exist"

	// broker
	BrokerInternalErrCode ErrorCode = "2090000"
	BrokerInternalErrMsg  ErrorMsg  = "%s"

	BrokerIllegalFunctionCode ErrorCode = "1090001"
	BrokerIllegalFunctionMsg  ErrorMsg  = "illegal funcs(%s), funcs should be (func,funcCb,funcRb)"

	BrokerIllegalIBTPToCode ErrorCode = "1090002"
	BrokerIllegalIBTPToMsg  ErrorMsg  = "ibtp.To(%s) is not chain service"

	BrokerNonexistentInterchainInvokeCode ErrorCode = "1090003"
	BrokerNonexistentInterchainInvokeMsg  ErrorMsg  = "not found interchain %s invoke"

	BrokerNonexistentOutMsgCode ErrorCode = "1090004"
	BrokerNonexistentOutMsgMsg  ErrorMsg  = "not found out message for key %s"

	// transaction
	TransactionInternalErrCode ErrorCode = "2100000"
	TransactionInternalErrMsg  ErrorMsg  = "%s"

	TransactionNoPermissionCode ErrorCode = "1100001"
	TransactionNoPermissionMsg  ErrorMsg  = "current caller %s is not allowed"

	TransactionNonexistentTxCode ErrorCode = "1100002"
	TransactionNonexistentTxMsg  ErrorMsg  = "transaction id %s does not exist"

	TransactionExistentChildTxCode ErrorCode = "1100003"
	TransactionExistentChildTxMsg  ErrorMsg  = "child tx %s of global tx %s exists"

	TransactionNonexistentGlobalTxCode ErrorCode = "1100004"
	TransactionNonexistentGlobalTxMsg  ErrorMsg  = "global tx %s of child tx %s does not exist"

	TransactionStateErrCode ErrorCode = "1100005"
	TransactionStateErrMsg  ErrorMsg  = "transaction state transition error: %s"

	TransactionNonexistentGlobalIdCode ErrorCode = "1100006"
	TransactionNonexistentGlobalIdMsg  ErrorMsg  = "cannot get global id of child tx id %s"

	// trust chain
	TrustInternalErrCode ErrorCode = "2110000"
	TrustInternalErrMsg  ErrorMsg  = "%s"

	TrustNonexistentTrustDataCode ErrorCode = "1110001"
	TrustNonexistentTrustDataMsg  ErrorMsg  = "not found target(%s) trust meta"

	// asset
	AssetInternalErrCode ErrorCode = "2120000"
	AssetInternalErrMsg  ErrorMsg  = "%s"

	AssetIllegalEscrowAddrFormatCode ErrorCode = "1120001"
	AssetIllegalEscrowAddrFormatMsg  ErrorMsg  = "escrow addr(%s) formot error"

	AssetNonexistentEscrowAddrCode ErrorCode = "1120002"
	AssetNonexistentEscrowAddrMsg  ErrorMsg  = "not found escrow addr"

	AssetNonexistentProxyAddrCode ErrorCode = "1120003"
	AssetNonexistentProxyAddrMsg  ErrorMsg  = "not found proxy contract addr"

	AssetNonexistentInterchainSwapAddrCode ErrorCode = "1120004"
	AssetNonexistentInterchainSwapAddrMsg  ErrorMsg  = "not found interchain swap contract addr"

	AssetNonexistentCurHeaderCode ErrorCode = "1120005"
	AssetNonexistentCurHeaderMsg  ErrorMsg  = "not found current header"

	AssetNonexistentHeaderCode ErrorCode = "1120006"
	AssetNonexistentHeaderMsg  ErrorMsg  = "not found header for hash %v"

	AssetNonexistentEthTxCode ErrorCode = "1120007"
	AssetNonexistentEthTxMsg  ErrorMsg  = "not found ethereum tx for the hash %v"

	AssetNoPermissionCode ErrorCode = "1120008"
	AssetNoPermissionMsg  ErrorMsg  = "caller is not an admin account"
)
