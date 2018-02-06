#!/bin/bash
set -e
start=`date +%s`
go test ./...

counter=0
for path in $(find .); do
    [ -d "${path}" ] || continue # if not a directory, skip
    if [[ "${path}" == *".git"* ]]; then
        continue
    fi
    #echo "Processing ${path}"
    files=$(ls ${path}/*.go 2> /dev/null | wc -l)
    #echo $files
    if [ $files != 0 ]; then
    	if [[ $path != */.git* ]]; then
    		#echo $path
    		go test -coverprofile=pkg${counter}.cover.out -coverpkg=./... $path
    		counter=$((counter+1))
    	fi
    fi
done

echo "mode: set" > coverage.out && cat *.cover.out | grep -v mode: | sort -r | \
awk '{if($1 != last) {print $0;last=$1}}' >> coverage.out

go tool cover -html=coverage.out
go tool cover -func=coverage.out > cover.txt
rm *.out

end=`date +%s`
runtime=$((end-start))
echo "Completed in ${runtime} seconds"