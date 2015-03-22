#include "stdio.h"
int main() {
int k = '\\';
int n = '\n';
int q = '\"';
char * s = "#include %cstdio.h%c%cint main() {%cint k = '%c%c';%cint n = '%cn';%cint q = '%c%c';%cchar * s = %c%s%c;%cprintf(s, q,q,n,n,k,k,n,k,n,k,q,n,q,s,q,n,n,n);%creturn 0;%c}";
printf(s, q,q,n,n,k,k,n,k,n,k,q,n,q,s,q,n,n,n);
return 0;
}
