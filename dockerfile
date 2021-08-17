FROM go
CMD ["go", "mod", "tidy"]
ENTRYPOINT ["go","run","main.go"]