---
description: Get started with the official IOTA Wallet Java library.
image: /img/logo/iota_mark_light.png
keywords:

- Java
- jar
- Maven
- Gradle

---
# IOTA Wallet Java Library

Get started with the official IOTA Wallet Java Library.

## Requirements

Minimum Java version: Java 8

## Use in your Android project (Android Studio)

1. Add following dependency to your `build.gradle` file:
```
implementation 'org.iota:iota-wallet:1.0.0-rc.1'
```

2. Download the `iota-wallet-1.0.0-rc.1-android-jni.zip` file from the GitHub release and unzip it.
3. Add the `jniLibs` folder with its contents to your Android Studio project as shown below:

```
project/
├──libs/
|  └── *.jar <-- if your library has jar files, they go here
├──src/
   └── main/
       ├── AndroidManifest.xml
       ├── java/
       └── jniLibs/ 
           ├── arm64-v8a/           <-- ARM 64bit
           │   └── libiota-wallet.so
           │   └── libc++_shared.so
           ├── armeabi-v7a/         <-- ARM 32bit
           │   └── libiota-wallet.so
           │   └── libc++_shared.so
           └── x86/                 <-- Intel 32bit
              └── libiota-wallet.so
              └── libc++_shared.so
```

## Use in your Java project (Linux, macOS, Windows)

Depending on your operating system, add one of the following dependencies to your `build.gradle` file:

#### linux-x86_64
```
implementation 'org.iota:iota-wallet:1.0.0-rc.1:linux-x86_64'
```

#### windows-x86_64
```
implementation 'org.iota:iota-wallet:1.0.0-rc.1:windows-x86_64'
```

#### aarch64-apple-darwin
```
implementation 'org.iota:iota-wallet:1.0.0-rc.1:aarch64-apple-darwin'
```

#### osx-x86_64
```
implementation 'org.iota:iota-wallet:1.0.0-rc.1:osx-x86_64'
```

## Use the Library

In order to use the library, you need to create a _Wallet_:

```java
// Copyright 2022 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import org.iota.Wallet;
import org.iota.types.AccountHandle;
import org.iota.types.ClientConfig;
import org.iota.types.CoinType;
import org.iota.types.WalletConfig;
import org.iota.types.exceptions.WalletException;
import org.iota.types.secret.StrongholdSecretManager;

public class CreateAccount {
    private static final String DEFAULT_DEVELOPMENT_MNEMONIC = "hidden enroll proud copper decide negative orient asset speed work dolphin atom unhappy game cannon scheme glow kid ring core name still twist actor";

    public static void main(String[] args) throws WalletException {
        // Build the wallet.
        Wallet wallet = new Wallet(new WalletConfig()
                .withClientOptions(new ClientConfig().withNodes("https://api.testnet.shimmer.network"))
                .withSecretManager(new StrongholdSecretManager("PASSWORD_FOR_ENCRYPTION", null, "example-wallet"))
                .withCoinType(CoinType.Shimmer)
        );
        wallet.storeMnemonic(DEFAULT_DEVELOPMENT_MNEMONIC);

        // Create an account.
        AccountHandle a = wallet.createAccount("Alice");

        // Print the account.
        System.out.println(a);
    }
}
```

## What's Next?

Now that you are up and running, you can get acquainted with the library using
its [how-to guides](../how_tos/run_how_tos.mdx) and the
repository's [code examples](https://github.com/iotaledger/wallet.rs/tree/develop/bindings/java/iota-wallet-java/examples/src).