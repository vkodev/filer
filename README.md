# Filer

# Документация

Для просмотра документации запустите:
```
$ go run docs.go
```

Документация будет доступна по адресу [http://localhost:1088](http://localhost:1088)

# Разработка

Добавьте pre-commit hook для git что-бы не забывать форматировать код перед комитом:

```
cp misc/git/hooks/pre-commit ./.git/hooks/pre-commit && chmod +x ./.git/hooks/pre-commit
```