#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void stream_cpy(FILE *in, FILE *out) {
  unsigned char buff[4096];

  size_t read; 
  while ((read = fread(buff, 1, 4096, in)) != 0)  {
    fprintf(out, "%s", buff);
  }
}

// TODOO: Write help!!

//
// {
//   "replase" : [
//
//   ]
//   
// }
//
//
//


int main(int argc, char **argv) {

  // _insert <temp-file> [--json <json-file>] [--insert <KEY> <arg-file>]

  

  FILE *fd_src = fopen(argv[1], "r");
  if (fd_src == NULL) {
      fprintf(stderr, "ERROR: Failed to open src file `%s`\n", argv[1]);
      exit(1);
  }

  FILE *fd_dat = fopen(argv[2], "r");
  if (fd_dat == NULL) {
      fprintf(stderr, "ERROR: Failed to open dat file `%s`\n", argv[2]);
      exit(1);
  }
  
  char *pch, buff[1024];
  size_t read, len;
  while ((read = fread(buff, 1, 1024, fd_src)) != 0)  {
    pch = strstr(buff,"<CSR>");
    if (pch != NULL) {
        pch += 5; len = pch-buff;
        fprintf(stdout, "%.*s", len, buff);
        stream_cpy(fd_dat, stdout);
        fprintf(stdout, "%.*s", read-len, pch);
    } else {
        fprintf(stdout, "%.*s", read, buff);
    }
  }
  
  fclose(fd_src);
  fclose(fd_dat);
}

