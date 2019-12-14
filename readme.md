# Tru-Auth
Tru-auth is an openid compliant oAuth2.0 authentication service.

## This is still a WIP

The TruAuth authentication service provides 3rd party service authentication to TruAuth resource agents, implementing the [authorization code grant](https://tools.ietf.org/html/rfc6749#section-1.3.1).

Tru-Auth micro-services are implemented in go, on top of gRPC and utilizes postgreSQL.

--------

## External Sources
- [Resource Service (Demo)](#resource-service)
- [Client Agent (Demo)](#client-agen)
- [Service Site](#service-site)

### Resource Service
__repo__: (https://github.com/truauth/resource-server)

The provided resource service is a demonstration of interaction between a client agent (3rd party service) back-end and the TruAuth authentication service.


### Client Agent
__repo__: (https://github.com/truauth/client)

The provided client agent is a demonstration of interaction between the 3rd party resource owner services & the TruAuth authentication service.  


### Service Site
__repo__: (https://github.com/truauth/service_site)

TruAuth Service registration site for clients & user accounts. 
