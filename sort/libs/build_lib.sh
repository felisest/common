#! /bin/bash

rm -R release
mkdir release && cd release

cmake ..
cmake --build . --target all --verbose