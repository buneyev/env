# env

Utilities for working with environment variables:

 - FillStructFromEnv - Fills struct fields from environment variables using "env" tags. Only works with strings, bool and int.
 - FillStructFromEnvFatal - Same as FillStructFromEnv, but using log.Fatal when an error occurs.