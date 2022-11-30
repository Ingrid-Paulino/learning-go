create schema if not exists personalidades;
use personalidades;

create table personalidades(
                               id int auto_increment primary key,
                               nome varchar(100),
                               historia varchar(200)
);

INSERT INTO personalidades(id, nome, historia) VALUES
                                                   (1, 'Deodato Petit Wertheimer', 'Deodato Petit Wertheimer foi um médico e político '),
                                                   (2, 'Carmela Dutra', 'Carmela Teles Leite Dutra foi a primeira-dama do Brasil....');

