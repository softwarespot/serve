#!/bin/bash

go build
mv ./serve ~/bin/serve
echo "Built and deployed to \"~/bin\""
