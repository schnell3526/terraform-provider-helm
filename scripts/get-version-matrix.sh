#!/bin/bash
# Copyright IBM Corp. 2017, 2026
# SPDX-License-Identifier: MPL-2.0


function get_latest_version() { 
    curl -s https://api.github.com/repos/hashicorp/terraform/git/refs/tags | \
        jq ".[] | .ref | split(\"/\") | .[2] | select(. | startswith(\"$1\"))" | \
            sort -V -r | head -1 
}

echo "matrix=[$(get_latest_version v1.0), $(get_latest_version v1.5), $(get_latest_version v1.9), $(get_latest_version v1.10), $(get_latest_version v1.11), $(get_latest_version v1.12), $(get_latest_version v1.13), $(get_latest_version v1.14), $(get_latest_version v1.15), $(get_latest_version v1.16), $(get_latest_version v1.17)]" >> "$GITHUB_OUTPUT"
