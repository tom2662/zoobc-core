package block

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zoobc/zoobc-core/common/chaintype"
	"github.com/zoobc/zoobc-core/common/crypto"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/transaction"
	"github.com/zoobc/zoobc-core/common/util"
	"github.com/zoobc/zoobc-core/observer"

	"github.com/zoobc/zoobc-core/common/query"
	"github.com/zoobc/zoobc-core/core/service"

	"github.com/zoobc/zoobc-core/common/database"

	"github.com/zoobc/zoobc-core/core/smith"

	"github.com/spf13/cobra"
)

var (
	blocksmith              *model.Blocksmith
	chainType               chaintype.ChainType
	blockProcessor          smith.BlockchainProcessorInterface
	blockService            service.BlockServiceInterface
	nodeRegistrationService service.NodeRegistrationServiceInterface
	sortedBlocksmiths       []model.Blocksmith
	queryExecutor           query.ExecutorInterface
	migration               database.Migration

	numberOfBlocks         int
	blocksmithSecretPhrase string
	outputPath             string

	blockCmd = &cobra.Command{
		Use:   "block",
		Short: "block command used to manipulate block of node",
		Long: `
			block command is use to manipulate block creation or broadcasting in the node
		`,
	}

	fakeBlockCmd = &cobra.Command{
		Use:   "fake-blocks",
		Short: "fake-blocks command used to create fake blocks",
		Run: func(cmd *cobra.Command, args []string) {
			generateBlocks(numberOfBlocks, blocksmithSecretPhrase, outputPath)
		},
	}
)

func init() {
	fakeBlockCmd.Flags().IntVar(
		&numberOfBlocks,
		"numberOfBlocks",
		100,
		"number of account to generate",
	)
	fakeBlockCmd.Flags().StringVar(
		&blocksmithSecretPhrase,
		"blocksmithSecretPhrase",
		"",
		"secret phrase of blocksmith | required",
	)
	fakeBlockCmd.Flags().StringVar(
		&outputPath,
		"out",
		"./testdata/zoobc.db",
		"output path of the database",
	)
	blockCmd.AddCommand(fakeBlockCmd)
}

func Commands() *cobra.Command {
	return blockCmd
}

func initialize(
	secretPhrase, outputPath string,
) {
	paths := strings.Split(outputPath, "/")
	dbPath, dbName := strings.Join(paths[:len(paths)-1], "/")+"/", paths[len(paths)-1]
	chainType = &chaintype.MainChain{}
	observerInstance := observer.NewObserver()
	blocksmith = model.NewBlocksmith(
		secretPhrase,
		util.GetPublicKeyFromSeed(secretPhrase),
	)
	// initialize/open db and queryExecutor
	dbInstance := database.NewSqliteDB()
	if err := dbInstance.InitializeDB(dbPath, dbName); err != nil {
		panic(err)
	}
	db, err := dbInstance.OpenDB(dbPath, dbName, 10, 20)
	if err != nil {
		panic(err)
	}
	queryExecutor = query.NewQueryExecutor(db)
	actionSwitcher := &transaction.TypeSwitcher{
		Executor: queryExecutor,
	}
	mempoolService := service.NewMempoolService(
		chainType,
		nil,
		queryExecutor,
		query.NewMempoolQuery(chainType),
		actionSwitcher,
		query.NewAccountBalanceQuery(),
		crypto.NewSignature(),
		query.NewTransactionQuery(chainType),
		observerInstance,
		log.New(),
	)
	blockService = service.NewBlockService(
		chainType,
		nil,
		queryExecutor,
		query.NewBlockQuery(chainType),
		query.NewMempoolQuery(chainType),
		query.NewTransactionQuery(chainType),
		query.NewMerkleTreeQuery(),
		crypto.NewSignature(),
		mempoolService,
		actionSwitcher,
		query.NewAccountBalanceQuery(),
		query.NewParticipationScoreQuery(),
		query.NewNodeRegistrationQuery(),
		observerInstance,
		&sortedBlocksmiths,
		log.New(),
	)
	nodeRegistrationService = service.NewNodeRegistrationService(
		queryExecutor,
		query.NewAccountBalanceQuery(),
		query.NewNodeRegistrationQuery(),
		query.NewParticipationScoreQuery(),
		log.New(),
	)

	migration = database.Migration{Query: queryExecutor}
}

func generateBlocks(numberOfBlocks int, blocksmithSecretPhrase, outputPath string) {
	fmt.Println("initializing dependency and database...")
	initialize(blocksmithSecretPhrase, outputPath)
	fmt.Println("done initializing database")
	blockProcessor = smith.NewBlockchainProcessor(
		chainType,
		blocksmith,
		blockService,
		nodeRegistrationService,
	)
	startTime := time.Now().UnixNano() / 1e6
	fmt.Printf("generating %d blocks\n", numberOfBlocks)
	fmt.Println("initializing database schema migration")
	if err := migration.Init(); err != nil {
		panic(err)
	}

	fmt.Println("applying database schema migration")
	if err := migration.Apply(); err != nil {
		panic(err)
	}
	fmt.Println("checking genesis...")
	if !blockService.CheckGenesis() { // Add genesis if not exist
		fmt.Println("genesis does not exist, adding genesis")
		// genesis account will be inserted in the very beginning
		if err := service.AddGenesisAccount(queryExecutor); err != nil {
			panic(err)
		}

		if err := blockService.AddGenesis(); err != nil {
			panic(err)
		}

		fmt.Println("begin generating blocks")
		if err := blockProcessor.FakeSmithing(numberOfBlocks, true); err != nil {
			panic(err)
		}
	} else {
		// start from last block's timestamp
		fmt.Println("continuing from last database...")
		if err := blockProcessor.FakeSmithing(numberOfBlocks, false); err != nil {
			panic("error in appending block to existing database")
		}
	}
	fmt.Printf("database generated in %s", outputPath)
	fmt.Printf("block generation success in %d miliseconds", (time.Now().UnixNano()/1e6)-startTime)
}
