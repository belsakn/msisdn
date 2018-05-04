# msisdn
1. Create an application with following requirements:

latest PHP or Golang
takes MSISDN as an input
returns MNO identifier, country dialling code, subscriber number and country identifier as defined with ISO 3166-1-alpha-2
do not care about number portability

2. Write all needed tests.

3. Expose solution through REST API.

4. Use swagger, GNU Make, docker and docker compose for bonus points.

5. Other:

a git repository with full commit history is expected to be part of the delivered solution
if needed, provide additional installation instructions, but there shouldn't be much more than running a simple command to set everything up
use best practices all around. For PHP, good source of that would be http://www.phptherightway.com

Important: Do not take this task lightly. You will be judged according to the quality, completion and perfection of the task.

## Requirements

- [Go](https://golang.org/)

## Usage 

Navigate to %GOPATH/src
```
git clone https://github.com/mali8/msisdn.git
cd msisdn
go get
go install
msisdn
```
Go to: localhost:8080
Api call example: localhost:8080/parse?phoneNumber=%2B38640579602

## Tests

Navigate to %GOPATH/src/msisdn
```
go test
```

Author: Nejc Belšak

E-mail: belsak@gmail.com
