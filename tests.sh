#!/usr/bin/env bash


. ./tests/install.sh

. ./tests/auth.sh
bats ./tests/test_users.sh