# caller
A simple dependency free portable command line tool for getting your current AWS identity.

You can wget or curl the bin for your platform and execute it without additional runtime dependencies such as Python, Node or a JVM.

Wraps the `aws sts get-caller-identity` api call and returns the arn, or exits 1 if there are no valid credentials, or the api call fails.