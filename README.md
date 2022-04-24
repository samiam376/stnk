# stnk
 CLI to fetch stock prices

 ## Usage

 ### Fetch Intraday Stock Prices
 Fetches intraday prices at 5 minute intervels. Defaults to last 5, set last=-1 to show all prices
 ```./stnk intraday -symbol IBM -last 10 -api-key AlphaVantageAPIKey```

 ### Fetch Current Quote
 Fetches the current quote for a stock.
 ```./stnk quote -symbol IBM -api-key AlphaVantageAPIKey```
