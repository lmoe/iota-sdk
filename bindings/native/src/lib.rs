// Copyright 2023 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

//! # Python binding implementation for the iota-sdk library.

mod client;
mod error;
mod secret_manager;
mod wallet;

use std::ffi::{CStr, CString, c_char};
use std::ptr::null;
use std::sync::Mutex;

use iota_sdk_bindings_core::{
    call_utils_method as rust_call_utils_method, init_logger as rust_init_logger,
    iota_sdk::client::stronghold::StrongholdAdapter, UtilsMethod,
};
use once_cell::sync::OnceCell;
use tokio::runtime::Runtime;
use crate::error::set_last_error;

use self::{
    client::*,
    error::{Error, Result},
    secret_manager::*,
    wallet::*,
};

/// Use one runtime.
pub(crate) fn block_on<C: futures::Future>(cb: C) -> C::Output {
    static INSTANCE: OnceCell<Mutex<Runtime>> = OnceCell::new();
    let runtime = INSTANCE.get_or_init(|| Mutex::new(Runtime::new().unwrap()));
    runtime.lock().unwrap().block_on(cb)
}

/// Init the Rust logger.
unsafe fn internal_init_logger(config_ptr: *const c_char) -> Result<()> {
    let method_str = CStr::from_ptr(config_ptr).to_str().unwrap();
    rust_init_logger(method_str.to_string()).map_err(|err| Error::from(format!("{:?}", err)))?;
    Ok(())
}

#[no_mangle]
unsafe extern "C" fn init_logger(config_ptr: *const c_char) -> bool {
    match internal_init_logger(config_ptr) {
        Ok(v) => true,
        Err(e) => { set_last_error(e); false }
    }
}

#[no_mangle]
unsafe fn internal_call_utils_method(method_ptr:  *const c_char) -> Result< *const c_char> {
    let method_str = CStr::from_ptr(method_ptr).to_str().unwrap();

    let method = serde_json::from_str::<UtilsMethod>(&method_str)?;
    let response = rust_call_utils_method(method);

    let response_string = serde_json::to_string(&response)?;
    let s = CString::new(response_string).unwrap();

    Ok(s.into_raw())
}

#[no_mangle]
pub unsafe extern "C" fn call_utils_method(config_ptr: *const c_char) -> *const c_char {
    match internal_call_utils_method(config_ptr) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}

/*
/// Migrates a stronghold snapshot from v2 to v3.
#[no_mangle]
pub extern "C" fn migrate_stronghold_snapshot_v2_to_v3(
    current_path: String,
    current_password: String,
    salt: &str,
    rounds: u32,
    new_path: Option<String>,
    new_password: Option<String>,
) -> Result<()> {
    Ok(StrongholdAdapter::migrate_snapshot_v2_to_v3(
        &current_path,
        current_password.into(),
        salt,
        rounds,
        new_path.as_ref(),
        new_password.map(Into::into),
    )
    .map_err(iota_sdk_bindings_core::iota_sdk::client::Error::Stronghold)?)
}*/
