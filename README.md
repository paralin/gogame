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

Every component has an ID (unsigned integer). The first few IDs are reserved for the built in components (transform, for example).

Creating Game Entities
======================

An entity is anything in your game. This could be:

 - a type of weapon
   - This would be composed of:
     - transform component, in parent-relative mode
     - "weapon" component, presumably with some kind of weapon system
     - parent (player entity)
 - a player

As an entity needs to be set up by someone, there is a mechanism in
place to build entities. This is called the EntityFactory interface.

An EntityFactory would be the generic prototype that knows how to
construct an instance of your entity. You would have one for each type
of entity. This could be, for example:

 - one for each type of weapon in your game
 - one for the generic "player"
 - one for each NPC in your game, or interactable object

An entity factory knows how to create an entity, by creating a New()
entity, and then building a tree of children entities using other
EntityPrototypes, or adding Components to the entity it has just
created.

Game Rules
=========

In gogame, a "Game" is an instance of the entire game logic tree. You
need a object to "tick" and decide what to do each frame.

There are also other events you need to handle, examples include:

 - Player connected
 - Player disconnected

You implement this logic in your own Game Rules type. This is passed to
GoGame when you construct the Game on the client or server.

A game doesn't have to be networked, you could make a single player game in the browser or just a sim in the server. Thus, the Game Rules and overall game has a concept of "operating mode".

Here are the operating modes:

 - Local - game is operating in local mode only
 - Remote - game is connected to a server and following sync

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


Renderer / Frontend
===================

The `gogame` system has no concept of actually displaying the game to the user, or taking input from the user. This must be implemented by something external. In Terram's case, this is done by TypeScript.

GoGame has a generic "Frontend" interface. When creating a game, a struct implementing the Frontend interface can be provided. This interface will be called to sync the internal game state with the frontend. Types of functions the frontend will have to implement might include:

 - Entity added, can return a FrontendEntity object which takes callbacks for entity events.
   - When an entity is added, the frontend entity object receives:
     - Init()
     - AddComponent() for each component, can return a FrontendComponent
     - InitLate()
     - And later: Destroy()
   - Frontend component receives similarly:
     - Init()
     - Destroy()
   - Frontend components can receive function calls from the Go component code.
     - Examples: set position, etc.

Main Update Tick
================

The physics engine and game logic in general needs to tick at a constant rate. Furthermore, we don't want to waste time iterating over every single component, if some of them don't need an Update() tick call.

GoGame uses a `time.Ticker` from Go to tick a main Update function. This update function calls in order:

 - GameRules.Update
 - Update on each entity with at least one update handler
  - This calls Update() on each component with an update handler.

This way, we only call Update() if it's going to do something with it. Also, in the frontend, we check if the Update() function exists in the beginning, and don't do the nil check again after. This is to save time.

Entity Lifecycle
================

An entity is created by an EntityFactory.

 - Entity is created in the factory with `&MyEntity{}`
 - Each component is added with `AddComponent(Component)`
 - Entity is returned from the factory.
 - Caller of factory calls `ent.InitComponents()`
   - `component.Init()` is called for each component
 - Caller of factory calls `g.Frontend.AddEntity`, sets frontend entity if any is returned.
 - `g.AddEntity` is called
   - `ent.InitFrontendEntity()` is called.
   - `ent.LateInitComponents` is called.
   - The value of `ent.HasUpdateTick` is checked.

When spawning one over network (remote entity):

 - `EntityFromNetInit()`: creates with `&Entity{}`
  - calls `comp.InitWithData` on each component
 - `g.AddEntity` is called
   - `ent.InitFrontendEntity()` is called.
   - `ent.LateInitComponents` is called.
   - The value of `ent.HasUpdateTick` is checked.
