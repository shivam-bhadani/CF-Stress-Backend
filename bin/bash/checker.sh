#!/bin/bash

g++ $1 -o $7
g++ $2 -o $8
g++ $3 -o $9

for((i = 1; ; ++i)); do
    $9 $i > $4 
    $7 <$4 > $5 
    $8 <$4 > $6 
    diff -w $5 $6 || break 
    echo "Passed test: "  $i
    # echo $?
done

# echo -e "\nWA on the following test:"
# cat "E:\Stress Test\stress\randomInput.txt"
# echo "Your answer is:"
# cat "E:\Stress Test\stress\myOutput.txt"
# echo "Correct answer is:"
# cat "E:\Stress Test\stress\correctOutput.txt"

# $1 = participant_solution
# $2 = jury_solution
# $3 = generator
# $4 = input_file
# $5 = participant_output_file
# $6 = jury_output_file
# $7 = participant_solution withoud cpp extension
# $8 = jury_solution without cpp extension
# $9 = generator withoud cpp extension