// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: hypurr/tools.proto

package hypurr

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HyperliquidSpotSniperConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TelegramId             int64   `protobuf:"varint,2,opt,name=telegram_id,json=telegramId,proto3" json:"telegram_id,omitempty"`
	Name                   string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetToken            string  `protobuf:"bytes,4,opt,name=target_token,json=targetToken,proto3" json:"target_token,omitempty"`
	MaxMarketCap           float64 `protobuf:"fixed64,5,opt,name=max_market_cap,json=maxMarketCap,proto3" json:"max_market_cap,omitempty"`
	MaxNotional            float64 `protobuf:"fixed64,6,opt,name=max_notional,json=maxNotional,proto3" json:"max_notional,omitempty"`
	MaxSlippage            int64   `protobuf:"varint,7,opt,name=max_slippage,json=maxSlippage,proto3" json:"max_slippage,omitempty"`
	MaxUserAllocation      int64   `protobuf:"varint,8,opt,name=max_user_allocation,json=maxUserAllocation,proto3" json:"max_user_allocation,omitempty"`
	MaxLiquidityAllocation int64   `protobuf:"varint,9,opt,name=max_liquidity_allocation,json=maxLiquidityAllocation,proto3" json:"max_liquidity_allocation,omitempty"`
	MaxAirdropAllocation   int64   `protobuf:"varint,10,opt,name=max_airdrop_allocation,json=maxAirdropAllocation,proto3" json:"max_airdrop_allocation,omitempty"`
	Enabled                bool    `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AntiRug                bool    `protobuf:"varint,12,opt,name=anti_rug,json=antiRug,proto3" json:"anti_rug,omitempty"`
}

func (x *HyperliquidSpotSniperConfig) Reset() {
	*x = HyperliquidSpotSniperConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypurr_tools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HyperliquidSpotSniperConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HyperliquidSpotSniperConfig) ProtoMessage() {}

func (x *HyperliquidSpotSniperConfig) ProtoReflect() protoreflect.Message {
	mi := &file_hypurr_tools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HyperliquidSpotSniperConfig.ProtoReflect.Descriptor instead.
func (*HyperliquidSpotSniperConfig) Descriptor() ([]byte, []int) {
	return file_hypurr_tools_proto_rawDescGZIP(), []int{0}
}

func (x *HyperliquidSpotSniperConfig) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetTelegramId() int64 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HyperliquidSpotSniperConfig) GetTargetToken() string {
	if x != nil {
		return x.TargetToken
	}
	return ""
}

func (x *HyperliquidSpotSniperConfig) GetMaxMarketCap() float64 {
	if x != nil {
		return x.MaxMarketCap
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetMaxNotional() float64 {
	if x != nil {
		return x.MaxNotional
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetMaxSlippage() int64 {
	if x != nil {
		return x.MaxSlippage
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetMaxUserAllocation() int64 {
	if x != nil {
		return x.MaxUserAllocation
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetMaxLiquidityAllocation() int64 {
	if x != nil {
		return x.MaxLiquidityAllocation
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetMaxAirdropAllocation() int64 {
	if x != nil {
		return x.MaxAirdropAllocation
	}
	return 0
}

func (x *HyperliquidSpotSniperConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HyperliquidSpotSniperConfig) GetAntiRug() bool {
	if x != nil {
		return x.AntiRug
	}
	return false
}

type HyperliquidWalletSpotTwapSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WalletId       int64                  `protobuf:"varint,2,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	PairId         int64                  `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	TelegramId     *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=telegram_id,json=telegramId,proto3" json:"telegram_id,omitempty"`
	TargetNotional float64                `protobuf:"fixed64,5,opt,name=target_notional,json=targetNotional,proto3" json:"target_notional,omitempty"`
	FilledQuantity float64                `protobuf:"fixed64,6,opt,name=filled_quantity,json=filledQuantity,proto3" json:"filled_quantity,omitempty"`
	FilledCost     float64                `protobuf:"fixed64,7,opt,name=filled_cost,json=filledCost,proto3" json:"filled_cost,omitempty"`
	Duration       int64                  `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Enabled        bool                   `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Buy            bool                   `protobuf:"varint,10,opt,name=buy,proto3" json:"buy,omitempty"`
	Ended          bool                   `protobuf:"varint,11,opt,name=ended,proto3" json:"ended,omitempty"`
}

func (x *HyperliquidWalletSpotTwapSession) Reset() {
	*x = HyperliquidWalletSpotTwapSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypurr_tools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HyperliquidWalletSpotTwapSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HyperliquidWalletSpotTwapSession) ProtoMessage() {}

func (x *HyperliquidWalletSpotTwapSession) ProtoReflect() protoreflect.Message {
	mi := &file_hypurr_tools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HyperliquidWalletSpotTwapSession.ProtoReflect.Descriptor instead.
func (*HyperliquidWalletSpotTwapSession) Descriptor() ([]byte, []int) {
	return file_hypurr_tools_proto_rawDescGZIP(), []int{1}
}

func (x *HyperliquidWalletSpotTwapSession) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetWalletId() int64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetPairId() int64 {
	if x != nil {
		return x.PairId
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetTelegramId() *wrapperspb.Int64Value {
	if x != nil {
		return x.TelegramId
	}
	return nil
}

func (x *HyperliquidWalletSpotTwapSession) GetTargetNotional() float64 {
	if x != nil {
		return x.TargetNotional
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetFilledQuantity() float64 {
	if x != nil {
		return x.FilledQuantity
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetFilledCost() float64 {
	if x != nil {
		return x.FilledCost
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *HyperliquidWalletSpotTwapSession) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HyperliquidWalletSpotTwapSession) GetBuy() bool {
	if x != nil {
		return x.Buy
	}
	return false
}

func (x *HyperliquidWalletSpotTwapSession) GetEnded() bool {
	if x != nil {
		return x.Ended
	}
	return false
}

var File_hypurr_tools_proto protoreflect.FileDescriptor

var file_hypurr_tools_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x79, 0x70, 0x75, 0x72, 0x72, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x79, 0x70, 0x75, 0x72, 0x72, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x03, 0x0a,
	0x1b, 0x48, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x53, 0x70, 0x6f, 0x74,
	0x53, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61,
	0x78, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x53, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d,
	0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x61,
	0x78, 0x5f, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x41,
	0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6e,
	0x74, 0x69, 0x5f, 0x72, 0x75, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x6e,
	0x74, 0x69, 0x52, 0x75, 0x67, 0x22, 0xf7, 0x02, 0x0a, 0x20, 0x48, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x70, 0x6f, 0x74, 0x54,
	0x77, 0x61, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x69, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x69, 0x72, 0x49, 0x64,
	0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x75, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x75, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79,
	0x70, 0x75, 0x72, 0x72, 0x2f, 0x68, 0x79, 0x70, 0x75, 0x72, 0x72, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x79, 0x70, 0x75, 0x72, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hypurr_tools_proto_rawDescOnce sync.Once
	file_hypurr_tools_proto_rawDescData = file_hypurr_tools_proto_rawDesc
)

func file_hypurr_tools_proto_rawDescGZIP() []byte {
	file_hypurr_tools_proto_rawDescOnce.Do(func() {
		file_hypurr_tools_proto_rawDescData = protoimpl.X.CompressGZIP(file_hypurr_tools_proto_rawDescData)
	})
	return file_hypurr_tools_proto_rawDescData
}

var file_hypurr_tools_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hypurr_tools_proto_goTypes = []any{
	(*HyperliquidSpotSniperConfig)(nil),      // 0: hypurr.HyperliquidSpotSniperConfig
	(*HyperliquidWalletSpotTwapSession)(nil), // 1: hypurr.HyperliquidWalletSpotTwapSession
	(*wrapperspb.Int64Value)(nil),            // 2: google.protobuf.Int64Value
}
var file_hypurr_tools_proto_depIdxs = []int32{
	2, // 0: hypurr.HyperliquidWalletSpotTwapSession.telegram_id:type_name -> google.protobuf.Int64Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hypurr_tools_proto_init() }
func file_hypurr_tools_proto_init() {
	if File_hypurr_tools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hypurr_tools_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HyperliquidSpotSniperConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hypurr_tools_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HyperliquidWalletSpotTwapSession); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hypurr_tools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hypurr_tools_proto_goTypes,
		DependencyIndexes: file_hypurr_tools_proto_depIdxs,
		MessageInfos:      file_hypurr_tools_proto_msgTypes,
	}.Build()
	File_hypurr_tools_proto = out.File
	file_hypurr_tools_proto_rawDesc = nil
	file_hypurr_tools_proto_goTypes = nil
	file_hypurr_tools_proto_depIdxs = nil
}
