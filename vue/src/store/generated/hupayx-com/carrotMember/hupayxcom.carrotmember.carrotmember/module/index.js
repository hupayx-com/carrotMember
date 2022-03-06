// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendReward } from "./types/carrotmember/tx";
import { MsgAddMember } from "./types/carrotmember/tx";
const types = [
    ["/hupayxcom.carrotmember.carrotmember.MsgSendReward", MsgSendReward],
    ["/hupayxcom.carrotmember.carrotmember.MsgAddMember", MsgAddMember],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendReward: (data) => ({ typeUrl: "/hupayxcom.carrotmember.carrotmember.MsgSendReward", value: MsgSendReward.fromPartial(data) }),
        msgAddMember: (data) => ({ typeUrl: "/hupayxcom.carrotmember.carrotmember.MsgAddMember", value: MsgAddMember.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
