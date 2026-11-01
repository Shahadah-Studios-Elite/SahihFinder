# SahihFinder

SahihFinder is a web-based platform designed to help Muslims verify the truth of information shared online by highlighting potentially unreliable claims and checking whether a text references trusted Islamic sources.

The project combines a lightweight static front end with a small Go-based verification API to provide a simple, accessible way to assess online content.

## Overview

SahihFinder was created for the purpose of:

- helping users evaluate suspicious or viral messages
- identifying claims that resemble chain-letter or rumor-style content
- checking whether a statement references a recognized hadith collection
- encouraging careful verification before sharing or acting on information

## Key Features

- Clean and responsive landing page
- Dark/light theme toggle
- PWA-style install support via service worker and manifest
- Hadith verification interface
- Wiki-style information page
- Lightweight Go backend for claim assessment
- Islamic-source awareness, including references to classical hadith collections

## Tech Stack

### Frontend
- HTML
- CSS
- JavaScript
- Progressive Web App (PWA) support via manifest.json and sw.js

### Backend
- Go
- HTTP handlers for verification logic

### Deployment / Hosting
- Vercel configuration included
- Static website structure suitable for simple hosting environments

## Project Structure

- index.html — main landing page
- checker.html — verification interface
- Wiki.html — informational content page
- chatbot.html — chatbot-focused experience
- api/check.go — Go handler used to evaluate submitted text
- manifest.json — PWA metadata
- sw.js — service worker
- package.json — frontend package metadata
- go.mod — Go module definition

## Useful Information

- The site is designed as a public-facing verification tool rather than a general-purpose content platform.
- Verification logic focuses on identifying warning patterns, such as chain-letter wording and statements attributed to the Prophet without trusted references.
- The project intentionally encourages cross-checking with authoritative Islamic sources and scholars.
- The repository is a front-end-first project with a very lightweight backend, making it easy to maintain and extend.

## Credits

SahihFinder is developed and maintained by Shahadah Studios Elite.

This project reflects a vision of combining digital tools with Islamic values, helping users think critically and verify information before sharing it.

## Notes

This project is intended to support informed, responsible engagement with online content and should be used alongside sound scholarship, trusted references, and careful personal judgment.

---

SahihFinder — Verify the Truth.
