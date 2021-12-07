#include <stdio.h>
#include <stdlib.h>
#include <string.h>


#define NUM_DAYS 256
#define INTERNAL_COUNT_SPAWN 8
#define INTERNAL_COUNT_BIRTH 6

int totalFish = 0;

void printFish(unsigned long long fish[]) {
    for (int i = 0; i <= INTERNAL_COUNT_SPAWN; i++) {
        printf("%llu ", fish[i]);
    }
    printf("\n");
}

void ageFishOneDay(unsigned long long fishCounts[]) {
    unsigned long long fishHavingChildren = fishCounts[0];
    for (int i = 1; i <= INTERNAL_COUNT_SPAWN; i++) {
        fishCounts[i - 1] = fishCounts[i]; 
    }

    fishCounts[INTERNAL_COUNT_BIRTH] += fishHavingChildren;
    fishCounts[INTERNAL_COUNT_SPAWN] = fishHavingChildren;
}

unsigned long long countFish(unsigned long long fishCounts[]) {
    unsigned long long totalFish = 0;
    for (int i = 0; i <= INTERNAL_COUNT_SPAWN; i++) {
        totalFish += fishCounts[i];
    }

    return totalFish;
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

    unsigned long long fishCounts[INTERNAL_COUNT_SPAWN + 1] = {0};

    int val;
    while (fscanf(f, "%d,", &val) != EOF) {
        fishCounts[val]++;
    }

    for (int i = 0; i < NUM_DAYS; i++) {
        ageFishOneDay(fishCounts);
    }

    printf("Number of fish: %llu\n", countFish(fishCounts));
}