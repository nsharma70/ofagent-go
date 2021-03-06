/*
   Copyright 2020 the original author or authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package ofagent

import (
	"context"
	"encoding/json"
	"net"

	ofp "github.com/donNewtonAlpha/goloxi/of13"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opencord/ofagent-go/internal/pkg/openflow"
	"github.com/opencord/voltha-lib-go/v3/pkg/log"
	"google.golang.org/grpc"
)

func (ofa *OFAgent) receiveChangeEvents(ctx context.Context) {
	logger.Debug("receive-change-events-started")
	// If we exit, assume disconnected
	defer func() {
		ofa.events <- ofaEventVolthaDisconnected
		logger.Debug("receive-change-events-finished")
	}()
	if ofa.volthaClient == nil {
		logger.Error("no-voltha-connection")
		return
	}
	opt := grpc.EmptyCallOption{}
	streamCtx, streamDone := context.WithCancel(context.Background())
	defer streamDone()
	stream, err := ofa.volthaClient.Get().ReceiveChangeEvents(streamCtx, &empty.Empty{}, opt)
	if err != nil {
		logger.Errorw("Unable to establish Receive Change Event Stream",
			log.Fields{"error": err})
		return
	}

top:
	for {
		select {
		case <-ctx.Done():
			break top
		default:
			ce, err := stream.Recv()
			if err != nil {
				logger.Errorw("error receiving change event",
					log.Fields{"error": err})
				break top
			}
			ofa.changeEventChannel <- ce
			logger.Debug("receive-change-event-queued")
		}
	}
}

func (ofa *OFAgent) handleChangeEvents(ctx context.Context) {
	logger.Debug("handle-change-event-started")

top:
	for {
		select {
		case <-ctx.Done():
			break top
		case changeEvent := <-ofa.changeEventChannel:
			deviceID := changeEvent.GetId()
			portStatus := changeEvent.GetPortStatus()
			logger.Debugw("received-change-event",
				log.Fields{
					"device-id":   deviceID,
					"port-status": portStatus})

			if portStatus == nil {
				if logger.V(log.WarnLevel) {
					js, _ := json.Marshal(changeEvent.GetEvent())
					logger.Warnw("Received change event that was not port status",
						log.Fields{"ChangeEvent": js})
				}
				break
			}
			ofPortStatus := ofp.NewPortStatus()
			ofPortStatus.SetXid(openflow.GetXid())
			ofPortStatus.SetVersion(4)

			ofReason := ofp.PortReason(portStatus.GetReason())
			ofPortStatus.SetReason(ofReason)
			ofDesc := ofp.NewPortDesc()

			desc := portStatus.GetDesc()
			ofDesc.SetAdvertised(ofp.PortFeatures(desc.GetAdvertised()))
			ofDesc.SetConfig(ofp.PortConfig(0))
			ofDesc.SetCurr(ofp.PortFeatures(desc.GetAdvertised()))
			ofDesc.SetCurrSpeed(desc.GetCurrSpeed())
			intArray := desc.GetHwAddr()
			var octets []byte
			for _, val := range intArray {
				octets = append(octets, byte(val))
			}
			addr := net.HardwareAddr(octets)
			ofDesc.SetHwAddr(addr)
			ofDesc.SetMaxSpeed(desc.GetMaxSpeed())
			ofDesc.SetName(openflow.PadString(desc.GetName(), 16))
			ofDesc.SetPeer(ofp.PortFeatures(desc.GetPeer()))
			ofDesc.SetPortNo(ofp.Port(desc.GetPortNo()))
			ofDesc.SetState(ofp.PortState(desc.GetState()))
			ofDesc.SetSupported(ofp.PortFeatures(desc.GetSupported()))
			ofPortStatus.SetDesc(*ofDesc)
			if err := ofa.getOFClient(deviceID).SendMessage(ofPortStatus); err != nil {
				logger.Errorw("handle-change-events-send-message", log.Fields{"error": err})
			}
		}
	}

	logger.Debug("handle-change-event-finsihed")
}
