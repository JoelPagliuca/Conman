name1="/conman/test/value1"
value1="ssm-test"

.PHONY: test
test: test.txt
	@echo "+ $@"
	go test

test.txt:
	@echo "+ configuring tests"
	@aws ssm put-parameter --name "${name1}" --value "${value1}" --type String --overwrite > /dev/null
	@touch test.txt

clean:
	@echo "+ $@"
	aws ssm delete-parameter --name ${name1} || true
	rm -f test.txt
