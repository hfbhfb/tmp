#!/bin/bash

container_pid="<container_pid>"

# Check if nsenter is available
if ! command -v nsenter &> /dev/null; then
    echo "nsenter not found. Please install it."
    exit 1
fi

# Enter the PID namespace of the container
nsenter --target $container_pid --pid --mount <<EOF
# Change to the /proc directory
cd /proc

# List all processes in the namespace
for pid in *; do
    if [ -d "$pid" ] && [[ "$pid" =~ ^[0-9]+$ ]]; then
        echo "Process ID: $pid"
        echo -n "Command Line: "
        cat "$pid/cmdline" 2>/dev/null || echo "<Not available>"
        echo "-----------------------------"
    fi
done
EOF