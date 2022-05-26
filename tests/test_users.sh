#!/usr/bin/env bats

load ./users


@test "list User" {
    # result=$(getUsers | jq '.users | length')
    result=$(getUsers '?page.pageSize=2')
    count=$(echo $result | jq '.users | length')
    totalItems=$(echo $result | jq '.page.totalItems')
    nextId=$(echo $result | jq '.page.nextId')
    prevId=$(echo $result | jq '.page.prevId')
    [ "$count" -gt 0 ]
    [ "$totalItems" -gt 0 ]
    [ ! "$nextId" = "" ]
    [ ! "$prevId" = "" ]
}

@test "list User with page" {
    # result=$(getUsers '?page.pageSize=2' | jq '.users | length')
    result=$(getUsers '?page.pageSize=2')
    count=$(echo $result | jq '.users | length')
    totalItems=$(echo $result | jq '.page.totalItems')
    nextId=$(echo $result | jq '.page.nextId')
    prevId=$(echo $result | jq '.page.prevId')
    [ "$count" = 2 ]
    [ "$totalItems" -gt 2 ]
}

@test "list User query next page" {
    result=$(getUsers '?page.pageSize=2')
    count=$(echo $result | jq '.users | length')
    totalItems=$(echo $result | jq '.page.totalItems')
    firstId=$(echo $result | jq '.page.nextId')
    echo $result > output.log
    # prevId=$(getUsers | jq '.page.prevId')
    result1=$(getUsers '?page.pageSize=2&page.offset=2')
    totalItems1=$(echo $result1 | jq '.page.totalItems')
    nextId=$(echo $result1 | jq '.page.nextId')
    echo $result1 >> output.log
    echo $firstId >> output.log
    echo $nextId >> output.log
    [ "$count" = 2 ]
    [ ! "$firstId" = "$nextId" ]
}

