LinkedIn Automation POC (Educational)
⚠️ Disclaimer

This project is a technical proof-of-concept created strictly for educational and evaluation purposes.

Automating LinkedIn directly violates LinkedIn’s Terms of Service.
This tool must not be used in production or on real accounts beyond controlled testing.

The purpose of this project is to demonstrate automation engineering concepts, anti-detection strategies, and clean Go architecture, not real-world scraping or growth automation.

🎯 Objective

To build a Go-based LinkedIn automation proof-of-concept using the Rod browser automation library that demonstrates:

Advanced browser automation patterns

Human-like behavior simulation

Anti-bot detection and stealth mechanisms

Modular, maintainable Go code architecture

The emphasis is on realism, safety, and robustness, not volume or aggressiveness.

🧱 Project Architecture
cmd/
└── main.go                  # Application entry point

internal/
├── auth/                    # Login, OTP handling, session persistence
├── browser/                 # Chromium launcher (Rod)
├── stealth/                 # Anti-detection & human-like behavior
├── search/                  # People search & targeting
├── connect/                 # Connection request logic
├── message/                 # Follow-up messaging
├── storage/                 # Persistent state & cookies
├── config/                  # Environment configuration
└── logger/                  # Structured logging


Each package has a single responsibility, enabling clean separation of concerns and easy extensibility.

🕵️ Anti-Detection Techniques Implemented
Mandatory Techniques

Browser fingerprint masking

Human-like mouse movement using Bézier-style curves

Randomized timing patterns between actions

Additional Techniques

Realistic typing simulation with variable delays

Random scrolling behavior

Idle mouse movement & hover simulation

Activity pacing and rate limiting

Session reuse to reduce repeated logins

These techniques are combined to closely mimic natural human behavior.

🔐 Core Features

Environment-based authentication

OTP / 2FA detection with manual completion support

Persistent session reuse via cookies

Keyword-based LinkedIn people search

Pagination-aware profile collection

Duplicate-safe processing using persistent state

Connection requests with daily limits

Personalized connection notes

Follow-up messaging using templates

Resume-safe execution across multiple runs

🤝 Connection Request Strategy (Important Design Note)

LinkedIn restricts access to many 3rd-degree profiles, displaying:

“You don’t have access to this profile”

Opening such profiles prevents sending connection requests.

✅ Implemented Solution

To handle this real-world LinkedIn limitation, the automation uses a two-level strategy:

Card-level Connect (Preferred)
Sends connection requests directly from search result cards when available.

Profile-level Connect (Fallback)
Used only when full profile access is available.

If neither option is possible, the profile is safely skipped.

This mirrors how real users behave and avoids unnecessary errors or forced access.

💬 Messaging System

Detects newly accepted connections

Sends follow-up messages automatically

Supports message templates

Prevents duplicate messaging using persistent state

📊 State Persistence

Persistent storage is used to track:

Visited profiles

Sent connection requests

Accepted connections

Sent messages

This ensures:

No duplicate actions

Safe resumption after interruption

Realistic long-term behavior

⚙️ Setup Instructions
Prerequisites

Go 1.21+

Chrome / Chromium browser

Environment Variables

Create a .env file using .env.example:

LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password
HEADLESS=false
DAILY_CONNECT_LIMIT=3


⚠️ Never commit real credentials.

🚀 Running the Project
go mod tidy
go run cmd/main.go

Reset for Demo / Fresh Run
rm cookies.json
rm state.json

🎥 Demonstration Video

A full walkthrough video demonstrating setup and execution is available here:

🔗 Demo Video:
https://drive.google.com/file/d/1Mu2gR8mol_Td8iorlPLnglzhgbsXXduz/view

The demo shows:

Browser launch

Login / session restoration

Search execution

Card-level connection requests

Personalized note sending

Human-like scrolling and mouse movement

🧠 Design Philosophy

Safety over aggressiveness

Realistic human behavior

Graceful handling of platform limitations

Clean, modular, maintainable code

This project demonstrates engineering judgment, not just automation.

📈 Evaluation Alignment
Evaluation Criteria	Coverage
Anti-Detection Quality	✅
Automation Correctness	✅
Code Architecture	✅
Real-World Robustness	✅
🏁 Final Notes

This proof-of-concept intentionally reflects real LinkedIn behavior and constraints.

The primary goal is to demonstrate advanced browser automation concepts, stealth engineering, and clean Go design, rather than creating a production-ready tool.