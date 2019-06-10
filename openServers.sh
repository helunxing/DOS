sudo apt-get install rabbitmq-server
sudo rabbitmq-plugins enable rabbitmq_management
wget localhost:15672/cli/rabbitmqadmin
python3 rabbitmqadmin declare exchange name=apiServers type=fanout
python3 rabbitmqadmin declare exchange name=dataServers type=fanout
sudo rabbitmqctl add_user test test
sudo rabbitmqctl set_permissions -p / test ".*" ".*" ".*"

# 以上为消息队列配置

for i in `seq 1 6`;do
    mkdir -p /tmp/$i/objects;done

export RABBITMQ_SERVER=amqp://test:test@localhost:5672

sudo ifconfig ens33:1 10.29.1.1/16
sudo ifconfig ens33:2 10.29.1.2/16
sudo ifconfig ens33:3 10.29.1.3/16
sudo ifconfig ens33:4 10.29.1.4/16
sudo ifconfig ens33:5 10.29.1.5/16
sudo ifconfig ens33:6 10.29.1.6/16
sudo ifconfig ens33:7 10.29.2.1/16
sudo ifconfig ens33:8 10.29.2.2/16

LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 go run dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.2.1:12345 go run apiServer/apiServer.go &
LISTEN_ADDRESS=10.29.2.2:12345 go run apiServer/apiServer.go &


curl -v 10.29.2.1:12345/objects/test2 -XPUT -d "this is object test2"

