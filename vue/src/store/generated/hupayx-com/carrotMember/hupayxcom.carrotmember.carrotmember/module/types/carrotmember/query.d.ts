import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../carrotmember/params";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Member } from "../carrotmember/member";
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
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of members items. */
    Members(request: QueryMembersRequest): Promise<QueryMembersResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Members(request: QueryMembersRequest): Promise<QueryMembersResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
