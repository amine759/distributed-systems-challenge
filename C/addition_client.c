/*
 * This is sample code generated by rpcgen.
 * These are only templates and you can use them
 * as a guideline for developing your own functions.
 */

#include "addition.h"


void
add_prog_1(char *host)
{
	CLIENT *clnt;
	struct ADDITION  *result_1;
	struct ADDITION  add_proc_1_arg;

#ifndef	DEBUG
	clnt = clnt_create (host, ADD_PROG, ADD_VER, "udp");
	if (clnt == NULL) {
		clnt_pcreateerror (host);
		exit (1);
	}
#endif	/* DEBUG */
	int a=10; 
	int b=9;

	add_proc_1_arg.arg1=a;
	add_proc_1_arg.arg2=b;

	result_1 = add_proc_1(&add_proc_1_arg, clnt);
	if (result_1 == (struct ADDITION *) NULL) {
		clnt_perror (clnt, "call failed");
	}
#ifndef	DEBUG
	clnt_destroy (clnt);
#endif	 /* DEBUG */
	printf("The Result is %.3f \n", result_1->result);
}


int
main (int argc, char *argv[])
{	
	char *host;

	if (argc < 2) {
		printf ("usage: %s server_host\n", argv[0]);
		exit (1);
	}
	host = argv[1];
	add_prog_1 (host);
exit (0);
}
