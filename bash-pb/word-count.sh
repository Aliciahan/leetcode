#!/bin/bash
cat world.txt | tr -cs "[a-z]" "\n" | sort  | uniq -c |sort -k1nr -k2 | awk '{print $2,$1}'