# CLI

- `scorekeep`: Gets the scores for all supported leagues
    - Flags:
        - `-l, --league={nhl, nba, nfl}`
        - `-d, --date=YYYY-mm-dd` (today if not specified)
- `scorekeep add favorites montreal-canadiens`: Adds team to favorites
- `scorekeep get favorites`: Get favorites by league
    - Flags: `-l, --league={nhl, nba, nfl}`
- `scorekeep get teams`: Get all teams by league
    - Flags: `-l, --league{nhl, nba, nfl}`
- `scorekeep get leagues`: Get supported leagues