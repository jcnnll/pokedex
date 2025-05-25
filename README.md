# pokedex

This is the Pokedex CLI created as a guided project at [boot.dev](https://boot.dev).

Although there are solution files available , I chose to deviate a little:
- Added a helper `doGetJSON` that uses generics to return the appropiate types from
the API
- Integrated caching into the `doGetJSON` to streamline and simplify the process of
caching results from the API
- Applied the principle of defininig an internal API that is meaningful, descriptive
and minimalistic by focussing only on modeling the fields that are required for the client to function


All-in-all, I learned a great deal about creating http clients, caching and generics and how employing
generics can greatly improve code reusability in a clean and safe way.
