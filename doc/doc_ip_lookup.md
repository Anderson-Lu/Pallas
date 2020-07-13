Google Public DNS is a free domain name resolution service (DNS) provided by Google on December 5, 2009. It's dns server is `8.8.8.8`, This is the source of the function of this tool.This is the source of the function of this toolThis is the source of the function of this tool

## Command

```shell
pallas iplookup [-ip specifiy ip address] [default localhost]
```

## Example

```
➜  pallas iplookup

+----------------------------------+-----------------------+
|            ATTRIBUTE             |         VALUE         |
+----------------------------------+-----------------------+
| Local IP                         | 10.0.2.15             |
|                                  | 172.17.0.1            |
|                                  | 172.18.0.1            |
|                                  | 172.19.0.1            |
|                                  | 172.20.0.1            |
+----------------------------------+-----------------------+
| Public IP(From myexternalip.com) | 213.116.15.31         |
+----------------------------------+-----------------------+
| Public IP(From Dns Server)       | 10.0.2.15             |
+----------------------------------+-----------------------+
| Mac                              | 32:42:62:ef:22:60     |
+----------------------------------+-----------------------+
| Mask                             | ffff0000              |
+----------------------------------+-----------------------+
| IP Range                         | 10.0.2.1 ~ 10.0.2.254 |
+----------------------------------+-----------------------+
```