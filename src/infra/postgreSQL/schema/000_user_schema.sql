/*
  Se quiser deletar essa tabela primeiro vai ter
  que deletar as outras tabelas primeiro
*/
DROP TABLE IF EXISTS list_item_schema;
DROP TABLE IF EXISTS list_schema;
DROP TABLE IF EXISTS user_schemas;

CREATE TABLE users(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

  login VARCHAR(30) UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password VARCHAR(20) NOT NULL,
  full_name VARCHAR(30) NOT NULL,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP
);

INSERT INTO user (username, fullname, email, passwd)
VALUES ('Hamilton', 'José Hamilton', 'hamilton@gmail.com', '12345');

INSERT INTO user (username, fullname, email, passwd)
VALUES ('vibbraneo', 'Vibbraneo da Cunha', 'vibbraneo@vibbra.com.br', '123456');