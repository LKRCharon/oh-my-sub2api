# Reporting security issues

Do not include live API keys, OAuth tokens, passwords, production databases, private keys, or unredacted request logs in issues, pull requests, or screenshots. Use synthetic input to demonstrate a problem.

Use GitHub's private vulnerability reporting on this repository when available. Do not disclose a working credential as a proof of concept. If a credential has already been published, revoke it at its provider.

The release workflow uses the repository's built-in `GITHUB_TOKEN` for GitHub releases and GHCR. It does not require production API credentials. Deployment secrets belong in local environment files or the deployment's secret store.
