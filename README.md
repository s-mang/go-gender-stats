# go-gender-stats
Pulls golang contributors list, predicts gender of each first name, prints gender stats

## training data
I took the name => gender data from [OpenGenderTracking/globalnamedata](https://github.com/OpenGenderTracking/globalnamedata)

Quote from their announcement [blog article](http://bocoup.com/weblog/global-name-data/) on 06/03/2013:

> Today, we are releasing Global Name Data, a dataset of birth name-gender mapping which we believe to be the most comprehensive in the world.

It is definitely huge.


## classifier
The classifier program in classifier/classifier.go is stolen from `github.com/hstove/gender` (thanks @hstove)

The `classifier.serialized` is the classifier I built on my machine. I am unsure if it will work on yours, as I'm not sure what the bytes in the file represent other than "a classifier".

If it does not work for you, you can recreate the file like yourself:

```bash
$ cd $GOPATH/src/github.com/adams-sarah/go-gender-stats/classifier
$ go run classifier.go
```

This will generate a new classifier file at `../classifier.serialized`, so be sure you're in the classifier dir when you run.


## run

```bash
$ go get github.com/adams-sarah/go-gender-stats
$ go-gender-stats
```

## output (06/26/2015, OpenGenderTracking Global Name Data)

```

Go Contributors by Gender:

  - Female: 5.75%

  - Male: 94.25%

  - Unknown: 0.00%


```

## output (06/26/2015, US Census data)

```

Go Contributors by Gender:

  - Female: 4.40%

  - Male: 84.94%

  - Unknown: 10.66%

```
