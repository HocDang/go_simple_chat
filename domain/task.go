package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type TaskRequest struct {
	Domain  string `form:"domain" binding:"required,url" json:"domain"`
	HookURL string `form:"hookUrl" binding:"required,url" json:"hookUrl"`
}

type Task struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Domain         string             `bson:"domain" form:"domain" binding:"required" json:"domain"`
	HookURL        string             `bson:"hook_url" form:"hookUrl" binding:"required" json:"hookUrl"`
	Status         int                `bson:"status" form:"status" json:"status" default:"0"`
	DomainRedirect *string            `bson:"domain_redirect,omitempty" form:"domain_redirect" json:"domainRedirect"`
	Message        *string            `bson:"message,omitempty" form:"message" json:"message"`
	ImagePath      *string            `bson:"image_path,omitempty" form:"image_path" json:"imagePath"`
	VideoPath      *string            `bson:"video_path,omitempty" form:"video_path" json:"videoPath"`
	ImageCloudPath *string            `bson:"image_cloud_path,omitempty" form:"image_cloud_path" json:"imageCloudPath"`
	VideoCloudPath *string            `bson:"video_cloud_path,omitempty" form:"video_cloud_path" json:"videoCloudPath"`
	UpdateAt       *time.Time         `bson:"update_at,omitempty" json:"updateAt"`
	CreatedAt      time.Time          `bson:"create_at" json:"createdAt"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	Fetch(c context.Context) ([]Task, error)
	GetByID(c context.Context, GetByID string) (*Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	Fetch(c context.Context) ([]Task, error)
	GetByID(c context.Context, GetByID string) (*Task, error)
}
