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


<a href="http://buymeacoff.ee/cdl" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

