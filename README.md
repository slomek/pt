# pt - CLI for Pivotal Tracker

## Installation

```
go get -u github.com/slomek/pt
```

## Usage

In order to use Pivotal Tracker API one need to generate a personal token and store it in `PIVOTAL_TOKEN` environmental variable.

### Listing tickets assigned to particular user

```
pt mine -p <project-id> -u <username>
```
