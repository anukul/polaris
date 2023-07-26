HOMEDIR="/root/.polard"
CHAINID="brickchain-666"
KEY="$1"
KEYRING="test"
KEYALGO="eth_secp256k1"


polard config set client keyring-backend $KEYRING --home "$HOMEDIR"

polard keys add "$KEY" --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"

polard genesis add-genesis-account "$KEY" 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"

polard genesis gentx "$KEY" 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
