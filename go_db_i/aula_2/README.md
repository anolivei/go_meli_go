# Aula 2

O arquivo sql foi passado pelo professor para rodarmos o banco na nossa
máquina.

O docker-compose sobe um banco de dados mariadb movies_db na porta 3306.

Para acessar usando o dBeaver é necessário especificar o banco de dados
(movies_db), a porta (3306) e a senha (root).

```SQL
SELECT
	id,
	title,
	rating,
	awards,
	release_date,
	`length`
FROM
	movies
#WHERE
	#awards >= 2 AND rating >= 6
WHERE
	title LIKE '%Guerra%'
ORDER BY
	release_date ASC,
	rating DESC
```