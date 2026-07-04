# 🍿 SnackYou Turret API Documentation

## Overview

This API controls a networked snack-launching turret with:

* 🎯 Real-time yaw/pitch control
* 🔒 Lock-based ownership system
* 🔥 Fire trigger system
* 🤖 ESP32 polling interface

All state is stored in-memory on the Go (Gin) backend.

Base URL:

```
http://<server>:8080/api
```

---

# 🧠 Core Concepts

## 1. Lock System

Only one user can control the turret at a time.

* You must acquire a lock before moving the turret.
* Lock expires automatically after 30 seconds unless refreshed.

---

## 2. State System

Turret state is:

```json
{
  "yaw": 0-180,
  "pitch": 0-90,
  "fire": false,
  "version": int
}
```

---

## 3. User Identity

Every request that modifies state must include:

```json
"user": "string"
```

This is used to validate lock ownership.

---

# 📡 API ENDPOINTS

---

# 🔐 LOCK SYSTEM

## POST /api/lock

Acquire or extend control of the turret.

### Request Body

```json
{
  "user": "ritam"
}
```

### Behavior

* If no lock exists → grants lock
* If same user → extends lock by 30s
* If different user owns lock → returns 409 Conflict

### Response (success)

```json
{
  "success": true,
  "expires": "2026-07-04T02:40:00Z"
}
```

### Response (conflict)

```json
{
  "success": false,
  "owner": "other_user"
}
```

---

## DELETE /api/lock

Release lock (only owner can release).

### Query Param OR Body

Preferred:

```json
{
  "user": "ritam"
}
```

### Behavior

* Clears lock if requester is owner
* Resets ownership and expiry

### Response

```json
{
  "released": true
}
```

---

## GET /api/lock

Get current lock state.

### Response

```json
{
  "owner": "ritam",
  "expires": "2026-07-04T02:40:00Z"
}
```

---

# 🎯 TURTLE STATE CONTROL

---

## GET /api/state

Returns full turret state + lock info.

### Response

```json
{
  "turret": {
    "yaw": 90,
    "pitch": 20,
    "fire": false,
    "version": 12
  },
  "lock": {
    "owner": "ritam",
    "expires": "2026-07-04T02:40:00Z"
  }
}
```

---

## PUT /api/state

Update turret orientation.

### Request Body

```json
{
  "user": "ritam",
  "yaw": 120,
  "pitch": 30
}
```

### Rules

* User MUST own lock
* Otherwise returns 403 Forbidden

### Response

```json
{
  "yaw": 120,
  "pitch": 30,
  "fire": false,
  "version": 13
}
```

---

# 🔥 FIRE SYSTEM

---

## POST /api/fire

Triggers snack launch event.

### Request Body

```json
{
  "user": "ritam"
}
```

### Rules

* Must own lock
* Sets `fire = true`
* ESP resets fire after ACK

### Response

```json
{
  "ok": true,
  "version": 14
}
```

---

# 🤖 ESP32 ENDPOINTS

---

## GET /api/esp/state

Used by ESP32 to poll turret commands.

### Response

```json
{
  "yaw": 120,
  "pitch": 30,
  "fire": false,
  "version": 14
}
```

### Purpose

* ESP reads desired state
* Executes movement / firing
* Compares version for updates

---

## POST /api/esp/ack

Acknowledges fire execution.

### Request Body

```json
{
  "version": 14
}
```

### Behavior

* If version matches → resets `fire = false`

### Response

```json
{
  "ok": true
}
```

---

# ⚠️ SYSTEM BEHAVIOR RULES

## Lock enforcement

* Only lock owner can move or fire turret

## Version system

* Increments on every state change or fire
* Prevents duplicate execution

## Fire handling

* Fire is a transient flag
* Must be ACKed by ESP

---

# 🔄 REQUEST FLOW EXAMPLE

## Typical session

1. User requests lock

```
POST /api/lock
```

2. User moves turret

```
PUT /api/state
```

3. User fires

```
POST /api/fire
```

4. ESP executes and ACKs

```
POST /api/esp/ack
```

5. User releases lock

```
DELETE /api/lock
```

---

# 🧠 SYSTEM SUMMARY

This system behaves like:

> A multiplayer real-time controller for a physical device with ownership-based access control.

---

If you want next, I can also generate:

* 📄 README.md version (GitHub ready)
* 🧩 architecture diagram (for judges)
* ⚡ ESP32 firmware pseudocode (so your teammate can implement fast)
* 🎮 UI state machine diagram (super helpful for debugging)

You're at the “this is actually a distributed system” stage now.
