all:
	go build
t:
	go build
	echo "build done"
	./server-manage t
m:
	go build
	./server-manage update -m /home/users/yanming02/workspace/server-manage/host_mem.info
s:
	go build
	./server-manage update -s /home/users/yanming02/workspace/server-manage/host_redis.info
r:
	go build
	./server-manage redis -m 1G
