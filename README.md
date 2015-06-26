# go-gender-stats
Pulls golang contributors list, predicts gender of each first name, prints gender stats

## classifier
All classification/prediction code lives in `github.com/hstove/gender/classifier`

The `classifier.serialized` is the classifier I built on my machine. I am unsure if it will work on yours, as I'm not sure what the bytes in the file represent.

If it does not work for you, you can find docs on how to re-create this file in `hstove`'s [README](https://github.com/hstove/gender/blob/master/README.md)

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
