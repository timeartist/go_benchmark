package main

type fixture interface {
    Set(string, string)
    Get(string) string
    Close()
}