# git hub documentation

git-hub is an opinionated *git* and *github* automation tool.

## Table of contents

- [Introduction](#introduction)
- [Detailed description](#detailed-description)
  - [The main branch](#the-main-branch)
  - [Issue branches](#issue-branches)
  - [Release branches](#release-branches)

## Introduction

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.

## Detailed description

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.

### The main branch

The *worfkow* implemented by *git-hub* uses only one eternal branch called `master`.

### Issue branches

Issue branches are where the day-to-day development work happens, there are used to develop new features and bugfixes for the upcoming release.

The naming convention for these branches is: `issue/<issue_number>-<issue_title>`

### Release branches

Releae branches are created to prepare the software to be released. Usually meaning a list (more or less complex) of steps to be done.

The important thing is all that happens on a separate branch, so that day-to-day development can continue as usual on the master branch.

The naming convention for these branches is: `release/<version_number>`

And `<version-number>` follows the latest [semver](http://semver.org) specification.

Release branches always start from the `master` branch.