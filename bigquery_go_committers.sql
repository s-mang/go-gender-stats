# There is probably a much more efficient way of doing this.
# NOTE: committers without an email and a first name are dropped.

# Project name denoted as @PROJECT
# Dataset name denoted as @DATASET

# dump (overwrite) to table 'commits'
SELECT author.email AS author_email,
  author.name AS author_name, 
  committer.name AS committer_name,
  committer.email AS committer_email,
  FIRST(repo_name) WITHIN RECORD AS repo_name # same commit, same code, repo doesn't matter
FROM [bigquery-public-data:github_repos.commits]
WHERE author.date > TIMESTAMP('2009-11-10 00:00:00') # go birthday

# dump (overwrite) to table 'committers'
SELECT committer_email AS email, committer_name AS name, repo_name AS repo
FROM [@PROJECT:@DATASET.commits]
WHERE committer_name <> ''
AND committer_email <> ''

# dump (append) to table 'committers'
SELECT author_email AS email, author_name AS name, repo_name AS repo
FROM [@PROJECT:@DATASET.commits]
WHERE author_name <> ''
AND author_email <> ''
AND (
  (committer_name = '' OR committer_email = '') OR
  (committer_email <> author_email)
)

# dump (overwrite) to table 'committers'
# we want unique committers.
SELECT email, name, repo
FROM [@PROJECT:@DATASET.committers]
GROUP BY email, name, repo

# dump (overwrite) to table 'go_repos'
SELECT repo_name AS repo
FROM [bigquery-public-data:github_repos.languages]
WHERE language.name = 'Go'

# dump (overwrite) to table 'go_committers'
SELECT name AS first_name, email
FROM [@PROJECT:@DATASET.committers] committers
JOIN [@PROJECT:@DATASET.go_repos] repos
ON repos.repo = committers.repo
GROUP BY first_name, email