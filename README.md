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
