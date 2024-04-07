#!/bin/bash

# Count the total number of tasks and the number completed
total_tasks=$(grep -c '\[.\]' README.md)
completed_tasks=$(grep -c '\[x\]' README.md)

# Calculate the tasks left and the percentage completed
tasks_left=$((total_tasks - completed_tasks))
percent_complete=$(echo "scale=2; $completed_tasks / $total_tasks * 100" | bc)

# Output the results
echo "Completed: $completed_tasks"
echo "Remaining: $tasks_left"
echo "Percent Complete: $percent_complete%"


