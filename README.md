# capital-gains

Nubank Code Challenge. Este projeto calcula taxas associadas aos ganhos de capital com base em operações de compra e venda de ações.

## Arquitetura e Código

Adotou-se na medida do possível clean architecture e princípios SOLID. Dado o contexto diminuído do sistema, evitamos rigor desnecessário que acabaria por nos afastar do objetivo principal do teste.

## Executando o Projeto Localmente

Para executar este projeto em sua máquina local, siga estas etapas:

1. **Configuração do Ambiente**
   - Certifique-se de ter o Go instalado em sua máquina numa versão compatível com a 1.22.3.

2. **Instalação de Dependências**
   Execute o seguinte comando para instalar as ferramentas necessárias e as dependências do projeto:
   ```
   make setup
   ```

3. **Compilação do Projeto**
   Para compilar o projeto, execute:
   ```
   make build
   ```

4. **Execução do Projeto**
   Para observar o aplicativo em execução de acordo com os casos de uso da avaliação, use:
   ```
   make run.sample
   ```

5. **Testes**
   Para executar os testes do projeto:
   ```
   make test
   ```

## Licença
Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.