# Cron Parser

## Overview

This command line application written in Go parses a cron string and expands each field to show the times at which it will run. It formats the output in a table-like structure with each field name followed by the corresponding times.

## Usage

To use this program, run it from the command line and provide a cron string as an argument. The cron string should follow the standard format with five fields: minute, hour, day of month, month, day of week, followed by a command.

## How It Works

CronParser Interface: Defines a way to parse and format cron strings.
cronParserImpl Struct: Implements the parsing logic for each cron field (minute, hour, etc.).
expandField Function: Expands a field from the cron string into a list of specific times or dates.
Main Function: Handles user input, parses the cron string, and prints the formatted output.
Running Tests
Unit tests are included to ensure that the parsing logic works correctly for different types of cron strings and edge cases.

To run the tests, use the following command:

### Example

```bash

$ go run CronParser.go '*/15 0 1,15 * 1-5 /usr/bin/find'

## The above command will produce the following output:

minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command        /usr/bin/find

```

### test cases
Unit tests are included to ensure that the parsing logic works correctly for different types of cron strings and edge cases.

To run the tests, use the following command:


```bash
$ go test
```
