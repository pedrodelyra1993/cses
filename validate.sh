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

for infile in tests/*.in; do
    if [ -f "$infile" ]; then
        # Extract test case number from filename (e.g., 1.in -> 1)
        i=$(basename "$infile" .in)
        outfile="tests/${i}.out"

        if [ -f "$outfile" ]; then
            echo "Running test case $i..."

            # Run the Go program with input from .in file
            output=$(go run main.go < "$infile")

            # Get expected output
            expected=$(cat "$outfile")

            # Compare outputs (trim trailing whitespace)
            if [ "$(echo "$output" | sed 's/[[:space:]]*$//')" = "$(echo "$expected" | sed 's/[[:space:]]*$//')" ]; then
                echo "Test $i: PASS"
            else
                echo "Test $i: FAIL"
                echo "Expected: $expected"
                echo "Got:      $output"
            fi
        else
            echo "Output file for test case $i missing"
        fi
    fi
done

echo "Validation complete for problem $PROBLEM_ID."
