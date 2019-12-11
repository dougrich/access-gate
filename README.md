# Access Gate

This will challenge a user to provide an __Access Code__ before allowing them access to a site. This is neither secure, robust, or a true AuthN system - instead, it is lightweight, easy to configure, and user-friendly. It is intended more to block access to short-lived sites rather than as a permenant solution. __Very Important__: see again that this is not secure. Do not make your access code a password to a different system. It will be stored in plain text on the client machine.

## Environment Variables

`ACCESS_CODE` is the actual access code the user needs to enter to gain access.

`PROXY_DEST` is the domain that requests will be proxy'd to after the user has provided the correct access code. This domain is not directly exposed to the user.

`PROXY_HOST` is the friendly host that is exposed to the user.