package post

import (
	"context"
	"time"

	"github.com/mberbero/go-microservice-template/pkg/dtos"
	"github.com/mberbero/go-microservice-template/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	Create(post *entities.Post) error
	Update(id string, post *dtos.PostDTO) error
	Delete(id string) error
	Get(id string) (*entities.Post, error)
	GetAll(page, perPage int64) (*dtos.PaginatedData, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) Create(post *entities.Post) error {
	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.TODO(), post)
	return err
}

func (r *repository) Update(id string, post *dtos.PostDTO) error {
	filter := primitive.M{
		"_id": id,
	}
	_, err := r.Collection.UpdateOne(context.Background(), filter, post)
	return err
}

func (r *repository) Delete(id string) error {
	_, err := r.Collection.DeleteOne(context.TODO(), id)
	return err
}

func (r *repository) Get(id string) (*entities.Post, error) {
	var post entities.Post
	err := r.Collection.FindOne(context.TODO(), id).Decode(&post)
	return &post, err
}

func (r *repository) GetAll(page, perPage int64) (*dtos.PaginatedData, error) {
	var post entities.Post
	var posts []*entities.Post
	var paginatedData *dtos.PaginatedData
	filter := bson.M{}
	findOptions := options.Find()

	total, _ := r.Collection.CountDocuments(context.TODO(), filter)
	totalPages := int(total / int64(perPage))

	findOptions.SetSkip((int64(page) - 1) * int64(perPage))
	findOptions.SetLimit(int64(perPage))

	cursor, _ := r.Collection.Find(context.TODO(), filter, findOptions)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		cursor.Decode(&post)
		posts = append(posts, &post)
	}
	paginatedData = &dtos.PaginatedData{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
		Data:       posts,
	}

	return paginatedData, nil
}
