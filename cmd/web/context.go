package main

type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
const triedToAccessContextKey = string("triedToAccess")
