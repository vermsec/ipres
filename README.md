# IPres
Simple IP resolver with maps for better external surface enumeration and tracking.

## Installation Instructions

```bash
▶ go get github.com/vermsec/ipres
```

## Help Menu
```bash
  -help
        usage info
  -map
        generates IPmap
  -o string
        output (default "ipres.map")
  -v    
        verbose
  -version
        current version
```

## Usage
Running with a list of domains
```bash
▶ cat domains.txt |ipres
```
Saving output 
```bash
▶ cat domains.txt |ipres |tee ipres.log
```
Generating ip-map 
```bash
▶ cat domains.txt 
google.com
amazon.com
tesla.com
domaindoesntexist.f
facebook.com

▶ cat domains.txt |ipres -map -o ipres.map |tee ipres.out
172.217.167.174
2404:6800:4009:82f::200e
176.32.103.205
205.251.242.103
54.239.28.85
199.66.11.62
157.240.16.35
2a03:2880:f12f:83:face:b00c:0:25de

▶ cat ipres.map 
google.com : [172.217.167.174 2404:6800:4009:82f::200e]
amazon.com : [176.32.103.205 205.251.242.103 54.239.28.85]
tesla.com : [199.66.11.62]
domaindoesntexist.f : []
facebook.com : [157.240.16.35 2a03:2880:f12f:83:face:b00c:0:25de]
```
