# DomainRecon 

A tool that can gather subdomains and spider urls on standard input.

Usage for gathering subdomains: 

```
cat rootDomains.txt | domainRecon 
```

Usage for spidering urls: 

```
cat urls.txt | domainRecon -c --depth 3
```

Flags: 
```
-c  --crawl     enable crawling
-d  --depth     choose depth for crawing
```
