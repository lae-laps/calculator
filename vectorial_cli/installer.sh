#!/usr/bin/bash
if ((`id -u` != 0))
then
    echo "Please run installer as root..."
    exit
fi

go build .
strip vectorial
mv vectorial /usr/local/bin/vectorial
echo "vectorial successfully installed..."
