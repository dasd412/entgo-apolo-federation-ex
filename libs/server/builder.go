package server

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"scheduler"
	"sync"
	"user/pkg/ent"
)

// todo *ent.Client에서 "user/pkg/ent" 의존성 제거 필요
type AppContainer struct {
	EntClient        *ent.Client
	RedisClient      *redis.Client
	S3Client         *s3.Client
	SchedulerManager *scheduler.ScheduleManager
}

var (
	once        sync.Once // 한 번만 실행되는 초기화 함수를 구현하는 데 사용됨 (싱글톤 패턴)
	appInstance *AppContainer
)

//func GetAppContainer(ctx context.Context) (*AppContainer, error) {
//	var err error
//
//	once.Do(func() { //최초 호출 시 한 번만 실행됨.
//		appInstance, err = initApp(ctx)
//	})
//
//	return appInstance, err
//}
//
//func initApp(ctx context.Context) (*AppContainer, error) {
//	config.LoadEnv(".env")
//
//	masterCfg := &db.DBConfig{
//		Username: os.Getenv("RDS_USERNAME"),
//		Password: os.Getenv("RDS_PASSWORD"),
//		Hostname: os.Getenv("RDS_MASTER"),
//		DBName:   os.Getenv("RDS_DBNAME"),
//		Port:     os.Getenv("RDS_PORT"),
//	}
//
//	replicaCfg := &db.DBConfig{
//		Username: os.Getenv("RDS_USERNAME"),
//		Password: os.Getenv("RDS_PASSWORD"),
//		Hostname: os.Getenv("RDS_REPLICA"),
//		DBName:   os.Getenv("RDS_DBNAME"),
//		Port:     os.Getenv("RDS_PORT"),
//	}
//
//	entClient := db.NewClients[ent.Client](
//		*masterCfg,
//		*replicaCfg,
//		client.NewEntClient,
//	)
//
//	if err := entClient.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
//		return nil, fmt.Errorf("failed to run migration: %w", err)
//	}
//
//	// todo entClient말고 다른 client도 들어갈 수 있게 해야 함
//	// client를 추상화한 것에 entClient, s3Client가 여러 개 들어갈 수도 있음.
//	scheduleManager, err := setupSchedulers(ctx, entClient)
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to setup schedulers: %w", err)
//	}
//
//	dbName, err := strconv.Atoi(os.Getenv("REDIS_DBNAME"))
//
//	redisClient := redis.NewClient(
//		&redis.Options{
//			Addr:     os.Getenv("REDIS_HOST"),
//			Password: os.Getenv("REDIS_PASSWORD"),
//			DB:       dbName,
//		},
//	)
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
//	}
//
//	s3Client, err := client.NewS3Client(
//		ctx, client.AWSConfig{
//			Region:          os.Getenv("AWS_DEFAULT_REGION"),
//			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
//			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
//		},
//	)
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to create S3 client: %w", err)
//	}
//
//	return &AppContainer{
//		EntClient:        entClient,
//		RedisClient:      redisClient,
//		S3Client:         s3Client,
//		SchedulerManager: scheduleManager,
//	}, nil
//}
//
//func setupSchedulers(
//	ctx context.Context,
//	entClient *ent.Client,
//) (*scheduler.ScheduleManager, error) {
//	scheduleManager := scheduler.NewSchedulerManager()
//
//	schedulers := []scheduler.Scheduler{}
//
//	for _, s := range schedulers {
//		scheduleManager.AddScheduler(s)
//	}
//
//	if err := scheduleManager.StartAll(ctx); err != nil {
//		return nil, fmt.Errorf("failed to start schedulers: %v", err)
//	}
//
//	return scheduleManager, nil
//}
