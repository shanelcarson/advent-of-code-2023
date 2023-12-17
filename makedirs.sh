#! /bin/bash

if [[ ($# -eq 1)  && ($1 =~ ^[0-9]+$)]]; then
  echo "Adding day$1 directory"
else
  echo "More than 1 argument was passed in."
  exit 1
fi

if [ -d ./day$1 ]; then
  echo "Directory day$1 already exists."
  exit 0
fi

echo "What language? (py/go)"

read ext

if [[ ! $ext =~ ^(py|go)$ ]]; then
  echo "Not a good extension bro."
  exit 0
fi


mkdir day$1; cd day$1;
if [[ $ext -eq "go" ]]; then
  touch go.mod
fi
touch problem1.$ext problem2.$ext sample.txt input1.txt input2.txt;
cd ../;

echo "Done.";
ls
