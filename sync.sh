#!/bin/bash

echo "sync starting"
cd `pwd`
git add .
read -p "comment:"
git commit -m"$REPLY"
git push
echo "sync finish"
