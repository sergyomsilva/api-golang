CREATE TABLE IF NOT EXISTS `products` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `sku`products varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `quantity` int(11) NOT NULL DEFAULT 0
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO personalidades(nome, historia) VALUES
('Deodato Petit Wertheimer', 'Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.'),
('Carmela Dutra', 'Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.');

