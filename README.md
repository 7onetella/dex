# dex
CLI version of dex, the chemical compound to treat engineering workplace ADHD 

### What

dex is a command line tool which is made up of collection of scripts. dex provides access to other command line tools that are written in golang. dex starts bunch of docker containers.

everything dex does you can do on command line as bash script.

think of dex as shortcuts, time saving scripts. I do not like typing long command line arguments or try to remember what they are.

I use [Cobra](htttps://github.com/spf13/cobra) to structure commands and subcommands.

### Why

there is no particular reason to use dex. however, I have seen bash scripts becoming really hard to maintain. you can write python script to do the same thing as dex. My experience with python is that as you add more functionality you required the users to install more depedent libraries. dex is written in go which means it's a single static binary file you can share with anyone.

### How to use dex

typing command name without any arguments will show help for that command.

```sh
$ dex [command]
```

### Current list of functionality
* dilbert     - Shows today's dilbert comic
* docker      - Docker automation
* generate    - Generates autocomplete scripts
* gitlog      - Git log with flattned view and formatting
* help        - Help about any command
* imgcat      - Renders image from local path on [iTerm](http://www.iterm.com) terminal
* mdcat       - Markdown automation
* mysql       - Runs mysql client
* psql        - Runs psql client
* redis       - Runs redis server locally
* redis-cli   - Runs redis client
* resolve     - Resolves container ip by container name
* siege       - Runs siege* 

