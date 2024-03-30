`curl -i -X GET  http://localhost:4000/snippet/view?id=1`

`curl -iL -X POST http://localhost:4000/snippet/create`

```
mysql -D snippetbox -u web -p
SELECT id, title, expires FROM snippets;
```
