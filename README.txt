Golang server for further projects

Uses simple cookie authentication and file-based data storage (will be adapting to a database system down the road).

Gorilla is used for mux and securecookie, but I don't lean heavily on either feature and plan on trying out other solutions.

Uses fcgi for deployment on dreamhost, or ListenAndServe for home testing.

The system is based off of a "UserData" struct from the auth package; this is my basic context solution, and each app must build its own wrappers to their handlers to get this context.  See package auth for more details.

Currently not many apps: roomgame and menagerie are both simple test apps.
