# go-string-generator
... is a library that generates random strings from patterns.

## Usage Example
Check out `example/main.go` for a usage example and run it: 
```sh
go run example/main.go
```

## The Basics
Let me guide you through a couple of examples, so you know the basics. And then we'll get to [The Cool Stuff](#the-cool-stuff).

### Building The Application
But first things first, let's build the application. 
```sh
CGO_ENABLED=0 go build -o gsg app/gsg/main.go

# run it once, so the data directory is created
./gsg

# some of the examples below require data 
# from the `example/data` directory, 
# copy it to the data directory
cp -R example/data/* ~/.rsg-data/
```
Note: you can also make `~/.rsg-data` a symlink to some other location. However, it will be automatically created as a directory upon first run, therefore you would have to create the symlink first.

### Random Strings
Let's start with something simple like generating random strings. The token given to the application describes the type (`str`) and the length (e.g. `4`, `8`, `16`, can be anything you want).
```sh
./gsg [str:4] [str:4] [str:8] [str:8] [str:16] [str:16]
```
```
ehsx
zhiv
ghanbwve
vehpivog
bdmfiyzkosszeezj
wunwzbrqqvfouhbm
```

You might have noticed that everything is lowercase which is not always desirable. No problem, you can also generate uppercase and mixed-case strings:
```sh
./gsg [str:4] [str:4] [strU:8] [strU:8] [strR:16] [strR:16]
```
```
jotu
omgh
BJBVSKER
OHDLNIUH
knXLNVeXwlfhMlih
pZOaKJJtjRKirfdJ
```

### Random Alphanumeric Characters
If you need alphanumeric characters the `[mix:N]` token can help you out:
```sh
./gsg [mix:4] [mix:4] [mixU:8] [mixU:8] [mixR:16] [mixR:16]
```
```
cii2
mmo4
T6QQJWT2
G31IMIZA
bVL9p3wNSC9LaWRD
xcph06IdrtjLr7h1
```

### Random Zero-Padded Integers
For stuff like serial numbers you might need zero-padded integers:
```sh
./gsg [int:4] [int:4] [int:8] [int:8] [int:16] [int:16]
```
```
3449
0016
00006590
65282560
0938605188205384
5128184623339865
```

### Random Integers
These generate integer values within a given range (inclusive):
```sh
./gsg [1-3] [67-120] [1-100]
```
```
→ ./gsg [1-3] [67-120] [1-100]
3
103
100
```

### Integer Lists
For those too lazy to type long integer lists by hand, hehe. Might not seem very useful at the moment, but wait for [The Cool Stuff](#the-cool-stuff).
```sh
./gsg [1..3] [67..120] [1..100]
```
```
1,2,3
67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120
1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100
```

### String Lists
Need a random string from a list? Then this is for you:
```sh
./gsg [a,b,c] [hello,world] [foo,bar]
```
```
b
hello
bar
```

### Hashes
Hashes are also pretty common, let's generate some random ones:
```sh
./gsg '[#4]' '[#8]' '[#16]'
```
```
be81
7f3eab3c
c4b3da6832290fba
```

### UUIDS
Need a random UUID? No problem:
```sh
./gsg '[#UUID]' '[#UUID]' '[#UUID]'
```
```
cb6c5f1f-3000-42ca-080e1fd9bf3a
61202220-2e65-79a9-c63d43243241
bba8a686-0229-94cd-0fd9ac8aa88a
```

### Encoding
Sometimes you'll have to encode (generated) data, currently `base64` and `url` encoding are supported. Feel free to submit PRs for other encodings!
Note that the arguments are wrapped in double quotes so we can use spaces.
```sh
./gsg "[b64:hello world]" "[url:hello world]"
```
```
aGVsbG8gd29ybGQ=
hello+world
```

If you've made it this far, you should now know all you need for the cool stuff:

## The Cool Stuff
### Combining 
The examples you've seen so far used just one token at a time, but you can form more complex patterns by combining tokens. You could, for example, generate random email addresses:
```sh
./gsg "[info,no-reply,hello]@[google,amazon,ebay].[com,co.uk,au]"
./gsg "[info,no-reply,hello]@[google,amazon,ebay].[com,co.uk,au]"
./gsg "[info,no-reply,hello]@[google,amazon,ebay].[com,co.uk,au]"
./gsg "[info,no-reply,hello]@[google,amazon,ebay].[com,co.uk,au]"
```
```
info@ebay.com
hello@google.co.uk
no-reply@ebay.co.uk
hello@google.au
```

### Nesting
Now we get to the real fun stuff: nesting tokens. 
```sh
./gsg "[[1..4],[8..12],[24..28]]"
./gsg "[[1..4],[8..12],[24..28]]"
./gsg "[[1..4],[8..12],[24..28]]"
```
```
4
25
11
```

Here the parsing order comes into play: the library parses inside-out, that is it will evaluate the inner-most token first and then work its way to the outer-most token. In the example, that means `[1..4]`, `[8..12]` and `[24..28]` will be evaluated first, resulting in the list `[1,2,3,4,8,9,10,11,12,24,25,26,27,28]` which will then be evaluated as a [string list](#string-lists). 

There is no nesting depth limit which allows for limitless recursion if you may choose so. Feel free to find a hardware supplier for infinite memory then :P 

#### Random Integer With Weight
Some of you might have already been inspired to use this to generate integers with different weights. For those who haven't, let me give you a nudge:
```sh
./gsg "[[1-10],[1-10],[100-200]]"
./gsg "[[1-10],[1-10],[100-200]]"
./gsg "[[1-10],[1-10],[100-200]]"
./gsg "[[1-10],[1-10],[100-200]]"
./gsg "[[1-10],[1-10],[100-200]]"
./gsg "[[1-10],[1-10],[100-200]]"
```
```
4
8
157
9
6
121
```

In this example values from the range 100-200 should only show up 1/3 of the time when repeating the command indefinitely. However, statistics is a %$@&%, so you might not see that distribution in a few runs.

### Files
Those among you paying close attention might wonder how recursion can be achieved in the first place. Good catch, time to introduce you to the `[:file]` token. 

Remember that we copied files before building the application? Here's why: when using the `[:file]` token data is being read from the `$HOME/.rsg-data` directory. Well, to be more precise: it is read from memory because when a generator is created the library will load all files from that directory into memory and from then on watch the directory for changes and update the in-memory cache accordingly. This produces overhead for each generator created, so you should reuse them where possible, especially when using a lot of data.

When loading files into memory the library will remove all blank lines and commented lines. Comments can be created by prefixing a line with `#` and as many leading spaces as you like. However, it is not possible to comment partial lines.

With that out of the way, let's look at how `[:file]` tokens work.

Assume the following directory structure:
```
$HOME
└─ .rsg-data
   ├─ a
   ├─ b
   │  └─ ull
   │     └─ shit
   └─ c
```

Let's add some content to the files:
```sh
# .rsg-data/a
hello
world
[:b]
```

```sh
# .rsg-data/b/ull/shit
[:a]:[:a]
[:a]:[:c]
[:c]:[:a]
[:c]:[:c]
```

```sh
# .rsg-data/c
foo
bar
[:[a,b]]
```

Now let's run it:
```sh
./gsg [:a] [:a] [:a] [:a] [:a] [:[a,b]]
```
```
world
foo:world
bar:world:foo:world
foo:bar:foo:hello:bar:bar:bar:bar:world:world:hello:hello:bar:hello:foo:foo:hello:hello:bar:foo:hello:world:world
hello:world:hello
world:foo:foo:hello:world
```

Did you notice that it doesn't matter whether `[:file]` references a file or directory? 
If it's a directory (like `[:b]`), the token will select a random file from that directory or its subdirectories and return a random line from it. Otherwise, a random line from the given file (like `[:a]` and `[:c]`) will be returned.

### Custom Tokens
As you can see files can contain tokens referencing other files. Using this you can abstract complex patterns into a couple of files. Let's take random email addresses as an example. We need a couple of files first:

```sh
# .rsg-data/name
info
no-reply
hello
```

```sh
# .rsg-data/domain
google
amazon
ebay
```

```sh
# .rsg-data/tld
com
co.uk
au
```

```sh
# .rsg-data/email
[:name]@[:domain].[:tld]
```

Now we have a new token that we can use:
```sh
./gsg "Here's a random email address: [:email]" [:email] [:email] [:email]
```
```
Here's a random email address: no-reply@ebay.au
no-reply@google.au
info@ebay.co.uk
no-reply@amazon.co.uk
```

### Recursion
Didn't you see it? Hint: tokens can be nested, also `[:file]` tokens.

The library does not keep track of recursion depth, so creating the file `.rsg-data/recursion` with `[:recursion]` as content and then running `./gsg [:recursion]` would crash your machine real quick. As a safeguard, you should have at least one case in each data file that does not lead to recursion.