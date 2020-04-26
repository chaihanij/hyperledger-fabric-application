#!/bn/bash

docker-compose -f docker-compose.yaml stop

rm -rf state/orderer*
rm -rf state/peer*

docker-compose -f docker-compose.yaml up
