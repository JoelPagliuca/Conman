#!/usr/bin/env bash

name1="/conman/test/value1"
value1="ssm-test"

function cleanup {
	echo '[*] cleaning up'
	aws ssm delete-parameter --name "${name1}"
	rm -f ./examples/examples
}
trap cleanup EXIT

echo "+ run-tests"

aws ssm put-parameter --name "${name1}" --value "${value1}" --type String --overwrite > /dev/null

go test
