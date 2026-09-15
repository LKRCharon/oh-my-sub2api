# Contributing

Report bugs and propose focused changes in this repository. Include the version, a minimal reproduction, and the relevant test results. Replace credentials, private hostnames, account identifiers, and personal data with synthetic values before sharing logs or fixtures.

Use `codex/` branches for maintained fork changes. Keep the upstream remote so a new upstream tag can be reviewed and merged independently of local deployment settings:

```sh
git remote add upstream https://github.com/Wei-Shaw/sub2api.git
git fetch upstream --tags
git switch -c codex/sync-upstream
git merge v<upstream-version>
```

Resolve conflicts, run the relevant checks, update `CHANGES.md` and `backend/cmd/server/VERSION`, then open a pull request here. Keep production configuration outside the repository.

```sh
pnpm --dir frontend install --frozen-lockfile
make test-frontend
make -C backend test-unit
make -C backend test-integration
make build
gitleaks git --redact=100 --log-opts=v0.2.5..HEAD
```

`make build` builds the frontend first and embeds it in `backend/bin/server`. Go, Node.js, and pnpm versions are documented in the README and release workflow. Container integration tests require Docker.

Contributions to this fork use LGPL-3.0-or-later. Preserve existing attribution and add a dated entry to `CHANGES.md`. The upstream `CLA.md` is retained as an upstream document; this fork does not require signing the upstream CLA.
