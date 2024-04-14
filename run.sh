#!/bin/bash
export $(grep -v '^#' TOOLS_ENV.env | xargs)
make run
