run:
	go run main.go

get_deparment:
	curl 127.0.0.1:8484/deparment/2 | jq

get_townships:
	curl 127.0.0.1:8484/township | jq

get_township:
	curl 127.0.0.1:8484/township/2 | jq

get_township_dep:
	curl 127.0.0.1:8484/township/deparment/11 | jq

get_branches:
	curl 127.0.0.1:8484/diaco | jq
	
get_main_complaint:
	curl 127.0.0.1:8484/complaint/main | jq
