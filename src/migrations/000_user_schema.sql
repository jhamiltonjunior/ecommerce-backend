-- /*
--   Se quiser deletar essa tabela primeiro vai ter
--   que deletar as outras tabelas primeiro
-- */
-- DROP TABLE IF EXISTS list_item_schema;
-- DROP TABLE IF EXISTS list_schema;
-- DROP TABLE IF EXISTS users;

-- CREATE TABLE users(
--   id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
--     --(START WITH 5 INCREMENT BY 5),
--   -- id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,

--   full_name VARCHAR(40) NOT NULL,
--   email VARCHAR(80) UNIQUE NOT NULL,
--   password VARCHAR(800) NOT NULL,
  
--   created_at TIMESTAMP DEFAULT NOW() NOT NULL,
--   updated_at TIMESTAMP,
--   deleted_at TIMESTAMP
-- );

-- INSERT INTO user (username, fullname, email, passwd)
-- VALUES ('Hamilton', 'José Hamilton', 'hamilton@gmail.com', '12345');

-- INSERT INTO user (username, fullname, email, passwd)
-- VALUES ('vibbraneo', 'Vibbraneo da Cunha', 'vibbraneo@vibbra.com.br', '123456');