#include <stdio.h>

int main() {
    float applePrice = 1.49;
    int appleQuantity = 23;
    const char appleLocation = 'F';
    double appleReview = 82.5;
    int appleReviewDisplay = (int) appleReview;



    printf("an apple costs: $%.2f, there are %d in inventory found in section: %c and your customers gave it an average review of %d%%!\n", applePrice, appleQuantity, appleLocation, appleReviewDisplay);
}
