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
package openflow

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/donNewtonAlpha/goloxi"
	ofp "github.com/donNewtonAlpha/goloxi/of13"
	"github.com/opencord/voltha-lib-go/v3/pkg/log"
	"github.com/opencord/voltha-protos/v3/go/openflow_13"
	"github.com/opencord/voltha-protos/v3/go/voltha"
)

func parseOxm(ofbField *openflow_13.OfpOxmOfbField) (goloxi.IOxm, uint16) {
	if logger.V(log.DebugLevel) {
		js, _ := json.Marshal(ofbField)
		logger.Debugw("parseOxm called",
			log.Fields{"ofbField": js})
	}

	switch ofbField.Type {
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_IN_PORT:
		ofpInPort := ofp.NewOxmInPort()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_Port)
		ofpInPort.Value = ofp.Port(val.Port)
		return ofpInPort, 4
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_ETH_TYPE:
		ofpEthType := ofp.NewOxmEthType()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_EthType)
		ofpEthType.Value = ofp.EthernetType(val.EthType)
		return ofpEthType, 2
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_IN_PHY_PORT:
		ofpInPhyPort := ofp.NewOxmInPhyPort()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_PhysicalPort)
		ofpInPhyPort.Value = ofp.Port(val.PhysicalPort)
		return ofpInPhyPort, 4
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_IP_PROTO:
		ofpIpProto := ofp.NewOxmIpProto()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_IpProto)
		ofpIpProto.Value = ofp.IpPrototype(val.IpProto)
		return ofpIpProto, 1
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_IPV4_DST:
		ofpIpv4Dst := ofp.NewOxmIpv4Dst()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_Ipv4Dst)
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, val.Ipv4Dst)
		if err != nil {
			logger.Errorw("error writing ipv4 address %v",
				log.Fields{"error": err})
		}
		ofpIpv4Dst.Value = buf.Bytes()
		return ofpIpv4Dst, 4
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_UDP_SRC:
		ofpUdpSrc := ofp.NewOxmUdpSrc()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_UdpSrc)
		ofpUdpSrc.Value = uint16(val.UdpSrc)
		return ofpUdpSrc, 2
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_UDP_DST:
		ofpUdpDst := ofp.NewOxmUdpDst()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_UdpDst)
		ofpUdpDst.Value = uint16(val.UdpDst)
		return ofpUdpDst, 2
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_VLAN_VID:
		ofpVlanVid := ofp.NewOxmVlanVid()
		val := ofbField.GetValue()
		if val == nil {
			ofpVlanVid.Value = uint16(0)
			return ofpVlanVid, 2
		}
		vlanId := val.(*openflow_13.OfpOxmOfbField_VlanVid)
		if ofbField.HasMask {
			ofpVlanVidMasked := ofp.NewOxmVlanVidMasked()
			valMask := ofbField.GetMask()
			vlanMask := valMask.(*openflow_13.OfpOxmOfbField_VlanVidMask)
			if vlanId.VlanVid == 4096 && vlanMask.VlanVidMask == 4096 {
				ofpVlanVidMasked.Value = uint16(vlanId.VlanVid)
				ofpVlanVidMasked.ValueMask = uint16(vlanMask.VlanVidMask)
			} else {
				ofpVlanVidMasked.Value = uint16(vlanId.VlanVid) | 0x1000
				ofpVlanVidMasked.ValueMask = uint16(vlanMask.VlanVidMask)

			}
			return ofpVlanVidMasked, 4
		}
		ofpVlanVid.Value = uint16(vlanId.VlanVid) | 0x1000
		return ofpVlanVid, 2
	case voltha.OxmOfbFieldTypes_OFPXMT_OFB_METADATA:
		ofpMetadata := ofp.NewOxmMetadata()
		val := ofbField.GetValue().(*openflow_13.OfpOxmOfbField_TableMetadata)
		ofpMetadata.Value = val.TableMetadata
		return ofpMetadata, 8
	default:
		if logger.V(log.WarnLevel) {
			js, _ := json.Marshal(ofbField)
			logger.Warnw("ParseOXM Unhandled OxmField",
				log.Fields{"OfbField": js})
		}
	}
	return nil, 0
}

func parseInstructions(ofpInstruction *openflow_13.OfpInstruction) (ofp.IInstruction, uint16) {
	if logger.V(log.DebugLevel) {
		js, _ := json.Marshal(ofpInstruction)
		logger.Debugw("parseInstructions called",
			log.Fields{"Instruction": js})
	}
	instType := ofpInstruction.Type
	data := ofpInstruction.GetData()
	switch instType {
	case ofp.OFPITWriteMetadata:
		instruction := ofp.NewInstructionWriteMetadata()
		instruction.Len = 24
		metadata := data.(*openflow_13.OfpInstruction_WriteMetadata).WriteMetadata
		instruction.Metadata = uint64(metadata.Metadata)
		return instruction, 24
	case ofp.OFPITMeter:
		instruction := ofp.NewInstructionMeter()
		instruction.Len = 8
		meter := data.(*openflow_13.OfpInstruction_Meter).Meter
		instruction.MeterId = meter.MeterId
		return instruction, 8
	case ofp.OFPITGotoTable:
		instruction := ofp.NewInstructionGotoTable()
		instruction.Len = 8
		gotoTable := data.(*openflow_13.OfpInstruction_GotoTable).GotoTable
		instruction.TableId = uint8(gotoTable.TableId)
		return instruction, 8
	case ofp.OFPITApplyActions:
		instruction := ofp.NewInstructionApplyActions()
		var instructionSize uint16
		instructionSize = 8
		//ofpActions := ofpInstruction.GetActions().Actions
		var actions []goloxi.IAction
		for _, ofpAction := range ofpInstruction.GetActions().Actions {
			action, actionSize := parseAction(ofpAction)
			actions = append(actions, action)
			instructionSize += actionSize

		}
		instruction.Actions = actions
		instruction.SetLen(instructionSize)
		if logger.V(log.DebugLevel) {
			js, _ := json.Marshal(instruction)
			logger.Debugw("parseInstructions returning",
				log.Fields{
					"size":               instructionSize,
					"parsed-instruction": js})
		}
		return instruction, instructionSize
	}
	//shouldn't have reached here :<
	return nil, 0
}

func parseAction(ofpAction *openflow_13.OfpAction) (goloxi.IAction, uint16) {
	if logger.V(log.DebugLevel) {
		js, _ := json.Marshal(ofpAction)
		logger.Debugw("parseAction called",
			log.Fields{"action": js})
	}
	switch ofpAction.Type {
	case openflow_13.OfpActionType_OFPAT_OUTPUT:
		ofpOutputAction := ofpAction.GetOutput()
		outputAction := ofp.NewActionOutput()
		outputAction.Port = ofp.Port(ofpOutputAction.Port)
		outputAction.MaxLen = uint16(ofpOutputAction.MaxLen)
		outputAction.Len = 16
		return outputAction, 16
	case openflow_13.OfpActionType_OFPAT_PUSH_VLAN:
		ofpPushVlanAction := ofp.NewActionPushVlan()
		ofpPushVlanAction.Ethertype = uint16(ofpAction.GetPush().Ethertype)
		ofpPushVlanAction.Len = 8
		return ofpPushVlanAction, 8
	case openflow_13.OfpActionType_OFPAT_POP_VLAN:
		ofpPopVlanAction := ofp.NewActionPopVlan()
		ofpPopVlanAction.Len = 8
		return ofpPopVlanAction, 8
	case openflow_13.OfpActionType_OFPAT_SET_FIELD:
		ofpActionSetField := ofpAction.GetSetField()
		setFieldAction := ofp.NewActionSetField()

		iOxm, _ := parseOxm(ofpActionSetField.GetField().GetOfbField())
		setFieldAction.Field = iOxm
		setFieldAction.Len = 16
		return setFieldAction, 16
	case openflow_13.OfpActionType_OFPAT_GROUP:
		ofpGroupAction := ofpAction.GetGroup()
		groupAction := ofp.NewActionGroup()
		groupAction.GroupId = ofpGroupAction.GroupId
		groupAction.Len = 8
		return groupAction, 8
	default:
		if logger.V(log.WarnLevel) {
			js, _ := json.Marshal(ofpAction)
			logger.Warnw("parseAction unknow action",
				log.Fields{"action": js})
		}
	}
	return nil, 0
}

func parsePortStats(port *voltha.LogicalPort) *ofp.PortStatsEntry {
	stats := port.OfpPortStats
	port.OfpPort.GetPortNo()
	var entry ofp.PortStatsEntry
	entry.SetPortNo(ofp.Port(port.OfpPort.GetPortNo()))
	entry.SetRxPackets(stats.GetRxPackets())
	entry.SetTxPackets(stats.GetTxPackets())
	entry.SetRxBytes(stats.GetRxBytes())
	entry.SetTxBytes(stats.GetTxBytes())
	entry.SetRxDropped(stats.GetRxDropped())
	entry.SetTxDropped(stats.GetTxDropped())
	entry.SetRxErrors(stats.GetRxErrors())
	entry.SetTxErrors(stats.GetTxErrors())
	entry.SetRxFrameErr(stats.GetRxFrameErr())
	entry.SetRxOverErr(stats.GetRxOverErr())
	entry.SetRxCrcErr(stats.GetRxCrcErr())
	entry.SetCollisions(stats.GetCollisions())
	entry.SetDurationSec(stats.GetDurationSec())
	entry.SetDurationNsec(stats.GetDurationNsec())
	return &entry
}
