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
