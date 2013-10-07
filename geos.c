#include <geos_c.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

#include "_cgo_export.h"


void notice(const char *fmt, ...) {
  va_list ap;

  fprintf(stdout, "[geos.notice] ");
  va_start(ap,fmt);
  vfprintf(stdout, fmt, ap);
  va_end(ap);
  fprintf(stdout, "\n");
}

void log_and_exit(const char *fmt, ...) {
  va_list ap;

  fprintf(stderr, "[geos.error] ");
  va_start(ap, fmt);
  vfprintf(stderr, fmt, ap);
  va_end(ap);
  fprintf(stderr, "\n");
  exit(1);
}

void initializeGEOS() {
  initGEOS(notice, notice);
}
