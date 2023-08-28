import { StakeEntry } from "../../codec/lavanet/lava/epochstorage/stake_entry";

export interface StateQuery {
  fetchPairing(): Promise<number>;
  getPairing(chainID: string): PairingResponse | undefined;
}

export interface PairingResponse {
  providers: StakeEntry[];
  maxCu: number;
  currentEpoch: number;
}
