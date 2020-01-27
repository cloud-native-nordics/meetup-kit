#!/bin/bash

SCRIPT_DIR=$( dirname "${BASH_SOURCE[0]}" )
cd ${SCRIPT_DIR}/..

RESOURCES="Meetup MeetupGroup Company Speaker"
CLIENT_NAME=MeetOpsInternal
OUT_DIR=pkg/apis/meetops/client
API_DIR="github.com/cloud-native-nordics/meetup-kit/pkg/apis/meetops"
mkdir -p ${OUT_DIR}
for Resource in ${RESOURCES}; do
    resource=$(echo "${Resource}" | awk '{print tolower($0)}')
    sed -e "s|Resource|${Resource}|g;s|resource|${resource}|g;/build ignore/d;s|API_DIR|${API_DIR}|g;s|*Client|*${CLIENT_NAME}Client|g" \
        vendor/github.com/weaveworks/gitops-toolkit/pkg/client/client_resource_template.go > \
        ${OUT_DIR}/zz_generated.client_${resource}.go
done
