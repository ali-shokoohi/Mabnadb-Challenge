
# Mabnadb-Challenge

APIs for the Mabnadb-Challengs

## Indices

* [Instruments](#instruments)

  * [Create A Instruments](#1-create-a-instruments)
  * [Delete A Instrument](#2-delete-a-instrument)
  * [Get A Instrument](#3-get-a-instrument)
  * [Get Instruments](#4-get-instruments)

* [Trades](#trades)

  * [Create A Trade](#1-create-a-trade)
  * [Delete A Trade](#2-delete-a-trade)
  * [Get A Trade](#3-get-a-trade)
  * [Get Trades](#4-get-trades)


--------


## Instruments
Instrument's APIs



### 1. Create A Instruments


Create A single new instrument


***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://127.0.0.1:8000/instruments
```



***Body:***

```js        
{
    "name": "Test"
}
```



### 2. Delete A Instrument


Delete a single instrument


***Endpoint:***

```bash
Method: DELETE
Type: 
URL: http://127.0.0.1:8000/instruments/{{$randomInt}}
```



### 3. Get A Instrument


Get a single instrument


***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:8000/instruments/{{$randomInt}}
```



### 4. Get Instruments


Get all of the Instruments


***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:8000/instruments
```



## Trades
Trade's APIs



### 1. Create A Trade


Create A single new trade


***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://127.0.0.1:8000/trades
```



***Body:***

```js        
{
    "open": 14,
    "close": 18,
    "high": 2,
    "low": 5,
    "date_en": "2020-04-07T00:00:00Z",
    "instrument_id": 3
}
```



### 2. Delete A Trade


Delete a single trade



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: http://127.0.0.1:8000/trades/{{$randomInt}}
```



### 3. Get A Trade


Get a single trade



***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:8000/trades/{{$randomInt}}
```



### 4. Get Trades


Get all of the Instruments


***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:8000/trades
```



***Available Variables:***

| Key | Value | Type |
| --- | ------|-------------|
| base_url | http://127.0.0.1:8000 |  |



---
[Back to top](#mabnadb-challenge)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-09-06 02:19:47 by [docgen](https://github.com/thedevsaddam/docgen)
