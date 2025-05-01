# Snake and Ladder Application

This is a command-line application for the classic Snake and Ladder game. The game takes input for the board configuration and players, simulates the gameplay, and prints the moves and results.

## Features

- Simulates a Snake and Ladder game with customizable snakes, ladders, and players.
- Prints each move in the format:  
  `<player_name> rolled a <dice_value> and moved from <initial_position> to <final_position>`
- Declares the winner when a player reaches position 100:  
  `<player_name> wins the game`

## Rules of the Game

1. The board has 100 cells numbered from 1 to 100.
2. The game uses a six-sided dice numbered from 1 to 6, which generates a random number on each roll.
3. Each player starts with their piece outside the board (at position 0).
4. Players take turns rolling the dice and move their piece forward by the dice value.
   - Example: If the dice value is 5 and the piece is at position 21, the piece moves to position 26 (21 + 5).
5. A player wins if they exactly reach position 100. The game ends immediately when a player wins.
6. If a dice roll causes a piece to move beyond position 100, the piece does not move.
7. The board contains snakes and ladders:
   - **Snakes**: If a piece lands on the head of a snake, it moves down to the tail of the snake.
   - **Ladders**: If a piece lands on the start of a ladder, it moves up to the end of the ladder.
8. If a piece lands on a position with another snake or ladder, it continues to move up or down accordingly.

## Assumptions

- There won’t be a snake at position 100.
- There won’t be multiple snakes or ladders starting at the same position.
- It is always possible to reach position 100 (i.e., the game is winnable).
- Snakes and ladders do not form an infinite loop.

## Input Format

The application takes input from the command line or a file in the following format:

1. **Number of snakes (`s`)** followed by `s` lines, each containing two numbers denoting the head and tail positions of the snake.
2. **Number of ladders (`l`)** followed by `l` lines, each containing two numbers denoting the start and end positions of the ladder.
3. **Number of players (`p`)** followed by `p` lines, each containing a player name.

### Example Input
