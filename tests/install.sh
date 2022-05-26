#!/usr/bin/env bash



if ! command -v http &> /dev/null
then
    echo "http could not be found"
    echo "please use below script install HTTPie"
    echo "for MacOSx"
    echo "brew install httpie"
    echo "for Linux"
    echo "sudo apt-get install httpie"
    echo "for python"
    echo "pip install httpie"
    exit
fi

