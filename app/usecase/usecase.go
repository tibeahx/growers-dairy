package usecase

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/tibeahx/growers-dairy/app/s3"
	"github.com/tibeahx/growers-dairy/app/storage"
	"github.com/tibeahx/growers-dairy/app/types"
)

var (
	errIdRequired                 = errors.New("id required")
	errNameAndDescriptionRequired = errors.New("name and discription required")
)

type GrowLogService interface {
	GrowLogs(ctx context.Context) ([]types.GrowLog, error)
	CreateNewGrowLog(ctx context.Context, req types.CreateGrowLogReq) (types.CreateGrowLogResp, error)
	GrowLogByID(ctx context.Context, id int) (types.GrowLog, error)
	UpdateGrowLog(ctx context.Context, req types.UpdateGrowLogReq) (types.UpdateGrowLogResp, error)
	DeleteGrowLogByID(ctx context.Context, id int) error
}

type StrainService interface {
	Strains(ctx context.Context) ([]types.Strain, error)
	CreateStrain(ctx context.Context, req types.CreateStrainReq) (types.CreateStrainReq, error)
	StrainByID(ctx context.Context, id int) (types.Strain, error)
	UpdateStrain(ctx context.Context, req types.UpdateStrainReq) (types.UpdateStrainResp, error)
	DeleteStrainByID(ctx context.Context, id int) error
}

type LogEntryService interface {
	LogEntries(ctx context.Context) ([]types.LogEntry, error)
	CreateLogEntry(ctx context.Context, req types.CreateLogEntryReq) (types.CreateLogEntryResp, error)
	EntryByID(ctx context.Context, id int) (types.LogEntry, error)
	UpdateEntry(ctx context.Context, req types.UpdateLogEntryReq) (types.UpdateLogEntryResp, error)
	DeleteEntryByID(ctx context.Context, id int) error
}

type S3Service interface {
	GetFileFromForm(ctx context.Context, r *http.Request) (multipart.File, *multipart.FileHeader)
	UploadFileToS3(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error)
	GetPresignedURL(ctx context.Context, bucketName string, objectName string) (string, error)
}

type ServiceProvider struct {
	GrowLogService
	StrainService
	LogEntryService
	S3Service
	db *storage.DB
}

func NewServiceProvider(db *storage.DB) *ServiceProvider {
	return &ServiceProvider{
		db: db,
	}
}

func (s *ServiceProvider) CreateStrain(ctx context.Context, req types.CreateStrainReq) (types.CreateStrainResp, error) {
	if len(req.Name) == 0 || len(req.Description) == 0 {
		return types.CreateStrainResp{}, errNameAndDescriptionRequired
	}
	if err := types.AssertStrainAttrsExist(req); err != nil {
		return types.CreateStrainResp{}, err
	}
	if err := types.ValidateFeminizedType(req.StrainAttrs.FeminizedType); err != nil {
		return types.CreateStrainResp{}, err
	}
	if err := types.ValidateStrainType(req.StrainAttrs.StrainType); err != nil {
		return types.CreateStrainResp{}, err
	}
	resp, err := s.db.CreateStrain(ctx, req)
	if err != nil {
		return types.CreateStrainResp{}, err
	}
	return resp, err
}

func (s *ServiceProvider) StrainByID(ctx context.Context, id int) (types.Strain, error) {
	if id <= 0 {
		return types.Strain{}, errIdRequired
	}
	strain, err := s.db.StrainByID(ctx, id)
	if err != nil {
		return types.Strain{}, err
	}
	return strain, nil
}

func (s *ServiceProvider) DeleteStrainByID(ctx context.Context, id int) error {
	if id <= 0 {
		return errIdRequired
	}
	if err := s.db.DeleteStrainByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *ServiceProvider) Strains(ctx context.Context) ([]types.Strain, error) {
	strains, err := s.db.Strains(ctx)
	if err != nil {
		return nil, err
	}
	return strains, nil
}

func (s *ServiceProvider) UpdateStrain(ctx context.Context, req types.UpdateStrainReq) (types.UpdateStrainResp, error) {
	if req.ID <= 0 {
		return types.UpdateStrainResp{}, errIdRequired
	}
	if err := types.AssertStrainAttrsExist(req); err != nil {
		return types.UpdateStrainResp{}, err
	}
	if err := types.ValidateFeminizedType(req.StrainAttrs.FeminizedType); err != nil {
		return types.UpdateStrainResp{}, err
	}
	if err := types.ValidateStrainType(req.StrainAttrs.StrainType); err != nil {
		return types.UpdateStrainResp{}, err
	}
	updated, err := s.db.UpdateStrain(ctx, req)
	if err != nil {
		return types.UpdateStrainResp{}, err
	}
	return updated, nil
}

func (s *ServiceProvider) CreateGrowLog(ctx context.Context, req types.CreateGrowLogReq) (types.CreateGrowLogResp, error) {
	if len(req.Name) == 0 || len(req.Description) == 0 {
		return types.CreateGrowLogResp{}, errNameAndDescriptionRequired
	}
	newGrowlog, err := s.db.CreateGrowLog(ctx, req)
	if err != nil {
		return types.CreateGrowLogResp{}, err
	}
	return newGrowlog, nil
}

func (s *ServiceProvider) GrowLogs(ctx context.Context) ([]types.GrowLog, error) {
	growLogs, err := s.db.GrowLogs(ctx)
	if err != nil {
		return nil, err
	}
	return growLogs, nil
}

func (s *ServiceProvider) GrowLogByID(ctx context.Context, id int) (types.GrowLog, error) {
	if id <= 0 {
		return types.GrowLog{}, errIdRequired
	}
	growLog, err := s.db.GrowLogByID(ctx, id)
	if err != nil {
		return types.GrowLog{}, err
	}
	return growLog, nil
}

func (s *ServiceProvider) DeleteGrowLogByID(ctx context.Context, id int) error {
	return s.db.DeleteGrowLogByID(ctx, id)
}

func (s *ServiceProvider) UpdateGrowLog(ctx context.Context, req types.UpdateGrowLogReq) (types.UpdateGrowLogResp, error) {
	if req.ID <= 0 {
		return types.UpdateGrowLogResp{}, errIdRequired
	}
	updatedGrowLog, err := s.db.UpdateGrowLog(ctx, req)
	if err != nil {
		return types.UpdateGrowLogResp{}, err
	}
	return updatedGrowLog, err
}

func (s *ServiceProvider) LogEntries(ctx context.Context) ([]types.LogEntry, error) {
	logEntries, err := s.db.LogEntries(ctx)
	if err != nil {
		return nil, err
	}
	return logEntries, nil
}

func (s *ServiceProvider) CreateLogEntry(ctx context.Context, req types.CreateLogEntryReq) (types.CreateLogEntryResp, error) {
	logEntries, err := s.db.CrateLogEntry(ctx, req)
	if err != nil {
		return types.CreateLogEntryResp{}, err
	}
	return logEntries, nil
}

func (s *ServiceProvider) EntryByID(ctx context.Context, id int) (types.LogEntry, error) {
	if id <= 0 {
		return types.LogEntry{}, errIdRequired
	}
	logEntry, err := s.db.EntryByID(ctx, id)
	if err != nil {
		return types.LogEntry{}, err
	}
	return logEntry, nil
}

func (s *ServiceProvider) UpdateEntry(ctx context.Context, req types.UpdateLogEntryReq) (types.UpdateLogEntryResp, error) {
	updatedEntry, err := s.db.UpdateEntry(ctx, req, "", "")
	if err != nil {
		return types.UpdateLogEntryResp{}, err
	}
	return updatedEntry, nil
}

func (s *ServiceProvider) DeleteEntryByID(ctx context.Context, id int) error {
	return s.db.DeleteEntryByID(ctx, id)
}

func (s *ServiceProvider) GetFileFromForm(r *http.Request) (multipart.File, *multipart.FileHeader) {
	file, header, _ := r.FormFile("photo")
	return file, header
}

func (s *ServiceProvider) UploadFileToS3(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	client := s3.NewS3Client()
	bucketName := "sjtorage"
	objectName := header.Filename

	_, err := client.PutObject(ctx,
		bucketName,
		objectName,
		file,
		header.Size,
		minio.PutObjectOptions{
			ContentType: header.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}

	url, err := client.PresignedGetObject(ctx, bucketName, objectName, time.Hour*168, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (s *ServiceProvider) LinkURLToEntry(ctx context.Context, r *http.Request, logEntryID int, url string) error {
	file, header := s.GetFileFromForm(r)
	defer file.Close()

	url, err := s.UploadFileToS3(ctx, file, header)
	if err != nil {
		return err
	}
	if err = s.db.InsertS3Link(ctx, logEntryID, url); err != nil {
		return err
	}
	return nil
}
