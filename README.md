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

## Features

|command|description|
|---|---|
|dnslookup|Query all records of a domain name, including A/AAAA/PTR/MX/NS and other records.|
|iplookup|Query IP info.|
|whois|Whois query the transfer protocol of the domain name's IP and owner information.|
|scan|ARP/ICMP Scanner.|


## Getting Started

### Lookup DNS

Dnslook is a dns query tool to help you query the dns records of a domain name. Use `dnslookup` to get dns infomations.

command:

```
pallas dnslookup [domain] [--recore_type] [--help]
```

example:

```
➜  pallas dnslookup qiniu.cn
Start fetching data from ICANN.
Target:qiniu.cn

+-------------+-----------------------------------------+
| RECORD TYPE |                 VALUES                  |
+-------------+-----------------------------------------+
| PTR         | []                                      |
+-------------+-----------------------------------------+
| AAAA        |                                         |
+-------------+-----------------------------------------+
| NS          | [dns32.hichina.com. dns31.hichina.com.] |
+-------------+-----------------------------------------+
| A           | [47.91.169.15]                          |
+-------------+-----------------------------------------+
| MX          | []                                      |
+-------------+-----------------------------------------+
| CNAME       | overdue.aliyun.com.                     |
+-------------+-----------------------------------------+
```

### Whois Tool

`whois` query domain name or IP attribution information.

command:

```
pallas whois [domain] [--beautify] [--help] [--version]
```

example:

```shell
➜  pallas whois baidu.com --beautify true
whois: start fetching information for  baidu.com
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
|   ATTRIBUTE    |                                                                      VALUE                                                                      |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| DNS服务器      | [ns1.baidu.com ns2.baidu.com ns3.baidu.com ns4.baidu.com ns7.baidu.com]                                                                         |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 所在省         | Beijing                                                                                                                                         |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 注册商         | MarkMonitor, Inc.                                                                                                                               |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 注册日期       | 1999-10-11T11:05:17Z                                                                                                                            |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 到期日期       | 2026-10-11T11:05:17Z                                                                                                                            |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 更新时间       | 2019-05-09T04:30:46Z                                                                                                                            |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 域名状态       | [clientdeleteprohibited clienttransferprohibited clientupdateprohibited serverdeleteprohibited servertransferprohibited serverupdateprohibited] |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 所有者         |                                                                                                                                                 |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 所有者联系邮箱 | select request email form at https://domains.markmonitor.com/whois/baidu.com                                                                    |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 注册机构       | Beijing Baidu Netcom Science Technology Co., Ltd.                                                                                               |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 所在国家       | CN                                                                                                                                              |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
| 所在城市       |                                                                                                                                                 |
+----------------+-------------------------------------------------------------------------------------------------------------------------------------------------+
```

### Lookup IP

IP query tool, including gathering local ip, public ip and detail for target ip.

command:

```
pallas iplookup [--ip]
```

example:

```shell
➜  Pallas git:(master) ✗ go run main.go iplookup
+----------------------------------+-----------------------+
|            ATTRIBUTE             |         VALUE         |
+----------------------------------+-----------------------+
| Local IP                         | 10.0.2.15             |
|                                  | 172.17.0.1            |
+----------------------------------+-----------------------+
| Public IP(From myexternalip.com) | 211.123.126.33         |
+----------------------------------+-----------------------+
| Public IP(From Dns Server)       | 10.0.2.15             |
+----------------------------------+-----------------------+
| Mac                              | 03:00:27:0e:61:8d     |
+----------------------------------+-----------------------+
| Mask                             | ffff0000              |
+----------------------------------+-----------------------+
| IP Range                         | 10.0.2.1 ~ 10.0.2.254 |
+----------------------------------+-----------------------+
```

# Run 

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