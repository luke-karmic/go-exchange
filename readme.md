# GoLang Orderbook (Simulation)

# Requirements
* Can create a orderbook for a single asset
  * Can fulfill a limit order, which is a collection of partial fill bids on asks
  * Price of asset should be impacts based on the supply and available orders
  * A MatchingEngine handles the price matching for the orderbook prices available using FIFO
  * When a Limit order is made, the collection of asks with their prices fulfilled, as well as the spread is output to the console
* Can create a OB Simulation
  * Create (x) asset, with initial supply (x) and value (x)
  * Creates (x) market participants who hold a random supply
  * Every (x) seconds, market participants place random Buys/Sells (bids/asks)

# TODO
- [x] Create Orderbook structure
- [] Create Limit, Asks, Bids structure
- [] Limit order market making
- [] Create market participants
- [] Create simulation