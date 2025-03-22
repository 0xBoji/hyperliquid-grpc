// @generated by protobuf-ts 2.8.2 with parameter long_type_number
// @generated from protobuf file "hypurr/evm_service.proto" (package "hypurr", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { EVM } from "./evm_service";
import type { UniV2CandlesResponse } from "./evm_service";
import type { UniV2CandlesRequest } from "./evm_service";
import type { ERC20ApprovalEventsResponse } from "./evm_service";
import type { ERC20ApprovalEventsRequest } from "./evm_service";
import type { ERC20TransferEventsResponse } from "./evm_service";
import type { ERC20TransferEventsRequest } from "./evm_service";
import type { UniV2SwapsResponse } from "./evm_service";
import type { UniV2SwapsRequest } from "./evm_service";
import type { UniV2SwapResponse } from "./evm_service";
import type { UniV2SwapRequest } from "./evm_service";
import type { UniV2PairsResponse } from "./evm_service";
import type { UniV2PairsRequest } from "./evm_service";
import type { UniV2PairResponse } from "./evm_service";
import type { UniV2PairRequest } from "./evm_service";
import type { ERC20TokensResponse } from "./evm_service";
import type { ERC20TokensRequest } from "./evm_service";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ERC20TokenResponse } from "./evm_service";
import type { ERC20TokenRequest } from "./evm_service";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service hypurr.EVM
 */
export interface IEVMClient {
    /**
     * Token related endpoints
     *
     * @generated from protobuf rpc: ERC20Token(hypurr.ERC20TokenRequest) returns (hypurr.ERC20TokenResponse);
     */
    eRC20Token(input: ERC20TokenRequest, options?: RpcOptions): UnaryCall<ERC20TokenRequest, ERC20TokenResponse>;
    /**
     * @generated from protobuf rpc: ERC20Tokens(hypurr.ERC20TokensRequest) returns (hypurr.ERC20TokensResponse);
     */
    eRC20Tokens(input: ERC20TokensRequest, options?: RpcOptions): UnaryCall<ERC20TokensRequest, ERC20TokensResponse>;
    /**
     * Pair related endpoints
     *
     * @generated from protobuf rpc: UniV2Pair(hypurr.UniV2PairRequest) returns (hypurr.UniV2PairResponse);
     */
    uniV2Pair(input: UniV2PairRequest, options?: RpcOptions): UnaryCall<UniV2PairRequest, UniV2PairResponse>;
    /**
     * @generated from protobuf rpc: UniV2Pairs(hypurr.UniV2PairsRequest) returns (hypurr.UniV2PairsResponse);
     */
    uniV2Pairs(input: UniV2PairsRequest, options?: RpcOptions): UnaryCall<UniV2PairsRequest, UniV2PairsResponse>;
    /**
     * Swap related endpoints
     *
     * @generated from protobuf rpc: UniV2Swap(hypurr.UniV2SwapRequest) returns (hypurr.UniV2SwapResponse);
     */
    uniV2Swap(input: UniV2SwapRequest, options?: RpcOptions): UnaryCall<UniV2SwapRequest, UniV2SwapResponse>;
    /**
     * @generated from protobuf rpc: UniV2Swaps(hypurr.UniV2SwapsRequest) returns (hypurr.UniV2SwapsResponse);
     */
    uniV2Swaps(input: UniV2SwapsRequest, options?: RpcOptions): UnaryCall<UniV2SwapsRequest, UniV2SwapsResponse>;
    /**
     * Event related endpoints
     *
     * @generated from protobuf rpc: ERC20TransferEvents(hypurr.ERC20TransferEventsRequest) returns (hypurr.ERC20TransferEventsResponse);
     */
    eRC20TransferEvents(input: ERC20TransferEventsRequest, options?: RpcOptions): UnaryCall<ERC20TransferEventsRequest, ERC20TransferEventsResponse>;
    /**
     * @generated from protobuf rpc: ERC20ApprovalEvents(hypurr.ERC20ApprovalEventsRequest) returns (hypurr.ERC20ApprovalEventsResponse);
     */
    eRC20ApprovalEvents(input: ERC20ApprovalEventsRequest, options?: RpcOptions): UnaryCall<ERC20ApprovalEventsRequest, ERC20ApprovalEventsResponse>;
    /**
     * Price candle related endpoints
     *
     * @generated from protobuf rpc: UniV2Candles(hypurr.UniV2CandlesRequest) returns (hypurr.UniV2CandlesResponse);
     */
    uniV2Candles(input: UniV2CandlesRequest, options?: RpcOptions): UnaryCall<UniV2CandlesRequest, UniV2CandlesResponse>;
}
/**
 * @generated from protobuf service hypurr.EVM
 */
export class EVMClient implements IEVMClient, ServiceInfo {
    typeName = EVM.typeName;
    methods = EVM.methods;
    options = EVM.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Token related endpoints
     *
     * @generated from protobuf rpc: ERC20Token(hypurr.ERC20TokenRequest) returns (hypurr.ERC20TokenResponse);
     */
    eRC20Token(input: ERC20TokenRequest, options?: RpcOptions): UnaryCall<ERC20TokenRequest, ERC20TokenResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ERC20TokenRequest, ERC20TokenResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ERC20Tokens(hypurr.ERC20TokensRequest) returns (hypurr.ERC20TokensResponse);
     */
    eRC20Tokens(input: ERC20TokensRequest, options?: RpcOptions): UnaryCall<ERC20TokensRequest, ERC20TokensResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ERC20TokensRequest, ERC20TokensResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Pair related endpoints
     *
     * @generated from protobuf rpc: UniV2Pair(hypurr.UniV2PairRequest) returns (hypurr.UniV2PairResponse);
     */
    uniV2Pair(input: UniV2PairRequest, options?: RpcOptions): UnaryCall<UniV2PairRequest, UniV2PairResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UniV2PairRequest, UniV2PairResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UniV2Pairs(hypurr.UniV2PairsRequest) returns (hypurr.UniV2PairsResponse);
     */
    uniV2Pairs(input: UniV2PairsRequest, options?: RpcOptions): UnaryCall<UniV2PairsRequest, UniV2PairsResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<UniV2PairsRequest, UniV2PairsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Swap related endpoints
     *
     * @generated from protobuf rpc: UniV2Swap(hypurr.UniV2SwapRequest) returns (hypurr.UniV2SwapResponse);
     */
    uniV2Swap(input: UniV2SwapRequest, options?: RpcOptions): UnaryCall<UniV2SwapRequest, UniV2SwapResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<UniV2SwapRequest, UniV2SwapResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UniV2Swaps(hypurr.UniV2SwapsRequest) returns (hypurr.UniV2SwapsResponse);
     */
    uniV2Swaps(input: UniV2SwapsRequest, options?: RpcOptions): UnaryCall<UniV2SwapsRequest, UniV2SwapsResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<UniV2SwapsRequest, UniV2SwapsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Event related endpoints
     *
     * @generated from protobuf rpc: ERC20TransferEvents(hypurr.ERC20TransferEventsRequest) returns (hypurr.ERC20TransferEventsResponse);
     */
    eRC20TransferEvents(input: ERC20TransferEventsRequest, options?: RpcOptions): UnaryCall<ERC20TransferEventsRequest, ERC20TransferEventsResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<ERC20TransferEventsRequest, ERC20TransferEventsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ERC20ApprovalEvents(hypurr.ERC20ApprovalEventsRequest) returns (hypurr.ERC20ApprovalEventsResponse);
     */
    eRC20ApprovalEvents(input: ERC20ApprovalEventsRequest, options?: RpcOptions): UnaryCall<ERC20ApprovalEventsRequest, ERC20ApprovalEventsResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<ERC20ApprovalEventsRequest, ERC20ApprovalEventsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Price candle related endpoints
     *
     * @generated from protobuf rpc: UniV2Candles(hypurr.UniV2CandlesRequest) returns (hypurr.UniV2CandlesResponse);
     */
    uniV2Candles(input: UniV2CandlesRequest, options?: RpcOptions): UnaryCall<UniV2CandlesRequest, UniV2CandlesResponse> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<UniV2CandlesRequest, UniV2CandlesResponse>("unary", this._transport, method, opt, input);
    }
}
