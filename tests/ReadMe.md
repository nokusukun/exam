# Tests

Put student submissions here following a specific format.

## Format

This should be the format of student submissions:

1. The first two lines are file metadata which should contain student's name and activity ID. Later on, this will help in generated automated reports.

2. Inputs should be taken from process arguments.

3. Main function. Can be any name. Final output should be printed on the console.

4. Main function invocation.

```
// Name: Von, Villamor E.
// Activity: 0-BN-ACT-22

const input1 = process.argv[2]
const input2 = process.argv[3]
const input3 = process.argv[4]

function main(arg1, arg2, arg3) {
    console.log(arg1 + arg2 + arg3)
}

main(input1, input2, input3)

```