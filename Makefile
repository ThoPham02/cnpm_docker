gen-service:
	goctl api go -api  account.api -dir .
gen-model:
	goctl model mysql ddl -src="./init-db.sql" -dir="./model"
run:
	go run account.go

build-image:
	docker build -t image-name .
run-container:
	docker run -d -it -p 8080:8080 --name container-name image-name
map-folder:
	docker run -ti -v /home/tholgbg/dev/school/cnpm/cnpm_docker:/app ubuntu
