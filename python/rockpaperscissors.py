
from random import choice, randint

num_of_games = int(input("How many games would you like to play?"));
num_of_played = 0;
num_of_player_wins = 0;
num_of_comp_wins = 0;

while num_of_played < num_of_games:
    # randomly assigns values to these variables for Rock, Paper, Scissors
    computer_choice = choice(["Rock", "Paper", "Scissors"]);
    print(computer_choice);
    #get player input
    player_input = input("Rock, Paper, Scissors ");
    print((f"The score is {num_of_player_wins} for the player, and {num_of_comp_wins} for the computer!"));
    if player_input == computer_choice:
            print(f"You tied with {player_input}! Try again");
    elif player_input == "Rock":
        if computer_choice == "Paper":
            print(f"Computer wins with {computer_choice} and you chose {player_input}!")
            num_of_played += 1;
            num_of_comp_wins += 1;
        else:
            print(f"You win with {player_input} and the computer chose {computer_choice}!")
            num_of_played += 1;
            num_of_player_wins += 1;
    elif player_input == "Paper":
        if computer_choice == "Scissors":
            print(f"Computer wins with {computer_choice} and you chose {player_input}!")
            num_of_played += 1;
            num_of_comp_wins += 1;
        else:
            print(f"You win with {player_input} and the computer chose {computer_choice}!")
            num_of_played += 1;
            num_of_player_wins += 1;
    elif player_input == "Scissors":
        if computer_choice == "Rock":
            print(f"Computer wins with {computer_choice} and you chose {player_input}!")
            num_of_played += 1;
            num_of_comp_wins += 1;
        else:
            print(f"You win with {player_input} and the computer chose {computer_choice}!")
            num_of_played += 1;
            num_of_player_wins += 1;
    else:
        print("Please enter valid values.");
print(f"Final score: Player {num_of_player_wins}, Computer {num_of_comp_wins}")
