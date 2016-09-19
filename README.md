GoGame
======

Using Golang as a cross-platform common ground for writing games.

This is intended to be run natively on the server, and transpiled to JavaScript for the browser.

The primary advantages to this:

 - Symmetric server <-> client codebase
 - Golang is fast on both sides
 - Rendering on the client is up to the JavaScript codebase.
 - Renderer agnostic
 - Fast networking + type safety

This is a base library that your game code should import.

Implementation
==============

This implements an Entity-Component model. Each "entity" in it's most basic form has nothing attached to it, and is merely an identity. You can then attach "Components" to the entity which give it functionality.

Some built in components are:

 - Transform (includes Orientation, Scale, Position)
   - Transform can be relative to parent or global.
 - Basic Physics (Velocity, Acceleration)

Some examples you could make:

 - Sprite (rendering information for renderer)
 - Input (controls)

Every entity has an ID (unsigned integer). This ID increments until int max at which point it wraps back around to 0 again. This means you can have a maximum of int_max entities before things will begin to be overwritten.

All of the types in GoGame are in Protobuf for fast serialization and transport.

Networking
==========

In your game you will have a number of entities in the world at any given time. These entities need to be synced over the network correctly, as well as all of their individual components.

As the user (you!) will be creating custom components, we can't just write all of the network sync code in this package.

We can however implement it for the basic built in components.

There are a few types of things we might want to send:

 - Property updates (position, orientation, health)
 - Events (on hit, on shoot, etc).

If your motion is predictable, there's no sense in streaming constant position updates.

Projectile motion is easily predictable:

 - Spawn with a mass, velocity, and time
 - Calculate forward from spawn time to now, continue from now -> onwards.
 - No other updates needed until impact.

So, it would make sense to somehow calculate when motion is easily predictable and does not need streaming. In this way you can avoid the nastiness of position sync and interpolation.

For player movement, we might send:

 - Started moving left @ time T starting at position (X, Y)
 - Stopped moving left @ time T ending at position (X, Y)

Movement includes acceleration and physics against the terrain. The client has the code required to figure out what happened, so we just simulate it client side, using the start and end positions to account for differences in simulation between the client and server.

Implementation of Networking
============================

It is the component's job to implement what is described above. Thus, the networking is purely an RPC mechanism.

The server-side Component can Emit a message to everyone who can see the Entity or to a specific subset of clients.

Interest based Networking
=========================

The client only need know about what he can see. The logic of what the client can see is up to the game to decide. The game's entity should call `AddClientVis` or `RemoveClientVis` when it comes into a client's visibility or leaves it.

Furthermore, we will want to shard the world into individual pieces.

Integration with JavaScript
===========================

This library will expose to JavaScript a public function to create a game client instance.
