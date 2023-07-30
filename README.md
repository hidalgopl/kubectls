# kubectls - one command to rule your all kubectl contexts

Are you operating in multi-cluster environment? Are you tired of constantly switching your kubectl contexts to execute the same command in different clusters?
Good news! `kubectls` is here to help you!

`kubectl` takes one kubernetes context as an input and executes the command in this context.
`kubectls` takes **a list of kubernetes context separated by comma** and simply executes the same command in each context!

... and it's less than 70 lines of code!


# Demo

[![asciicast](https://asciinema.org/a/sM4ogGmfBnk0wBC31K5RK3n88.svg)](https://asciinema.org/a/sM4ogGmfBnk0wBC31K5RK3n88)

# Commands not supported
- exec
- port-forward
- attach
- edit

# Install
Build it from source:

    $ git clone 