#!/bin/sh

# get parent directory
source="${BASH_SOURCE[0]}"
while [[ -h "${source}" ]]; do 
    source="$(readlink "${source}")"
done
dir="$(cd -P "$(dirname "${source}")/.." && pwd )"

while getopts "sx" opt; do
    case ${opt} in
        s)
            printf "Starting eve...\n"
            nohup ${dir}/bin/eve -config-file ${dir)/etc/eve.toml >> ${dir}/log/eve.log 2>&1 & echo $! > ${dir}/.eve.pid
            ;;
        x)
            printf "Stopping eve...\n"
            kill -9 $(cat ${dir}/.eve.pid)
            rm -f ${dir}/.eve.pid
            ;;
        \?)
            printf "Invalid option: -%s\n" "$OPTARG" >&2
            exit 1
            ;;
    esac
done

