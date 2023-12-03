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

mkdir day$1; cd day$1; 
touch problem1.py problem2.py input1.txt input2.txt;
cd ../;

echo "Done.";
ls