// 1. Definimos a interface do DTO idêntica ao UserRequest do Go
interface UserRequest {
    name: string;
    email: string;
    senha: string;
}

// 2. Capturamos os elementos do DOM com tipagem do TypeScript
const form = document.getElementById('cadastroForm') as HTMLFormElement;
const msgDiv = document.getElementById('mensagem') as HTMLDivElement;

// 3. Ouvimos o evento de submit do formulário
form.addEventListener('submit', async (event: Event) => {
    event.preventDefault(); // Impede a página de recarregar

    // Captura os valores dos inputs
    const nomeInput = document.getElementById('nome') as HTMLInputElement;
    const emailInput = document.getElementById('email') as HTMLInputElement;
    const senhaInput = document.getElementById('senha') as HTMLInputElement;

    // Monta o objeto de requisição seguindo o contrato do DTO
    const dadosUsuario: UserRequest = {
        name: nomeInput.value,
        email: emailInput.value,
        senha: senhaInput.value
    };

    try {
        // Faz o disparo elétrico (HTTP POST) para a API em Go
        const response = await fetch('http://localhost:8080/usuary', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dadosUsuario) // Converte o objeto TS para string JSON
        });

        // Se a API retornar sucesso (status 201)
        if (response.ok) {
            const usuarioCriado = await response.json();
            
            exibirMensagem(`✅ Usuário criado com sucesso! ID: ${usuarioCriado.id}`, 'success');
            form.reset(); // Limpa o formulário
        } else {
            // Se der erro (ex: nome vazio retornado pelo service)
            const erroTexto = await response.text();
            exibirMensagem(`❌ Erro: ${erroTexto}`, 'error');
        }
    } catch (error) {
        // Erro de rede (ex: se a API em Go estiver desligada)
        exibirMensagem('❌ Erro ao conectar com o servidor backend.', 'error');
    }
});

// Função utilitária para exibir alertas na tela
function exibirMensagem(texto: string, tipo: 'success' | 'error') {
    msgDiv.innerText = texto;
    msgDiv.className = tipo; // Aplica a classe CSS correspondente
}