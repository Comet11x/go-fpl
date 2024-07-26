#!/bin/sh
# This pre-commit hook script which runs all tests

make run_tests
exit $?