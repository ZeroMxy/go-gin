#!/bin/bash

# The name of your application
APP_NAME="main"
# Path of the application program
APP_PATH="/data/app"
# Log file output path
LOG="${APP_PATH}/server.log"
# Path to the application pid file
PID_FILE="${APP_PATH}/server.pid"

# Start program
start () {
    if [ -f $PID_FILE ]; then 
        echo "The service is running with PID $(cat ${PID_FILE})"
        exit 1
    fi

    nohup "$APP_PATH/$APP_NAME" > "$LOG_FILE" 2>&1 &
    echo $! > "$PID_FILE"

    echo "The service started successfully with PID $(cat ${PID_FILE})"
}

# Stop program
stop () {
    if [ ! -f $PID_FILE ]; then 
        echo "The service is not running"
        exit 1
    fi

    kill -TERM $(cat "${PID_FILE}")
    rm "$PID_FILE"

    echo "The service has been stopped"
}

# Restart program
restart () {
    stop
    start
}

# Usage guide
usage () {
    echo "Usage: $0 {start|stop|restart}" >&2
    exit 1
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    *)
        usage
        ;;
esac
