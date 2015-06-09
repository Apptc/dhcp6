// generated by stringer -output=string.go -type=DUIDType,MessageType,Status,OptionCode; DO NOT EDIT

package dhcp6

import "fmt"

const _DUIDType_name = "DUIDTypeLLTDUIDTypeENDUIDTypeLL"

var _DUIDType_index = [...]uint8{0, 11, 21, 31}

func (i DUIDType) String() string {
	i -= 1
	if i >= DUIDType(len(_DUIDType_index)-1) {
		return fmt.Sprintf("DUIDType(%d)", i+1)
	}
	return _DUIDType_name[_DUIDType_index[i]:_DUIDType_index[i+1]]
}

const _MessageType_name = "MessageTypeSolicitMessageTypeAdvertiseMessageTypeRequestMessageTypeConfirmMessageTypeRenewMessageTypeRebindMessageTypeReplyMessageTypeReleaseMessageTypeDeclineMessageTypeReconfigureMessageTypeInformationRequestMessageTypeRelayForwardMessageTypeRelayReply"

var _MessageType_index = [...]uint8{0, 18, 38, 56, 74, 90, 107, 123, 141, 159, 181, 210, 233, 254}

func (i MessageType) String() string {
	i -= 1
	if i >= MessageType(len(_MessageType_index)-1) {
		return fmt.Sprintf("MessageType(%d)", i+1)
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}

const _Status_name = "StatusSuccessStatusUnspecFailStatusNoAddrsAvailStatusNoBindingStatusNotOnLinkStatusUseMulticast"

var _Status_index = [...]uint8{0, 13, 29, 47, 62, 77, 95}

func (i Status) String() string {
	if i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}

const (
	_OptionCode_name_0 = "OptionClientIDOptionServerIDOptionIANAOptionIATAOptionIAAddrOptionOROOptionPreferenceOptionElapsedTimeOptionRelayMsg"
	_OptionCode_name_1 = "OptionAuthOptionUnicastOptionStatusCodeOptionRapidCommitOptionUserClassOptionVendorClassOptionVendorOptsOptionInterfaceIDOptionReconfMsgOptionReconfAccept"
)

var (
	_OptionCode_index_0 = [...]uint8{0, 14, 28, 38, 48, 60, 69, 85, 102, 116}
	_OptionCode_index_1 = [...]uint8{0, 10, 23, 39, 56, 71, 88, 104, 121, 136, 154}
)

func (i OptionCode) String() string {
	switch {
	case 1 <= i && i <= 9:
		i -= 1
		return _OptionCode_name_0[_OptionCode_index_0[i]:_OptionCode_index_0[i+1]]
	case 11 <= i && i <= 20:
		i -= 11
		return _OptionCode_name_1[_OptionCode_index_1[i]:_OptionCode_index_1[i+1]]
	default:
		return fmt.Sprintf("OptionCode(%d)", i)
	}
}
