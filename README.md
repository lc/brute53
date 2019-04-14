# brute53


A tool to bruteforce nameservers when working with subdomain delegations to AWS. Based off Frans Rosén's talk ["DNS hijacking using cloud providers - no verification needed"](https://youtu.be/FXCzdWm2qDg?t=1132)


### Installation:

```
go get -u github.com/lc/brute53
```

### Usage:
```
root@doggos:~# brute53 -c ~/.aws/credentials -delay 2 -t vulnerable.example.com
```

