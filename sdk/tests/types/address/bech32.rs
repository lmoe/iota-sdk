// Copyright 2020-2023 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use iota_sdk::types::block::address::{Address, Bech32Address, Ed25519Address};

const ED25519_ADDRESS: &str = "0xebe40a263480190dcd7939447ee01aefa73d6f3cc33c90ef7bf905abf8728655";
const ED25519_BECH32: &str = "rms1qr47gz3xxjqpjrwd0yu5glhqrth6w0t08npney8000ust2lcw2r92j5a8rt";

#[test]
fn ctors() {
    let ed25519_address = ED25519_ADDRESS.parse::<Ed25519Address>().unwrap();
    let address = Address::Ed25519(ed25519_address);
    let bech32_address = Bech32Address::try_new("rms", address).unwrap();
    assert_eq!(bech32_address.inner(), &address);
    assert_eq!(bech32_address.hrp(), "rms");

    // This makes sure that the custom `try_from_str` method does the same as `FromStr::from_str` trait impl.
    assert_eq!(bech32_address, ED25519_BECH32.parse::<Bech32Address>().unwrap());
    assert_eq!(bech32_address, Bech32Address::try_from_str(ED25519_BECH32).unwrap());
}
