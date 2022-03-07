import { Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "hupayxcom.carrotmember.carrotmember";
export interface Member {
    creator: string;
    memberAddress: string;
    totalCnt: number;
    remainCnt: number;
    amount: Coin | undefined;
    createTime: Date | undefined;
}
export declare const Member: {
    encode(message: Member, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Member;
    fromJSON(object: any): Member;
    toJSON(message: Member): unknown;
    fromPartial(object: DeepPartial<Member>): Member;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
