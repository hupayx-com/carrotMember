/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";
export const protobufPackage = "hupayxcom.carrotmember.carrotmember";
const baseMsgAddMember = {
    creator: "",
    memberAddress: "",
    totalCnt: 0,
};
export const MsgAddMember = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddMember };
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
                    message.totalCnt = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = { ...baseMsgAddMember };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.memberAddress !== undefined && object.memberAddress !== null) {
            message.memberAddress = String(object.memberAddress);
        }
        else {
            message.memberAddress = "";
        }
        if (object.totalCnt !== undefined && object.totalCnt !== null) {
            message.totalCnt = Number(object.totalCnt);
        }
        else {
            message.totalCnt = 0;
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Coin.fromJSON(object.amount);
        }
        else {
            message.amount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.memberAddress !== undefined &&
            (obj.memberAddress = message.memberAddress);
        message.totalCnt !== undefined && (obj.totalCnt = message.totalCnt);
        message.amount !== undefined &&
            (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddMember };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.memberAddress !== undefined && object.memberAddress !== null) {
            message.memberAddress = object.memberAddress;
        }
        else {
            message.memberAddress = "";
        }
        if (object.totalCnt !== undefined && object.totalCnt !== null) {
            message.totalCnt = object.totalCnt;
        }
        else {
            message.totalCnt = 0;
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Coin.fromPartial(object.amount);
        }
        else {
            message.amount = undefined;
        }
        return message;
    },
};
const baseMsgAddMemberResponse = {};
export const MsgAddMemberResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddMemberResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgAddMemberResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddMemberResponse };
        return message;
    },
};
const baseMsgSendReward = { creator: "" };
export const MsgSendReward = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSendReward };
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
    fromJSON(object) {
        const message = { ...baseMsgSendReward };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSendReward };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseMsgSendRewardResponse = {};
export const MsgSendRewardResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSendRewardResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgSendRewardResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSendRewardResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    AddMember(request) {
        const data = MsgAddMember.encode(request).finish();
        const promise = this.rpc.request("hupayxcom.carrotmember.carrotmember.Msg", "AddMember", data);
        return promise.then((data) => MsgAddMemberResponse.decode(new Reader(data)));
    }
    SendReward(request) {
        const data = MsgSendReward.encode(request).finish();
        const promise = this.rpc.request("hupayxcom.carrotmember.carrotmember.Msg", "SendReward", data);
        return promise.then((data) => MsgSendRewardResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
