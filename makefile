createdev:
	soda create -e development
	soda migrate up
createprod:
	soda create -e production
	soda migrate up
run:
	go run main.go
reset:
	soda drop -a
	soda migrate  up
teardown:
	soda drop -e development

