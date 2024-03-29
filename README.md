# go-gender-stats
1. pulls golang contributors list, predicts gender of each first name, prints gender stats
2. pulls gophers slack members list, predicts gender, prints gender stats
3. pulls Go committers from bigquery github archive, predicts gender, prints gender stats

# Disclaimer
The purpose of this project is to try to give some (any) statistics to track improvement of the M/F gender ratio
in the Go community over time.

Please note that heuristic and probabilistic gender classification by first name is horribly imperfect.
Please also note that gender is not binary, and it is ultimately up to each individual to determine how they identify.

# About
## training data
I took the name => gender data from [OpenGenderTracking/globalnamedata](https://github.com/OpenGenderTracking/globalnamedata)

Quote from their announcement [blog article](http://bocoup.com/weblog/global-name-data/) on 06/03/2013:

> Today, we are releasing Global Name Data, a dataset of birth name-gender mapping which we believe to be the most comprehensive in the world.

# Run
## install

```bash
$ go get github.com/adams-sarah/go-gender-stats
```

## re-create classifier
The classifier program in classifier/classifier.go is stolen from `github.com/hstove/gender` (thanks @hstove)

You will need to recreate the classifier before running `go-gender-stats`:

```bash
$ cd $GOPATH/src/github.com/adams-sarah/go-gender-stats/classifier
$ go run classifier.go
```

This will generate a new classifier file at `../classifier.serialized`, so be sure you're in the classifier dir when you run.


## run

```bash
$ go-gender-stats
```

# Output over time
## output (03/15/2017, OpenGenderTracking Global Name Data)

```

Go Contributors by Gender:

  - Female: 5.57%

  - Male: 94.43%

  # Total: 1149

-------------

Slack Gophers by Gender:

  - Female: 6.92%

  - Male: 93.08%

  # Total: 14562

-------------

Github Go Committers by Gender:

  - Female: 5.99%

  - Male: 94.01%

  # Total: 74582
  
```

## output (03/09/2017, OpenGenderTracking Global Name Data)

```

Go Contributors by Gender:

  - Female: 5.57%

  - Male: 94.43%

-------------

Slack Gophers by Gender:

  - Female: 6.95%

  - Male: 93.05%

```

## output (07/24/2016, OpenGenderTracking Global Name Data)

```

Go Contributors by Gender:

  - Female: 5.18%

  - Male: 94.82%

-------------

Slack Gophers by Gender:

  - Female: 6.39%

  - Male: 93.61%


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
