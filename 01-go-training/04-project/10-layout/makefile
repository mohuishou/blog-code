proto_build:
	# 根据proto文件编译生成对应的go文件
	protoc -I ./api -I ./third_party/googleapis \
	--openapiv2_out ./api --openapiv2_opt logtostderr=true \
	--openapiv2_opt json_names_for_fields=false \
	--go_out ./api --go_opt=paths=source_relative \
	--go-gin_out ./api --go-gin_opt=paths=source_relative \
	api/product/app/v1/*.proto
	protoc-go-inject-tag -input=./api/product/app/v1/v1.pb.go

mockgen:
	cd ./internal && for file in `egrep -rnl "type.*?interface" ./domain | grep -v "_test" `; do \
		echo $$file ; \
		cd .. && mockgen -destination="./internal/pkg/mock/$$file" -source="./internal/$$file" && cd ./internal ; \
	done