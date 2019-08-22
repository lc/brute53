# brute53


A tool to bruteforce nameservers when working with subdomain delegations to AWS. Based off Frans Rosén's talk ["DNS hijacking using cloud providers - no verification needed"](https://youtu.be/FXCzdWm2qDg?t=1132)

⚠️ Note: this tool is currently not working and is on the list to be fixed at somepoint


### Pre-requisites:
- golang
- AWS IAM User with access to [Route53](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html)

### Installation:

```
go get -u github.com/lc/brute53
```

### Usage:
```
root@doggos:~# brute53 -c ~/.aws/credentials -delay 2 -t vulnerable.example.com
```


<a href="http://buymeacoff.ee/cdl" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

