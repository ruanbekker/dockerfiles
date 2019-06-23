#!/bin/sh
cat > /etc/motd << EOF

+ Container Info:

    Name: $(hostname)

    CPU: `cat /proc/cpuinfo | grep 'model name' | head -1 | cut -d':' -f2`
    Memory: `free -m | head -n 2 | tail -n 1 | awk {'print $2'}`M
    Disk: `df -h / | awk '{ a = $2 } END { print a }'`
    Distro: Alpine:$(cat /etc/os-release | tr '\n' ' '  | awk '{print $4}' | cut -d '=' -f2)

    CPU Load: `cat /proc/loadavg | awk '{print $1 ", " $2 ", " $3}'`
    Free Memory: `free -m | head -n 2 | tail -n 1 | awk {'print $4'}`M
    Free Disk: `df -h / | awk '{ a = $2 } END { print a }'`

    Container Address: `ifconfig eth0 | grep "inet addr" | awk -F: '{print $2}' | awk '{print $1}'`
    Public Address: $(curl -s ip.ruanbekker.com)
    SSH User: ${SSH_USER}

EOF
