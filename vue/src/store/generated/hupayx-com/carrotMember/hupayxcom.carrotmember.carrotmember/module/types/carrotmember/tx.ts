/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "hupayxcom.carrotmember.carrotmember";

export interface MsgAddMember {
  creator: string;
  memberAddress: string;
  totalCnt: number;
  amount: Coin | undefined;
}

export interface MsgAddMemberResponse {}

export interface MsgSendReward {
  creator: string;
}

export interface MsgSendRewardResponse {}

const baseMsgAddMember: object = {
  creator: "",
  memberAddress: "",
  totalCnt: 0,
};

export const MsgAddMember = {
  encode(message: MsgAddMember, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.memberAddress !== "") {
      writer.uint32(18).string(message.memberAddress);
    }
    if (message.totalCnt !== 0) {
      writer.uint32(24).uint64(message.totalCnt);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddMember {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddMember } as MsgAddMember;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.memberAddress = reader.string();
          break;
        case 3:
          message.totalCnt = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddMember {
    const message = { ...baseMsgAddMember } as MsgAddMember;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.memberAddress !== undefined && object.memberAddress !== null) {
      message.memberAddress = String(object.memberAddress);
    } else {
      message.memberAddress = "";
    }
    if (object.totalCnt !== undefined && object.totalCnt !== null) {
      message.totalCnt = Number(object.totalCnt);
    } else {
      message.totalCnt = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddMember): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.memberAddress !== undefined &&
      (obj.memberAddress = message.memberAddress);
    message.totalCnt !== undefined && (obj.totalCnt = message.totalCnt);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddMember>): MsgAddMember {
    const message = { ...baseMsgAddMember } as MsgAddMember;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.memberAddress !== undefined && object.memberAddress !== null) {
      message.memberAddress = object.memberAddress;
    } else {
      message.memberAddress = "";
    }
    if (object.totalCnt !== undefined && object.totalCnt !== null) {
      message.totalCnt = object.totalCnt;
    } else {
      message.totalCnt = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },
};

const baseMsgAddMemberResponse: object = {};

export const MsgAddMemberResponse = {
  encode(_: MsgAddMemberResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddMemberResponse } as MsgAddMemberResponse;
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

  fromJSON(_: any): MsgAddMemberResponse {
    const message = { ...baseMsgAddMemberResponse } as MsgAddMemberResponse;
    return message;
  },

  toJSON(_: MsgAddMemberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddMemberResponse>): MsgAddMemberResponse {
    const message = { ...baseMsgAddMemberResponse } as MsgAddMemberResponse;
    return message;
  },
};

const baseMsgSendReward: object = { creator: "" };

export const MsgSendReward = {
  encode(message: MsgSendReward, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendReward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendReward } as MsgSendReward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendReward {
    const message = { ...baseMsgSendReward } as MsgSendReward;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgSendReward): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendReward>): MsgSendReward {
    const message = { ...baseMsgSendReward } as MsgSendReward;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgSendRewardResponse: object = {};

export const MsgSendRewardResponse = {
  encode(_: MsgSendRewardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendRewardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendRewardResponse } as MsgSendRewardResponse;
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

  fromJSON(_: any): MsgSendRewardResponse {
    const message = { ...baseMsgSendRewardResponse } as MsgSendRewardResponse;
    return message;
  },

  toJSON(_: MsgSendRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSendRewardResponse>): MsgSendRewardResponse {
    const message = { ...baseMsgSendRewardResponse } as MsgSendRewardResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddMember(request: MsgAddMember): Promise<MsgAddMemberResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendReward(request: MsgSendReward): Promise<MsgSendRewardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddMember(request: MsgAddMember): Promise<MsgAddMemberResponse> {
    const data = MsgAddMember.encode(request).finish();
    const promise = this.rpc.request(
      "hupayxcom.carrotmember.carrotmember.Msg",
      "AddMember",
      data
    );
    return promise.then((data) =>
      MsgAddMemberResponse.decode(new Reader(data))
    );
  }

  SendReward(request: MsgSendReward): Promise<MsgSendRewardResponse> {
    const data = MsgSendReward.encode(request).finish();
    const promise = this.rpc.request(
      "hupayxcom.carrotmember.carrotmember.Msg",
      "SendReward",
      data
    );
    return promise.then((data) =>
      MsgSendRewardResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
