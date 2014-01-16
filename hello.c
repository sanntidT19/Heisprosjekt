// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i = 0;

pthread_mutex_t mutx;

// Note the return type: void*
void* adder(){
    for(int x = 0; x < 1000000; x++){
        pthread_mutex_lock(&mutx);
        i++;
		pthread_mutex_unlock(&mutx);
    }
    return NULL;
}

void* subtracter(){
    for(int x = 0; x < 1000000; x++){
		pthread_mutex_lock(&mutx);
        i--;
		pthread_mutex_unlock(&mutx);
    }
    return NULL;
}


int main(){
	pthread_mutex_init(&mutx, NULL);
    pthread_t adder_thr;
    pthread_create(&adder_thr, NULL, adder, NULL);
    for(int x = 0; x < 50; x++){
        printf("%i\n", i);
    }

    pthread_t subtracter_thr;
    pthread_create(&subtracter_thr, NULL, subtracter, NULL);
    for(int x = 0; x < 50; x++){
        printf("%i\n", i);
    }

    pthread_join(adder_thr, NULL);
    pthread_join(subtracter_thr, NULL); //Waiting for the thread to complete?

    printf("Done: %i\n", i);
    return 0;
    
}
