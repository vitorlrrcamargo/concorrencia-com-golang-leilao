package integration

import (
	"context"
	"os"
	"testing"
	"time"

	auctionEntity "fullcycle-auction_go/internal/entity/auction_entity"
	auctionRepo "fullcycle-auction_go/internal/infra/database/auction"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionAutoClosure(t *testing.T) {
	// Define um intervalo curto de fechamento (para teste)
	_ = os.Setenv("AUCTION_INTERVAL", "5s")

	// Conecta ao MongoDB de teste com autenticação
	clientOpts := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/auctions?authSource=admin")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		t.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Usa um banco temporário para teste
	db := client.Database("auction_testdb")
	defer db.Drop(context.TODO()) // limpa após o teste

	repo := auctionRepo.NewAuctionRepository(db)

	// Cria um leilão novo
	newAuction, ierr := auctionEntity.CreateAuction(
		"Produto de Teste",
		"Categoria Teste",
		"Descrição do leilão teste com mais de 10 caracteres",
		auctionEntity.New,
	)
	if ierr != nil {
		t.Fatalf("Erro ao criar o leilão: %v", ierr)
	}

	ierr = repo.CreateAuction(context.TODO(), newAuction)
	if ierr != nil {
		t.Fatalf("Erro ao inserir o leilão no banco: %v", ierr)
	}

	// Aguarda o tempo de fechamento automático + margem
	time.Sleep(7 * time.Second)

	// Busca o leilão atualizado
	closedAuction, ierr := repo.FindAuctionById(context.TODO(), newAuction.Id)
	if ierr != nil {
		t.Fatalf("Erro ao buscar o leilão por ID: %v", ierr)
	}

	assert.Equal(t, auctionEntity.Completed, closedAuction.Status, "O leilão deveria estar com status 'Completed'")
}
