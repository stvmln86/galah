# Galah

Galah is a minimalist terminal-based maze runner, written in Go 1.25 by Stephen Malone. On launch, the program generates a large 250x250 tile maze with a randomly placed origin and destination. After a prompt, the user is placed at the origin and permitted to explore and solve the maze.

Scattered around the maze are points the player can collect by walking over, which can then be spent on cheats that help the player solve the maze. Cheats include marking the player's current position, destroying walls around the player in a circle and "breadcrumbs" that guide the player to the destination.

The game's interface are simple: use WASD or arrow keys to move, `B` to enter the buy screen and `Q` to quit. An optional line-of-sight feature allows users to only see their immediate surroundings (instead of the whole maze), which multiples any points earned by ten.

Once the destination is reached, the player is shown a victory screen showing the amount of steps taken and points collected. The user is then prompted to restart with a new maze or quit the game.
