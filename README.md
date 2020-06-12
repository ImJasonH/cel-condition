# `cel` Condition

This Tekton `Condition` can be used in a Pipeline to gate execution of a `Task`
on the successful evaluation and result of a CEL expression, passed as an argument.

## Usage

Install the Condition:

```
KO_DOCKER_REPO=gcr.io/<my-project> ko apply -f condition.yaml
```

Use the Condition in a Pipeline:

```
apiVersion: tekton.dev/v1alpha1
kind: Pipeline
metadata:
  name: say-hello-if-300-is-even
spec:
  tasks:
  - name: say-hello
    conditions:
    - conditionRef: 'cel'
      params:
      - name: expression
        value: '300 % 2 == 0'
    taskSpec:
      steps:
      - image: busybox
        script: echo hello
```

When the Pipeline is run, the Task will only run and echo `hello` if the number
300 is even (which it always is).

## Limitations

* There are no custom methods provided. Some could be added in the future.

* At this time, the expression must be fully specified upfront, and can't
  reference any other Pipeline parameters (i.e., `expression`'s value can't be
  dependent on the resolved value of another parameter). This limits its
  usefulness.

* Tekton will spawn a Pod just to evaluate the expression, when it's perfectly
  safe to evaluate the expression in the control plane.

# CLI example

This can also be run as a simple Go binary:

```
$ go run . '3 == 3'
2020/06/11 20:23:22 evaluating "3 == 3"
2020/06/11 20:23:22 result is true

$ go run . '3 == 4'
2020/06/11 20:23:28 evaluating "3 == 4"
2020/06/11 20:23:28 result is false
exit status 1

$ go run . 'true'
2020/06/11 20:23:35 evaluating "true"
2020/06/11 20:23:35 result is true

$ go run . 'false'
2020/06/11 20:23:38 evaluating "false"
2020/06/11 20:23:38 result is false
exit status 1

$ go run . 'TOTALLY INVALID SYNTAX'
2020/06/11 20:23:49 evaluating "TOTALLY INVALID SYNTAX"
2020/06/11 20:23:49 ERROR: <input>:1:9: Syntax error: mismatched input 'INVALID' expecting <EOF>
 | TOTALLY INVALID SYNTAX
 | ........^
exit status 1
```
