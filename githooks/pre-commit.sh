#!/bin/sh
# This pre-commit hook script runs all tests

make run_tests
exit $?