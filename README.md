# SahihFinder - Verify the truth

SahihFinder is a web app that helps Muslims investigate Islamic claims and find hadith references. It brings together a hadith search tool, a text checker, curated resource links, and a tasbih counter.

> **Important:** SahihFinder is an aid for further research... Just remember at the moment its not fully released to the world

## Features

- **Hadith Searcher** — search a small curated set of sample narrations and view source links.
- **Text Checker** — flag certain chain-letter phrases and claims that mention the Prophet or a similar authority without naming a recognized collection. It can also identify when a collection name is present; that does not verify the narration.
- **Wiki / resources** — browse links to Islamic reference sites and a Madhab section.
- **Tasbih counter** — count selected dhikr phrases, choose a target, undo, and reset.
- **Responsive pages and theme toggle** — a lightweight HTML, CSS, and JavaScript frontend with light/dark theme controls.
- **Installable web app support** — web app manifest, service worker, and an install prompt where supported.

## Technology

- HTML and CSS
- JavaScript, with a small TypeScript entry file
- Go 1.21 HTTP handlers for checker and search endpoints
- Vercel configuration for clean URLs and serverless API routes

## Project layout

| Path | Purpose |
| --- | --- |
| `index.html`, `index.css` | Home page and its styles |
| `checker.html` | Text-checking interface; sends requests to `/api/check` |
| `searcher.html`, `searcher.js`, `search.css` | Hadith search interface |
| `tasbih.html`, `tasbih.js`, `tasbih.css` | Tasbih counter |
| `Wiki.html` | Resource and Madhab information page |
| `chatbot.html` | Chatbot page |
| `api/check.go` | Text-checking API handler |
| `api/search.go` | Search API handler and built-in sample index |
| `manifest.json`, `sw.js` | PWA metadata and service worker |
| `vercel.json` | Vercel clean-URL configuration |
| `go.mod`, `package.json` | Go module and frontend package metadata |

## API overview

### `POST /api/check`

Request JSON:

```json
{
  "text": "Text to check"
}
```

Returns a JSON object with a `status` and explanatory `message`. The result is a pattern-based prompt for manual checking, not an authenticity verdict.

### `POST /api/search`

Request JSON:

```json
{
  "query": "actions are by intentions"
}
```

Returns a `results` array containing matching entries from the handler's built-in sample index. Results should be checked against the linked primary source.

Both endpoints accept `POST` requests with JSON bodies. The handlers reject other HTTP methods and malformed requests.

## Development and deployment

The repository includes Vercel configuration and Go API handlers intended for deployment as Vercel serverless functions. Deploy the repository to Vercel to use the `/api/check` and `/api/search` endpoints with the static pages.

`package.json` currently defines package metadata only; it does not provide npm scripts or dependencies. The repository does not currently document a separate local development server command.

## Credits

Developed and maintained by Shahadah Studios Elite.

---

**SahihFinder — Verify the Truth.**
