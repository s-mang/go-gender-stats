# go-gender-stats
Pulls golang contributors list, predicts gender of each first name, prints gender stats

## classifier
All classification/prediction code lives in `github.com/hstove/gender/classifier`

The `classifier.serialized` file is the classifier I built on my machine from the data in `github.com/hstove/gender/classifier/names/`. I am unsure if this file will be helpful on your machine, as I don't know what the bytes in the file represent, other than "a classifier".

So, if you get stuck, the docs on how to re-create this file are in `hstove`'s [README](https://github.com/hstove/gender/blob/master/README.md). 

## run

```bash
$ go get github.com/adams-sarah/go-gender-stats

$ go-gender-stats
```

## output (06/26/2015)

```

Go Contributors by Gender:

  - Female: 4.40%

  - Male: 84.94%

  - Unknown: 10.66%

```
