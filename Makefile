test:
# exclude TestAll_
	go test -v -race -parallel=16 -run='^Test_'

test/all:
	go test -v -race -parallel=16

