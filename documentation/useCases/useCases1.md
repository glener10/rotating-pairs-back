Usar bdd (Behavior driven development) para descrever as funcionalidades. O mesmo usa de linguagem natural para orientar o desenvolvimento ou até mesmo testes (não é exclusivo mas é útil).

Por causa da linguagem natural fica fácil de pessoas que não são da área entender a funcionalidade e seus caminhos e também os desenvolvedores.

MODO DESCRITIVO (RECOMENDADO)
```
Cenário: Cadastro de usuário realaizado com sucesso
  Dado que estou na tela de cadastro de usuário
  Quando o cadastro for realizado
  Então o usuário deverá ter acesso a aplicaçãao
```

MODO IMPERATIVO (NÃO RECOMENDADO PQ SE MUDA ALGO NO CÓDIGO A DOCUMENTAÇÃO NÃO É MAIS VÁLIDA OU DEPENDE DE ALTERAÇÃO TAMBÉM)
```
Cenário: Cadastro de usuário realaizado com sucesso
  Dado que estou na tela de cadastro de usuário
  Quando eu clico no botão "Novo usuário"
  E preencho o campo "email"
  E preencho o campo "senha"
  E clico no botão "salvar"
  Então receberei uma mensagem de "Usuário cadastrado com sucesso"
```