# go-gender-stats
1. pulls golang contributors list, predicts gender of each first name, prints gender stats
2. pulls gophers slack members list, predicts gender, prints gender stats

# Disclaimer
The purpose of this project is to try to give some (any) statistics to track improvement of the M/F gender ratio
in the Go community over time.

Please note that heuristic and probabilistic gender classification by first name is horribly imperfect.
Please also note that gender is not binary, and it is ultimately up to each individual to determine how they identify.

## training data
I took the name => gender data from [OpenGenderTracking/globalnamedata](https://github.com/OpenGenderTracking/globalnamedata)

Quote from their announcement [blog article](http://bocoup.com/weblog/global-name-data/) on 06/03/2013:

> Today, we are releasing Global Name Data, a dataset of birth name-gender mapping which we believe to be the most comprehensive in the world.


## classifier
The classifier program in classifier/classifier.go is stolen from `github.com/hstove/gender` (thanks @hstove)

You will need to recreate the classifier before running `go-gender-stats`:

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

## output (03/32/2017, GitHub all languages, OpenGenderTracking Global Name Data))
See [https://eyskens.me/how-many-women-actually-go-c-rust-js..../](https://eyskens.me/how-many-women-actually-go-c-rust-js..../)

Used dataset SQL dump available at [Google Drive](https://drive.google.com/file/d/0B2DN0LlE0FSlaDBsWk1PeDljbDA/view?usp=sharing)

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
