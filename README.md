# ip-upvote

demo: http://ip-upvote-demo.herokuapp.com

I reverse engineered the UrbanDictionary upvote system, in golang. And I don't really care what the original upvoter is built with.
## Why Do I do this type of stuff ?
* *I love to waste my time.*

## How the UrbanDictionary upvote system works ?
* UrbanDictionary is a very popular website where anyone can post definitions/meanings of words/phrases.
* The correct meaning selected by the number of upvotes. [People can flag sick meanings but that is a different thing]
* [An Example](https://www.urbandictionary.com/define.php?term=Simp).
## So What's Wrong ?
* The system allows the users to upvote with their IP addresses.
* This results having a single upvote from your network, and nothing more.

I noticed this when I tried to upvote with an Incongnito [private] window open on my laptop. And after that when I tried to upvote with my phone, [which was on the same network as laptop] but it was already upvoted. Which is may sound like a great idea but is meaningless, **just like my life**.

As a wise man once said,
>Saving critical entities with IP addresses is a bad idea because IP addresses will mostly be dynamic and a new one can be obtained simply by restarting the router or toggling the airplane mode on a mobile device, leading to multiple entities, then to confusion of the highest level to a literal spam as **now you can vote for a word 30 times a minute** (by everytime restarting the network) but a user was only supposed to do once.

This works because of the fact that the IP address of a device changes after restarting. And hence the user is allowed to upvote multiple times which is worrying.

## My Project
This project is built with GO and SQLite which I have never used together before for a demo app. Maybe I used them as I like to suffer. But I'm happy with the functionality.

* My project is the clone of the exact same system but has much worse styles *because I suck at styling*
* You can add items and vote for them once.
* But if you restart your router or toggle airplane mode, you will be assigned a new IP so you can vote for that again, lol.

**Please don't do this to the main UrbanDictionary website, I will soon reach out those guys about this issue**

## Some real threat
I'm just saying, I think someone can use bunch of automated VPN proxies to make literally thousands of requests to the upvote system. And making it undetectable will be like a walk in the park.





