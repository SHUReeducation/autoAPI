# CONTRIBUTE.md

This file is for bootstrap for newcomer. 

## git

	- Fork this repo under your own github account.
    - `git clone https://github.com/<YOURNAME>/AutoAPI"`
	
## Install dependencies

    You'll need `go`, `make`, etc. If you meet up with a problem, welcome to open an issue or pull request.
    ```shell
    go get -u github.com/valyala/quicktemplate
    go get -u github.com/valyala/quicktemplate/qtc
    ```
    After install `qtc`, You may need a new terminal session to fluash your $PATH, or you may have to add /Users/<YOURNAME>/go/bin:$PATH to your $PATH

## Create a new branch
    
    You need to `new` and `checkout` to a branch with a meaningful name.
    
    Do some coding.

## Build

    ```shell
    make build
    ```
    pass your local build 

    ```shell
    git add <new file>
    git commit -m "<a meaningful message>"
    git push
    ```

## Create a pull request to this repo

    `git push` and create pr to `SHUReeducation/master` branch.
    check if you can pass all tests and request a review.

