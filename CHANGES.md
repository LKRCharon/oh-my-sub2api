# Fork changes

## 0.2.5-ohmy.1 — 2026-09-15

Based on [Sub2API v0.2.5](https://github.com/Wei-Shaw/sub2api/releases/tag/v0.2.5).

- Preserve named standalone Codex `function_call_output` input without a `call_id`, including scheduled automation and heartbeat input. Keep paired tool-result validation for ordinary continuations.
- Add a desktop sidebar resize handle with a 180–400 px range, saved width, keyboard controls, and cleanup on pointer cancellation, window blur, and component removal. Mobile layout follows the viewport breakpoint.
- Use Apple Blue (`#007AFF`) for the primary theme and a blue, cyan, indigo, and purple chart palette. Codex chart series use teal (`#14B8A6`). Keep light and dark modes.
- Set Oh My Sub2API as the default site name and add an original fork logo. Existing configured site names and logos take precedence.
- Read updates and rollback releases from this fork. Compare numeric fork revisions such as `ohmy.2` and `ohmy.10` correctly.
- Publish binaries, GHCR images, checksums, and corresponding source from version tags. Installation scripts use the fork's release channels.
- Preserve upstream copyright and legal notices, include GPLv3 alongside LGPLv3, and exclude local credentials and runtime data from source and container build contexts.

The backend module path, binary name (`sub2api`), database schema, and service names retain their upstream identifiers for compatibility. The Git diff against `v0.2.5` records every modified file.
