#!/bin/bash

# ls -Ra | grep -v '/$' | wc -l 

takprosto=$(find . -type f -o -type d | wc -l)
echo "$takprosto"