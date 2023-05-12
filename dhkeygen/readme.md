# Understanding the Diffie Hillman key exchange

**Why learn about the Diffie Hillman Key exchange:**

I am using the topic of cryptography, a topic I am very interested in, to hone my writing Go skills. After writing a Caesar Cipher package, and a Vigenere Cipher package, the Diffie Hillman key exhange seemed like the next logical adventure.

**Easy math:**

The beauty of this key exchange is probably how simple the math is. All that is needed are two protagonists, let's call them Alice and Bob, because every exchange on the internet involes an Alice and a Bob.
Alice and Bob choose a random secret integer, and do not share it with anyone.Alice and Bob both know two other integers, and use these other integers to calculate a public number, a number that they can share with the world.

Both Alice and Bob use the same formula to calculate a public number.

Alice: A = g^a mod p
Bob: B=g^b mod p

Alice sends Bob her public number A, Bob sends to Alice his public number B.

Then both Bob and Alice calculate their secret key:
B^a mod p (Alice)
A^b mod p (Bob)

Here,'a' represents Alice's secret number, 'b', representd Bob's private number, A is Alice's public Number, and B is Bob's public number. g represents the base and mod the modulus ( remember the two integers I mention above).

The strength of the scheme comes from the fact that g^ab mod p = g^ba mod p take a long time to compute by any known algorithm

Once Alice and Bob compute the shared secret they can use it as an encryption key, known only to them, for sending messages across the same open communications channel.
