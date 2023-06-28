// Copyright 2023 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::sync::Arc;

use iota_sdk_bindings_core::{
    call_secret_manager_method as rust_call_secret_manager_method,
    iota_sdk::client::secret::{SecretManager as RustSecretManager, SecretManagerDto},
    SecretManagerMethod,
};
use tokio::sync::RwLock;

use std::{
    cell::RefCell,
    error::Error,
    ffi::{CStr, CString, c_char},
    ptr, slice,
};
use std::ptr::null;


use crate::error::{Result, set_last_error};

pub struct SecretManager {
    pub secret_manager: Arc<RwLock<RustSecretManager>>,
}

/// Create secret_manager for python-side usage.
unsafe fn internal_create_secret_manager(options: *const c_char) -> Result<*const SecretManager> {
    let options_string = CStr::from_ptr(options);

    let secret_manager_dto = serde_json::from_str::<SecretManagerDto>(options_string.to_str().unwrap())?;
    let secret_manager = RustSecretManager::try_from(secret_manager_dto)?;

    let secret_manager_wrap = SecretManager {
        secret_manager: Arc::new(RwLock::new(secret_manager)),
    };

    let secret_manager_ptr = Box::into_raw(Box::new(secret_manager_wrap));

    Ok(secret_manager_ptr)
}

#[no_mangle]
pub unsafe extern "C" fn create_secret_manager(options: *const c_char) -> *const SecretManager {
    match internal_create_secret_manager(options) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}


unsafe fn internal_destroy_secret_manager(secret_manager_ptr: *mut SecretManager) -> Result<()> {
    log::debug!("[Rust] Secret Manager destroy called");

    if secret_manager_ptr.is_null() {
        log::error!("[Rust] Secret Manager pointer was null!");
        return Ok(());
    }

    let _ = Box::from_raw(secret_manager_ptr);

    log::debug!("[Rust] Destroyed Secret Manager");
    Ok(())
}

#[no_mangle]
pub unsafe extern "C" fn destroy_secret_manager(secret_manager_ptr: *mut SecretManager) -> bool {
    match internal_destroy_secret_manager(secret_manager_ptr) {
        Ok(v) => true,
        Err(e) => { set_last_error(e); false }
    }
}

unsafe fn internal_call_secret_manager_method(secret_manager: &SecretManager, method: *const c_char) -> Result<*const c_char> {
    let method_string = CStr::from_ptr(method);

    let method = serde_json::from_str::<SecretManagerMethod>(method_string.to_str().unwrap())?;
    let response =
        crate::block_on(async { rust_call_secret_manager_method(&secret_manager.secret_manager, method).await });

    let response_string = serde_json::to_string(&response)?;
    let s = CString::new(response_string).unwrap();

    Ok(s.into_raw())
}

#[no_mangle]
pub unsafe extern "C"  fn call_secret_manager_method(secret_manager: &SecretManager, method: *const c_char) -> *const c_char {
    match internal_call_secret_manager_method(secret_manager, method) {
        Ok(v) => v,
        Err(e) => { set_last_error(e); null() }
    }
}