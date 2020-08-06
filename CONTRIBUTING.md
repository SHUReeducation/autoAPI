# CONTRIBUTE.md

You're very welcomed to be a contributor to autoAPI! 🎉

This artical will show you the best practice of doing some contribution to this repo.

## Get the code

- Fork this repo under your own GitHub account.
- Clone it to your computer
  ```shell
  git clone https://github.com/<your GitHub username>/autoAPI
  ```
	
## Install dependencies

We presume you have installed [`go`](https://golang.google.cn/) and [`make`](https://www.gnu.org/software/make/) on your computer, if you don't, click the links or use your favourite package manager to install them.

And then you'll have to follow [these instructions](https://github.com/valyala/quicktemplate#quick-start) to install quicktemplate and qtc.

After install `qtc`, you have to make sure it is in your `$PATH`.

## Create a new branch
    
You need to `new` a branch with a meaningful name and `checkout` it.
    
## Do Some Coding

You can modify whatever you want!

## Build

```shell
make build
```
    
## Run an example

```shell
make run-example
```

## Commit and push

```shell
git add <new file>
git commit -m "<a meaningful message>"
git push
```

## Create a pull request to this repo

Create a pull request to `SHUReeducation/master`, and check if your code can pass all the tests.

We'll check the pull request now and then, you may also request a review manually if you like.
    
## Need help?
    
Feel free to [open an issue here](https://github.com/SHUReeducation/autoAPI/issues)!

