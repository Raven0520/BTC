# Chain

> This document provides information on the Chain platform's source code, including links to relevant repositories. It also addresses several questions and requirements related to the platform's functionality and security, such as the best way for users to provide seed and whether the generated address is correct and safe to use. The document suggests adding features to improve security, such as modifying generation rules and mapping user mnemonics.
> 

## Source Code

---

gitub: https://github.com/Raven0520/BTC.git

### Main package

Grace grpc: https://github.com/facebookarchive/grace

BTCsuite: 

[Bitcoin in Go](https://github.com/btcsuite)

### Code structure

- app
    - app - Start the application
    - base - Set base configuration
    - consul - Consul call example
    - grpc - Elegant start and stop GRPC
    - node - Set node configuration
    - viper - Parse the configuration file
- bip39 - Bip39 protocol implementation
- bip44 - Bip44 protocol implementation
- config - Configuration file
- core - Generate HD segwit bitcoin address & Generate Multiple signatures
- format -Format the response data
- logic -Logical code
- proto - Protobuf
- router - GRPC service startup & stop
- service - Implement rpc services
- util - Global variables and general functions
- wordlists - Word list used to generate mnemonics

### Questions

- What is the best way(s) to provide the seed onto this server? Please justify the approach(es) with reasons.
    
    The best way for users to provide seed to the platform is to only provide seed of non-main wallets, because in the blockchain network, any behavior of exposing mnemas, seed and private keys is unsafe. It is more reasonable to only use non-main wallets to interact with the platform, and then transfer assets to the main wallet addresses. At the same time, users can also use multi-signature technology to protect assets in non-primary wallets.
    

### Requirements

- Is the generated address correct?
    
    Importing the wallet address into the appropriate blockchain network can verify whether the address is correct. At present, only manual import verification is carried out, and it is directly implemented without unit testing.
    
- Is it safe for users to use?
    
    Not very safe for now. Users should use non-primary wallets.
    
- Is the documentation easy to read?
    
    Because the function is not complicated, the document is also very simple.
    
- Does it follow best practices?
    
    I think so at present.
    
- Add features you can think of on top of the basic requirements?
    
    The requirements document does not further elaborate on the usage scenario after the generated address. Assuming that the generated address is only used in the platform, BTC PRC service can add the generated rules to ensure that the third party cannot generate the same address and private key after the mnie is stolen.
    
    ### Improve security
    
    - Modify the generation rules
    - Map the mnemonics provided by users. This will cause users to use the same mnemonic and path cannot correctly export the wallet address to other platforms.