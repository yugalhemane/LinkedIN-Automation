# LinkedIn Automation POC (Educational)

## ⚠️ Disclaimer
This project is a technical proof-of-concept created strictly for **educational and evaluation purposes**.

Automating LinkedIn violates their Terms of Service.  
This tool **must not be used in production** or on real accounts beyond controlled testing.

---

## 🎯 Objective
Demonstrate advanced browser automation, anti-detection techniques, and clean Go architecture using the **Rod** library.

The focus is on:
- Human-like behavior simulation
- Stealth mechanisms
- Maintainable, modular Go code

---

## 🧱 Project Architecture

cmd/
└── main.go # Application entry point

internal/
├── auth/ # Login, OTP handling, session persistence
├── browser/ # Chromium launcher
├── stealth/ # Anti-detection techniques
├── search/ # People search & targeting
├── connect/ # Connection request logic
├── message/ # Follow-up messaging
├── storage/ # Persistent state & cookies
├── config/ # Environment configuration
└── logger/ # Structured logging


---

## 🕵️ Anti-Detection Techniques Implemented

Mandatory:
- Browser fingerprint masking
- Human-like mouse movement (Bezier curves)
- Randomized timing patterns

Additional:
- Realistic typing simulation
- Random scrolling behavior
- Idle mouse & hover movement
- Activity pacing & rate limiting
- Session reuse to reduce suspicious logins

---

## 🔐 Core Features

- Environment-based authentication
- OTP / 2FA detection with manual completion
- Persistent session reuse via cookies
- Keyword-based LinkedIn profile search
- Duplicate-safe profile collection
- Connection requests with daily limits
- Follow-up messaging with templates
- Resume-safe execution using persistent state

---

## ⚙️ Setup Instructions

### Prerequisites
- Go 1.21+
- Chrome / Chromium browser

### Environment Variables

Create a `.env` file (see `.env.example`):

```env
LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password
HEADLESS=false
DAILY_CONNECT_LIMIT=3

Run the Project
go run cmd/main.go

```
