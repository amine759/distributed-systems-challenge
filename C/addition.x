/* 
 * filename addition.x 
   * function: Define constants, non-standard data types and the calling process in remote calls
 */

const ADD = 0;

struct ADDITION
{
    float arg1;
    float arg2;
    float result;
};

program ADD_PROG
{
    version ADD_VER
    {
        struct ADDITION ADD_PROC(struct ADDITION) = 1;
         } = 1; /* Version number=1 */
 } = 0x20000001; /* RPC program number */
