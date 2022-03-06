import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "hupayxcom.carrotmember.carrotmember";
export interface MsgAddMember {
    creator: string;
    memberAddress: string;
    totalCnt: number;
    amount: Coin | undefined;
}
export interface MsgAddMemberResponse {
}
export interface MsgSendReward {
    creator: string;
}
export interface MsgSendRewardResponse {
}
export declare const MsgAddMember: {
    encode(message: MsgAddMember, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddMember;
    fromJSON(object: any): MsgAddMember;
    toJSON(message: MsgAddMember): unknown;
    fromPartial(object: DeepPartial<MsgAddMember>): MsgAddMember;
};
export declare const MsgAddMemberResponse: {
    encode(_: MsgAddMemberResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddMemberResponse;
    fromJSON(_: any): MsgAddMemberResponse;
    toJSON(_: MsgAddMemberResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddMemberResponse>): MsgAddMemberResponse;
};
export declare const MsgSendReward: {
    encode(message: MsgSendReward, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendReward;
    fromJSON(object: any): MsgSendReward;
    toJSON(message: MsgSendReward): unknown;
    fromPartial(object: DeepPartial<MsgSendReward>): MsgSendReward;
};
export declare const MsgSendRewardResponse: {
    encode(_: MsgSendRewardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendRewardResponse;
    fromJSON(_: any): MsgSendRewardResponse;
    toJSON(_: MsgSendRewardResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendRewardResponse>): MsgSendRewardResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    AddMember(request: MsgAddMember): Promise<MsgAddMemberResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SendReward(request: MsgSendReward): Promise<MsgSendRewardResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    AddMember(request: MsgAddMember): Promise<MsgAddMemberResponse>;
    SendReward(request: MsgSendReward): Promise<MsgSendRewardResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
