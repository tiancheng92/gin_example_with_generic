package generic

import (
	"context"
	"gin_example_with_generic/types/paginate"
)

func NewService[R RequestInterface, M ModelInterface](repo RepositoryInterface[M]) *Service[R, M] {
	return &Service[R, M]{repo}
}

type Service[R RequestInterface, M ModelInterface] struct {
	RepositoryInterface[M]
}

func (s *Service[R, M]) Get(ctx context.Context, pk any) (*M, error) {
	return s.RepositoryInterface.Get(ctx, pk)
}

func (s *Service[R, M]) List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error) {
	return s.RepositoryInterface.List(ctx, pq)
}

func (s *Service[R, M]) Update(ctx context.Context, pk any, request *R) (*M, error) {
	return s.RepositoryInterface.Update(ctx, pk, (*request).FormatToModel().(M))
}

func (s *Service[R, M]) Create(ctx context.Context, request *R) (*M, error) {
	return s.RepositoryInterface.Create(ctx, (*request).FormatToModel().(M))
}

func (s *Service[R, M]) Delete(ctx context.Context, pk any) error {
	return s.RepositoryInterface.Delete(ctx, pk)
}

type ServiceInterface[R RequestInterface, M ModelInterface] interface {
	Get(ctx context.Context, pk any) (*M, error)
	List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error)
	Update(ctx context.Context, pk any, request *R) (*M, error)
	Create(ctx context.Context, request *R) (*M, error)
	Delete(ctx context.Context, pk any) error
}

func NewReadOnlyService[M ModelInterface](repo RepositoryInterface[M]) *ReadOnlyService[M] {
	return &ReadOnlyService[M]{repo}
}

type ReadOnlyService[M ModelInterface] struct {
	RepositoryInterface[M]
}

func (ros *ReadOnlyService[M]) Get(ctx context.Context, pk any) (*M, error) {
	return ros.RepositoryInterface.Get(ctx, pk)
}

func (ros *ReadOnlyService[M]) List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error) {
	return ros.RepositoryInterface.List(ctx, pq)
}

type ReadOnlyServiceInterface[M ModelInterface] interface {
	Get(ctx context.Context, pk any) (*M, error)
	List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error)
}
