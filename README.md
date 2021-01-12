# ip-upvote

demo: http://ip-upvote-demo.herokuapp.com

I reverse engineered the UrbanDictionary upvote system, written in golang. And I don't really care what the original upvoter is built with.
## Why Do I do this type of stuff ?
* *I love to waste my time.*

## How the UrbanDictionary upvote system works ?
* UrbanDictionary is a very popular website where anyone can post definitions/meanings of words.
* The correct meaning selected by the number of upvotes. [People can flag sick meanings but that is a different thing]
* [An Example](https://www.urbandictionary.com/define.php?term=Simp).
## So What's Wrong ?
* The system allows the users to upvote with their IP addresses.
* This results having a single upvote from your network, and nothing more.

I noticed this when I tried to upvote with an Incongnito [private] window open on my laptop. And after that when I tried to upvote with my phone, [which was on the same network as laptop] but it was already upvoted. Which is may sound like a great idea but is pointless, **just like my life**.

As, a wise man said,
>Saving critical entities with IP addresses is a bad idea because IP addresses will mostly be dynamic and a new one can be obtained simply by restarting the router or toggling the airplane mode on a mobile device, leading to multiple entities, then to confusion of the highest level to a literal spam as **now you can vote for a word 30 times a minute** (by everytime restarting the network) but a user was only supposed to do once.

This works because of the fact that the IP address of a device changes after restarting.



