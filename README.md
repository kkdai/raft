A Raft algorithm concept prove application 
==================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/raft/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/raft?status.svg)](https://godoc.org/github.com/kkdai/raft) 

It is the [Raft](https://raftconsensus.github.io/) Consensus algorithm Raft implement in Golang.



## What is Raft?

Raft is a consensus algorithm designed as an alternative to Paxos. It was meant to be more understandable than Paxos by means of separation of logic, but it is also formally proven safe and offers some new features. (from [wiki](https://en.wikipedia.org/wiki/Raft_(computer_science)))


## What is this application?

This applicaition is a concept prove application to understand about Raft algorithm. 

It use a `go channel` to simulator to network, so that it could work on single machine and understand how it works.

## How it works?

This application provide you a overview about "Raft Algorithm", you can run a `go test` to check all testing message about how raft process related message.

## Raft Algorithm Explanation

This par stil to be confirm after all functions been completed

### Major Role

TBD

### Resposibility for each role

TBD

### How to handle exceptions

TBD

## Current Progress:

- ~~Basic Architecture~~
    - ~~Network~~ 
    - ~~Messaging~~
- Role Play
    - Follower
        - ~~Handle request vote~~
        - ~~Got action from client change to Candidate~~
        - Handle hearbeat from leader
    - Candidate
        - Request vote broadcast
        - Handle accept vote reedback
        - Win election broadcast
    - Leader
- Two major problem in Raft
    - Leader Election
    - Log Replication
    - Safety


Inspired By
---------------

- [Raft paper](http://nil.csail.mit.edu/6.824/2015/papers/raft-atc14.pdf)
- [Goraft](https://github.com/goraft/raft)
    - Inspired by programming style and logic.
- [Raft wiki](https://goo.gl/jrEs0a)
- [大数据日知录：架构与算法](http://product.dangdang.com/23561651.html)

License
---------------

This package is licensed under MIT license. See LICENSE for details.


