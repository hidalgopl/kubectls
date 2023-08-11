# kubectls - one command to rule your all kubectl contexts

Are you operating in multi-cluster environment? Are you tired of constantly switching your kubectl contexts to execute the same command in different clusters?
Good news! `kubectls` is here to help you!

`kubectl` takes one kubernetes context as an input and executes the command in this context.
`kubectls` takes **a list of kubernetes context separated by comma** and simply executes the same command in each context!

... and it's less than 70 lines of code!


# Demo


# Commands not supported
- exec
- port-forward
- attach
- edit

# Install
Download a binary for your OS/ARCH from releases page: https://github.com/hidalgopl/kubectls/releases
Add it to your `$PATH`.

Try it with `kubectls --context=my-context-1,my-context-2 get pods -A` !
