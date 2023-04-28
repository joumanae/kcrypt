# Understanding the Diffie Hillman key exchange

There are two protagonists, Alice (A) and Bob (B)
They both agree to use a modulus p and a base g
Alice then chooses a secret integer a, and sends Bob A=g^a mod p

Bob chooses a secret integer b and sends Alice B= g^b mod p

Alice computes s = B^a mod p

Bob computes s = A^b mod p

Alice and Bob now share a secret number

The strength of the scheme comes from the fact that g^ab mod p = g^ba mod p take a long time to compute by any known algorithm

Once Alice and Bob compute the shared secret they can use it as an encryption key, known only to them, for sending messages across the same open communications channel
