run:
	go run main.go

get_deparment:
	curl 127.0.0.1:8484/deparment/2 | jq

get_townships:
	curl 127.0.0.1:8484/township | jq

get_township:
	curl 127.0.0.1:8484/township/2 | jq
