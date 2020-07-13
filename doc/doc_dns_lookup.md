DNS maintains the mapping relationship between domain names and IP addresses, and the IP address where it is located can be found through the domain name.

## Introduction

Pallas dnslookup tool aims to provide related dns resolution tools, such as lookup DNS records.

Pallas dnslookup tool use the functions provided by golang's internal net package to implement DNS query.

## Command

```
pallas dnsloolup [--record-type] domain
```

support `record-type`:

|record-type|param value|description|
|--|--|--|
|A|A|Point the domain name to an IPv4 address|
|AAAA|AAAA|Point the host name (or domain name) to an IPv6 address|
|CNAME|CNAME|If you point the domain name to a domain name to achieve the same access effect <br> as the domain name being pointed to, you need to add CNAME records|
|NS|NS|Domain name resolution server records, if you want to specify a domain name server<br> for subdomain name resolution, you need to set up NS records|
|MX|MX|Create an email service that will point to the mail server address, and you need to set up MX records|
|PTR|PTR|The PTR record is the reverse record of the A record, also known as the IP reverse search record <br> or pointer record, which is responsible for reversely resolving the IP to the domain name|
|SOA|SOA|SOA is called the initial authority record, NS is used to identify multiple domain name resolution servers|

## Example

Specify parameters:

```shell
➜  pallas dnslookup --record-type=A baidu.com

[Info] Start fetching data from ICANN.
[Info] Target:baidu.com

+-------------+-------------------------------+
| RECORD TYPE |            VALUES             |
+-------------+-------------------------------+
| A           | [220.181.38.148 39.156.69.79] |
+-------------+-------------------------------+

```

Check all:

```shell
➜  pallas dnslookup baidu.com          

[Info] Start fetching data from ICANN.
[Info] Target:baidu.com

+-------------+-------------------------------------------------------------------------------------------------------+
| RECORD TYPE |                                                VALUES                                                 |
+-------------+-------------------------------------------------------------------------------------------------------+
| A           | [39.156.69.79 220.181.38.148]                                                                         |
+-------------+-------------------------------------------------------------------------------------------------------+
| MX          | [mx.maillb.baidu.com. 10 mx.n.shifen.com. 15 mx1.baidu.com. 20 jpmx.baidu.com. 20 mx50.baidu.com. 20] |
+-------------+-------------------------------------------------------------------------------------------------------+
| CNAME       | baidu.com.                                                                                            |
+-------------+-------------------------------------------------------------------------------------------------------+
| PTR         | []                                                                                                    |
+-------------+-------------------------------------------------------------------------------------------------------+
| AAAA        |                                                                                                       |
+-------------+-------------------------------------------------------------------------------------------------------+
| NS          | [ns4.baidu.com. ns1.baidu.com. ns3.baidu.com. ns7.baidu.com. ns2.baidu.com.]                          |
+-------------+-------------------------------------------------------------------------------------------------------+

```