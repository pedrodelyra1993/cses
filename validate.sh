#!/bin/bash

# Generic script to validate Go solution against CSES test cases for a given problem

if [ -z "$1" ]; then
    echo "Usage: $0 <problem_id>"
    echo "Example: $0 1068"
    exit 1
fi

PROBLEM_ID="$1"

if [ ! -d "$PROBLEM_ID" ]; then
    echo "Directory $PROBLEM_ID not found"
    exit 1
fi

cd "$PROBLEM_ID"

echo "Running validation for problem $PROBLEM_ID"

for i in {1..14}; do
    if [ -f "tests/${i}.in" ] && [ -f "tests/${i}.out" ]; then
        echo "Running test case $i..."

        # Run the Go program with input from .in file
        output=$(go run main.go < "tests/${i}.in")

        # Get expected output
        expected=$(cat "tests/${i}.out")

        # Compare outputs (trim trailing whitespace)
        if [ "$(echo "$output" | sed 's/[[:space:]]*$//')" = "$(echo "$expected" | sed 's/[[:space:]]*$//')" ]; then
            echo "Test $i: PASS"
        else
            echo "Test $i: FAIL"
            echo "Expected: $expected"
            echo "Got:      $output"
        fi
    else
        echo "Test case $i files missing"
    fi
done

echo "Validation complete for problem $PROBLEM_ID."
