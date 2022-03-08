import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../carrotmember/params";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Member } from "../carrotmember/member";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "hupayxcom.carrotmember.carrotmember";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryMembersRequest {
    /** Adding pagination to request */
    pagination: PageRequest | undefined;
}
export interface QueryMembersResponse {
    /** Returning a list of posts */
    Member: Member[];
    /** Adding pagination to response */
    pagination: PageResponse | undefined;
}
export interface QueryNextRewardTimeRequest {
}
export interface QueryNextRewardTimeResponse {
    nextTime: string;
}
export interface QueryRewardPoolRequest {
}
export interface QueryRewardPoolResponse {
    amount: Coin[];
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryMembersRequest: {
    encode(message: QueryMembersRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryMembersRequest;
    fromJSON(object: any): QueryMembersRequest;
    toJSON(message: QueryMembersRequest): unknown;
    fromPartial(object: DeepPartial<QueryMembersRequest>): QueryMembersRequest;
};
export declare const QueryMembersResponse: {
    encode(message: QueryMembersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryMembersResponse;
    fromJSON(object: any): QueryMembersResponse;
    toJSON(message: QueryMembersResponse): unknown;
    fromPartial(object: DeepPartial<QueryMembersResponse>): QueryMembersResponse;
};
export declare const QueryNextRewardTimeRequest: {
    encode(_: QueryNextRewardTimeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryNextRewardTimeRequest;
    fromJSON(_: any): QueryNextRewardTimeRequest;
    toJSON(_: QueryNextRewardTimeRequest): unknown;
    fromPartial(_: DeepPartial<QueryNextRewardTimeRequest>): QueryNextRewardTimeRequest;
};
export declare const QueryNextRewardTimeResponse: {
    encode(message: QueryNextRewardTimeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryNextRewardTimeResponse;
    fromJSON(object: any): QueryNextRewardTimeResponse;
    toJSON(message: QueryNextRewardTimeResponse): unknown;
    fromPartial(object: DeepPartial<QueryNextRewardTimeResponse>): QueryNextRewardTimeResponse;
};
export declare const QueryRewardPoolRequest: {
    encode(_: QueryRewardPoolRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryRewardPoolRequest;
    fromJSON(_: any): QueryRewardPoolRequest;
    toJSON(_: QueryRewardPoolRequest): unknown;
    fromPartial(_: DeepPartial<QueryRewardPoolRequest>): QueryRewardPoolRequest;
};
export declare const QueryRewardPoolResponse: {
    encode(message: QueryRewardPoolResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryRewardPoolResponse;
    fromJSON(object: any): QueryRewardPoolResponse;
    toJSON(message: QueryRewardPoolResponse): unknown;
    fromPartial(object: DeepPartial<QueryRewardPoolResponse>): QueryRewardPoolResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of members items. */
    Members(request: QueryMembersRequest): Promise<QueryMembersResponse>;
    /** Queries a list of nextRewardTime items. */
    NextRewardTime(request: QueryNextRewardTimeRequest): Promise<QueryNextRewardTimeResponse>;
    /** Queries a list of rewardPool items. */
    RewardPool(request: QueryRewardPoolRequest): Promise<QueryRewardPoolResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Members(request: QueryMembersRequest): Promise<QueryMembersResponse>;
    NextRewardTime(request: QueryNextRewardTimeRequest): Promise<QueryNextRewardTimeResponse>;
    RewardPool(request: QueryRewardPoolRequest): Promise<QueryRewardPoolResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
