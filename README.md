# Blockchain
Blockchain electoral voting system implementation

# How does it works?
The blockchain is implemented in a fairly rudmentary fashion, where proof-of-work has to be shown in order
to add another block onto the chain. 

The intersting voting protococl however revolves around a central agency (The Government) acting as an identity provider service.
The protocol follows as such.

1. A person will register to vote through the central agency, providing identity information.
2. If identity is confirmed by the IDP (Identity Provider), then they will return a JWT token, 
which encapsulates an anonymous string that can be referenced by the providers internal systems 
(this is stored to deal with a case of a voter being forced to vote, and later wanting to revoke this) and will expire after the voting ballot ends.
3. Now a voter can register their vote, directly to the blockchain system, by supplying the JWT token as well as their chosen representative. This is then stored on the blockchain.
4. Counting votes can be done using the IDP's public key, and iterating through the blockchain. The public key is used to verify whether the voter was actually registered with the IDP and was allowed to vote.

