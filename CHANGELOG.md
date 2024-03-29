# Changelog

All notable changes to this package will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2024-02-15
### Breaking Changes
Command Registration API Update: The RegisterCommand function has been renamed to Command to improve readability and simplicity. The previous usage pattern of cli.RegisterCommand(...) should now be updated to cli.Command(...). The parameters and their order remain unchanged.

### Change
Renamed RegisterCommand() to Command() for a more intuitive API.

## [1.0.0] - 2024-02-03
### Feature
- Build command accepts custom output path

## [0.1.2] - 2024-02-03
## Change
- Fix build method and accept custom output path


## [0.1.1] - 2024-02-03
### Change
- Readme file
- Fix build error

## [0.1.0] - 2024-02-03
### Added
- Initial of Callamary CLI.
- Implement `about` command to show application information.
- Implement `build` command for compiling the application.
- Implement `list` command to display available commands.
- Implement `update` command for updating project dependencies.