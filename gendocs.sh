#!/bin/bash

#删除之前的
rm "routes/docs.go"
	swagger -apiPackage="gin/controllers" -mainApiFile="gin/routes/router.go" -output="routes" -basePath="http://localhost:8010/docs"
