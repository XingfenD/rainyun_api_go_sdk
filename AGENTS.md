# AGENTS.md

## Safety Rules
- **Branch safety**: If the current branch is NOT a dev branch (name does NOT contain `feat`, `refactor`, `docs`, `fix`), do NOT commit directly — warn the user first.
- **Commit SOP**: Before every `git commit`, run `git branch --show-current`. If the branch is not a dev branch, ask the user for a branch name, create it, and commit there.
- **Branch naming**: Dev branches must contain `feat`, `refactor`, `docs`, or `fix`.
- **CHANGELOG**: Core changes → `docs/CHANGELOG.md` only. Higher versions on TOP, lower versions below.
