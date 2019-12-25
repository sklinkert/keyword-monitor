# keyword-monitor
Monitor your keyword positions in Google Search.

`keyword-monitor` uses [Google Custom Search (CSE)](https://developers.google.com/custom-search/). It takes a list of keywords and your domain. It will query CSE for your keywords and checks the first pages to find your domain results.

## Usage

```ID=0093493... DOMAIN=etfs24.de KEYWORDS='keyword1,keyword2,keyword3' keyword-monitor```

INFO[0005] Found                                         Domain=etfs24.de Keyword="trade republic etf" Position=50 Title="trade republic - Moderne und unkomplizierte Geldanlage mit ETFs ..." TotalKeywords=33 URL="https://etfs24.de/tag/trade-republic/"
INFO[0006] Found                                         Domain=etfs24.de Keyword="trade republic sparplan" Position=32 Title="Der große Depot-Vergleich für ETF-Sparer 2019" TotalKeywords=33 URL="https://etfs24.de/der-grosse-etf-depot-vergleich/"
INFO[0010] Not found                                     Domain=etfs24.de Keyword="deka etf" TotalKeywords=33
INFO[0015] Not found                                     Domain=etfs24.de Keyword="etf msci world" TotalKeywords=33
INFO[0020] Not found                                     Domain=etfs24.de Keyword="dividenden etf" TotalKeywords=33

INFO[0162] Found                                         Domain=etfs24.de Keyword="vl etf sparplan" Position=28 Title="Vermögenswirksame Leistungen ETF - Geldanlag mit attraktiver ..." TotalKeywords=33 URL="https://etfs24.de/vermoegenswirksame-leistungen-etf/"
INFO[0163] Found                                         Domain=etfs24.de Keyword="msci world etf sparplan" Position=24 Title="ETF Sparplan MSCI World - die Allzweckwaffe für jeden Anleger?" TotalKeywords=33 URL="https://etfs24.de/etf-sparplan-msci-world-die-allzweckwaffe-fur-jeden-anleger/"

[...]

**ID**: The ID from your [https://cse.google.com/cse/all](https://cse.google.com/cse/all) 

**KEYWORDS**: comma separated list of keywords

**DOMAIN**: your website domain