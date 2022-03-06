/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "hupayxcom.carrotmember.carrotmember";

export interface Member {
  creator: string;
  memberAddress: string;
  totalCnt: number;
  remainCnt: number;
  /** google.protobuf.Timestamp create_time = 6; */
  amount: Coin | undefined;
}

const baseMember: object = {
  creator: "",
  memberAddress: "",
  totalCnt: 0,
  remainCnt: 0,
};

export const Member = {
  encode(message: Member, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.memberAddress !== "") {
      writer.uint32(18).string(message.memberAddress);
    }
    if (message.totalCnt !== 0) {
      writer.uint32(24).uint64(message.totalCnt);
    }
    if (message.remainCnt !== 0) {
      writer.uint32(32).uint64(message.remainCnt);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Member {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMember } as Member;
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
          message.remainCnt = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Member {
    const message = { ...baseMember } as Member;
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
    if (object.remainCnt !== undefined && object.remainCnt !== null) {
      message.remainCnt = Number(object.remainCnt);
    } else {
      message.remainCnt = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },

  toJSON(message: Member): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.memberAddress !== undefined &&
      (obj.memberAddress = message.memberAddress);
    message.totalCnt !== undefined && (obj.totalCnt = message.totalCnt);
    message.remainCnt !== undefined && (obj.remainCnt = message.remainCnt);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Member>): Member {
    const message = { ...baseMember } as Member;
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
    if (object.remainCnt !== undefined && object.remainCnt !== null) {
      message.remainCnt = object.remainCnt;
    } else {
      message.remainCnt = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },
};

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
