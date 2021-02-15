# IOTA Wallet Library - Python binding

Python binding to the IOTA wallet library.

## Requirements

Ensure you have first installed the required dependencies for the library [here](https://github.com/iotaledger/wallet.rs/blob/develop/README.md).

## Build the library

Currently the package isn't published so you'd need to build by yourself.

- Go to `bindings/python/native`
- $ cargo build --release
- The built library is located in `target/release/` On MacOS, rename `libiota_wallet.dylib` to `iota_wallet.so`, on Windows, use `iota_wallet.dll` directly, and on Linux, rename `libiota_wallet.so` to `iota_wallet.so`.
- Copy your renamed library to `bindings/python/examples/``
- Go to `binding/python/examples`
- `$ python account_operations.py`


## Getting Started

The following shows an exmaple to use the python library.

### Example 

```python
# Copyright 2020 IOTA Stiftung
# SPDX-License-Identifier: Apache-2.0


import iota_wallet

# Create a AccountManager
manager = iota_wallet.AccountManager(
    storage='Stronghold', storage_path='./storage')
manager.set_stronghold_password("password")
manager.store_mnemonic("Stronghold")

# First we'll create an example account and store it
client_options = {'node': 'https://api.lb-0.testnet.chrysalis2.com'}
account_initialiser = manager.create_account(client_options)
account_initialiser.alias('alias')
account = account_initialiser.initialise()

# Update alias
account.set_alias('the new alias')

# Get unspent addresses
unspend_addresses = account.list_unspent_addresses()
print(f'Unspend addresses: {unspend_addresses}')

# Get spent addresses
spent_addresses = account.list_spent_addresses()
print(f'Spent addresses: {spent_addresses}')

# Get all addresses
addresses = account.addresses()
print(f'All addresses: {addresses}')

# Generate a new unused address
new_address = account.generate_address()
print(f'New address: {new_address}')

# List messages
messages = account.list_messages(5, 0, message_type='Failed')
print(f'Messages: {messages}')
```

## API Reference

Note that in the following APIs, the corresponding exception will be returned if an error occurs.
Also for all the optional values, the default values are the same as the ones in the Rust version.

### AccountManager

#### constructor(storage_path (optional), storage (optional), password (optional), polling_interval (optional)): [AccountManager](#accountmanager)

Creates a new instance of the AccountManager.

| Param              | Type             | Default                   | Description                                                           |
| ------------------ | ---------------- | ------------------------- | --------------------------------------------------------------------- |
| [storage_path]     | <code>str</code> | <code>`./storage`</code>  | The path where the database file will be saved                        |
| [storage]          | <code>str</code> | <code>`Stronghold`</code> | The storage implementation to use. Should be `Stronghold` or `Sqlite` |
| [storagePassword]  | <code>str</code> | <code>undefined</code>    | The storage password to encrypt/decrypt accounts                      |
| [polling_interval] | <code>int</code> | <code>30000</code>        | The polling interval in milliseconds                                  |

Note: if the `storage_path` is set, then the `storage` needs to be set too. An exception will be thrown when errors happened.

**Returns** The constructed [AccountManager](#accountmanager).

#### stop_background_sync(): void

Stops the background polling and MQTT monitoring.

#### set_storage_password(password): void

Sets the password used for encrypting the storage.

| Param    | Type             | Default                | Description          |
| -------- | ---------------- | ---------------------- | -------------------- |
| password | <code>str</code> | <code>undefined</code> | The storage password |

#### set_stronghold_password(password): void

Sets the stronghold password.

| Param    | Type             | Default                | Description          |
| -------- | ---------------- | ---------------------- | -------------------- |
| password | <code>str</code> | <code>undefined</code> | The storage password |

#### is_latest_address_unused(): bool

Determines whether all accounts have the latest address unused.

**Returns** `true` if the latest address is unused.

#### store_mnemonic(signer_type, mnemonic (optional)): bool

Stores a mnemonic for the given signer type.
If the mnemonic is not provided, we'll generate one.

| Param       | Type             | Default                         | Description                                                    |
| ----------- | ---------------- | ------------------------------- | -------------------------------------------------------------- |
| signer_type | <code>str</code> | <code>undefined</code>          | Should be `Stronghold`, `LedgerNano`, or `LedgerNanoSimulator` |
| mnemonic    | <code>str</code> | <code>randomly generated</code> | The provided mnemonic or the randomly generated one            |

#### generate_mnemonic(): str

Generates a new mnemonic.

**Returns** The generated mnemonic string.

#### verify_mnemonic(mnemonic): void

Checks is the mnemonic is valid. If a mnemonic was generated with `generate_mnemonic()`, the mnemonic here should match the generated.

| Param    | Type             | Default                | Description           |
| -------- | ---------------- | ---------------------- | --------------------- |
| mnemonic | <code>str</code> | <code>undefined</code> | The provided mnemonic |

#### create_account(client_options): [AccountInitialiser](#accountinitialiser)

Creat a new account.

| Param          | Type                                         | Default                | Description        |
| -------------- | -------------------------------------------- | ---------------------- | ------------------ |
| client_options | <code>[ClientOptions](#clientoptions)</code> | <code>undefined</code> | The client options |

**Returns** A constructed [AccountInitialiser](#accountinitialiser).

#### remove_account(account_id): void

Deletes an account.

| Param      | Type             | Default                | Description                            |
| ---------- | ---------------- | ---------------------- | -------------------------------------- |
| account_id | <code>str</code> | <code>undefined</code> | The account with this id to be deleted |

#### sync_accounts(): list[SyncedAccount]

Syncs all accounts.

**Returns** A promise resolving to an array of [SyncedAccount](#syncedaccount).

#### internal_transfer(from_account_id, to_account_id, amount): WalletMessage

Transfers an amount from an account to another.

| Param           | Type             | Default                | Description                                      |
| --------------- | ---------------- | ---------------------- | ------------------------------------------------ |
| from_account_id | <code>str</code> | <code>undefined</code> | The source of account id in the transfering      |
| to_account_id   | <code>str</code> | <code>undefined</code> | The destination of account id in the transfering |
| amount          | <code>int</code> | <code>undefined</code> | The transfer amount                              |

**Returns** A promise resolving to the transfer's [WalletMessage](#walletmessage).

#### backup(destination): str

Backups the storage to the given destination.

| Param       | Type             | Default                | Description                 |
| ----------- | ---------------- | ---------------------- | --------------------------- |
| destination | <code>str</code> | <code>undefined</code> | The path to the backup file |

**Returns** The full path to the backup file.

#### import_accounts(source, stronghold_password): void

Imports a database file.

| Param               | Type             | Default                | Description                    |
| ------------------- | ---------------- | ---------------------- | ------------------------------ |
| source              | <code>str</code> | <code>undefined</code> | The path to the backup file    |
| stronghold_password | <code>str</code> | <code>undefined</code> | The backup stronghold password |

#### get_account(account_id): [AccountHandle](#accounthandle)

Gets the account with the given identifier or index.

| Param      | Type             | Default                | Description                                          |
| ---------- | ---------------- | ---------------------- | ---------------------------------------------------- |
| account_id | <code>str</code> | <code>undefined</code> | The account id, alias, index or one of its addresses |

**Returns** the associated AccountHandle object or undefined if the account wasn't found.

#### get_accounts(): list[[AccountHandle](#accounthandle)]

Gets all stored accounts.

**Returns** an list of [AccountHandle](#accounthandle).

#### retry(account_id, message_id): [WalletMessage](#walletmessage)

Retries (promotes or reattaches) the given message.

| Param      | Type             | Default                | Description                                          |
| ---------- | ---------------- | ---------------------- | ---------------------------------------------------- |
| account_id | <code>str</code> | <code>undefined</code> | The account id, alias, index or one of its addresses |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier                             |

**Returns** the retried [WalletMessage](#walletmessage).

#### reattach(account_id, message_id): [WalletMessage](#walletmessage)

Reattach the given message.

| Param      | Type             | Default                | Description                                          |
| ---------- | ---------------- | ---------------------- | ---------------------------------------------------- |
| account_id | <code>str</code> | <code>undefined</code> | The account id, alias, index or one of its addresses |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier                             |

**Returns** the reattached [WalletMessage](#walletmessage).

#### promote(account_id, message_id): [WalletMessage](#walletmessage)

Promote the given message.

| Param      | Type             | Default                | Description                                          |
| ---------- | ---------------- | ---------------------- | ---------------------------------------------------- |
| account_id | <code>str</code> | <code>undefined</code> | The account id, alias, index or one of its addresses |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier                             |

**Returns** the promoted [WalletMessage](#walletmessage).

### AccountSynchronizer

#### gap_limit(limit): void

Set the number of address indexes that are generated.

| Param | Type             | Default                | Description                                      |
| ----- | ---------------- | ---------------------- | ------------------------------------------------ |
| limit | <code>int</code> | <code>undefined</code> | The number of address indexes that are generated |

#### skip_persistance(): void

Skip saving new messages and addresses on the account object.
The found [SyncedAccount](#syncedaccount) is returned on the `execute` call but won't be persisted on the database.

#### address_index(address_index): void

Set the initial address index to start syncing.

| Param         | Type             | Default                | Description                                |
| ------------- | ---------------- | ---------------------- | ------------------------------------------ |
| address_index | <code>int</code> | <code>undefined</code> | The initial address index to start syncing |

#### execute(): [SyncedAccount](#syncedaccount)

Syncs account with the tangle.
The account syncing process ensures that the latest metadata (balance, transactions) associated with an account is fetched from the tangle and is stored locally.

### Transfer

#### constructor(amount, address, bench32_hrp, indexation (optional), remainder_value_strategy: str): [Transfer](#transfer)

The `Transfer` object used in [SyncedAccount](#syncedaccount)

| Param                    | Type                                   | Default                | Description                                 |
| ------------------------ | -------------------------------------- | ---------------------- | ------------------------------------------- |
| amount                   | <code>int</code>                       | <code>undefined</code> | The amount to transfer                      |
| address                  | <code>str</code>                       | <code>undefined</code> | The addree to send                          |
| bench32_hrp              | <code>str</code>                       | <code>undefined</code> | the bench32_hrp of the address              |
| indexation               | <code>[Indexation](#indexation)</code> | <code>undefined</code> | The indexation payload                      |
| remainder_value_strategy | <code>str</code>                       | <code>undefined</code> | Should be `ReuseAddress` or `ChangeAddress` |

### SyncedAccount

#### transfer(transfer_obj): [WalletMessage](#walletmessage)

Transfer tokens.

| Param        | Type                               | Default                | Description                  |
| ------------ | ---------------------------------- | ---------------------- | ---------------------------- |
| transfer_obj | <code>[Transfer](#transfer)</code> | <code>undefined</code> | The transfer we want to make |

**Returns** the [WalletMessage](#walletmessage) which makes the transfering.

#### retry(message_id): [WalletMessage](#walletmessage)

Retries (promotes or reattaches) the given message.

| Param      | Type             | Default                | Description              |
| ---------- | ---------------- | ---------------------- | ------------------------ |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier |

**Returns** the retried [WalletMessage](#walletmessage).

#### reattach(message_id): [WalletMessage](#walletmessage)

Reattach the given message.

| Param      | Type             | Default                | Description              |
| ---------- | ---------------- | ---------------------- | ------------------------ |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier |

**Returns** the reattached [WalletMessage](#walletmessage).

#### promote(message_id): [WalletMessage](#walletmessage)

Promote the given message.

| Param      | Type             | Default                | Description              |
| ---------- | ---------------- | ---------------------- | ------------------------ |
| message_id | <code>str</code> | <code>undefined</code> | The message's identifier |

**Returns** the promoted [WalletMessage](#walletmessage).

### AccountHandle

#### sync(): [AccountSynchronizer](#accountsynchronizer)

**Returns** the [AccountSynchronizer](#accountsynchronizer) to setup the process to synchronize this account with the Tangle.

#### generate_address(): [Address](#address)

**Returns** a new unused address and links it to this account.

#### get_unused_address(): [Address](#address)

Synchronizes the account addresses with the Tangle and returns the latest address in the account, which is an address without balance.

**Returns** the latest address in the account.

#### is_latest_address_unused(): bool

Syncs the latest address with the Tangle and determines whether it's unused or not.
An unused address is an address without balance and associated message history.
Note that such address might have been used in the past, because the message history might have been pruned by the node.

**Returns** `true` if the latest address in the account is unused.

#### latest_address(): [Address](#address)

**Returns** the most recent address of the account.

#### addresses(): list[[Address](#address)]

**Returns** a list of [Address](#address) in the account.

#### balance(): [AccountBalance](#accountbalance)

Gets the account balance information.

**Returns** the [AccountBalance](#accountbalance) in this account.

#### set_alias(alias): void

Updates the account alias.

| Param | Type             | Default                | Description              |
| ----- | ---------------- | ---------------------- | ------------------------ |
| alias | <code>str</code> | <code>undefined</code> | The account alias to set |

#### set_client_options(options): void

Updates the account's client options.

| Param   | Type                                         | Default                | Description               |
| ------- | -------------------------------------------- | ---------------------- | ------------------------- |
| options | <code>[ClientOptions](#clientoptions)</code> | <code>undefined</code> | The client options to set |

#### list_messages(count, from, message_type (optional)): list([WalletMessage](#walletmessage))

Get the list of messages of this account.

| Param        | Type             | Default                | Description                                                       |
| ------------ | ---------------- | ---------------------- | ----------------------------------------------------------------- |
| count        | <code>int</code> | <code>undefined</code> | The count of the messages to get                                  |
| from         | <code>int</code> | <code>undefined</code> | The iniital address index                                         |
| message_type | <code>str</code> | <code>undefined</code> | Should be `Received`, `Sent`, `Failed`, `Unconfirmed`, or `Value` |

#### list_spent_addresses(): list[[Address](#address)]

**Returns** the list of spent [Address](#address) in the account.

#### get_message(message_id): WalletMessage](#walletmessage) (optional)

Get the [WalletMessage](#walletmessage) by the message identifier in the account if it exists.

### AccountInitialiser

#### signer_type(signer_type): void

Sets the account type.

| Param       | Type             | Default                  | Description                                                    |
| ----------- | ---------------- | ------------------------ | -------------------------------------------------------------- |
| signer_type | <code>str</code> | <code>signer_type</code> | Should be `Stronghold`, `LedgerNano`, or `LedgerNanoSimulator` |

#### alias(alias): void

Defines the account alias. If not defined, we'll generate one.

| Param | Type             | Default                | Description       |
| ----- | ---------------- | ---------------------- | ----------------- |
| alias | <code>str</code> | <code>undefined</code> | The account alias |

#### created_at(created_at): void

Time of account creation.

| Param      | Type             | Default                | Description               |
| ---------- | ---------------- | ---------------------- | ------------------------- |
| created_at | <code>u64</code> | <code>undefined</code> | The account creation time |


#### messages(messages): void

Messages associated with the seed.
The account can be initialised with locally stored messages.

| Param    | Type                                               | Default                | Description                 |
| -------- | -------------------------------------------------- | ---------------------- | --------------------------- |
| messages | <code>list([WalletMessage](#walletmessage))</code> | <code>undefined</code> | The locally stored messages |

#### addresses(addresses): list([WalletAddress](#walletaddress))

Address history associated with the seed.
The account can be initialised with locally stored address history.

| Param     | Type                                               | Default                | Description              |
| --------- | -------------------------------------------------- | ---------------------- | ------------------------ |
| addresses | <code>list([WalletAddress](#walletaddress))</code> | <code>undefined</code> | The historical addresses |


#### skip_persistance(): void

Skips storing the account to the database.

#### initialise(): [AccountHandle](#accounthandle)

Initialises the account.

**Returns** the initilized [AccountHandle](#accounthandle)

### WalletAddress

A dict with the following key/value pairs.

```python
wallet_address = {
    'address': str,
    'balance': int,
    'key_index': int,
    'internal': bool,
    'outputs': list[WalletAddressOutput],
}
```

Please refer to [WalletAddressOutput](#walletaddressoutput) for the details of this type.

### WalletAddressOutput

A dict with the following key/value pairs.

```python
wallet_address_output = {
    'transaction_id': str,
    'message_id': str,
    'index': int,
    'amount': int,
    'is_spent': bool,
    'address': str,
    'kind': str,
}
}
```

### Address

A dict with the following key/value pairs.

```python
address = {
    'address': AddressWrapper,
    'balance': int,
    'key_index': int,
    'internal': bool,
    'outputs': list[AddressOutput],
}
```

Please refer to [AddressWrapper](#addresswrapper) and [AddressOutput](#addressoutput) for the details of this type.

### AddressWrapper

A dict with the following key/value pairs.

```python
address_wrapper = {
    'inner': str
}
```

### AddressOutput

A dict with the following key/value pairs.

```python
address_output = {
    'transaction_id': str,
    'message_id': str,
    'index': int,
    'amount': int,
    'is_spent': bool,
    'address': AddressWrapper,
}
```

Please refer to [AddressWrapper](#addresswrapper) for the details of this type.

### AccountBalance

A dict with the following key/value pairs.

```python
account_balance = {
    'total': int,
    'available': int,
    'incoming': int,
    'outgoing': int,
}
```

### ClientOptions

A dict with the following key/value pairs.

```python
client_options = {
    'nodes': list[str] (opitonal),
    'node_pool_urls': list[str] (opitonal),
    'network': str (optional),
    'mqtt_broker_options': [BrokerOptions](#brokeroptions) (optional), 
    'local_pow': bool (optional),
    'node_sync_interval': int (optional), # in milliseconds
    'node_sync_enabled': bool (optional),
    'request_timeout': int (optional), # in milliseconds
    'api_timeout': {
        'GetTips': int (optional) # in milliseconds
        'PostMessage': int (optional) # in milliseconds
        'GetOutput': int (optional) # in milliseconds
    } (optional)
}
```

Note that this message object in `wallet.rs` is not the same as the message object in `iota.rs`.

### BrokerOptions

A dict with the following key/value pairs.

```python
broker_options = {
    'automatic_disconnect': bool (optional),
    'timeout': int (optional),
    'use_websockets': bool (optional)
}
```

### WalletMessage

A dict with the following key/value pairs.

```python
wallet_message = {
    'id': str,
    'version': u64,
    'parents': list[str],
    'payload_length': int,
    'payload': Payload,
    'timestamp': int,
    'nonce': int,
    'confirmed': bool (optional),
    'broadcasted': bool,
    'incoming': bool,
    'value': int,
    'remainder_value': int
}
```

Please refer to [Payload](#payload) for the details of this type.

### Payload

A dict with the following key/value pairs.

```python
payload = {
    'transaction': list[Transaction] (optional),
    'milestone': list[Milestone] (optional),
    'indexation': list[Indexation] (optional),
}
```

Please refer to [Transaction](#transaction), [Milestone](#milestone), and [Indexation](#indexation) for the details of these types.

### Transaction

A dict with the following key/value pairs.

```python
transaction = {
    'essence': TransactionPayloadEssence,
    'unlock_blocks': list[UnlockBlock],
}
```

Please refer to [TransactionPayloadEssence](#transactionpayloadessence) and [UnlockBlock](#unlockblock) for the details of these types.

### Milestone

A dict with the following key/value pairs.

```python
milestone = {
    'essence': MilestonePayloadEssence,
    'signatures': list[bytes],
}
```

Please refer to [MilestonePayloadEssence](#milestonepayloadessence) for the details of this type.

### MilestonePayloadEssence

A dict with the following key/value pairs.

```python
milestone_payload_essence = {
    'index': int,
    'timestamp': int,
    'parents': list[str],
    'merkle_proof': bytes,
    'public_keys': bytes
}
```

### Indexation

A dict with the following key/value pairs.

```python
indexation = {
    'index': str,
    'data': bytes
}
```

### TransactionPayloadEssence

A dict with the following key/value pairs.

```python
transaction_payload_essence = {
    'inputs': list[Input],
    'outputs': list[Output],
    'payload': Payload (optional),
}
```
Please refer to [Input](#input), [Output](#output), and [Payload](#payload) for the details of these types.

### Output

A dict with the following key/value pairs.

```python
output = {
    'address': str,
    'amount': int
}
```

### Input

A dict with the following key/value pairs.

```python
input = {
    'transaction_id': str,
    'index': int
}
```

### UnlockBlock

A dict with the following key/value pairs.

```python
unlock_block = {
    'signature': Ed25519Signature (optional),
    'reference': int (optional)
}
```

Please refer to [Ed25519Signature](#ed25519signature) for the details of this type.

### Ed25519Signature

A dict with the following key/value pairs.

```python
ed25519_signature = {
    'public_key': bytes,
    'public_key': bytes
}
```