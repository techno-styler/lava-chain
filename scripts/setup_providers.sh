#!/bin/bash
__dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
. ${__dir}/vars/variables.sh
LOGS_DIR=${__dir}/../testutil/debugging/logs
mkdir -p $LOGS_DIR
rm $LOGS_DIR/*.log

echo "---------------Setup Providers------------------"
killall screen
screen -wipe

EXTRA_PROVIDER_FLAGS="$EXTRA_PROVIDER_FLAGS --chain-id=lava"
EXTRA_PORTAL_FLAGS="$EXTRA_PORTAL_FLAGS --chain-id=lava"

PROVIDER1_LISTENER="127.0.0.1:2221"
PROVIDER2_LISTENER="127.0.0.1:2222"
PROVIDER3_LISTENER="127.0.0.1:2223"

screen -d -m -S provider1 bash -c "source ~/.bashrc; lavap rpcprovider \
$PROVIDER1_LISTENER ETH1 jsonrpc '$ETH_RPC_WS' \
$PROVIDER1_LISTENER GTH1 jsonrpc '$GTH_RPC_WS' \
$PROVIDER1_LISTENER FTM250 jsonrpc '$FTM_RPC_HTTP' \
$PROVIDER1_LISTENER CELO jsonrpc '$CELO_HTTP' \
$PROVIDER1_LISTENER ALFAJORES jsonrpc '$CELO_ALFAJORES_HTTP' \
$PROVIDER1_LISTENER ARB1 jsonrpc '$ARB1_HTTP' \
$PROVIDER1_LISTENER APT1 rest '$APTOS_REST' \
$PROVIDER1_LISTENER STRK jsonrpc '$STARKNET_RPC' \
$PROVIDER1_LISTENER POLYGON1 jsonrpc '$POLYGON_MAINNET_RPC' \
$PROVIDER1_LISTENER OPTM jsonrpc '$OPTIMISM_RPC' \
$PROVIDER1_LISTENER BASET jsonrpc '$BASE_GOERLI_RPC' \
$PROVIDER1_LISTENER BSC jsonrpc '$BSC_RPC' \
$PROVIDER1_LISTENER SOLANA jsonrpc '$SOLANA_RPC' \
$PROVIDER1_LISTENER SUIT jsonrpc '$SUI_RPC' \
$PROVIDER1_LISTENER COS3 rest '$OSMO_REST' \
$PROVIDER1_LISTENER COS3 tendermintrpc '$OSMO_RPC,$OSMO_RPC' \
$PROVIDER1_LISTENER COS3 grpc '$OSMO_GRPC' \
$PROVIDER1_LISTENER LAV1 rest '$LAVA_REST' \
$PROVIDER1_LISTENER LAV1 tendermintrpc '$LAVA_RPC,$LAVA_RPC' \
$PROVIDER1_LISTENER LAV1 grpc '$LAVA_GRPC' \
$PROVIDER1_LISTENER COS5 rest '$GAIA_REST' \
$PROVIDER1_LISTENER COS5 tendermintrpc '$GAIA_RPC,$GAIA_RPC' \
$PROVIDER1_LISTENER COS5 grpc '$GAIA_GRPC' \
$PROVIDER1_LISTENER JUN1 rest '$JUNO_REST' \
$PROVIDER1_LISTENER JUN1 tendermintrpc '$JUNO_RPC,$JUNO_RPC' \
$PROVIDER1_LISTENER JUN1 grpc '$JUNO_GRPC' \
$PROVIDER1_LISTENER EVMOS jsonrpc '$EVMOS_RPC' \
$PROVIDER1_LISTENER EVMOS tendermintrpc '$EVMOS_TENDERMINTRPC,$EVMOS_TENDERMINTRPC' \
$PROVIDER1_LISTENER EVMOS rest '$EVMOS_REST' \
$PROVIDER1_LISTENER EVMOS grpc '$EVMOS_GRPC' \
$PROVIDER1_LISTENER CANTO jsonrpc '$CANTO_RPC' \
$PROVIDER1_LISTENER CANTO tendermintrpc '$CANTO_TENDERMINT,$CANTO_TENDERMINT' \
$PROVIDER1_LISTENER CANTO rest '$CANTO_REST' \
$PROVIDER1_LISTENER CANTO grpc '$CANTO_GRPC' \
$PROVIDER1_LISTENER AXELAR tendermintrpc '$AXELAR_RPC_HTTP,$AXELAR_RPC_HTTP' \
$PROVIDER1_LISTENER AXELAR rest '$AXELAR_REST' \
$PROVIDER1_LISTENER AXELAR grpc '$AXELAR_GRPC' \
$PROVIDER1_LISTENER AVAX jsonrpc '$AVALANCH_PJRPC' \
$PROVIDER1_LISTENER FVM jsonrpc '$FVM_JRPC' \
$EXTRA_PROVIDER_FLAGS --metrics-listen-address ":7780" --geolocation 1 --log_level debug --from servicer1 2>&1 | tee $LOGS_DIR/PROVIDER1.log" && sleep 0.25
# $PROVIDER1_LISTENER MANTLE jsonrpc '$MANTLE_JRPC' \

screen -d -m -S provider2 bash -c "source ~/.bashrc; lavap rpcprovider \
$PROVIDER2_LISTENER ETH1 jsonrpc '$ETH_RPC_WS' \
$PROVIDER2_LISTENER GTH1 jsonrpc '$GTH_RPC_WS' \
$PROVIDER2_LISTENER FTM250 jsonrpc '$FTM_RPC_HTTP' \
$PROVIDER2_LISTENER CELO jsonrpc '$CELO_HTTP' \
$PROVIDER2_LISTENER ALFAJORES jsonrpc '$CELO_ALFAJORES_HTTP' \
$PROVIDER2_LISTENER ARB1 jsonrpc '$ARB1_HTTP' \
$PROVIDER2_LISTENER APT1 rest '$APTOS_REST' \
$PROVIDER2_LISTENER STRK jsonrpc '$STARKNET_RPC' \
$PROVIDER2_LISTENER POLYGON1 jsonrpc '$POLYGON_MAINNET_RPC' \
$PROVIDER2_LISTENER OPTM jsonrpc '$OPTIMISM_RPC' \
$PROVIDER2_LISTENER BASET jsonrpc '$BASE_GOERLI_RPC' \
$PROVIDER2_LISTENER BSC jsonrpc '$BSC_RPC' \
$PROVIDER2_LISTENER SOLANA jsonrpc '$SOLANA_RPC' \
$PROVIDER2_LISTENER SUIT jsonrpc '$SUI_RPC' \
$PROVIDER2_LISTENER COS3 rest '$OSMO_REST' \
$PROVIDER2_LISTENER COS3 tendermintrpc '$OSMO_RPC,$OSMO_RPC' \
$PROVIDER2_LISTENER COS3 grpc '$OSMO_GRPC' \
$PROVIDER2_LISTENER LAV1 rest '$LAVA_REST' \
$PROVIDER2_LISTENER LAV1 tendermintrpc '$LAVA_RPC,$LAVA_RPC' \
$PROVIDER2_LISTENER LAV1 grpc '$LAVA_GRPC' \
$PROVIDER2_LISTENER COS5 rest '$GAIA_REST' \
$PROVIDER2_LISTENER COS5 tendermintrpc '$GAIA_RPC,$GAIA_RPC' \
$PROVIDER2_LISTENER COS5 grpc '$GAIA_GRPC' \
$PROVIDER2_LISTENER JUN1 rest '$JUNO_REST' \
$PROVIDER2_LISTENER JUN1 tendermintrpc '$JUNO_RPC,$JUNO_RPC' \
$PROVIDER2_LISTENER JUN1 grpc '$JUNO_GRPC' \
$PROVIDER2_LISTENER EVMOS jsonrpc '$EVMOS_RPC' \
$PROVIDER2_LISTENER EVMOS tendermintrpc '$EVMOS_TENDERMINTRPC,$EVMOS_TENDERMINTRPC' \
$PROVIDER2_LISTENER EVMOS rest '$EVMOS_REST' \
$PROVIDER2_LISTENER EVMOS grpc '$EVMOS_GRPC' \
$PROVIDER2_LISTENER CANTO jsonrpc '$CANTO_RPC' \
$PROVIDER2_LISTENER CANTO tendermintrpc '$CANTO_TENDERMINT,$CANTO_TENDERMINT' \
$PROVIDER2_LISTENER CANTO rest '$CANTO_REST' \
$PROVIDER2_LISTENER CANTO grpc '$CANTO_GRPC' \
$PROVIDER2_LISTENER AXELAR tendermintrpc '$AXELAR_RPC_HTTP,$AXELAR_RPC_HTTP' \
$PROVIDER2_LISTENER AXELAR rest '$AXELAR_REST' \
$PROVIDER2_LISTENER AXELAR grpc '$AXELAR_GRPC' \
$PROVIDER2_LISTENER AVAX jsonrpc '$AVALANCH_PJRPC' \
$PROVIDER2_LISTENER FVM jsonrpc '$FVM_JRPC' \
$EXTRA_PROVIDER_FLAGS --geolocation 1 --log_level debug --from servicer2 --chain-id lava 2>&1 | tee $LOGS_DIR/PROVIDER2.log" && sleep 0.25
# $PROVIDER2_LISTENER MANTLE jsonrpc '$MANTLE_JRPC' \

screen -d -m -S provider3 bash -c "source ~/.bashrc; lavap rpcprovider \
$PROVIDER3_LISTENER ETH1 jsonrpc '$ETH_RPC_WS' \
$PROVIDER3_LISTENER GTH1 jsonrpc '$GTH_RPC_WS' \
$PROVIDER3_LISTENER FTM250 jsonrpc '$FTM_RPC_HTTP' \
$PROVIDER3_LISTENER CELO jsonrpc '$CELO_HTTP' \
$PROVIDER3_LISTENER ALFAJORES jsonrpc '$CELO_ALFAJORES_HTTP' \
$PROVIDER3_LISTENER ARB1 jsonrpc '$ARB1_HTTP' \
$PROVIDER3_LISTENER APT1 rest '$APTOS_REST' \
$PROVIDER3_LISTENER STRK jsonrpc '$STARKNET_RPC' \
$PROVIDER3_LISTENER POLYGON1 jsonrpc '$POLYGON_MAINNET_RPC' \
$PROVIDER3_LISTENER OPTM jsonrpc '$OPTIMISM_RPC' \
$PROVIDER3_LISTENER BASET jsonrpc '$BASE_GOERLI_RPC' \
$PROVIDER3_LISTENER BSC jsonrpc '$BSC_RPC' \
$PROVIDER3_LISTENER SOLANA jsonrpc '$SOLANA_RPC' \
$PROVIDER3_LISTENER SUIT jsonrpc '$SUI_RPC' \
$PROVIDER3_LISTENER COS3 rest '$OSMO_REST' \
$PROVIDER3_LISTENER COS3 tendermintrpc '$OSMO_RPC,$OSMO_RPC' \
$PROVIDER3_LISTENER COS3 grpc '$OSMO_GRPC' \
$PROVIDER3_LISTENER LAV1 rest '$LAVA_REST' \
$PROVIDER3_LISTENER LAV1 tendermintrpc '$LAVA_RPC,$LAVA_RPC' \
$PROVIDER3_LISTENER LAV1 grpc '$LAVA_GRPC' \
$PROVIDER3_LISTENER COS5 rest '$GAIA_REST' \
$PROVIDER3_LISTENER COS5 tendermintrpc '$GAIA_RPC,$GAIA_RPC' \
$PROVIDER3_LISTENER COS5 grpc '$GAIA_GRPC' \
$PROVIDER3_LISTENER JUN1 rest '$JUNO_REST' \
$PROVIDER3_LISTENER JUN1 tendermintrpc '$JUNO_RPC,$JUNO_RPC' \
$PROVIDER3_LISTENER JUN1 grpc '$JUNO_GRPC' \
$PROVIDER3_LISTENER EVMOS jsonrpc '$EVMOS_RPC' \
$PROVIDER3_LISTENER EVMOS tendermintrpc '$EVMOS_TENDERMINTRPC,$EVMOS_TENDERMINTRPC' \
$PROVIDER3_LISTENER EVMOS rest '$EVMOS_REST' \
$PROVIDER3_LISTENER EVMOS grpc '$EVMOS_GRPC' \
$PROVIDER3_LISTENER CANTO jsonrpc '$CANTO_RPC' \
$PROVIDER3_LISTENER CANTO tendermintrpc '$CANTO_TENDERMINT,$CANTO_TENDERMINT' \
$PROVIDER3_LISTENER CANTO rest '$CANTO_REST' \
$PROVIDER3_LISTENER CANTO grpc '$CANTO_GRPC' \
$PROVIDER3_LISTENER AXELAR tendermintrpc '$AXELAR_RPC_HTTP,$AXELAR_RPC_HTTP' \
$PROVIDER3_LISTENER AXELAR rest '$AXELAR_REST' \
$PROVIDER3_LISTENER AXELAR grpc '$AXELAR_GRPC' \
$PROVIDER3_LISTENER AVAX jsonrpc '$AVALANCH_PJRPC' \
$PROVIDER3_LISTENER FVM jsonrpc '$FVM_JRPC' \
$EXTRA_PROVIDER_FLAGS --geolocation 1 --log_level debug --from servicer3 --chain-id lava 2>&1 | tee $LOGS_DIR/PROVIDER3.log" && sleep 0.25
# $PROVIDER3_LISTENER MANTLE jsonrpc '$MANTLE_JRPC' \

# Setup Portal
screen -d -m -S portals bash -c "source ~/.bashrc; lavap rpcconsumer \
127.0.0.1:3333 ETH1 jsonrpc \
127.0.0.1:3334 GTH1 jsonrpc \
127.0.0.1:3335 FTM250 jsonrpc \
127.0.0.1:3346 CELO jsonrpc \
127.0.0.1:3347 ALFAJORES jsonrpc \
127.0.0.1:3348 ARB1 jsonrpc \
127.0.0.1:3349 STRK jsonrpc \
127.0.0.1:3350 APT1 rest \
127.0.0.1:3351 POLYGON1 jsonrpc \
127.0.0.1:3352 OPTM jsonrpc \
127.0.0.1:3353 BASET jsonrpc \
127.0.0.1:3354 COS3 rest 127.0.0.1:3355 COS3 tendermintrpc 127.0.0.1:3356 COS3 grpc \
127.0.0.1:3357 COS4 rest 127.0.0.1:3358 COS4 tendermintrpc 127.0.0.1:3359 COS4 grpc \
127.0.0.1:3360 LAV1 rest 127.0.0.1:3361 LAV1 tendermintrpc 127.0.0.1:3362 LAV1 grpc \
127.0.0.1:3363 COS5 rest 127.0.0.1:3364 COS5 tendermintrpc 127.0.0.1:3365 COS5 grpc \
127.0.0.1:3366 JUN1 rest 127.0.0.1:3367 JUN1 tendermintrpc 127.0.0.1:3368 JUN1 grpc \
127.0.0.1:3369 EVMOS jsonrpc 127.0.0.1:3370 EVMOS rest 127.0.0.1:3371 EVMOS tendermintrpc 127.0.0.1:3372 EVMOS grpc \
127.0.0.1:3373 CANTO jsonrpc 127.0.0.1:3374 CANTO rest 127.0.0.1:3375 CANTO tendermintrpc 127.0.0.1:3376 CANTO grpc \
127.0.0.1:3377 AXELAR rest 127.0.0.1:3378 AXELAR tendermintrpc 127.0.0.1:3379 AXELAR grpc \
127.0.0.1:3380 BSC jsonrpc \
127.0.0.1:3381 SOLANA jsonrpc \
127.0.0.1:3382 SUIT jsonrpc \
127.0.0.1:3383 AVAX jsonrpc \
127.0.0.1:3384 FVM jsonrpc \
$EXTRA_PORTAL_FLAGS --metrics-listen-address ":7779" --geolocation 1 --log_level debug --from user1 --chain-id lava --allow-insecure-provider-dialing 2>&1 | tee $LOGS_DIR/PORTAL.log" && sleep 0.25
# 127.0.0.1:3385 MANTLE jsonrpc \

echo "--- setting up screens done ---"
screen -ls