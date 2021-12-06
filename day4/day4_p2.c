#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BOARD_SIZE 5
#define ARR_SIZE 1000

typedef struct {
    int value;
    int marked;
} Square;

typedef struct {
    Square board[BOARD_SIZE][BOARD_SIZE];
} Board;

void printBoards(Board *boards, int numBoards) {
    for (int i = 0; i < numBoards; i++) {
        for (int j = 0; j < BOARD_SIZE; j++) {
            for (int k = 0; k < BOARD_SIZE; k++) {
                printf("%d ", boards[i].board[j][k].value);
            }
            printf("\n");
        }
        printf("\n");
    }
}

void callBingoNumber(Board *boards, int numBoards, int numCalled) {
    for (int j = 0; j < numBoards; j++) {
            for (int k = 0; k < BOARD_SIZE; k++) {
                for (int l = 0; l < BOARD_SIZE; l++) {
                    if (boards[j].board[k][l].value == numCalled) {
                        boards[j].board[k][l].marked = 1;
                    }
                }
            }
        }
}

int checkIfWon(Board board) {
    int won = 1;
    // row win
    for (int k = 0; k < BOARD_SIZE; k++) {
        won = 1;
        for (int l = 0; l < BOARD_SIZE; l++) {
            if (!board.board[k][l].marked) {
                won = 0;
            }
        }
        
        if (won) {
            int sum = 0;
            for (int k = 0; k < BOARD_SIZE; k++) {
                for (int l = 0; l < BOARD_SIZE; l++) {
                    if (!board.board[k][l].marked) {
                        sum += board.board[k][l].value;
                    }
                }
            }
            return sum; 
        }
    }

    // col win
    for (int k = 0; k < BOARD_SIZE; k++) {
        won = 1;
        for (int l = 0; l < BOARD_SIZE; l++) {
            if (!board.board[l][k].marked) {
                won = 0;
            }
        }

        if (won) {
            int sum = 0;
            for (int k = 0; k < BOARD_SIZE; k++) {
                for (int l = 0; l < BOARD_SIZE; l++) {
                    if (!board.board[k][l].marked) {
                        sum += board.board[k][l].value;
                    }
                }
            }

            return sum;
        }
    }

    return 0;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Invalid number of command line arguments\n");
        return 1;
    }

    FILE *f = fopen(argv[1], "r");
    if (f == NULL) {
        printf("Invalid file name\n");
        return 1;
    }

    int numsL = ARR_SIZE;
    int numBoards = ARR_SIZE;
    int *nums = malloc(ARR_SIZE * sizeof(int));
    Board *boards = malloc(ARR_SIZE * sizeof(Board));

    int read = 0;
    char buffer[1000];
    while (1) {
        char *retVal = fgets(buffer, 1000, f);
        char *token = strtok(buffer, ",");
   
        while (token != NULL) {
            if (numsL < read + 1) {
                nums = realloc(nums, sizeof(int) * numsL * 2);
                numsL *= 2;
            }

            nums[read] = atoi(token);
            read++;
            token = strtok(NULL, ",");
        }

        if (buffer[strlen(buffer) - 1] == '\n') {
            break;
        }
    }

    int val;
    int count = 0;
    int board = 0;
    int row = 0;
    int col = 0;
    while (fscanf(f, "%d", &val) != EOF) {
        boards[board].board[row][col].value = val;
        boards[board].board[row][col].marked = 0;
        count++;
        
        if (col != 4) {
            col++;
        } else if (row != 4) {
            row++;
            col = 0;
        } else {
            board++;
            row = 0;
            col = 0;
        }
        
    }

    int numBingos = 0;
    int bingos[board];
    for (int i = 0; i < board; i++) {
        bingos[i] = 0;
    }

    int lastOne = 0;
    for (int i = 0; i < read; i++) {
        callBingoNumber(boards, board, nums[i]);

        if (lastOne) {
            int won = checkIfWon(boards[lastOne]);
            if (won) {
                printf("Sum, Num, Total: %d, %d, %d\n", won, nums[i], won * nums[i]);
                free(nums);
                free(boards);
                exit(1);
            }
            continue;
        }

        for (int j = 0; j < board; j++) {
            int won = checkIfWon(boards[j]);
            if (won && !bingos[j]) {
                bingos[j] = won;
                numBingos++;
            } 
        }

        if (numBingos == board - 1) {
            for (int j = 0; j < board; j++) {
                if (bingos[j] == 0) {
                    lastOne = j;
                }
            }
        } 
    }
}