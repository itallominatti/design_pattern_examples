interface Veiculo {
    buscarCliente(): void;
}

class CarroLuxo implements Veiculo {
    buscarCliente() {
        console.log('Carro de luxo buscando cliente');
    }
}

class CarroPopular implements Veiculo {
    buscarCliente() {
        console.log('Carro popular buscando cliente');
    }
}

class Moto implements Veiculo {
    buscarCliente() {
        console.log('Moto buscando cliente');
    }
}

class VeiculoFactory {
    private carro: Veiculo;

    constructor(tipo: string) {
        this.carro = this.getCarro(tipo);
    }

    private getCarro(tipo: string): Veiculo {
        switch (tipo) {
            case 'luxo':
                return new CarroLuxo();
            case 'popular':
                return new CarroPopular();
            case 'moto':
                return new Moto();
            default:
                throw new Error('Tipo de veículo desconhecido');
        }
    }

    buscarCliente() {
        this.carro.buscarCliente();
    }
}

function main() {
    // Coloque no parâmetro do VeiculoFactory o tipo de veículo que deseja
    const veiculo = new VeiculoFactory('luxo');
    veiculo.buscarCliente();
}

main();