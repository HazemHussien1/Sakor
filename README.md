# Sakor 

A tool that can gather subdomains and spider urls on standard input.

Usage for gathering subdomains: 

```
cat rootDomains.txt | Sakor 
```

Usage for spidering urls: 

```
cat urls.txt | Sakor -c --depth 3
```

Flags: 
```
-c  --crawl     enable crawling
-d  --depth     choose depth for crawing
-v  --verbose   enable verbose mode
```

Sources requested for subdomains
1- Crtsh (https://crt.sh)
2- Certspotter (https://cerspotter.com)
3- Virus Total (https://virustotal.com)

##Todos:
1- Add concurrency to subdomain enumeration
2- Add more sources for subdomain enumeration
