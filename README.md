# Pallas

```
____   __    __    __      __    ___ 
(  _ \ /__\  (  )  (  )    /__\  / __)
 )___//(__)\  )(__  )(__  /(__)\ \__ \
(__) (__)(__)(____)(____)(__)(__)(___/ 

Pallas is a collection of network security tools.
Use Pallas to help you solve most network security problems.
For suggestions or comments, please visit https://github.com/Anderson-Lu/Pallas
```

## Introduction

Pallas is a network toolbox that helps you troubleshoot some network security issues, including information gathering, detection, scanning and even intrusion.

## Commands

|command|description|
|---|---|
|[dnslookup](./doc/doc_dns_lookup.md)|Query all records of a domain name, including A/AAAA/PTR/MX/NS and other records.|
|[iplookup](./doc/doc_ip_lookup.md)|Query IP info.|
|[whois](./doc/doc_whois.md)|Whois query the transfer protocol of the domain name's IP and owner information.|
|scan|ARP/ICMP Scanner.|

## Features

- [DNS] Find DNS records(A/AAAA/PTR/MX/CNAME)
- [IP] Find the IP location
- [Whois] Find registration information for ip or domain name   
- [ARP] ARP intranet scanning
- [ICMP] ICMP host online detection

## Installation

#### Running from source code

```
➜ git clone https://github.com/Anderson-Lu/Pallas.git
➜ cd Pallas
➜ go run main.go --help
```

#### Running from dockerfile

```
➜ git clone https://github.com/Anderson-Lu/Pallas.git
➜ cd Pallas
```

#### Running from release packges

```
➜ wget xxx
➜ ./pallas --help
```


# Acknowledgement

- [Cobra](https://github.com/spf13/cobra): A Commander for modern Go CLI interactions
- [TableWriter](https://github.com/olekukonko/tablewriter): ASCII table in golang
- [Spinner](https://github.com/briandowns/spinner): Go (golang) package with 70+ configurable terminal spinner/progress indicators.