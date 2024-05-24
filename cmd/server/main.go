package main

import (
	"log"
	"net/http"
	"os"

	"graphql-comments/internal/app/comment"
	"graphql-comments/internal/app/post"
	"graphql-comments/internal/config"
	"graphql-comments/internal/graph"
	"graphql-comments/internal/storage/memory"
	"graphql-comments/internal/storage/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cfg := config.LoadConfig()

	var postRepo post.Repository
	var commentRepo comment.Repository

	if cfg.Storage == "memory" {
		postRepo = memory.NewPostRepository()
		commentRepo = memory.NewCommentRepository()
	} else {
		db, err := postgres.InitDB(cfg)
		if err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}
		postRepo = postgres.NewPostgresPostRepository(db)
		commentRepo = postgres.NewPostgresCommentRepository(db)
	}

	postService := post.NewService(postRepo)
	commentService := comment.NewService(commentRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		PostService:    postService,
		CommentService: commentService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
