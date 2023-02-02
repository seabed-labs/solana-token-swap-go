#!/bin/bash

VERSION="v$1";
PKG="./pkg/$VERSION"; 
mkdir $PKG;
cp $2 "$PKG/idl.json";
cd $PKG && anchor-go --src=./idl.json --codec=bin --dst=./ --remove-account-suffix=false