package tests

import (
	"strings"
	"testing"

	"github.com/mfridman/tparse/parse"
)

func TestPrescan(t *testing.T) {

	t.Parallel()

	t.Run("pass", func(t *testing.T) {

		t.Parallel()

		intput := `git.checkout
		0.55s$ git clone --depth=50 --branch=master https://github.com/mfridman/tparse.git mfridman/tparse
		Cloning into 'mfridman/tparse'...
		remote: Enumerating objects: 485, done.
		remote: Counting objects: 100% (485/485), done.
		remote: Compressing objects: 100% (310/310), done.
		remote: Total 485 (delta 209), reused 442 (delta 168), pack-reused 0
		Receiving objects: 100% (485/485), 1.17 MiB | 20.58 MiB/s, done.
		Resolving deltas: 100% (209/209), done.
		$ cd mfridman/tparse
		$ git checkout -qf d7a97f658463e3abd90357d9988575e8138d9a86
		4.38s$ GIMME_OUTPUT="$(gimme 1.10 | tee -a ${TRAVIS_HOME}/.bashrc)" && eval "$GIMME_OUTPUT"
		go version go1.10 linux/amd64
		$ export GOPATH=${TRAVIS_HOME}/gopath
		$ export PATH=${TRAVIS_HOME}/gopath/bin:$PATH
		$ mkdir -p ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ tar -Pczf ${TRAVIS_TMPDIR}/src_archive.tar.gz -C ${TRAVIS_BUILD_DIR} . && tar -Pxzf ${TRAVIS_TMPDIR}/src_archive.tar.gz -C ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ export TRAVIS_BUILD_DIR=${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ cd ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		0.00s
		$ gimme version
		v1.5.3
		$ go version
		go version go1.10 linux/amd64
		go.env
		$ go env
		GOARCH="amd64"
		GOBIN=""
		GOCACHE="/home/travis/.cache/go-build"
		GOEXE=""
		GOHOSTARCH="amd64"
		GOHOSTOS="linux"
		GOOS="linux"
		GOPATH="/home/travis/gopath"
		GORACE=""
		GOROOT="/home/travis/.gimme/versions/go1.10.linux.amd64"
		GOTMPDIR=""
		GOTOOLDIR="/home/travis/.gimme/versions/go1.10.linux.amd64/pkg/tool/linux_amd64"
		GCCGO="gccgo"
		CC="gcc"
		CXX="g++"
		CGO_ENABLED="1"
		CGO_CFLAGS="-g -O2"
		CGO_CPPFLAGS=""
		CGO_CXXFLAGS="-g -O2"
		CGO_FFLAGS="-g -O2"
		CGO_LDFLAGS="-g -O2"
		PKG_CONFIG="pkg-config"
		GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build714003296=/tmp/go-build -gno-record-gcc-switches"
		before_install
		{"Time":"2018-10-24T13:13:35.445816-04:00","Action":"run","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447209-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== RUN   TestStack\n"}
		{"Time":"2018-10-24T13:13:35.44729-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== PAUSE TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447302-04:00","Action":"pause","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447311-04:00","Action":"cont","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447317-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== CONT  TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447348-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"--- PASS: TestStack (0.00s)\n"}
		{"Time":"2018-10-24T13:13:35.447394-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Elapsed":0}
		{"Time":"2018-10-24T13:13:35.447481-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"PASS\n"}
		{"Time":"2018-10-24T13:13:35.447759-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"ok  \tgithub.com/mfridman/tparse/tests\t0.014s\n"}
		{"Time":"2018-10-24T13:13:35.454438-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Elapsed":0.021}`

		if _, err := parse.Start(strings.NewReader(intput)); err != nil {
			t.Fatalf("failed to scan over non-event lines, 50 max: %v", err)
		}
	})

	t.Run("fail01", func(t *testing.T) {

		t.Parallel()

		intput := `git.checkout
		0.55s$ git clone --depth=50 --branch=master https://github.com/mfridman/tparse.git mfridman/tparse
		Cloning into 'mfridman/tparse'...
		remote: Enumerating objects: 485, done.
		remote: Counting objects: 100% (485/485), done.
		remote: Compressing objects: 100% (310/310), done.
		remote: Total 485 (delta 209), reused 442 (delta 168), pack-reused 0
		Receiving objects: 100% (485/485), 1.17 MiB | 20.58 MiB/s, done.
		Resolving deltas: 100% (209/209), done.
		$ cd mfridman/tparse
		$ git checkout -qf d7a97f658463e3abd90357d9988575e8138d9a86
		4.38s$ GIMME_OUTPUT="$(gimme 1.10 | tee -a ${TRAVIS_HOME}/.bashrc)" && eval "$GIMME_OUTPUT"
		go version go1.10 linux/amd64
		$ export GOPATH=${TRAVIS_HOME}/gopath
		$ export PATH=${TRAVIS_HOME}/gopath/bin:$PATH
		$ mkdir -p ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ tar -Pczf ${TRAVIS_TMPDIR}/src_archive.tar.gz -C ${TRAVIS_BUILD_DIR} . && tar -Pxzf ${TRAVIS_TMPDIR}/src_archive.tar.gz -C ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ export TRAVIS_BUILD_DIR=${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		$ cd ${TRAVIS_HOME}/gopath/src/github.com/mfridman/tparse
		0.00s
		$ gimme version
		v1.5.3
		$ go version
		go version go1.10 linux/amd64
		go.env
		$ go env
		GOARCH="amd64"
		GOBIN=""
		GOCACHE="/home/travis/.cache/go-build"
		GOEXE=""
		GOHOSTARCH="amd64"
		GOHOSTOS="linux"
		GOOS="linux"
		GOPATH="/home/travis/gopath"
		GORACE=""
		GOROOT="/home/travis/.gimme/versions/go1.10.linux.amd64"
		GOTMPDIR=""
		GOTOOLDIR="/home/travis/.gimme/versions/go1.10.linux.amd64/pkg/tool/linux_amd64"
		GCCGO="gccgo"
		CC="gcc"
		CXX="g++"
		CGO_ENABLED="1"
		CGO_CFLAGS="-g -O2"
		CGO_CPPFLAGS=""
		CGO_CXXFLAGS="-g -O2"
		CGO_FFLAGS="-g -O2"
		CGO_LDFLAGS="-g -O2"
		PKG_CONFIG="pkg-config"
		GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build714003296=/tmp/go-build -gno-record-gcc-switches"
		before_install
		pirate ipsum!
		{"Time":"2018-10-24T13:13:35.445816-04:00","Action":"run","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447209-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== RUN   TestStack\n"}
		{"Time":"2018-10-24T13:13:35.44729-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== PAUSE TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447302-04:00","Action":"pause","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447311-04:00","Action":"cont","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447317-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== CONT  TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447348-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"--- PASS: TestStack (0.00s)\n"}
		{"Time":"2018-10-24T13:13:35.447394-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Elapsed":0}
		{"Time":"2018-10-24T13:13:35.447481-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"PASS\n"}
		{"Time":"2018-10-24T13:13:35.447759-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"ok  \tgithub.com/mfridman/tparse/tests\t0.014s\n"}
		{"Time":"2018-10-24T13:13:35.454438-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Elapsed":0.021}`

		if _, err := parse.Start(strings.NewReader(intput)); err == nil {
			t.Fatalf("want failure when stream contains >50 lines of non-parsable events: %v", err)
		}
	})

	t.Run("fail02", func(t *testing.T) {

		t.Parallel()

		// logic: unparsable event(s), good event(s), at least one event = fail.
		// Once we get a good event, we expect only good events to follow until EOF.
		intput := `git.checkout
		0.55s$ git clone --depth=50 --branch=master https://github.com/mfridman/tparse.git mfridman/tparse
		Cloning into 'mfridman/tparse'...
		remote: Enumerating objects: 485, done.
		remote: Counting objects: 100% (485/485), done.
		{"Time":"2018-10-24T13:13:35.447209-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== RUN   TestStack\n"}
		remote: Compressing objects: 100% (310/310), done.
		{"Time":"2018-10-24T13:13:35.447317-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack",
		remote: Total 485 (delta 209), reused 442 (delta 168), pack-reused 0
		Receiving objects: 100% (485/485), 1.17 MiB | 20.58 MiB/s, done.
		Resolving deltas: 100% (209/209), done.
		$ cd mfridman/tparse
		$ git checkout -qf d7a97f658463e3abd90357d9988575e8138d9a86
		{"Time":"2018-10-24T13:13:35.445816-04:00","Action":"run","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447209-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== RUN   TestStack\n"}
		{"Time":"2018-10-24T13:13:35.44729-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== PAUSE TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447302-04:00","Action":"pause","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447311-04:00","Action":"cont","Package":"github.com/mfridman/tparse/tests","Test":"TestStack"}
		{"Time":"2018-10-24T13:13:35.447317-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"=== CONT  TestStack\n"}
		{"Time":"2018-10-24T13:13:35.447348-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Output":"--- PASS: TestStack (0.00s)\n"}
		{"Time":"2018-10-24T13:13:35.447394-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Test":"TestStack","Elapsed":0}
		{"Time":"2018-10-24T13:13:35.447481-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"PASS\n"}
		{"Time":"2018-10-24T13:13:35.447759-04:00","Action":"output","Package":"github.com/mfridman/tparse/tests","Output":"ok  \tgithub.com/mfridman/tparse/tests\t0.014s\n"}
		{"Time":"2018-10-24T13:13:35.454438-04:00","Action":"pass","Package":"github.com/mfridman/tparse/tests","Elapsed":0.021}`

		if _, err := parse.Start(strings.NewReader(intput)); err == nil {
			t.Fatalf("want failure when stream contains a bad event(s) -> good event(s) -> bad event: %v", err)
		}
	})

}
