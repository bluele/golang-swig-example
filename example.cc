
#include "example.h"

Example::Example() {

}

Example::~Example() {

}

void Example::setVal(int x) {
  val = x;
}

int Example::getVal() {
  return val;
}

int Example::add(int x, int y) {
  return x + y;
}