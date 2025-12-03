# Chess Program

A console-based chess game written in Go.

---

## 📋 Table of Contents

- [Features](#-features)
- [Quick Start](#-quick-start)
- [Project Structure](#-project-structure)
- [Game Rules](#-game-rules)
- [Usage Examples](#-usage-examples)
- [Testing](#-testing)
- [Design Decisions](#-design-decisions)

---

## Features

- Standard 8x8 chessboard with proper piece placement
- Move validation for all piece types (pawn, rook, bishop, queen, knight, king)
- Turn-based gameplay with alternating white and black moves
- Path obstruction checking for sliding pieces
- Input parsing using algebraic notation (e.g. `e2 e4`)
- Game ends when king is captured

---

## 🚀 Quick Start

### Prerequisites

```bash
go version # Requires Go 1.20+
```

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd chessProgram

# Install dependencies
go mod download
```

### Running the Game

```bash
# Start the chess game
go run ./cmd/app

# Or build and run
go build -o chess ./cmd/app
./chess
```

### Playing the Game

```plaintext
   a  b  c  d  e  f  g  h
8  r  n  b  q  k  b  n  r
7  p  p  p  p  p  p  p  p
6  .  .  .  .  .  .  .  .
5  .  .  .  .  .  .  .  .
4  .  .  .  .  .  .  .  .
3  .  .  .  .  .  .  .  .
2  P  P  P  P  P  P  P  P
1  R  N  B  Q  K  B  N  R

White's turn
Enter move (e.g., 'e2 e4'): e2 e4
✓ Move successful!
```

---

## 📁 Project Structure

```plaintext
chessProgram/
│
├── cmd/
│   └── app/
│       └── main.go           # Entry point
│
├── board.go                  # Board representation & initialization
├── piece.go                  # Piece definitions & rules
├── move.go                   # Move validation logic
├── input.go                  # Input parsing & validation
├── game.go                   # Game state & loop management
├── utils.go                  # Helper utilities
│
├── tests/
│   ├── board_test.go         # Board initialization tests
│   ├── move_test.go          # Move validation tests
│   └── game_test.go          # Game flow tests
│
├── go.mod
├── go.sum
└── README.md
```

---

## 🎮 Game Rules

### Movement Rules

| Piece | Movement | Special Rules |
|-------|----------|---------------|
| ♟️ Pawn | 1-2 squares forward (first move), 1 square forward after | Diagonal capture only |
| ♜ Rook | Any number of squares horizontally/vertically | Cannot jump over pieces |
| ♝ Bishop | Any number of squares diagonally | Cannot jump over pieces |
| ♛ Queen | Rook + Bishop combined | Cannot jump over pieces |
| ♞ Knight | L-shape (2+1 or 1+2 squares) | Can jump over pieces |
| ♚ King | 1 square in any direction | Game ends if captured |

### Win Condition

The game ends immediately when either king is captured.

### Intentional Simplifications

This implementation focuses on core chess mechanics. **Not implemented**:
- ❌ Castling
- ❌ En passant
- ❌ Pawn promotion
- ❌ Check/checkmate detection
- ❌ Stalemate rules

These omissions keep the codebase focused and maintainable while demonstrating solid engineering principles.

---

## 📝 Usage Examples

### Opening Moves

```bash
# King's Pawn Opening
e2 e4

# Knight Development
b1 c3

# Queen's Gambit
d2 d4
c2 c4
```

### Sample Game Sequence

```plaintext
1. e2 e4    # White pawn to e4
2. e7 e5    # Black pawn to e5
3. g1 f3    # White knight to f3
4. b8 c6    # Black knight to c6
5. f1 c4    # White bishop to c4
```

---

## 🧪 Testing

### Run All Tests

```bash
# Basic test run
go test ./...

# Verbose output
go test -v ./...

# With coverage report
go test -cover ./...

# Detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Coverage

The test suite covers:
- ✅ Board initialization and display
- ✅ All piece movement rules
- ✅ Path obstruction detection
- ✅ Boundary validation
- ✅ Input parsing edge cases
- ✅ Game state transitions
- ✅ Win condition detection

---

## 🏗️ Design Decisions

### Architecture Principles

1. **Separation of Concerns**
   - Board logic separated from game logic
   - Move validation independent of input parsing
   - Display logic isolated for easy modification

2. **Testability**
   - Pure functions for move validation
   - Minimal state mutation
   - Clear interfaces between modules

3. **Simplicity Over Completeness**
   - Focus on core chess mechanics
   - Omit complex rules to maintain clarity
   - Prioritize code readability

### Key Design Choices

- **Coordinate System**: Row/column integers (0-7) internally, algebraic notation for I/O
- **Board Representation**: 2D array for simplicity and direct indexing
- **Piece Identification**: Character-based (uppercase/lowercase) for easy visualization
- **Turn Management**: Simple boolean toggle between white/black

---

## 🛠️ Future Enhancements

Potential improvements for extended versions:

- [ ] Implement check/checkmate detection
- [ ] Add castling support
- [ ] Implement en passant capture
- [ ] Add pawn promotion
- [ ] Move history and undo functionality
- [ ] Save/load game state
- [ ] AI opponent (minimax algorithm)
- [ ] Web-based UI
- [ ] Multiplayer support

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

Created as part of a backend engineering take-home assessment.

For questions or feedback, please open an issue in the repository.

---

## 🙏 Acknowledgments

- Chess rules based on FIDE regulations (simplified for scope)
- Inspired by classic console chess implementations
- Built with Go's standard library for minimal dependencies
