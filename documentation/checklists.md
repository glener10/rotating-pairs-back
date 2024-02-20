# **Checklists**

🏁 Table of Contents

===================

<!--ts-->

- 👉 [Inicializar Projeto](#inicializar)

- 👉 [Manter projeto](#manter)

  - ▶️ [Manter projeto - Back](#manter-back)

- 👉 [Segurança](#seguranca)

- 👉 [Geral Back-End](#backend)

<!--te-->

===================

<div id="inicializar"></div>

## **Inicializar Projeto**

[ ] Cria um projeto o mais simples possível, e mesmo assim tire/limpe o que vem de "extra". Pode-se usar uma base/esqueleto pronta para evitar trabalho

[ ] Adicione na documentação o _changelog.md_ e _useCases.md_

[ ] Crie um arquivo de _LICENÇA_

[ ] Inicialize o arquivo de documentação geral _README.md_

[ ] Inicialize e use um verificador de código (ex: _eslint_), necessário inicializar e deixar as regras bem restritas. Use o **.eslintignore** para tornar o verificador mais leve

[ ] Use linguagem tipada de preferência (TypeScript para frameworks JavaScript), sempre bem estrito (observação: usar o _strictNullChecks_ e _strictMode_ como **true**)

- Sincronize o _ESLint_ com o _Typescript_, adicionando plugins específicos no arquivo _eslintrc_ para utilizar as vantagens de detecção do _lint_ em problemas de _Typescript_

[ ] Inicialize e use um formatador de código para equipe seguir mesmo padrão, exemplo do _Prettier_ que disponibiliza criar um arquivo local para ter a mesma configuração para qualquer desenvolvedor

[ ] Sincronize o _ESLint_ com o _Prettier_. A integração do ESLint com o Prettier garante que a formatação seja aplicada de acordo com as regras do verificador de código, evitando conflitos entre as regras de formatação e as regras de linting

[ ] (OPCIONAL) Usar script para não deixar commit se houver erros nos code checkers e testes, validadores e build. No _Node_ existe a biblioteca _Husky_ para essa funcionalidade

[ ] (OPCIONAL) Usar Biblioteca para padrão de _commits_, no _Node_ existe a [commitlint](https://github.com/conventional-changelog/commitlint/#what-is-commitlint)

[ ] Defina a arquitetura e a organização de pastas que será utilizada em seu projeto de acordo com os requisitos. **Utilizar os princípios do Clean Arch, SOLID e Scream Architecture é recomendado!**

[ ] Use/configure o protocole Git e configure o .gitignore

[ ] _Dockerize_ a aplicação, importante conferir se todos os serviços estão funcionando e se comunicando como deveriam, para desenvolvimento existe a extensão dev contêiner

<div id="manter"></div>

## **Manter projeto**

[ ] Deve possuir testes e executá-los constantemente de forma automatizada, alguns tipos de testes diferentes dos básicos Unit, integration e E2E:

[ ] Deve ter uma boa observabilidade

[ ] Deve ser seguro

- Estude DevOps da plataforma de versionamento que irá utilizar. GitHub Actions é uma sugestão de estudo, permite protejer as branchs principais do seu projeto no Git contra merges indesejáveis, com ações de executar build, criação de jobs, testes automatizados e verificações por alguém com mais permissões, etc

[ ] Deve ser otimizado e escalável

[ ] Tem que ser legível, sustentável e organizado

[ ] Deve possuir uma boa documentação, um bom _Readme_, sobre o versionamento (_changelog.md_), rotas, componentes, arquitetura, banco de dados (migrations), etc.

[ ] Atualizado: Manutenção de dependências e atualização de versões, mantenha as dependências do projeto atualizadas, incluindo bibliotecas, frameworks e plugins utilizados. Isso é importante para manter a segurança e a compatibilidade

[ ] Ser versionado: utilize o conceito de versionamento semântico em conjuntos com o releases/tags e mantenha um arquivo _changelog.md_ atualizado do projeto sobre as atualizações

[ ] Deve estar por dentro das leis de proteção dos dados (Brasil LGPD), atribuir os créditos necessários a fontes externas, informar usuário sobre _copyright_, GDPR (Regulamento Geral de Proteção de Dados) ou outras regulamentações aplicáveis

<div id="manter-back"></div>

### **Manter projeto - Back**

[ ] Gerenciamento de logs

- (Sistema): Implemente um sistema de gerenciamento de logs para registrar eventos importantes e monitorar o funcionamento/desempenho do back-end. Isso pode ajudar na detecção de problemas, depuração, análise de desempenho e a disponibilidade do serviço.

- (Segurança): Implemente um sistema de gerenciamento de logs para registrar eventos importantes relacionados a segurança, como tentativas de login, cadastro de dispositivos para 2FA, Datas de logins, Localização, Bloqueio/Desbloqueio de conta, etc.

**OBS**: Os loggers devem ser tanto para casos de sucesso quanto para erros. Tanto em persistência de dados (banco de dados) quanto para desenvolvimento no console, quem está tentando acessar tal rota, qual método chamou, data da chamada, tempo de duração, statusCode da resposta, etc.

[ ] Utilize interceptadores de requisições padrão, para adicionar um comportamento comum a todas as requisições que passam pelo aplicativo. Esses interceptadores permitem interceptar as requisições antes que elas sejam tratadas pelos controladores ou serviços e executar ações adicionais, como monitoramento de erros, rastreamento, registro de métricas, autenticação, validação, entre outros

[ ] Utilize filtros padrões de exceções, para sempre retornar na mesma estrutura

<div id="seguranca"></div>

## **Segurança**

### **DevOps**

[ ] Defina e ative a proteção das branchs principais de desenvolvimento e produção, desabilitando o '_force pushing_', a funcionalidade de remover a branch e a execução de ações antes de merges, como execução de testes, build, etc.

[ ] Ative todas as funcionalidades de segurança disponíveis na plataforma de hospedagem, por exemplo: '_Security Policy_', '_Security advisories_', '_Dependabot alerts_' e '_Code scanning alerts_'

### **JWT (JSON Web Token)**

[ ] Não use **Basic Auth**. Em vez disso, use a autenticação padrão (por exemplo, **JWT**, **OAuth**).

[ ] Não reinvente a roda em **Autenticação**, **geração de token**, **armazenamento de senha**. Use os padrões.

[ ] Use uma chave complicada aleatória (**JWT Secret**) para dificultar muito a força bruta do token.

[ ] Não extraia o algoritmo do cabeçalho. Forçar o algoritmo no back-end (**HS256** ou **RS256**).

[ ] Torne a expiração do token (**TTL**, **RTTL**) o mais curta possível.

[ ] Não armazene dados confidenciais na carga JWT, eles podem ser decodificados facilmente.

[ ] Evite armazenar muitos dados. O JWT geralmente é compartilhado em cabeçalhos e eles têm um limite de tamanho.

### **OAuth**

[ ] Sempre valide **redirect_uri** do lado do servidor para permitir apenas URLs da lista de permissões.

[ ] Sempre tente trocar por código e não tokens (não permita **response_type=token**).

[ ] Use o parâmetro de estado com um hash aleatório para evitar CSRF no processo de autenticação OAuth.

[ ] Defina o escopo padrão e valide os parâmetros do escopo para cada aplicativo.

### **Senhas**

Ao armazenar senhas em um sistema de login, é crucial adotar boas práticas de segurança para proteger as informações confidenciais dos usuários. Abaixo estão algumas recomendações para lidar com senhas e logins de forma segura em programação:

- **Hashing de senha**: Em vez de armazenar senhas em texto simples , é recomendado usar uma função de hash para transformar as senhas em uma sequência de caracteres irreversível. Algoritmos populares de hash, como bcrypt, PBKDF2 ou Argon2, são recomendados porque adicionam camadas adicionais de segurança, como "salt" (um valor aleatório único adicionado à senha antes de ser hash) e custo computacional ajustável.

- **Salt**: Um "salt" é uma sequência de caracteres aleatórios única gerada para cada usuário. O "salt" é concatenado à senha antes de ser hash. Isso impede que os ataques de dicionário ou de tabela arco-íris (rainbow table) sejam eficazes, pois cada senha terá um hash diferente, mesmo que as senhas sejam idênticas.

- **Iterações de hash**: Para aumentar a segurança, é recomendável executar várias iterações do algoritmo de hash. Quanto mais iterações, mais lento será o processo de hashing, dificultando o ataque de força bruta. No entanto, é necessário encontrar um equilíbrio entre segurança e desempenho.

- **Armazenamento seguro**: Os hashes de senha resultantes devem ser armazenados em um local seguro, como um banco de dados, protegidos contra acesso não autorizado.

- **Autenticação de dois fatores (2FA)**: Considere a implementação de autenticação de dois fatores para adicionar uma camada extra de segurança. O 2FA exige que os usuários forneçam uma segunda forma de autenticação, como um código enviado para um dispositivo móvel, além da senha.

- **Transmissão segura**: Ao enviar senhas pela rede, certifique-se de que a comunicação esteja protegida por meio de criptografia, como o uso de HTTPS (SSL/TLS).

- **Políticas de senha fortes**: Estabeleça requisitos para as senhas dos usuários, como comprimento mínimo, uso de caracteres especiais, letras maiúsculas e minúsculas e números. Isso incentiva os usuários a escolher senhas mais seguras.

- **Proteção contra ataques de força bruta**: Implemente medidas para impedir ataques de força bruta, como bloquear temporariamente contas após várias tentativas de login malsucedidas ou introduzir um mecanismo de captchas.

- **Criptografia de dados em repouso**: Além de proteger as senhas, é importante criptografar os dados em repouso, ou seja, os dados armazenados em disco ou em um banco de dados. Isso ajuda a garantir que, mesmo que alguém tenha acesso físico aos dados, eles não possam ser lidos sem a chave de criptografia correta.

- **Auditoria de logs**: Mantenha registros de log detalhados como endereço IP, dispositivo e localização de eventos relacionados ao login, como tentativas de login, alterações de senha e outras atividades relacionadas à autenticação. Isso pode ser útil para identificar possíveis ataques ou comportamentos suspeitos.

- **Alertas de login**: Enviar notificações por e-mail ou mensagem para o usuário informando sobre qualquer atividade de login suspeita, como um novo dispositivo ou local desconhecido.

- **Práticas de desenvolvimento seguro**: Siga boas práticas de desenvolvimento seguro ao criar sistemas de login, como evitar injeção de SQL, XSS (Cross-Site Scripting) e outros tipos comuns de vulnerabilidades.

- **Testes de segurança**: Realize testes regulares de segurança em seu sistema de login para identificar possíveis vulnerabilidades. Isso pode incluir testes de penetração, análise de código e verificação de conformidade com as melhores práticas de segurança.

<div id="backend"></div>

## **Geral Back-End**

#### Sugestão estrutural de cadastro

```ts
interface Autenticacao {
  idAutenticacao: number;
  tokenAcesso?: string;
  dthExpiracaoTokenAcesso?: DateTime | null;
  tokenAtualizacao?: string | null;
  tokenRedefinicaoSenha?: string | null;
  dthUltimaSolicitacaoRedefinicaoSenha?: DateTime | null;
  tokenDesbloquearConta?: string | null;
  dthUltimoBloqueioConta?: DateTime | null;
  codigoVerificacaoEmail?: string | null;
  emailVerificado?: DateTime | null;
  dthUltimaAutenticacao?: DateTime | null;
  provedorAutenticacao?: string | null; // Google, Facebook, Local
  ipUsual?: string | null;
  metodoAutenticacaoPreferencial?: string | null;
  usuario: Usuario;
  idUsuario: string;
  dispositivos2FA: Dispositivo2FA[];
}

interface Usuario {
  idUsuario: string;
  login: string;
  senha: string;
  salt: string;
  iteracoesHash: number;
  nome: string;
  email: string;
  telefone?: string | null;
  dthUltimaAlteracao: DateTime;
  usuarioUltimaAlteracao: string;
  dthUltimoLogin?: DateTime | null;
  dthCriacao: DateTime;
  usuarioCriacao: string;
  tentativasLogin: number;
  contaBloqueada: boolean;
  dthBloqueioTemporario?: DateTime | null;
  razaoBloqueioTemporario: string | null;
  limiteTentativasLogin: number;
  logins: HistoricoLogins[];
  atividades: HistoricoAtividades[];
  erros: HistoricoErros[];
  autenticacao?: Autenticacao;
}

interface Dispositivo2FA {
  idDispositivo: number;
  codigoVerificacao: string;
  tipoDispositivo: string;
  nomeDispositivo: string;
  identificador: string;
  autenticacao: Autenticacao;
  idAutenticacao: number;
}

interface HistoricoLogins {
  idLogin: number;
  idUsuario: string;
  usuario: Usuario;
  dthLogin: DateTime;
  loginBemSucedido: boolean;
  enderecoIP: string;
  userAgent?: string | null;
  nomeDispositivo?: string | null;
  cidade?: string | null;
  estado?: string | null;
  cep?: string | null;
}

interface HistoricoAtividades {
  idAtividade: number;
  idUsuario: string;
  usuario: Usuario;
  acao: string;
  dthAtividade: DateTime;
}

interface HistoricoErros {
  idErro: number;
  idUsuario?: string | null;
  usuario?: Usuario | null;
  codigoErro: number;
  tipoErro?: string | null;
  mensagemErro: string;
  dthErro: DateTime;
}
```

#### Entenda Decorators, Middlewares, Guards, Pipes e Cacheamento

- Decorators: Decoradores são uma característica comum em linguagens de programação que permitem adicionar funcionalidade a funções ou métodos existentes de forma transparente. Eles são usados para modificar ou estender o comportamento de uma função sem modificar seu código subjacente. Em essência, um decorator é uma função que envolve outra função e executa algum código antes ou depois dela.

- Middlewares: Middlewares são componentes frequentemente usados em estruturas de desenvolvimento web para tratar solicitações HTTP antes que elas alcancem a parte principal do aplicativo. Eles podem realizar tarefas como **autenticação**, **autorização**, **manipulação de erros**, **registro de logs**, entre outras. Middlewares são uma forma de separar preocupações em um aplicativo e facilitar a reutilização de código. OBS: Vem antes dos guards, então autenticação e autorização pode ser feito nos _guards_, enquanto o middleware pode garantir se os dados da requisição de autorização é de fato um e-mail por exemplo, ou seja, testar se são do tipo e estão dentro do padrão certo, caso não for já corta a requisição sem nem passar pelo Guard.

- Guards: Guardas são uma abordagem comum em frameworks de desenvolvimento web para proteger rotas ou endpoints de acesso não autorizado. Eles verificam certas condições antes de permitir que uma solicitação alcance a rota desejada. Por exemplo, um guarda pode verificar se um usuário está autenticado ou tem permissões adequadas para acessar determinada página.

- Pipes Transformadores: Imagine um servidor web que recebe solicitações HTTP e precisa processar essas solicitações antes de enviá-las para um banco de dados. Um pipeline de transformação poderia incluir etapas como autenticação, validação de entrada, registro de logs e conversão de dados antes de enviar a solicitação ao banco de dados.

- Pipes de Validação: Em um sistema de processamento de pagamentos, você pode ter um pipeline de validação que verifica cada transação antes de processá-la. Esse pipeline poderia incluir etapas como verificação do número do cartão de crédito, verificação de saldo na conta do cliente e autenticação do cliente antes de aprovar a transação.

- Cacheamento: Cacheamento é uma técnica de armazenamento temporário de dados frequentemente acessados em um local de fácil acesso, como memória RAM ou armazenamento rápido, para acelerar o acesso subsequente a esses dados. Um exemplo, quando uma requisição é feita e seu resultado é armazenado em cache, subsequentes requisições dentro de um intervalo de tempo X podem ser atendidas diretamente a partir da cache, economizando tempo e recursos, em vez de realizar a operação novamente.
