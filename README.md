# Kaleodoscopic Cryptography

## Description 

This project contains various cryptographic methods. It starts with the ones that are straight forward and the ones that are a little more complex. This project is meant to be full of colors and nuances, thus the reference to a kaleodoscope.

## Substitute Ciphers 

The first cipher I ever wrote was the caesar cipher. It takes a single character key (a letter), and uses it to substitute each character in the ciphered text by another one. 

The 2nd cipher is the vigenere cipher. It is a little more sophisticated. It takes a multi-character key and uses it to cipher or decipher a message. 

## Understanding the Diffie Hillman key exchange

**Why learn about the Diffie Hillman Key exchange:**

I am interested in cryptography, and would like to write descent code in Go. I am using cryptography projects to improve my skills writing Go. After writing a Caesar Cipher package, and a Vigenere Cipher package, the Diffie Hillman key exhange seemed like the next logical adventure.

**Easy math:**

The beauty of this key exchange is probably how simple the math is. All that is needed are two protagonists, let's call them Alice and Bob, because every exchange on the internet involes an Alice and a Bob.
Alice and Bob choose a random secret integer, and do not share it with anyone.Alice and Bob both know two other integers, and use these other integers to calculate a public number, a number that they can share with the world.

First, Both Alice and Bob use the same formula to calculate a public number. 

Alice: A = g^a mod p
Bob: B=g^b mod p
_Here,'a' represents Alice's secret number, 'b', representd Bob's private number, **A** is Alice's public number, and **B** is Bob's public number. **g** represents the base and **mod **the modulus (remember, the two integers known by both Alice & Bob that I mention above)._

The strength of the scheme comes from the fact that _g^ab mod p = g^ba mod p_ take a long time to compute by any known algorithm

Once Alice and Bob compute the shared secret, they can use it as an encryption key, known only to them, for sending messages across the same open communications channel.

Alice sends Bob her public number A, Bob sends to Alice his public number B.

Then both Bob and Alice calculate their secret key:
B^a mod p (Alice)
A^b mod p (Bob)

**Here's a step by step example of what the exchange would look like using this package: **

**Alice:** has the base and the modulus and has a secret key generated for her, let's name it s1 
**Alice types:** dhkeygen -modulus=13 -base=10 
**Alice gets:** 
```
"This is your public key: 9, & this is your secret key 2.", pn1, secretKey 
```
**Bob types:** dhkeygen -modulus=13 -base=10 

**Bob gets:** 
```
"This is your public key: 12, & this is your secret key 3.", pn2, secretKey 
```

**Alice types:** dhkeygen -modulus=13 -publicKey=12 -secret=2 

**Alice gets:**
```
This is your shared key: 1.

```

**Bob types:** dhkeygen -modulus=13 -publickey=9 -secret=3 

**Bob gets:** 

```
This is your shared key: 1. 

``` 

**The beauty of big.Int**
It is always interesting to have to use a package I never used before when writing code. If you take a look at my code you will notice that I am using the math/big package. If you already know about it, you can skip the following paragraph.

When generating a public key, you will need more than 64 bytes. Indeed, it is safer ( yes, it is going overboard, but why not, sometimes). The package big allows you to implement arbitrary-precision arithmetic a.k.a big numbers of types Int, Rat, or Float. In this case, I am only interested in using the Int type.
