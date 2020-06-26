clang -c uid.mm
ar -r libuid.a uid.o
#  -L./ Foundation.tbd  ==> -framework Foundation
clang -shared -fPIC uid.o -o libuid.dylib -framework Foundation
rm uid.o
