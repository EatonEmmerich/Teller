### Teller App
_It's not a computer, It's a Teller_

## Prerequisites
postgresql:
```shell
psql -d postgres -c "CREATE DATABASE test;"
psql -d test -c "ALTER ROLE postgres WITH LOGIN CREATEDB;"
```

# DB
* Catalogue
* Stock (optional)
* Invoice
* Payments
* Reports (optional?)
 
# Payment Options
* LN
* Cash
* BTC (optional)

# Price Definitions......
* LND & BTC lookup price on exchange
* Flippening mode (Cash lookup price on exchange)

# LN Node
* Remote?
* Local? (too risky?)

# UI
* Outputs
  * lowest priority, just log for now
  * TBD
* Inputs
  * Sales
    * CMDLine
    * Scanners
    * Network?
  * Reports and admin
    * CSV?
    * CMDLine
    * Scheduled?
    * Abstract required (network endpoints)