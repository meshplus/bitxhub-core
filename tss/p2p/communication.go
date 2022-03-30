package p2p

import (
	"sync"

	peer_mgr "github.com/meshplus/bitxhub-core/peer-mgr"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/sirupsen/logrus"
)

// Communication use p2p to send messages among all the TSS nodes
type Communication struct {
	peer_mgr.OrderPeerManager

	SendMsgChan chan *message.SendMsgChan
	stopChan    chan struct{} // channel to indicate whether we should stop

	wg     *sync.WaitGroup
	logger logrus.FieldLogger
}

// NewCommunication create a new instance of Communication
func NewCommunication(peerMgr peer_mgr.OrderPeerManager, logger logrus.FieldLogger) (*Communication, error) {
	return &Communication{
		OrderPeerManager: peerMgr,
		logger:           logger,
		wg:               &sync.WaitGroup{},
		stopChan:         make(chan struct{}),
		SendMsgChan:      make(chan *message.SendMsgChan, 1024),
	}, nil
}

// Start will start the communication
func (c *Communication) Start() {
	c.wg.Add(1)
	go c.ProcessBroadcast()
}

// Stop communication
func (c *Communication) Stop() error {
	close(c.stopChan)
	c.wg.Wait()
	return nil
}

func (c *Communication) ProcessBroadcast() {
	c.logger.Debug("start to process send message channel")
	defer c.logger.Debug("stop process broadcast message channel")
	defer c.wg.Done()
	for {
		select {
		case msg := <-c.SendMsgChan:
			c.logger.Debugf("=================== send %s message to %+v", msg.P2PMsg.Type, msg.PartiesID)
			if len(msg.PartiesID) == 0 {
				err := c.Broadcast(msg.P2PMsg)
				if err != nil {
					c.logger.WithFields(logrus.Fields{
						"to":   msg.PartiesID,
						"type": msg.P2PMsg.Type,
						"err":  err.Error(),
					}).Warnf("broadcast error")
				}
			} else {
				for _, id := range msg.PartiesID {
					err := c.AsyncSend(id, msg.P2PMsg)
					if err != nil {
						c.logger.WithFields(logrus.Fields{
							"to":   id,
							"type": msg.P2PMsg.Type,
							"err":  err.Error(),
						}).Warnf("AsyncSend error")
					}
				}
			}
		case <-c.stopChan:
			return
		}
	}
}
