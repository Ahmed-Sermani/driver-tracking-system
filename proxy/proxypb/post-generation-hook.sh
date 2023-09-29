#!/bin/bash
set -e

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
gomodifytype -file "$parent_path/proxy.pb.go" -all -w -from "[]byte" -to "Raw"
gomodifytags -file "$parent_path/proxy.pb.go" -field User -struct RefreshRequest -all -w -remove-options json=omitempty >/dev/null
gomodifytags -file "$parent_path/proxy.pb.go" -field User -struct SubscribeRequest -all -w -remove-options json=omitempty >/dev/null
gomodifytags -file "$parent_path/proxy.pb.go" -field User -struct PublishRequest -all -w -remove-options json=omitempty >/dev/null
gomodifytags -file "$parent_path/proxy.pb.go" -field User -struct RPCRequest -all -w -remove-options json=omitempty >/dev/null
gomodifytags -file "$parent_path/proxy.pb.go" -field User -struct SubRefreshRequest -all -w -remove-options json=omitempty >/dev/null
