<div align="center">
  <img src="assets/ohmy-logo.svg" alt="Oh My Sub2API" width="96" />
  <h1>Oh My Sub2API</h1>
  <p>An unofficial Sub2API fork with Codex automation fixes, a resizable sidebar, and an Apple Blue theme.</p>

[中文](README_CN.md) · [Releases](https://github.com/LKRCharon/oh-my-sub2api/releases) · [Changes](CHANGES.md) · [Upstream](https://github.com/Wei-Shaw/sub2api)
</div>

## This release

`0.2.5-ohmy.1` is based on **Sub2API v0.2.5**. This fork maintains its own source, releases, and update channel.

| Change | Behavior |
| --- | --- |
| Codex automation input | Preserves named standalone tool output without a `call_id`, so scheduled task instructions survive request normalization. |
| Sidebar resizing | Drag the desktop sidebar edge between 180 and 400 px. Width persists in the browser. The focused handle also supports arrow keys, Home, and End. |
| Theme | Apple Blue primary color, coordinated chart colors, and light/dark modes. Codex chart series use teal. |
| Own releases | The updater and installation scripts target this fork; numeric `ohmy` revisions are compared correctly. |

Existing configured site names and logos take precedence over the fork defaults. The executable remains `sub2api` for deployment compatibility.

## Run with Docker

Copy the deployment files, choose strong local credentials in `.env`, then start the services:

```sh
git clone https://github.com/LKRCharon/oh-my-sub2api.git
cd oh-my-sub2api/deploy
cp .env.example .env
# Edit .env before starting; do not commit it.
docker compose -f docker-compose.local.yml up -d
```

The image is `ghcr.io/lkrcharon/oh-my-sub2api:0.2.5-ohmy.1` for Linux amd64/arm64. See the [deployment guide](deploy/README.md) for configuration, backups, and binary installation. The inherited compose files use `latest`; pin the version above for a reproducible deployment.

## Build from source

Requires **Go 1.27.0+**, **Node.js 20+**, and **pnpm 9.15.9**. PostgreSQL and Redis are required to run the gateway.

```sh
pnpm --dir frontend install --frozen-lockfile
make build
./backend/bin/server -version
```

`make build` embeds the frontend into the Go binary. Configure your own deployment using `deploy/.env.example` or `deploy/config.example.yaml`. For a complete container build:

```sh
docker build -t oh-my-sub2api:local .
```

Version tags publish platform binaries, multi-architecture GHCR images, SHA-256 checksums, and a corresponding source archive. Build definitions live in [the release workflow](.github/workflows/release.yml), [.goreleaser.yaml](.goreleaser.yaml), and [Dockerfile](Dockerfile).

## Development

See [CONTRIBUTING.md](CONTRIBUTING.md) for tests and upstream synchronization. Keep API keys, SMTP/IMAP credentials, production configurations, and runtime databases outside Git. Examples and regression tests use synthetic values.

## License and attribution

[LGPL-3.0-or-later](LICENSE), following the upstream README declaration. [COPYING](COPYING) includes the GPLv3 terms incorporated by LGPLv3. Preserve copyright, licenses, and notices when redistributing; provide corresponding source for modified binary and container releases.

This is an unofficial fork of [Wei-Shaw/sub2api](https://github.com/Wei-Shaw/sub2api), with no endorsement from the upstream project or API providers. See [NOTICE](NOTICE) for attribution and [the original upstream README](README_UPSTREAM.md) for the retained upstream disclaimers. Its references to sponsors and services describe the upstream project. The upstream README's commercial-use wording differs from the standard license; this fork does not claim an additional commercial authorization from upstream. API providers' terms continue to apply.
