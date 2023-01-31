#!/bin/bash

source ~/.nvm/nvm.sh

for i in "$@"
do
case $i in
    --force*)
    FORCE="true"
    shift # past argument=value
    ;;
    --build=*)
    BUILD="${i#*=}"
    shift # past argument=value
    ;;
    *)
      # unknown option: ignore
    ;;
esac
done

cd $(dirname $BASH_SOURCE)

PROG_NAME=$(awk -F'"' '/^const progname =/ {print $2}' main.go)


if [ -z "$(git status --porcelain)" ]; then
  # Working directory clean
  GIT_COMMIT=$(git rev-parse HEAD)
  LAST_BUILD=$(cat build/commit.dat)
  if [[ $GIT_COMMIT == $LAST_BUILD ]]; then
    echo "${PROG_NAME}: Nothing new."
    exit 1
  fi
elif [[ "$FORCE" == "true" ]]; then
  echo "Force build requested."
  GIT_COMMIT=$(date +%Y-%m-%d-%H-%M-%S)
else
  echo "${PROG_NAME}: Project must be clean before you can build"
  exit 1
fi

rm -rf build

echo "${PROG_NAME}: Building..."

cd webapp
rm -rf build
npm i
npm run build
cd ..

if [[ $BUILD == "" ]]; then
  FILENAME=${PROG_NAME}_${GIT_COMMIT}
else
  FILENAME=${PROG_NAME}_${GIT_COMMIT}_${BUILD}
fi


mkdir -p build/linux/$FILENAME
mkdir -p build/dist


# Linux
env GOOS=linux GOARCH=amd64 go build --trimpath -o build/linux/$FILENAME/${PROG_NAME} -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
cp assets/linux/start.sh assets/linux/stop.sh build/linux/$FILENAME
cd build/linux
tar cvfz ../dist/${PROG_NAME}-linux.tar.gz ./$FILENAME
cd ../.. 

if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  echo "${PROG_NAME}: Built $FILENAME"

  cp assets/index.html ./build/dist
  cp -r assets/media ./build/dist

  cd build/dist
  tar cvfz ../../$FILENAME.tar.gz ./*
  
  echo "${PROG_NAME}: Artifact $FILENAME.tar.gz"
else
  echo "${PROG_NAME}: Build FAILED"
fi
