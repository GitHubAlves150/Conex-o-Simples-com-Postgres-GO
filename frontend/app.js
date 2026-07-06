"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
// 2. Capturamos os elementos do DOM com tipagem do TypeScript
const form = document.getElementById('cadastroForm');
const msgDiv = document.getElementById('mensagem');
// 3. Ouvimos o evento de submit do formulário
form.addEventListener('submit', (event) => __awaiter(void 0, void 0, void 0, function* () {
    event.preventDefault(); // Impede a página de recarregar
    // Captura os valores dos inputs
    const nomeInput = document.getElementById('nome');
    const emailInput = document.getElementById('email');
    const senhaInput = document.getElementById('senha');
    // Monta o objeto de requisição seguindo o contrato do DTO
    const dadosUsuario = {
        name: nomeInput.value,
        email: emailInput.value,
        senha: senhaInput.value
    };
    try {
        // Faz o disparo elétrico (HTTP POST) para a API em Go
        const response = yield fetch('http://localhost:8080/usuary', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dadosUsuario) // Converte o objeto TS para string JSON
        });
        // Se a API retornar sucesso (status 201)
        if (response.ok) {
            const usuarioCriado = yield response.json();
            exibirMensagem(`✅ Usuário criado com sucesso! ID: ${usuarioCriado.id}`, 'success');
            form.reset(); // Limpa o formulário
        }
        else {
            // Se der erro (ex: nome vazio retornado pelo service)
            const erroTexto = yield response.text();
            exibirMensagem(`❌ Erro: ${erroTexto}`, 'error');
        }
    }
    catch (error) {
        // Erro de rede (ex: se a API em Go estiver desligada)
        exibirMensagem('❌ Erro ao conectar com o servidor backend.', 'error');
    }
}));
// Função utilitária para exibir alertas na tela
function exibirMensagem(texto, tipo) {
    msgDiv.innerText = texto;
    msgDiv.className = tipo; // Aplica a classe CSS correspondente
}
