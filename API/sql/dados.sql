insert into usuarios (nome, nick, email, senha)
values
    ("usuario1", "usuario_1", "usuario1@email.com", "$2a$10$pneQGBxzkJUDvwlNiI0DNO9C..qnGgtqCBb6XZs8RrehMj7kl4SfC"),
    ("usuario2", "usuario_2", "usuario2@email.com", "$2a$10$pneQGBxzkJUDvwlNiI0DNO9C..qnGgtqCBb6XZs8RrehMj7kl4SfC"),
    ("usuario3", "usuario_3", "usuario3@email.com", "$2a$10$pneQGBxzkJUDvwlNiI0DNO9C..qnGgtqCBb6XZs8RrehMj7kl4SfC");


insert into seguidores(usuario_id, seguidor_id)
values
    (1,2),
    (3,1),
    (1,3);


insert into publicacoes(titulo, conteudo, autor_id)
values
    ("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
    ("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
    ("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);