# Mew
## your private stock assistant
## Publish
* Mac
```
set GOOS=darwin
set GOARCH=amd64
go build
```
* Windows
```
set GOOS=windows
set GOARCH=386
go build
```
## Usage
### Auth
* You need to auth the mew to be able to let it work on your robinhood account
* We currently only supports the account with 2fa enabled
```
Y:\Projects\Mew>mew auth -u <username> -p <password> -m <2fa_code>
time="2020-05-16T23:57:10-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-16T23:57:10-07:00" level=info msg="Creating config file for <username>"
```
Above will create config.yml, which stores your info to call robinhood
### Version
```
Y:\Projects\Mew>mew -v
time="2020-05-16T23:57:46-07:00" level=info msg="welcome to use the Mew stock assistant"
Mew version v0.1.2_build_2020-05-14
```
### Buy
#### Market
```
Y:\Projects\Mew>mew b -t QQQ -v 300
time="2020-05-17T00:34:35-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:34:35-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:34:35-07:00" level=info msg="Loading config..."
Please confirm the order details.
Order type: Market Buy  Security: QQQ   Quantity: 1     price: 222.74 [y/n]y
time="2020-05-17T00:34:37-07:00" level=info msg="About to place Market Buy for 1 shares of QQQ at $222.74"
time="2020-05-17T00:34:38-07:00" level=info msg="Order placed for QQQ ID dba0080c-6888-49ea-afee-c77446249681"
```
#### Limit
```
Y:\Projects\Mew>mew lb -t QQQ -l 99.5 -v 300
time="2020-05-17T00:42:54-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:42:54-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:42:54-07:00" level=info msg="Loading config..."
time="2020-05-17T00:42:56-07:00" level=info msg="quoted price is 222.740000"
time="2020-05-17T00:42:56-07:00" level=info msg="limit price is 221.630000"
Please confirm the order details.
Order type: Limit Buy   Security: QQQ   Quantity: 1     price: 221.63 [y/n]y
time="2020-05-17T00:42:58-07:00" level=info msg="About to place Limit Buy for 1 shares of QQQ at $221.63"
time="2020-05-17T00:42:58-07:00" level=info msg="Order placed with order ID 5948a01a-e491-480d-9355-bc7c1ca6a33d"
```
#### Batch
```
Y:\Projects\Mew>mew lb -t @QQQ_SPY -l 99.5 -v 300
time="2020-05-17T00:43:33-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:43:33-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:43:33-07:00" level=info msg="Loading config..."
time="2020-05-17T00:43:34-07:00" level=info msg="quoted price is 222.740000"
time="2020-05-17T00:43:34-07:00" level=info msg="limit price is 221.630000"
Please confirm the order details.
Order type: Limit Buy   Security: QQQ   Quantity: 1     price: 221.63 [y/n]y
time="2020-05-17T00:43:36-07:00" level=info msg="About to place Limit Buy for 1 shares of QQQ at $221.63"
time="2020-05-17T00:43:36-07:00" level=info msg="Order placed with order ID b792c3a7-5ceb-4867-84a9-48876e2c988e"
time="2020-05-17T00:43:36-07:00" level=info msg="quoted price is 285.690000"
time="2020-05-17T00:43:36-07:00" level=info msg="limit price is 284.260000"
Please confirm the order details.
Order type: Limit Buy   Security: SPY   Quantity: 1     price: 284.26 [y/n]y
time="2020-05-17T00:43:37-07:00" level=info msg="About to place Limit Buy for 1 shares of SPY at $284.26"
time="2020-05-17T00:43:38-07:00" level=info msg="Order placed with order ID 37b4b3df-ed60-4daf-b262-ad1a5a17dbc2"
```
### Trailing Stop
```
D:\Projects\Mew>mew tsb -pt 10 -v 2000 -t QQQ
time="2020-06-09T22:36:32-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-06-09T22:36:32-07:00" level=info msg="Creating rhClient..."
time="2020-06-09T22:36:32-07:00" level=info msg="Loading config..."
time="2020-06-09T22:36:33-07:00" level=info msg="Quote price is 243.820000"
time="2020-06-09T22:36:33-07:00" level=info msg="The final price is 243.820000"
time="2020-06-09T22:36:33-07:00" level=info msg="Preparing trailing stop buy order of 7 shares with stop price 268.20"
Please confirm the order details.
Order type: Trailing Stop Buy   Security: QQQ   Quantity: 7     Market price: 243.82    Stop price: 268.20 [y/n]:y
time="2020-06-09T22:36:34-07:00" level=info msg="Order placed for QQQ with order ID: 69d5ea64-7fa0-4aec-b1f1
```
### Sell
#### Market
```
Y:\Projects\Mew>mew s -t QQQ -v 300
time="2020-05-17T00:44:55-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:44:55-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:44:55-07:00" level=info msg="Loading config..."
Please confirm the order details.
Order type: Market Sell Security: QQQ   Quantity: 1     price: 222.74 [y/n]y
time="2020-05-17T00:44:59-07:00" level=info msg="About to place Market Sell for 1 shares of QQQ at $222.74"
time="2020-05-17T00:44:59-07:00" level=info msg="Order placed for QQQ ID 008e41b7-a1bf-47a0-97d7-96f8be3f7621"
```
#### Limit
```
Y:\Projects\Mew>mew ls -t QQQ -v 300 -ls 101.5
time="2020-05-17T00:45:32-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:45:32-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:45:32-07:00" level=info msg="Loading config..."
Please confirm the order details.
Order type: Limit Sell  Security: QQQ   Quantity: 1     price: 226.08 [y/n]y
time="2020-05-17T00:45:34-07:00" level=info msg="About to place Limit Sell for 1 shares of QQQ at $226.08"
time="2020-05-17T00:45:34-07:00" level=info msg="Order placed for QQQ ID f9d279ed-146e-4a5e-868b-e90272158461"
```
#### Limit with Percent
```
D:\Projects\Mew>mew ls -ls 101 -ps 50 -t SPHD
time="2020-05-31T18:44:14-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-31T18:44:14-07:00" level=info msg="Creating rhClient..."
time="2020-05-31T18:44:14-07:00" level=info msg="Loading config..."
time="2020-05-31T18:44:15-07:00" level=info msg="Quote price is 0.000000"
time="2020-05-31T18:44:15-07:00" level=info msg="Updated price is 0.000000"
time="2020-05-31T18:44:15-07:00" level=error msg="Error returned from API: missing_instruments: [\"LNKD\"]"
time="2020-05-31T18:44:15-07:00" level=error msg="Error returned from API: inactive_instruments: [\"SYMC\"]"
time="2020-05-31T18:44:15-07:00" level=error msg="Error returned from API: missing_instruments: [\"ZU\"]"
time="2020-05-31T18:44:15-07:00" level=error msg="Error returned from API: inactive_instruments: [\"PNRA\"]"
time="2020-05-31T18:44:15-07:00" level=error msg="Error returned from API: inactive_instruments: [\"XIV\"]"
time="2020-05-31T18:44:15-07:00" level=info msg="processing 50.00 percent of current holding of SPHD with a total of 63 shares, which is 31 shares"
Please confirm the order details.
Order type: Limit Sell  Security: SPHD  Quantity: 31    price: 0.00 [y/n]
```
#### Batch
```
Y:\Projects\Mew>mew ls -t @QQQ_MSFT -v 300 -ls 101.5
time="2020-05-17T00:47:12-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-17T00:47:12-07:00" level=info msg="Creating rhClient..."
time="2020-05-17T00:47:12-07:00" level=info msg="Loading config..."
Please confirm the order details.
Order type: Limit Sell  Security: QQQ   Quantity: 1     price: 226.08 [y/n]y
time="2020-05-17T00:47:17-07:00" level=info msg="About to place Limit Sell for 1 shares of QQQ at $226.08"
time="2020-05-17T00:47:17-07:00" level=info msg="Order placed for QQQ ID 58cda674-7b7b-48f9-bc11-258808f41b43"
Please confirm the order details.
Order type: Limit Sell  Security: MSFT  Quantity: 1     price: 185.19 [y/n]y
time="2020-05-17T00:47:20-07:00" level=info msg="About to place Limit Sell for 1 shares of MSFT at $185.19"
time="2020-05-17T00:47:20-07:00" level=error msg="Execute() for MSFT error : Error returned from API: detail: \"Not enough shares to sell.\""
time="2020-05-17T00:47:20-07:00" level=fatal msg="Error returned from API: detail: \"Not enough shares to sell.\""
```
#### Batch with Percent
```
D:\Projects\Mew>mew ls -ls 101 -ps 50 -t @JETS_SPHD
time="2020-05-31T18:45:12-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-05-31T18:45:12-07:00" level=info msg="Creating rhClient..."
time="2020-05-31T18:45:12-07:00" level=info msg="Loading config..."
time="2020-05-31T18:45:13-07:00" level=info msg="Quote price is 15.340000"
time="2020-05-31T18:45:13-07:00" level=info msg="Updated price is 15.490000"
time="2020-05-31T18:45:14-07:00" level=error msg="Error returned from API: inactive_instruments: [\"SYMC\"]"
time="2020-05-31T18:45:14-07:00" level=error msg="Error returned from API: missing_instruments: [\"LNKD\"]"
time="2020-05-31T18:45:14-07:00" level=error msg="Error returned from API: missing_instruments: [\"ZU\"]"
time="2020-05-31T18:45:14-07:00" level=error msg="Error returned from API: inactive_instruments: [\"XIV\"]"
time="2020-05-31T18:45:14-07:00" level=error msg="Error returned from API: inactive_instruments: [\"PNRA\"]"
time="2020-05-31T18:45:14-07:00" level=info msg="processing 50.00 percent of current holding of JETS with a total of 227 shares, which is 113 shares"
Please confirm the order details.
Order type: Limit Sell  Security: JETS  Quantity: 113   price: 15.49 [y/n]
```
#### Trailing Stop
```
D:\Projects\Mew>mew tss -pt 10 -v 2000 -t QQQ
time="2020-06-09T22:35:05-07:00" level=info msg="welcome to use the Mew stock assistant"
time="2020-06-09T22:35:05-07:00" level=info msg="Creating rhClient..."
time="2020-06-09T22:35:05-07:00" level=info msg="Loading config..."
time="2020-06-09T22:35:07-07:00" level=info msg="Quote price is 243.820000"
time="2020-06-09T22:35:07-07:00" level=info msg="The final price is 243.820000"
time="2020-06-09T22:35:07-07:00" level=info msg="Preparing trailing stop sell order of 9 shares with stop price 219.44"
Please confirm the order details.
Order type: Trailing Stop Sell  Security: QQQ   Quantity: 9     Market price: 243.82    Stop price: 219.44 [y/n]:y
time="2020-06-09T22:35:14-07:00" level=info msg="Order placed for QQQ with order ID: 00d962dd-d98c-4aaa-bd91
```
