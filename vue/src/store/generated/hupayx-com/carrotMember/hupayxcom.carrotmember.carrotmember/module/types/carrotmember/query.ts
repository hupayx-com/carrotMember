/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../carrotmember/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Member } from "../carrotmember/member";

export const protobufPackage = "hupayxcom.carrotmember.carrotmember";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

export interface QueryNextRewardTimeRequest {}

export interface QueryNextRewardTimeResponse {
  nextTime: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryMembersRequest: object = {};

export const QueryMembersRequest = {
  encode(
    message: QueryMembersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMembersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMembersRequest } as QueryMembersRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembersRequest {
    const message = { ...baseQueryMembersRequest } as QueryMembersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMembersRequest>): QueryMembersRequest {
    const message = { ...baseQueryMembersRequest } as QueryMembersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMembersResponse: object = {};

export const QueryMembersResponse = {
  encode(
    message: QueryMembersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Member) {
      Member.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMembersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMembersResponse } as QueryMembersResponse;
    message.Member = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Member.push(Member.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembersResponse {
    const message = { ...baseQueryMembersResponse } as QueryMembersResponse;
    message.Member = [];
    if (object.Member !== undefined && object.Member !== null) {
      for (const e of object.Member) {
        message.Member.push(Member.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembersResponse): unknown {
    const obj: any = {};
    if (message.Member) {
      obj.Member = message.Member.map((e) =>
        e ? Member.toJSON(e) : undefined
      );
    } else {
      obj.Member = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMembersResponse>): QueryMembersResponse {
    const message = { ...baseQueryMembersResponse } as QueryMembersResponse;
    message.Member = [];
    if (object.Member !== undefined && object.Member !== null) {
      for (const e of object.Member) {
        message.Member.push(Member.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryNextRewardTimeRequest: object = {};

export const QueryNextRewardTimeRequest = {
  encode(
    _: QueryNextRewardTimeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryNextRewardTimeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryNextRewardTimeRequest,
    } as QueryNextRewardTimeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryNextRewardTimeRequest {
    const message = {
      ...baseQueryNextRewardTimeRequest,
    } as QueryNextRewardTimeRequest;
    return message;
  },

  toJSON(_: QueryNextRewardTimeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryNextRewardTimeRequest>
  ): QueryNextRewardTimeRequest {
    const message = {
      ...baseQueryNextRewardTimeRequest,
    } as QueryNextRewardTimeRequest;
    return message;
  },
};

const baseQueryNextRewardTimeResponse: object = { nextTime: "" };

export const QueryNextRewardTimeResponse = {
  encode(
    message: QueryNextRewardTimeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nextTime !== "") {
      writer.uint32(10).string(message.nextTime);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryNextRewardTimeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryNextRewardTimeResponse,
    } as QueryNextRewardTimeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextTime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryNextRewardTimeResponse {
    const message = {
      ...baseQueryNextRewardTimeResponse,
    } as QueryNextRewardTimeResponse;
    if (object.nextTime !== undefined && object.nextTime !== null) {
      message.nextTime = String(object.nextTime);
    } else {
      message.nextTime = "";
    }
    return message;
  },

  toJSON(message: QueryNextRewardTimeResponse): unknown {
    const obj: any = {};
    message.nextTime !== undefined && (obj.nextTime = message.nextTime);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryNextRewardTimeResponse>
  ): QueryNextRewardTimeResponse {
    const message = {
      ...baseQueryNextRewardTimeResponse,
    } as QueryNextRewardTimeResponse;
    if (object.nextTime !== undefined && object.nextTime !== null) {
      message.nextTime = object.nextTime;
    } else {
      message.nextTime = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of members items. */
  Members(request: QueryMembersRequest): Promise<QueryMembersResponse>;
  /** Queries a list of nextRewardTime items. */
  NextRewardTime(
    request: QueryNextRewardTimeRequest
  ): Promise<QueryNextRewardTimeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "hupayxcom.carrotmember.carrotmember.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Members(request: QueryMembersRequest): Promise<QueryMembersResponse> {
    const data = QueryMembersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "hupayxcom.carrotmember.carrotmember.Query",
      "Members",
      data
    );
    return promise.then((data) =>
      QueryMembersResponse.decode(new Reader(data))
    );
  }

  NextRewardTime(
    request: QueryNextRewardTimeRequest
  ): Promise<QueryNextRewardTimeResponse> {
    const data = QueryNextRewardTimeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "hupayxcom.carrotmember.carrotmember.Query",
      "NextRewardTime",
      data
    );
    return promise.then((data) =>
      QueryNextRewardTimeResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
