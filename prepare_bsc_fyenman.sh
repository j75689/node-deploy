basedir=$(
    cd $(dirname $0)
    pwd
)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

cd /mnt/efs/bc-fusion/node-deploy/

current_time=$(date +%s)
hardfork_time=$((${current_time} + ${BSC_WAIT_SEC_FOR_FYENMAN}))
bash ./setup_bsc_node.sh prepare_fyenman_hardfork $hardfork_time

cd /mnt/efs/bc-fusion/node-deploy/tmp/bsc-feynman/bsc
docker build . -t bsc:master
docker image save -o /mnt/efs/bc-fusion/images/bsc.tar bsc:master

echo "hardfork_time: ${hardfork_time}"

# Get the current timestamp
current_timestamp=$(date +%s)

# Calculate the waiting time in seconds
wait_seconds=$((hardfork_time - current_timestamp))

# Check if waiting is necessary
if [ $wait_seconds -le 0 ]; then
  echo "The hardfork timestamp has already passed. Please build again."
  exit 1
fi

echo "Still have $wait_seconds seconds until the hardfork timestamp $hardfork_time"