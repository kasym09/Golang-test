#!/bin/bash
ls -l | awk 'NR % 2 == 0 || NR == 1 {print}' | sed 1d