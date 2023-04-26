# SQLite

### Criando um novo banco de dados

Para criar um novo banco de dados SQLite, utilize o seguinte comando:

```sql
sqlite3 nome_do_banco.db
```

### Inserindo e selecionando dados
Para inserir dados em uma tabela, utilize o seguinte comando:

```sql
INSERT INTO nome_da_tabela (coluna1, coluna2, coluna3) VALUES (valor1, valor2, valor3);
```
Para selecionar dados de uma tabela, utilize o seguinte comando:

```sql
SELECT coluna1, coluna2, coluna3 FROM nome_da_tabela;
```

### Atualizando e deletando dados
Para atualizar dados em uma tabela, utilize o seguinte comando:

```sql
UPDATE nome_da_tabela SET coluna1 = valor1 WHERE coluna2 = valor2;
```
Para deletar dados de uma tabela, utilize o seguinte comando:

```sql
DELETE FROM nome_da_tabela WHERE coluna1 = valor1;
```

### Saindo do SQLite
Para sair do SQLite, utilize o seguinte comando:

```sql
.quit
```

### Selecionando colunas de uma tabela:
Suponha que temos uma tabela chamada clientes com as colunas id, nome, idade, email, cidade e estado. Queremos selecionar apenas as colunas nome e cidade da tabela:

```sql
SELECT nome, cidade FROM clientes;
```

### Selecionando linhas de uma tabela:
Suponha que queremos selecionar apenas os clientes com idade maior ou igual a 30 anos:

```sql
SELECT * FROM clientes WHERE idade >= 30;
```

### Ordenando os resultados:
Suponha que queremos selecionar os clientes da tabela clientes ordenados por idade de forma decrescente:

```sql
SELECT * FROM clientes ORDER BY idade DESC;
```

### Fazendo junções de tabelas:
Suponha que temos duas tabelas: clientes e pedidos, onde a tabela pedidos tem as colunas id, id_cliente, produto, quantidade e preco. Queremos selecionar os nomes dos clientes e os produtos que eles compraram:

```sql
SELECT clientes.nome, pedidos.produto FROM clientes JOIN pedidos ON clientes.id = pedidos.id_cliente;
```
