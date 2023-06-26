// Copyright 2023 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::{ffi::c_void, sync::Arc};
use std::ffi::{c_char, CStr, CString};
use std::ptr::null;

use iota_sdk_bindings_core::{
    call_wallet_method as rust_call_wallet_method,
    iota_sdk::wallet::{events::types::WalletEventType, Wallet as RustWallet},
    Response, WalletMethod, WalletOptions,
};
use tokio::sync::RwLock;
use iota_sdk_bindings_core::iota_sdk::types::block::ConvertTo;

use crate::{
    client::Client,
    error::{Error, Result},
    SecretManager,
};
use crate::error::set_last_error;

pub struct Wallet {
    pub wallet: Arc<RwLock<Option<RustWallet>>>,
}

/// Destroys the wallet instance.
unsafe fn internal_destroy_wallet(wallet_ptr: *mut Wallet) -> Result<()> {
    let wallet = {
        assert!(!wallet_ptr.is_null());
        &mut *wallet_ptr
    };

    crate::block_on(async {
        *wallet.wallet.write().await = None;
    });
    Ok(())
}

#[no_mangle]
pub unsafe extern "C" fn destroy_wallet(wallet_ptr: *mut Wallet) -> bool {
    match internal_destroy_wallet(wallet_ptr) {
        Ok(v) => true,
        Err(e) => { set_last_error(e); false }
    }
}

#[no_mangle]
pub unsafe extern "C" fn GO_EXPORT() {

}

/// Create wallet handler for python-side usage.
unsafe fn internal_create_wallet(options_ptr: *const c_char) -> Result<*const Wallet> {
    let options_string = CStr::from_ptr(options_ptr).to_str().unwrap();

    let wallet_options = serde_json::from_str::<WalletOptions>(options_string)?;
    let wallet = crate::block_on(async { wallet_options.build().await })?;

    let wallet_ptr = &Wallet {
        wallet: Arc::new(RwLock::new(Some(wallet))),
    };

    std::mem::forget(wallet_ptr);

    Ok(wallet_ptr)
}

#[no_mangle]
pub unsafe extern "C" fn create_wallet(options_ptr: *const c_char) -> *const Wallet {
    match internal_create_wallet(options_ptr) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}

/// Call a wallet method.
unsafe fn internal_call_wallet_method(wallet_ptr: *mut Wallet, method_ptr: *const c_char) -> Result<*const c_char> {
    let wallet = {
        assert!(!wallet_ptr.is_null());
        &mut *wallet_ptr
    };

    let method_string = CStr::from_ptr(method_ptr).to_str().unwrap();

    let method = serde_json::from_str::<WalletMethod>(&method_string)?;
    let response = crate::block_on(async {
        match wallet.wallet.read().await.as_ref() {
            Some(wallet) => rust_call_wallet_method(wallet, method).await,
            None => Response::Panic("wallet got destroyed".into()),
        }
    });

    let response_string = serde_json::to_string(&response)?;
    let s = CString::new(response_string).unwrap();

    Ok(s.into_raw())
}


#[no_mangle]
pub unsafe extern "C" fn call_wallet_method(wallet_ptr: *mut Wallet, method_ptr: *const c_char) -> *const c_char  {
    match internal_call_wallet_method(wallet_ptr, method_ptr) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}

/// Listen to wallet events.
unsafe fn internal_listen_wallet(wallet_ptr: *mut Wallet, events: Vec<String>, handler: extern "C" fn()) {
    let wallet = {
        assert!(!wallet_ptr.is_null());
        &mut *wallet_ptr
    };

    let mut rust_events = Vec::with_capacity(events.len());

    for event in events {
        let event = match serde_json::from_str::<WalletEventType>(&event) {
            Ok(event) => event,
            Err(e) => {
                panic!("Wrong event to listen! {e:?}");
            }
        };
        rust_events.push(event);
    }

    crate::block_on(async {
        wallet
            .wallet
            .read()
            .await
            .as_ref()
            .expect("wallet got destroyed")
            .listen(rust_events, move |_| handler())
            .await;
    });
}

#[no_mangle]
pub unsafe extern "C" fn listen_wallet(wallet_ptr: *mut Wallet, events: Vec<String>, handler: extern "C" fn()) -> bool {
    internal_listen_wallet(wallet_ptr, events, handler);
    true
}

/// Get the client from the wallet.
unsafe fn internal_get_client_from_wallet(wallet_ptr: *mut Wallet) -> Result<*const Client> {
    let wallet = {
        assert!(!wallet_ptr.is_null());
        &mut *wallet_ptr
    };

    let client = crate::block_on(async {
        wallet
            .wallet
            .read()
            .await
            .as_ref()
            .map(|w| w.client().clone())
            .ok_or_else(|| {
                Error::from(
                    serde_json::to_string(&Response::Panic("wallet got destroyed".into()))
                        .expect("json to string error")
                        .as_str(),
                )
            })
    })?;

    let client_ptr = &Client { client };
    std::mem::forget(client_ptr);

    Ok(client_ptr)
}

#[no_mangle]
pub unsafe extern "C" fn get_client_from_wallet(wallet_ptr: *mut Wallet) -> *const Client {
    match internal_get_client_from_wallet(wallet_ptr) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}

/// Get the secret manager from the wallet.
unsafe fn internal_get_secret_manager_from_wallet(wallet_ptr: *mut Wallet) -> Result<*const SecretManager> {
    let wallet = {
        assert!(!wallet_ptr.is_null());
        &mut *wallet_ptr
    };

    let secret_manager = crate::block_on(async {
        wallet
            .wallet
            .read()
            .await
            .as_ref()
            .map(|w| w.get_secret_manager().clone())
            .ok_or_else(|| {
                Error::from(
                    serde_json::to_string(&Response::Panic("wallet got destroyed".into()))
                        .expect("json to string error")
                        .as_str(),
                )
            })
    })?;

    let secret_manager_ptr = &SecretManager { secret_manager };
    std::mem::forget(secret_manager_ptr);

    Ok(secret_manager_ptr)
}

#[no_mangle]
pub unsafe extern "C" fn get_secret_manager_from_wallet(wallet_ptr: *mut Wallet) -> *const SecretManager {
    match internal_get_secret_manager_from_wallet(wallet_ptr) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}