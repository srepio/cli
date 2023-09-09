#!/bin/bash

updateDockerFrom() {
    file=$1
    latest=$2

    echo "Updating $file to $latest"

    sed -Ei "s~^FROM ghcr.io/henrywhitaker3/srep/base:[0-9]+\.[0-9]+\.[0-9]+$~FROM ghcr.io/henrywhitaker3/srep/base:$latest~g" $file
}

updateScenarioVersion() {
    file=$1

    current=$(cat $file | jq -r '.version')
    updated=$(echo ${current} | awk -F. -v OFS=. '{$NF += 1 ; print}')

    echo "Updating version from $current to $updated"
    jq --indent 4 --arg version "$updated" '.version = $version' $file > /tmp/srep-metadata.json
    mv /tmp/srep-metadata.json "$file"
}

latest=$(cat containers/base/metadata.json | jq -r .version)

for scenario in containers/*; do
    if [ $scenario != "containers/base" ] && [ $scenario != "containers/k3s" ]; then
        updateDockerFrom "$scenario/Dockerfile" $latest
        updateScenarioVersion "$scenario/metadata.json"
    fi
done

updateDockerFrom ".templates/scenario/Dockerfile" $latest
