## Install, Build, Run Convention.
* Install go file from

* One single directory for all go codes, aka *workspace*. This is a `$GOPATH`.
    * This single directory defaults to `/Users/LOGED_IN_USER/go` eg `/Users/S/go`
    * `~/go/src` folder is Go src folder for all Go source code.
    * `~/go/bin` folder is compiled Go binary folder from `~/go/src`
    
* To set `$GOPATH`, write below in `.bash_profile`
    ```bash
    export GOPATH=$(go env GOPATH)
    # This is to run built binary files easily with single command.
    # eg) instead of '$GOPATH/bin/hello_world' to run binary, we can write just 'hello_world'
    PATH=${PATH}:${GOPATH}/bin 
    ```

* Compiler uses $GOPATH as a root for all imports so it is *required* to put everything in '$GOPATH/src'

* `$GOPATH` must not be the same path as your Go installation path which is `user/local/go`.

* Import path looks through either `~/go/src` or remote repo

    * Standard library with short path `"fmt"`, `"net/http"`
    
    * So if we have our own `~/go/src/fmt` it will collide with stdlib `"fmt"`, so convention is to use source repo base path.
        * eg) `github.com/imskojs/fmt` should be the folder name, and import using `import "githumb.com/imskojs/fmt"`.
        * Note we cane import repo within `github.com/imskojs` for all my packages. 
        * `git clone https://githumb.com/imskojs/fmt.git` within, `~/src/github.com/imskojs`.
        
    * If package is not available at local directory, it will fetch from, in this case, github.com.
    This is why we name a folder with `github.com/whatever` so it will be easier to get and post to github if need to be in the future.
    
    * To manually make remote stuff a local stuff `go get githumb.com/whatever` can be used. 
    Note though this is the exact command used when import statements get package from remote source
    Goland IDE will give a note if imported remote package is not currently available in the local repo
    
    * `go get` will download and install the remote repo (install as in generate binary file)
        
* `go install github.com/imskojs/fmt` builds binary file in `~/go/bin/` from entry of `package main`
    * `go` command uses $GOPATH as the base path so `go intall` can be run from any folder
   
* `go build` does not create *binary* file, the build will be cached.
    * it is to check whether a package gets compiled.
    
    * As a library gets imported and not compiled `go install` is not used with a library.
    Library is imported with `import lib/path/libName`, and importing executable gets built using `install`
    
* Executable commands must always use `package main`

## Testing
* Test have a name suffixed with `_test.go`, and a functions prefixed with `Test` with signature `func (t *testing.T)`

* Test have the same package name as the package function we are testing.

* `go test github.com/imskojs/fmt`, to run test for `fmt` package.
