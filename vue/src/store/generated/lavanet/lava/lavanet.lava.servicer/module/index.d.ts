import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgStakeServicer } from "./types/servicer/tx";
import { MsgProofOfWork } from "./types/servicer/tx";
import { MsgUnstakeServicer } from "./types/servicer/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgStakeServicer: (data: MsgStakeServicer) => EncodeObject;
    msgProofOfWork: (data: MsgProofOfWork) => EncodeObject;
    msgUnstakeServicer: (data: MsgUnstakeServicer) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
