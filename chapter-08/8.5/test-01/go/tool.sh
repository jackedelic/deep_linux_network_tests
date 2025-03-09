#!/bin/bash

#---------------------------  begin -------------------------
# Note: Adjust these settings according to your experimental environment

# 1. Client IP list: Choose 20, and they should not exist in the LAN
IPS=(
    "192.168.1.200"
    "192.168.1.201"
)

# 2. Subnet mask for client IPs
NETMASK="255.255.248.0"

# 3. Server IP and port
SERVERIP="192.168.1.100"
SERVERPORT="8090"
#---------------------------   end  -------------------------

TYPE=$1

exec_ping(){
    for i in "${!IPS[@]}"; do
        ping -c 1 -W 1 ${IPS[$i]} > /dev/null
        if [[ $? != 0 ]]; then 
            echo ${IPS[$i]}" false"
        else
            echo ${IPS[$i]}" true"
        fi
    done 
}

exec_ifup(){
    for i in "${!IPS[@]}"; do
        echo ifconfig eth0:$i ${IPS[$i]} netmask $NETMASK up
        ifconfig eth0:$i ${IPS[$i]} netmask $NETMASK up
    done
}

exec_ifdown(){
    for i in "${!IPS[@]}"; do
        echo ifconfig eth0:$i down
        ifconfig eth0:$i down
    done
}

exec_runcli(){
    CLIENT=$2
    for i in "${!IPS[@]}"; do
        echo $CLIENT ${IPS[$i]} $SERVERIP $SERVERPORT &
        $CLIENT ${IPS[$i]} $SERVERIP $SERVERPORT &
    done
}

exec_stopcli(){
    CLIENT=$2
    ps -ef | grep $CLIENT | awk '{print $2}' | xargs kill -9 > /dev/null
}

exec_runsrv(){
    SERVER=$2
    echo $SERVER 0.0.0.0 $SERVERPORT
    $SERVER 0.0.0.0 $SERVERPORT
}

case $TYPE in
    "ping")  exec_ping;;
    "ifup")  exec_ifup;;
    "ifdown")  exec_ifdown;;
    "runcli")  exec_runcli $2;;
    "stopcli")  exec_stopcli $2;;
    "runsrv")  exec_runsrv $2;;
    *)  echo "get unknown type $TYPE"; exit ;;
esac