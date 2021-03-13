# YAUS (yet another url shortener)

## Why YAUS?
Because the other url shortener services seemed over complicated.

## Concept
The concept of YAUS is simple. The `config.yaml` drives how the serivce works.
In the config file there will be port configuration, storage, and all of that.
Then another section with hardcoded redirection values.

## Short Endpoints
- `/i/` will contain the hardcoded *information* links.
- `/e/` will contain the hardcoded *error* links.
- `/d/` will be where the dynamic short links are.
