# eva-database
eva-database

    Jogadores (Players):
        ID do Jogador (PlayerID) - Chave primária.
        Nome (Name).
        Contato (Contact) - Email, telefone, etc.
        Data de Cadastro (SignupDate).

    Personagens (Characters):
        ID do Personagem (CharacterID) - Chave primária.
        Nome (Name).
        Raça (Race).
        Classe (Class).
        Nível (Level).
        Atributos (Attributes) - Pode ser uma estrutura JSON para armazenar força, destreza, etc.
        ID do Jogador (PlayerID) - Chave estrangeira referenciando a tabela Jogadores.
        História do Personagem (Backstory) - Opcional, para armazenar a história do personagem.

    Aventuras (Adventures):
        ID da Aventura (AdventureID) - Chave primária.
        Título (Title).
        Descrição (Description).
        Data de Início (StartDate).
        Data de Término (EndDate) - Opcional, se a aventura já terminou.
        ID do Mestre de Jogo (GameMasterID) - Chave estrangeira referenciando a tabela Jogadores.

    Sessões (Sessions):
        ID da Sessão (SessionID) - Chave primária.
        ID da Aventura (AdventureID) - Chave estrangeira referenciando a tabela Aventuras.
        Data da Sessão (SessionDate).
        Notas da Sessão (SessionNotes) - Resumo do que aconteceu na sessão.
        ID do Mestre de Jogo (GameMasterID) - Chave estrangeira referenciando a tabela Jogadores.
        ID dos Participantes (PlayerIDs) - Lista ou referência à tabela de Participantes da Sessão.

    Participantes da Sessão (SessionParticipants):
        ID da Sessão (SessionID) - Chave estrangeira referenciando a tabela Sessões.
        ID do Jogador (PlayerID) - Chave estrangeira referenciando a tabela Jogadores.
        ID do Personagem (CharacterID) - Chave estrangeira referenciando a tabela Personagens.

Relações entre as Tabelas:

    Jogadores podem ter múltiplos Personagens.
    Cada Personagem pertence a um Jogador.
    Aventuras são conduzidas por um Mestre de Jogo (também um Jogador).
    Sessões estão ligadas a uma Aventura e têm vários Participantes (Jogadores e seus Personagens).
    Participantes da Sessão faz a ligação entre Sessões, Jogadores, e Personagens.