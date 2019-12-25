# keyword-monitor
Monitor your keyword positions in Google Search.

`keyword-monitor` uses [Google Custom Search (CSE)](https://developers.google.com/custom-search/). It takes a list of keywords and your domain. It will query CSE for your keywords and checks the first pages to find your domain results.

## Usage

```
ID=0093493... DOMAIN=etfs24.de KEYWORDS='trade republic etf,trade republic sparplan,deka etf,msci world etf sparplan' keyword-monitor
INFO[0005] Found                                         Domain=etfs24.de Keyword="trade republic etf" Position=50  TotalKeywords=33 URL="https://etfs24.de/tag/trade-republic/"
INFO[0006] Found                                         Domain=etfs24.de Keyword="trade republic sparplan" Position=32  TotalKeywords=33 URL="https://etfs24.de/der-grosse-etf-depot-vergleich/"
INFO[0010] Not found                                     Domain=etfs24.de Keyword="deka etf" TotalKeywords=33
INFO[0163] Found                                         Domain=etfs24.de Keyword="msci world etf sparplan" Position=24 TotalKeywords=33 URL="https://etfs24.de/etf-sparplan-msci-world-die-allzweckwaffe-fur-jeden-anleger/"

[...]
```

**ID**: The ID from your [https://cse.google.com/cse/all](CSE) 

**KEYWORDS**: comma separated list of keywords

**DOMAIN**: your website domain
