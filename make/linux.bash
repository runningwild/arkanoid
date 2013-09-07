# Run from magnus root directory.
USER=runningwild
PACKAGE=arkanoid

cd $GOPATH/src/github.com/$USER/$PACKAGE

if CGO_LDFLAGS=-L../glop/gos/linux/lib go build . ; then
		rm -rf bin
		mkdir bin
		mkdir bin/data

		cp $PACKAGE bin/base_binary
		cp data/* bin/data
		cp ../glop/gos/linux/lib/libglop.so bin/libglop.so
		echo "LD_LIBRARY_PATH=$LD_LIBRARY_PATH:. ./base_binary" > bin/$PACKAGE
		chmod 776 bin/$PACKAGE

		cd bin
		./$PACKAGE
fi
