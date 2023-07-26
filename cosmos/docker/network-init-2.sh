# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

CONTAINER0="polard-node0"
CONTAINER1="polard-node1"
HOMEDIR="/root/.polard"
SCRIPTS="/scripts"

# init nodes
docker exec $CONTAINER0 bash -c "$SCRIPTS/init.sh $CONTAINER0"
docker exec $CONTAINER1 bash -c "$SCRIPTS/init.sh $CONTAINER1"

# create validators
docker exec $CONTAINER0 bash -c "$SCRIPTS/create-validator.sh $CONTAINER0"
docker exec $CONTAINER1 bash -c "$SCRIPTS/create-validator.sh $CONTAINER1"

# copy gentx to CONTAINER0
docker cp $CONTAINER1:$HOMEDIR/config/gentx ./temp/gentx
docker cp ./temp/gentx $CONTAINER0:$HOMEDIR/config

# update genesis file using gentx
docker exec $CONTAINER0 bash -c "$SCRIPTS/dump-genesis.sh"

# copy genesis file to all nodes
docker cp $CONTAINER0:$HOMEDIR/config/genesis.json ./temp/genesis.json
docker cp ./temp/genesis.json $CONTAINER1:$HOMEDIR/config/genesis.json

# start
# docker exec -it $CONTAINER0 bash -c "$SCRIPTS/seed-start.sh"
# docker exec -it $CONTAINER1 bash -c "$SCRIPTS/seed-start.sh"
