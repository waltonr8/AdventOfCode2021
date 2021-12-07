#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ARR_SIZE 1
#define NUM_DAYS 80

void printFish(int fish[], int numFish) {
    for (int i = 0; i < numFish; i++) {
        printf("%d ", fish[i]);
    }
    printf("\n");
}

int *createNewLanternFish(int fish[], int *numFish) {
    int currentNumFish = *numFish;

    int *temp = fish;
    fish = (int *) realloc(fish, (currentNumFish + 1) * sizeof(int));

    fish[currentNumFish] = 8;
    *numFish = currentNumFish + 1;

    return fish;
}

int *ageFishOneDay(int fish[], int *numFish) {
    int numFishToday = *numFish;
    for (int i = 0; i < numFishToday; i++) {
        if (fish[i] == 0) {
            fish = createNewLanternFish(fish, numFish);
            fish[i] = 6;
        } else {
            fish[i]--;
        }
    }

    return fish;
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

    int size = ARR_SIZE;
    int numFish = 0;
    int *fish = malloc(size * sizeof(int));

    int val;
    while (fscanf(f, "%d,", &val) != EOF) {
        fish[numFish] = val;
        numFish++;

        if (numFish >= size) {
            size++;
            
            int *temp = fish;
            fish = realloc(fish, size * sizeof(int));
            if (fish == NULL) {
                free(temp);
                return 1;
            }
        }
    }

    //printFish(fish, numFish);

    for (int i = 0; i < NUM_DAYS; i++) {
        fish = ageFishOneDay(fish, &numFish);
    }

    printf("Number of fish: %d\n", numFish);
}