import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStoreCode } from "./types/cosmwasm/wasm/v1/tx";
import { MsgIBCSend } from "./types/cosmwasm/wasm/v1/ibc";
import { MsgExecuteContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgClearAdmin } from "./types/cosmwasm/wasm/v1/tx";
import { MsgUpdateAdmin } from "./types/cosmwasm/wasm/v1/tx";
import { MsgInstantiateContract2 } from "./types/cosmwasm/wasm/v1/tx";
import { MsgInstantiateContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgMigrateContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgIBCCloseChannel } from "./types/cosmwasm/wasm/v1/ibc";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cosmwasm.wasm.v1.MsgStoreCode", MsgStoreCode],
    ["/cosmwasm.wasm.v1.MsgIBCSend", MsgIBCSend],
    ["/cosmwasm.wasm.v1.MsgExecuteContract", MsgExecuteContract],
    ["/cosmwasm.wasm.v1.MsgClearAdmin", MsgClearAdmin],
    ["/cosmwasm.wasm.v1.MsgUpdateAdmin", MsgUpdateAdmin],
    ["/cosmwasm.wasm.v1.MsgInstantiateContract2", MsgInstantiateContract2],
    ["/cosmwasm.wasm.v1.MsgInstantiateContract", MsgInstantiateContract],
    ["/cosmwasm.wasm.v1.MsgMigrateContract", MsgMigrateContract],
    ["/cosmwasm.wasm.v1.MsgIBCCloseChannel", MsgIBCCloseChannel],
    
];

export { msgTypes }