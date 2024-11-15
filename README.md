# CC-Uniq - Command-Line Unique Line Processor

This project is a custom implementation of the classic Unix `uniq` command. It was developed as part of a coding challenge [here](https://codingchallenges.fyi/challenges/challenge-uniq). The utility is designed to filter and manipulate duplicate lines from input data, providing detailed control over the output displayed.

## Features

- Filter out duplicate lines or exclusively show them.
- Count the occurrences of each line.
- Output results to standard output or a specified file.
- Handle input from both files and standard input (piping).
- Efficiently process large text data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.15 or later is recommended).
- You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-uniq.git
cd cc-uniq
```

### Building

Compile the project using:

```bash
go build -o ccuniq
```

### Usage

To run the utility, you can either pass a file name as an argument or pipe text into it via standard input.

#### Print unique lines from a file:

```bash
./ccuniq -u filename.txt
```

#### Print only repeated lines:

```bash
./ccuniq -d filename.txt
```

#### Count occurrences of each line in a file:

```bash
./ccuniq -c filename.txt
```

#### Using piped input to filter unique lines:

```bash
cat filename.txt | ./ccuniq -u
```

#### Running the Tests

```bash
go test ./...
```
