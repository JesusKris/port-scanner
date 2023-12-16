<!-- ctrl + shift + v to preview -->
# port-scanner

## Table of Contents
- [port-scanner](#port-scanner)
  - [Table of Contents](#table-of-contents)
  - [General Information](#general-information)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Team \& My Work](#team--my-work)
  - [Main Learnings](#main-learnings)
  - [Setup](#setup)

## General Information
This project was made as a school project in [kood/Jõhvi](https://kood.tech/) (12.1.2023)

The project required me to create a TCP/UDP port scanner which allows users to quickly scan TCP or UDP ports openess on a selected host IP. It also allows to bulk scan ranges of ports.

**DISCLAIMER**: This task was for educational purposes only. Do not use any of the methods described to cause harm. I do not take any accountability on any harm caused using these tools.

**NB! Different source control platform was used hence no commit history.**

## Features
- Scan TCP or UDP ports on a given IP, either singular or ranges.

## Technologies Used

[Golang](https://go.dev/)


## Team & My Work
This was a solo project.

Everything was done by me.

## Main Learnings
- Inner workings of a port scanner
- TCP/UDP protocols
- Importance of port scanners in pen-testing

## Setup
Clone the repository
```
git clone https://github.com/JesusKris/port-scanner.git
```
To see all available flags and arguments:
```
go run active.go --help
```

To perform a TCP port scan:
```
go run active.go -t 161.107.32.8 -p 8080
```

To perform a UDP port scan:
```
go run active.go -u 161.107.32.8 -p 8080
```

To perform a scan on ranges of ports:

TCP:
```
go run active.go -t 161.107.32.8 -p 8080-8100
``` 
UDP:
```
go run active.go -u 161.107.32.8 -p 8080-8100
``` 