#!/bin/bash
set -e
start=`date +%s`
GITCODEGEN=https://github.com/swagger-api/swagger-codegen.git
if [ -d "swagger-codegen" ]; then	
	cd swagger-codegen
	#echo "Updating ${GITCODEGEN} to swagger-codegen/..."
	#git pull .
else
	echo "Cloning ${GITCODEGEN} to swagger-codegen/... (This will take a long time!)"
	git clone ${GITCODEGEN}
	cd swagger-codegen
	echo "running `./run-in-docker.sh mvn package`"
	./run-in-docker.sh mvn package
fi

if [ ! -d "out/html" ]; then 
	mkdir -p out/html
fi

cp ../xegony.yml tmp.yml
echo "Creating documentation in docs/..."
./run-in-docker.sh generate -i tmp.yml -l html2 -o /gen/out/html -Dpackagename=xegony
echo "Creating SDKs in sdk/..."
./run-in-docker.sh generate -i tmp.yml -l go -o /gen/out/sdk/go/ -Dpackagename=xegony
echo "Creating Go Server baseline to go/tmp/..."
./run-in-docker.sh generate -i tmp.yml -l go-server -o /gen/out/go/tmp/ -Dpackagename=xegony
rm tmp.yml

mv out/html/index.html ../docs/
mv out/sdk/* ../sdk
mv out/go/tmp/* ../go/tmp


end=`date +%s`
runtime=$((end-start))
echo "Completed in ${runtime} seconds"